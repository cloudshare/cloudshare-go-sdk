package cloudshare

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"
)

func TestBuildURL(t *testing.T) {
	require.Equal(t, "https://use.cloudshare.com/api/v3/projects",
		buildURL("projects", nil).String(), "failed to build url")

	require.Equal(t, "https://use.cloudshare.com/api/v3/projects/",
		buildURL("/projects/", nil).String(), "failed to build url")

	params := &url.Values{}
	params.Set("key", "value_with/_in_it")
	require.Equal(t, "https://use.cloudshare.com/api/v3/path?key=value_with%2F_in_it", buildURL("path", params).String(), "url param encoding failed")
}

type PingResponse struct {
	Result string `json:"result"`
}

const testEnvName = "go-sdk-test-env"

var apikey, apiid, allowTestCreate = os.Getenv("CLOUDSHARE_API_KEY"), os.Getenv("CLOUDSHARE_API_ID"), os.Getenv("ALLOW_TEST_CREATE")

var c = &Client{
	APIKey: apikey,
	APIID:  apiid,
	Tags:   "go_sdk_test",
}

func skipNoAPIKeys(t *testing.T) {
	if apikey == "" || apiid == "" {
		t.Skipf("test only runs with actual credentials")
	}
}

func skipResourceCreation(t *testing.T) {
	if allowTestCreate != "true" {
		t.Skipf("Test that creates resources ($$$) skipped unless ALLOW_TEST_CREATE=true is defined")
	}
}

func TestPing(t *testing.T) {
	skipNoAPIKeys(t)

	res, apierr := c.Request("GET", "ping", nil, nil)
	require.Nil(t, apierr, "failed to ping")
	var parsed PingResponse
	err := json.Unmarshal(res.Body, &parsed)
	require.NoError(t, err, "Failed to parse json")
	require.Equal(t, "Pong", parsed.Result)
}

func requireGreaterThan(t *testing.T, left int, right int) {
	if left <= right {
		t.Error("Expecting %d > %d", left, right)
	}
}

func TestBadKeys(t *testing.T) {
	badClient := &Client{
		APIKey: "not valid",
		APIID:  "invalid",
	}
	regions := []Region{}
	require.NotNil(t, badClient.GetRegions(&regions))
}

func TestGetBlueprints(t *testing.T) {
	skipNoAPIKeys(t)

	var projects = []Project{}
	apierr := c.GetProjects(&projects)
	require.Nil(t, apierr, "failed to fetch projects")
	requireGreaterThan(t, len(projects), 0)
	proj1 := projects[0]

	var policies = []Policy{}
	apierr = c.GetPolicies(proj1.ID, &policies)
	require.Nil(t, apierr, "failed to fetch policies")

	var proj1Details = ProjectDetails{}
	apierr = c.GetProjectDetails(proj1.ID, &proj1Details)
	require.Nil(t, apierr, "failed to fetch project details ")

	var blueprints = []Blueprint{}
	apierr = c.GetBlueprints(proj1.ID, &blueprints)
	require.Nil(t, apierr, "failed to fetch blueprints")
	requireGreaterThan(t, len(blueprints), 0)

	var blue1 = BlueprintDetails{}
	apierr = c.GetBlueprintDetails(proj1.ID, blueprints[0].ID, &blue1)
	require.Nil(t, apierr, "failed to fetch blueprint details")
	require.NotEmpty(t, blue1.Name)
}

func TestGetProjectsByFilter(t *testing.T) {
	skipNoAPIKeys(t)

	var projects = []Project{}
	apierr := c.GetProjectsByFilter([]string{"WhereUserIsProjectManager"}, &projects)
	require.Nil(t, apierr, "failed to fetch projects")
}

func TestGetEnvs(t *testing.T) {
	skipNoAPIKeys(t)
	var envs = Environments{}
	apierr := c.GetEnvironments(true, "allvisible", &envs)
	require.Nil(t, apierr, "failed to fetch envs")

	var envID = envs[0].ID
	var env1 = Environment{}
	apierr = c.GetEnvironment(envID, "view", &env1)
	require.Nil(t, apierr, "failed to fetch env by ID")

}

func TestGetEnvDetails(t *testing.T) {
	skipNoAPIKeys(t)
	env, apierr := c.GetEnvironmentByName(testEnvName)
	require.Nil(t, apierr, "failed to fetch env by ID")
	require.NotNil(t, env, "failed to find test env. possibly this test suite hasn't been run with ALLOW_TEST_CREATE?")
	require.NotNil(t, env.ID)
	var envEx EnvironmentExtended = EnvironmentExtended{}
	apierr = c.GetEnvironmentExtended(env.ID, &envEx)
	require.Nil(t, apierr, "failed to fetch extended env info")
	require.Equal(t, 1, len(envEx.Vms))
}

func TestEnvResume(t *testing.T) {
	skipNoAPIKeys(t)
	env, apierr := c.GetEnvironmentByName(testEnvName)
	require.Nil(t, apierr, "failed to fetch env by name")
	require.NotNil(t, env)
	apierr = c.EnvironmentResume(env.ID)
	require.Nil(t, apierr, "failed to resume env")
}

func TestEnvExtend(t *testing.T) {
	skipNoAPIKeys(t)
	env, apierr := c.GetEnvironmentByName(testEnvName)
	require.NotNil(t, env)
	require.Nil(t, apierr, "failed to fetch env by name")
	apierr = c.EnvironmentExtend(env.ID)
	require.Nil(t, apierr, "failed to extend env")
}

func TestDeleteEnv(t *testing.T) {
	skipNoAPIKeys(t)
	if os.Getenv("TEST_DELETE_ENV") != "true" {
		t.Skip("Not running delete-env test unless TEST_DELETE_ENV is true")
	}
	env, apierr := c.GetEnvironmentByName(testEnvName)
	require.Nil(t, apierr, "failed to fetch env by name")
	require.NotNil(t, env)
	c.EnvironmentDelete(env.ID)
}

func TestFindTemplateByName(t *testing.T) {
	skipNoAPIKeys(t)
	templates := []VMTemplate{}
	require.Nil(t, c.GetTemplates(nil, &templates))
	for _, template := range templates {
		if template.Name == "Docker - Ubuntu 14.04 Server" {
			t.Logf("Found docker template: %s", template.ID)
			return
		}
	}
	t.Errorf("Docker template not found in list of templates")
}

func waitForEnvStatus(t *testing.T, envID string, code EnvironmentStatusCode) EnvironmentStatusCode {
	details := EnvironmentExtended{}
	for i := 0; i < 10; i++ {
		require.Nil(t, c.GetEnvironmentExtended(envID, &details))
		if details.StatusCode == code {
			return code
		}
		t.Logf("Status is still %d, waiting for %d", details.StatusCode, code)
		time.Sleep(time.Second)
	}

	return details.StatusCode
}

func TestWaitForEnvironment(t *testing.T) {
	env, apierr := c.GetEnvironmentByName(testEnvName)
	require.Nil(t, apierr, "failed to fetch envs")
	require.NotNil(t, env, "Test env not found")
	envID := env.ID
	// Suspend and wait for suspended status
	require.Nil(t, c.EnvironmentSuspend(envID))
	require.Equal(t, StatusSuspended, waitForEnvStatus(t, envID, StatusSuspended))
	require.Nil(t, c.EnvironmentResume(envID))
	require.Equal(t, StatusReady, waitForEnvStatus(t, envID, StatusReady))
}

func TestCreateEnv(t *testing.T) {
	skipNoAPIKeys(t)
	skipResourceCreation(t)
	env, apierr := c.GetEnvironmentByName(testEnvName)

	require.Nil(t, apierr, "failed to fetch envs")
	if env != nil {
		t.Skipf("Env with name %s already exists. Skipping this test.", testEnvName)
	}

	var regions = []Region{}
	apierr = c.GetRegions(&regions)
	require.Nil(t, apierr, "failed to fetch envs")

	region1 := regions[0].ID

	var templates = []VMTemplate{}
	var params = GetTemplateParams{TemplateType: "1", RegionID: region1}
	apierr = c.GetTemplates(&params, &templates)
	require.Nil(t, apierr, "failed to fetch templates")

	var projects = []Project{}
	apierr = c.GetProjects(&projects)
	require.Nil(t, apierr, "failed to fetch projects")
	requireGreaterThan(t, len(projects), 0)
	proj1 := projects[0]

	// Find Ubuntu template
	var ubuntuTemplateID string
	for _, t := range templates {
		if strings.Contains(t.Name, "Ubuntu 16.04") {
			ubuntuTemplateID = t.ID
			break
		}
	}

	var request = EnvironmentTemplateRequest{
		Environment: Environment{
			Name:        testEnvName,
			Description: "not super important",
			ProjectID:   proj1.ID,
			RegionID:    region1,
		},
		ItemsCart: []VM{{
			Type:         2,
			Name:         "vm1",
			TemplateVMID: ubuntuTemplateID,
			Description:  "my little vm",
		}},
	}

	var envCreateResponse CreateTemplateEnvResponse = CreateTemplateEnvResponse{}

	apierr = c.EnvironmentCreateFromTemplate(&request, &envCreateResponse)
	require.Nil(t, apierr, "failed to create env from template")

}

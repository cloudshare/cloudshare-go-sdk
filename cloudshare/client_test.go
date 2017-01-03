package cloudshare

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/url"
	"os"
	"strings"
	"testing"
)

func TestBuildURL(t *testing.T) {
	assert.Equal(t, "https://use.cloudshare.com/api/v3/projects",
		buildURL("projects", nil).String(), "failed to build url")

	assert.Equal(t, "https://use.cloudshare.com/api/v3/projects/",
		buildURL("/projects/", nil).String(), "failed to build url")

	params := &url.Values{}
	params.Set("key", "value_with/_in_it")
	assert.Equal(t, "https://use.cloudshare.com/api/v3/path?key=value_with%2F_in_it", buildURL("path", params).String(), "url param encoding failed")
}

type PingResponse struct {
	Result string `json:"result"`
}

const testEnvName = "go-sdk-test-env"

var apikey, apiid, allowTestCreate = os.Getenv("CLOUDSHARE_API_KEY"), os.Getenv("CLOUDSHARE_API_ID"), os.Getenv("ALLOW_TEST_CREATE")

var c = &Client{
	APIKey: apikey,
	APIID:  apiid,
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
	assert.Nil(t, apierr, "failed to ping")
	var parsed PingResponse
	err := json.Unmarshal(res.Body, &parsed)
	assert.NoError(t, err, "Failed to parse json")
	assert.Equal(t, "Pong", parsed.Result)
}

func assertGreaterThan(t *testing.T, left int, right int) {
	if left <= right {
		t.Error("Expecting %d > %d", left, right)
	}
}

func TestGetBlueprints(t *testing.T) {
	skipNoAPIKeys(t)

	var projects = []Project{}
	apierr := c.GetProjects(&projects)
	assert.Nil(t, apierr, "failed to fetch projects")
	assertGreaterThan(t, len(projects), 0)
	proj1 := projects[0]

	var policies = []Policy{}
	apierr = c.GetPolicies(proj1.ID, &policies)
	assert.Nil(t, apierr, "failed to fetch policies")

	var proj1Details = ProjectDetails{}
	apierr = c.GetProjectDetails(proj1.ID, &proj1Details)
	assert.Nil(t, apierr, "failed to fetch project details ")

	var blueprints = []Blueprint{}
	apierr = c.GetBlueprints(proj1.ID, &blueprints)
	assert.Nil(t, apierr, "failed to fetch blueprints")
	assertGreaterThan(t, len(blueprints), 0)

	var blue1 = BlueprintDetails{}
	apierr = c.GetBlueprintDetails(proj1.ID, blueprints[0].ID, &blue1)
	assert.Nil(t, apierr, "failed to fetch blueprint details")
	assert.NotEmpty(t, blue1.Name)
}

func TestGetProjectsByFilter(t *testing.T) {
	skipNoAPIKeys(t)

	var projects = []Project{}
	apierr := c.GetProjectsByFilter([]string{"WhereUserIsProjectManager"}, &projects)
	assert.Nil(t, apierr, "failed to fetch projects")
}

func TestGetEnvs(t *testing.T) {
	skipNoAPIKeys(t)
	var envs = Environments{}
	apierr := c.GetEnvironments(true, "allvisible", &envs)
	assert.Nil(t, apierr, "failed to fetch envs")

	var envID = envs[0].ID
	var env1 = Environment{}
	apierr = c.GetEnvironment(envID, "view", &env1)
	assert.Nil(t, apierr, "failed to fetch env by ID")

}

func TestGetEnvDetails(t *testing.T) {
	skipNoAPIKeys(t)
	env, apierr := c.GetEnvironmentByName(testEnvName)
	assert.Nil(t, apierr, "failed to fetch env by ID")
	assert.NotNil(t, env, "failed to find test env. possibly this test suite hasn't been run with ALLOW_TEST_CREATE?")
	assert.NotNil(t, env.ID)
	var envEx EnvironmentExtended = EnvironmentExtended{}
	apierr = c.GetEnvironmentExtended(env.ID, &envEx)
	assert.Nil(t, apierr, "failed to fetch extended env info")
	assert.Equal(t, 1, len(envEx.Vms))
	assert.Equal(t, "Ready", envEx.StatusText)
	assert.Equal(t, StatusReady, envEx.StatusCode)
}

func TestEnvResume(t *testing.T) {
	skipNoAPIKeys(t)
	env, apierr := c.GetEnvironmentByName(testEnvName)
	assert.Nil(t, apierr, "failed to fetch env by name")
	apierr = c.EnvironmentResume(env.ID)
	assert.Nil(t, apierr, "failed to resume env")
}

func TestEnvExtend(t *testing.T) {
	skipNoAPIKeys(t)
	env, apierr := c.GetEnvironmentByName(testEnvName)
	assert.Nil(t, apierr, "failed to fetch env by name")
	apierr = c.EnvironmentExtend(env.ID)
	assert.Nil(t, apierr, "failed to extend env")
}

func TestDeleteEnv(t *testing.T) {
	skipNoAPIKeys(t)
	if os.Getenv("TEST_DELETE_ENV") == "true" {
		env, apierr := c.GetEnvironmentByName(testEnvName)
		assert.Nil(t, apierr, "failed to fetch env by name")
		c.EnvironmentDelete(env.ID)
	}
}

func TestCreateEnv(t *testing.T) {
	skipNoAPIKeys(t)
	skipResourceCreation(t)

	var regions = []Region{}
	apierr := c.GetRegions(&regions)
	assert.Nil(t, apierr, "failed to fetch envs")

	region1 := regions[0].ID

	var templates = []VMTemplate{}
	var params = GetTemplateParams{templateType: "1", regionID: region1}
	apierr = c.GetTemplates(&params, &templates)
	assert.Nil(t, apierr, "failed to fetch templates")

	var projects = []Project{}
	apierr = c.GetProjects(&projects)
	assert.Nil(t, apierr, "failed to fetch projects")
	assertGreaterThan(t, len(projects), 0)
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
	assert.Nil(t, apierr, "failed to create env from template")

}

package cloudshare

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/url"
	"os"
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

var apikey, apiid = os.Getenv("CLOUDSHARE_API_KEY"), os.Getenv("CLOUDSHARE_API_ID")

var c = &Client{
	APIKey: apikey,
	APIID:  apiid,
}

func skipNoAPIKeys(t *testing.T) {
	if apikey == "" || apiid == "" {
		t.Skipf("test only runs with actual credentials")
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
	var envs = []Environment{}
	apierr := c.GetEnvironments(true, "allvisible", &envs)
	assert.Nil(t, apierr, "failed to fetch envs")

	var envID = envs[0].ID
	var env1 = Environment{}
	apierr = c.GetEnvironment(envID, "view", &env1)
	assert.Nil(t, apierr, "failed to fetch env by ID")

}

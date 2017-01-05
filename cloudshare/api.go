package cloudshare

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

func (c *Client) makeRequest(method string, path string, response interface{}, params *url.Values, jsonable interface{}) error {
	var body *string
	if jsonable != nil {
		buffer, err := json.Marshal(&jsonable)
		if err != nil {
			return APIError{
				InnerError: err,
				Message:    "Failed to serialize request object to JSON",
			}
		}
		bodyString := string(buffer)
		body = &bodyString
	}
	res, err := c.Request(method, path, params, body)
	if err != nil {
		return err
	}
	if response != nil {
		e := json.Unmarshal(res.Body, &response)
		//// NOCOMMIT
		//fmt.Println(path)
		// fmt.Println(string(res.Body))
		//fmt.Println("---------------------------------")
		////
		if e != nil {
			return APIError{InnerError: e}
		}
	}
	return nil
}

func (c *Client) makeGetRequest(path string, response interface{}, params *url.Values) error {
	return c.makeRequest("GET", path, response, params, nil)
}

func (c *Client) makePostRequest(path string, response interface{}, params *url.Values, jsonable interface{}) error {
	return c.makeRequest("POST", path, response, params, jsonable)
}

// GetBlueprintDetails returns details about a blueprint
func (c *Client) GetBlueprintDetails(projectID string, blueprintID string, ret *BlueprintDetails) error {
	path := fmt.Sprintf("projects/%s/blueprints/%s", projectID, blueprintID)
	return c.makeGetRequest(path, ret, nil)
}

/*
GetProjectsByFilter returns a list of projects for the user according to the filter strings
	"WhereUserIsProjectManager"    Returns only projects where user is a project manager
	"WhereUserIsProjectMember"     Returns only projects where user is a project member
	"WhereUserCanCreateClass"      Returns only projects where user can create a class in

	example:
		var projects = []Project{}
		client.GetProjectsByFilter(["WhereUserIsProjectManager", "WhereUserCanCreateClass"], &projects)
*/
func (c *Client) GetProjectsByFilter(filters []string, ret *[]Project) error {
	query := url.Values{}
	for _, filter := range filters {
		query.Add(filter, "true")
	}
	return c.makeGetRequest("projects", ret, &query)
}

// GetProjects returns a list of projects for the user
func (c *Client) GetProjects(ret *[]Project) error {
	return c.makeGetRequest("projects", ret, nil)
}

// GetProjectDetails returns project details by id
func (c *Client) GetProjectDetails(projectID string, ret *ProjectDetails) error {
	path := fmt.Sprintf("projects/%s", projectID)
	return c.makeGetRequest(path, ret, nil)
}

// GetBlueprints returns the blueprints available for a project
func (c *Client) GetBlueprints(projectID string, ret *[]Blueprint) error {
	path := fmt.Sprintf("projects/%s/blueprints", projectID)
	return c.makeGetRequest(path, ret, nil)
}

// GetPolicies returns a list of all policies by project id
func (c *Client) GetPolicies(projectID string, ret *[]Policy) error {
	path := fmt.Sprintf("projects/%s/policies", projectID)
	return c.makeGetRequest(path, ret, nil)
}

// GetEnvironments returns a list of environments, either in brief or full details
// Possible criteria: allowed | allvisible
func (c *Client) GetEnvironments(brief bool, criteria string, ret *Environments) error {
	query := url.Values{}
	query.Add("brief", strconv.FormatBool(brief))
	query.Add("criteria", criteria)
	return c.makeGetRequest("envs", ret, &query)
}

// GetEnvironment returns a specific environment by ID
// permission can be view|edit|owner
func (c *Client) GetEnvironment(id string, permission string, ret *Environment) error {
	path := fmt.Sprintf("envs/%s", id)
	query := url.Values{}
	query.Add("permission", permission)
	return c.makeGetRequest(path, ret, &query)
}

/* GetEnvironmentExtended returns extended information about an environment.
See http://docs.cloudshare.com/rest-api/v3/environments/envs/actions-getextended/ */
func (c *Client) GetEnvironmentExtended(id string, ret *EnvironmentExtended) error {
	query := url.Values{}
	query.Add("envId", id)
	return c.makeGetRequest("envs/actions/getextended", ret, &query)
}

// EnvironmentCreateFromTemplate creates a new environment based on a VM template
func (c *Client) EnvironmentCreateFromTemplate(request *EnvironmentTemplateRequest, response *CreateTemplateEnvResponse) error {
	return c.makePostRequest("envs", response, nil, request)
}

func (c *Client) envPutAction(action string, params *url.Values) error {
	return c.makeRequest("PUT", action, nil, params, nil)
}

func (c *Client) envPutActionByID(action string, id string) error {
	query := url.Values{}
	query.Add("envId", id)
	return c.envPutAction("envs/actions/"+action, &query)
}

func (c *Client) EnvironmentDelete(envID string) error {
	return c.makeRequest("DELETE", fmt.Sprintf("envs/%s", envID), nil, nil, nil)
}

// EnvironmentResume resumes a suspended environment
func (c *Client) EnvironmentResume(envID string) error {
	return c.envPutActionByID("resume", envID)
}

func (c *Client) RebootVM(vmID string) error {
	query := url.Values{}
	query.Add("vmId", vmID)
	return c.envPutAction("vms/actions/reboot", &query)
}

// EnvironmentSuspend suspends a running environment
func (c *Client) EnvironmentSuspend(envID string) error {
	return c.envPutActionByID("suspend", envID)
}

// EnvironmentExtend extends the lifetime of an environment
func (c *Client) EnvironmentExtend(envID string) error {
	return c.envPutActionByID("extend", envID)
}

/* GetTemplates returns a list of available templates that can be filtered by GetTemplateParams
 */
func (c *Client) GetTemplates(params *GetTemplateParams, ret *[]VMTemplate) error {
	query := url.Values{}
	if params != nil {
		if params.skip != 0 {
			query.Add("skip", fmt.Sprintf("%d", params.skip))
		}
		if params.take != 0 {
			query.Add("take", fmt.Sprintf("%d", params.take))
		}
		if params.regionID != "" {
			query.Add("regionId", params.regionID)
		}
		if params.projectID != "" {
			query.Add("projectId", params.projectID)
		}
		if params.templateType != "" {
			query.Add("templateType", params.templateType)
		}
	}
	return c.makeGetRequest("templates", ret, &query)
}

func (c *Client) GetRegions(ret *[]Region) error {
	return c.makeGetRequest("regions", ret, nil)
}

package cloudshare

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

func (c *Client) makeRequest(path string, response interface{}, params *url.Values) *APIError {
	res, err := c.Request("GET", path, params, nil)
	if err != nil {
		return err
	}
	e := json.Unmarshal(res.Body, &response)
	//// NOCOMMIT
	//fmt.Println(path)
	//fmt.Println(string(res.Body))
	//fmt.Println("---------------------------------")
	////
	if e != nil {
		return &APIError{Error: &e}
	}
	return nil
}

// GetBlueprintDetails returns details about a blueprint
func (c *Client) GetBlueprintDetails(projectID string, blueprintID string, ret *BlueprintDetails) *APIError {
	path := fmt.Sprintf("projects/%s/blueprints/%s", projectID, blueprintID)
	return c.makeRequest(path, ret, nil)
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
func (c *Client) GetProjectsByFilter(filters []string, ret *[]Project) *APIError {
	query := url.Values{}
	for _, filter := range filters {
		query.Add(filter, "true")
	}
	return c.makeRequest("projects", ret, &query)
}

// GetProjects returns a list of projects for the user
func (c *Client) GetProjects(ret *[]Project) *APIError {
	return c.makeRequest("projects", ret, nil)
}

// GetProjectDetails returns project details by id
func (c *Client) GetProjectDetails(projectID string, ret *ProjectDetails) *APIError {
	path := fmt.Sprintf("projects/%s", projectID)
	return c.makeRequest(path, ret, nil)
}

// GetBlueprints returns the blueprints available for a project
func (c *Client) GetBlueprints(projectID string, ret *[]Blueprint) *APIError {
	path := fmt.Sprintf("projects/%s/blueprints", projectID)
	return c.makeRequest(path, ret, nil)
}

// GetPolicies returns a list of all policies by project id
func (c *Client) GetPolicies(projectID string, ret *[]Policy) *APIError {
	path := fmt.Sprintf("projects/%s/policies", projectID)
	return c.makeRequest(path, ret, nil)
}

// GetEnvironments returns a list of environments, either in brief or full details
// Possible criteria: allowed | allvisible
func (c *Client) GetEnvironments(brief bool, criteria string, ret *[]Environment) *APIError {
	query := url.Values{}
	query.Add("brief", strconv.FormatBool(brief))
	query.Add("criteria", criteria)
	return c.makeRequest("envs", ret, &query)
}

// GetEnvironment returns a specific environment by ID
// permission can be view|edit|owner
func (c *Client) GetEnvironment(id string, permission string, ret *Environment) *APIError {
	path := fmt.Sprintf("envs/%s", id)
	query := url.Values{}
	query.Add("permission", permission)
	return c.makeRequest(path, ret, &query)
}

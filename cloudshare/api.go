package cloudshare

import (
	"encoding/json"
	"fmt"
)

func (c *Client) makeRequest(path string, response interface{}) *APIError {
	res, err := c.Request("GET", path, nil, nil)
	if err != nil {
		return err
	}
	e := json.Unmarshal(res.Body, &response)
	if e != nil {
		return &APIError{Error: &e}
	}
	return nil
}

// GetBlueprintDetails returns details about a blueprint
func (c *Client) GetBlueprintDetails(projectID string, blueprintID string, ret *BlueprintDetails) *APIError {
	path := fmt.Sprintf("projects/%s/blueprints/%s", projectID, blueprintID)
	return c.makeRequest(path, ret)
}

// GetProjects returns a list of projects for the user
func (c *Client) GetProjects(ret *[]Project) *APIError {
	return c.makeRequest("projects", ret)
}

// GetBlueprints returns the blueprints available for a project
func (c *Client) GetBlueprints(projectID string, ret *[]Blueprint) *APIError {
	path := fmt.Sprintf("projects/%s/blueprints", projectID)
	return c.makeRequest(path, &ret)
}

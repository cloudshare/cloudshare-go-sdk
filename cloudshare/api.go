package cloudshare

import (
	"encoding/json"
	"fmt"
)

// GetBlueprints returns the blueprints available for a project
func (c *Client) GetBlueprints(projectID string) (*[]Blueprint, *APIError) {
	path := fmt.Sprintf("projects/%s/blueprints", projectID)
	res, err := c.Request("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}
	var ret []Blueprint
	e := json.Unmarshal(res.Body, &ret)
	if e != nil {
		return nil, &APIError{Error: &e}
	}
	return &ret, nil
}

// GetBlueprintDetails returns details about a blueprint
func (c *Client) GetBlueprintDetails(projectID string, blueprintID string) (*BlueprintDetails, *APIError) {
	path := fmt.Sprintf("projects/%s/blueprints/%s", projectID, blueprintID)
	res, err := c.Request("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}
	var ret BlueprintDetails
	e := json.Unmarshal(res.Body, &ret)
	if e != nil {
		return nil, &APIError{Error: &e}
	}
	return &ret, nil
}

// GetProjects returns a list of projects for the user
func (c *Client) GetProjects() (*[]Project, *APIError) {
	res, err := c.Request("GET", "projects", nil, nil)
	if err != nil {
		return nil, err
	}
	var ret []Project
	e := json.Unmarshal(res.Body, &ret)
	if e != nil {
		return nil, &APIError{Error: &e}
	}
	return &ret, nil
}

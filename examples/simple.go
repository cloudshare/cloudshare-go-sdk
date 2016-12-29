package main

import "fmt"
import "encoding/json"
import "github.com/cloudshare/go-sdk/cloudshare"

func main() {
	fmt.Println("vim-go")
	c := cloudshare.Client{
		APIKey: "your API key here",
		APIID:  "your API id here",
	}
	response, apierr := c.Request("GET", "projects", nil, nil)
	if apierr != nil {
		panic(apierr.Error)
	}
	type Projects []struct {
		Name     string `json:"name"`
		IsActive bool   `json:"isActive"`
		ID       string `json:"id"`
	}
	var projects Projects
	json.Unmarshal(response.Body, &projects)
	fmt.Printf("Project 1: name: %s, id: %s\n", projects[0].Name, projects[0].ID)
}

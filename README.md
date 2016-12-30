# CloudShare Go SDK

## Install

`go get github.com/cloudshare/go-sdk/cloudshare`

Fetch your API key and ID from the [user details page](https://use.cloudshare.com/Ent/Vendor/UserDetails.aspx).

## Example - getting a list of projects

```
package main

import "fmt"
import "github.com/cloudshare/go-sdk/cloudshare"

func main() {

	c := cloudshare.Client{
		APIKey: "your API key here",
		APIID:  "your API id here",
	}

	// Get the list of projects for the user account
	var projects = []Project{}
	apierr := c.GetProjects(&projects)
	if apierr != nil {
		panic(apierr.Error)
	}
	fmt.Printf("Project 1: name: %s, id: %s\n", projects[0].Name, projects[0].ID)
}
```

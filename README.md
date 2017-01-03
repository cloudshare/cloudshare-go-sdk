# CloudShare Go SDK

## Install

`go get github.com/cloudshare/go-sdk/cloudshare`

Fetch your API key and ID from the [user details page](https://use.cloudshare.com/Ent/Vendor/UserDetails.aspx).


## Example - generic REST API calls

Use the `Client` struct to execute any REST API call as defined [in the docs](http://docs.cloudshare.com/rest-api/v3/environments/envs/)

```
package main

import "fmt"
import "github.com/cloudshare/go-sdk/cloudshare"
import "net/url"

func main() {

	c := cloudshare.Client{
		APIKey: "your API key here",
		APIID:  "your API id here",
	}

	// Get the list of projects for the user account
	apiresponse, apierror := c.Request("GET", "envs", nil, nil)
    
    // Suspend a running environment
    queryParams = &url.Values{}
    queryParams.Add("envId", "my-env-id-here")
	apiresponse, apierror = c.Request("PUT", "envs/actions/suspend", nil, nil)
}
    
```

## Example - typed API functions

We provide friendly, typed wrappers for the most common API operations.

Have a look at the go docs for the package (`godoc -http=:6060` in the repository directory)
to see the types and wrapper functions.

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

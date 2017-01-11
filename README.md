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

# cscurl

The Go SDK ships with a command line utility called `cscurl` that lets you invoke REST API calls, somewhat like `curl`.

## Installing

Pre-built binaries are available in the [releases page](https://github.com/cloudshare/go-sdk/releases).

- Download and place it somewhere in your `PATH`.
- If you don't want to pass the API Key & ID for every call, define them as environment variables:
    - CLOUDSHARE_API_KEY
    - CLOUDSHARE_API_ID

## Examples

### GET request - getting the list of regions

```
$ cscurl https://use.cloudshare.com/api/v3/regions| jq
[
  {
    "id": "REKolD1-ab84YIxODeMGob9A2",
    "name": "Miami",
    "friendlyName": "US East (Miami)",
    "cloudName": "CloudShare"
  },
  {
    "id": "RE0YOUV7_lTmgb0X8D1UjM3g2",
    "name": "VMware_Singapore",
    "friendlyName": "Asia Pacific (Singapore)",
    "cloudName": "CloudShare"
  },
  {
    "id": "RE6OEZs-y-mkK1mEMGwIgZiw2",
    "name": "VMware_Amsterdam",
    "friendlyName": "EU (Amsterdam)",
    "cloudName": "CloudShare"
  }
]
```

### PUT request with JSON body - setting the number of CPUs of a VM

```
$ cscurl -m put https://use.cloudshare.com/api/v3/vms/actions/editvmhardware -d \
   '{"vmId": "[my vm id...]", "numCpus": 2}' | jq
{
    "conflictsFound": false,
    "conflicts": ""
}
```





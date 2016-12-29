package cloudshare

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
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

func TestPing(t *testing.T) {
	apikey, apiid := os.Getenv("CLOUDSHARE_API_KEY"), os.Getenv("CLOUDSHARE_API_ID")
	if apikey == "" || apiid == "" {
		t.Skipf("test only runs with actual credentials")
	}
	c := &Client{
		APIKey: apikey,
		APIID:  apiid,
	}

	res, apierr := c.Request("GET", "ping", nil, nil)
	assert.Nil(t, apierr, "failed to ping")
	body, err := ioutil.ReadAll(res.Body)
	var parsed PingResponse
	err = json.Unmarshal(body, &parsed)
	assert.NoError(t, err, "Failed to parse json")
	assert.Equal(t, "Pong", parsed.Result)
}

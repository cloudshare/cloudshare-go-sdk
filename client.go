package cloudshare

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Client holds the API credentials can be found in your User Details page.
type Client struct {
	APIKey string
	APIID  string
}

func buildURL(path string, params *url.Values) *url.URL {
	u := &url.URL{
		Scheme: "https",
		Host:   "use.cloudshare.com",
		Path:   "/api/v3/" + strings.TrimLeft(path, "/"),
	}

	if params != nil {
		u.RawQuery = url.Values.Encode(*params)
	}

	return u
}

// ErrorResponse is returned by CheckAPI
type ErrorResponse struct {
	Code    string
	Message string
}

// CheckAPI tests the HTTP response returned by Client.Request and return nil
// if the call was successful.
func CheckAPI(response *http.Response) *ErrorResponse {
	if response.StatusCode/100 != 2 {
		buffer, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return &ErrorResponse{
				Code:    "unknown error",
				Message: "failed to parse http response body",
			}
		}
		var ret ErrorResponse
		json.Unmarshal(buffer, &ret)
		return &ret
	}
	return nil
}

// Request invokes an API call
// Example:
// client.Request("get", "projects", nil)
//
func (c *Client) Request(method string, path string, queryParams *url.Values, content *string) (*http.Response, error) {
	client := http.Client{}
	url := buildURL(path, queryParams)

	headers := &http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Accept", "application/json")

	token := authToken(c.APIKey, c.APIID, url.String())
	headers.Set("Authorization", "cs_sha1 "+token)
	request := &http.Request{
		Method: method,
		URL:    url,
		Header: *headers,
	}

	return client.Do(request)
}

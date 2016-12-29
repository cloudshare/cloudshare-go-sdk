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

// APIError is returned by client.Request in case of a failure.
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Error   *error
}

func (e *APIError) String() string {
	return e.Message
}

// Request invokes an API call
//
// Example:
//
// 		client.Request("get", "projects", nil)
func (c *Client) Request(method string, path string, queryParams *url.Values, content *string) (*http.Response, *APIError) {
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

	response, err := client.Do(request)
	if err != nil {
		return nil, &APIError{Error: &err}
	}

	if response.StatusCode/100 != 2 {
		buffer, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return response, &APIError{
				Code:    "unknown error",
				Message: "failed to parse http response body",
				Error:   &err,
			}
		}
		var ret APIError
		json.Unmarshal(buffer, &ret)
		return response, &ret
	}
	return response, nil
}
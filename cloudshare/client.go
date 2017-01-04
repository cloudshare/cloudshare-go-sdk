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
	Code       string `json:"code"`
	Message    string `json:"message"`
	InnerError *error
}

// APIResponse is returned by client.Request in case of success.
//
// The response body is a JSON buffer
type APIResponse struct {
	StatusCode int
	Body       []byte
}

func (e *APIError) Error() string {
	s := e.Message
	if e.InnerError != nil {
		s += "\n" + (*e.InnerError).Error()
	}
	return s
}

/*

Request invokes any API call

Example:

		client.Request("get", "projects", nil)

		method: the HTTP method to use. e.g. "GET", "PUT"
		path: the path relative to the API version.
		for example for this method: http://docs.cloudshare.com/rest-api/v3/environments/envs/actions-getextended/
			the path should be "envs/actions/getextended" (i.e. what goes after the v3/ prefix and before the query params (?)
		queryParams: url query params
		content: optional JSON body
*/
func (c *Client) Request(method string, path string, queryParams *url.Values, content *string) (*APIResponse, *APIError) {
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

	if content != nil {
		bodyReader := strings.NewReader(*content)
		request.Body = ioutil.NopCloser(bodyReader)

		// TODO: Test this with Unicode
		request.ContentLength = int64(len(*content))
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, &APIError{InnerError: &err}
	}

	body, err := ioutil.ReadAll(response.Body)
	// fmt.Println(path, string(body)) // NOCOMMIT
	if response.StatusCode/100 != 2 {
		if err != nil {
			return nil, &APIError{
				Code:       "unknown error",
				Message:    "failed to parse http response body",
				InnerError: &err,
			}
		}
		var ret APIError
		json.Unmarshal(body, &ret)
		return &APIResponse{StatusCode: response.StatusCode, Body: body}, &ret
	}

	if err != nil {
		return nil, &APIError{
			Message:    "Unable to read HTTP response body",
			InnerError: &err,
		}
	}

	return &APIResponse{
		Body:       body,
		StatusCode: response.StatusCode,
	}, nil
}

package cloudshare

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Client holds the API credentials can be found in your User Details page.
// APIKey & APIID are mandatory, and you can get your keys on the user details page.
// Tags is optional, and defaults to "go_sdk". It's for internal analytics, so feel free to ignore it.
type Client struct {
	APIKey  string
	APIID   string
	Tags    string
	APIHost string
}

func (c *Client) buildURL(path string, params *url.Values) *url.URL {

	host := c.APIHost
	if host == "" {
		host = "use.cloudshare.com"

	}

	u := &url.URL{
		Scheme: "https",
		Host:   host,
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
	InnerError error
}

// APIResponse is returned by client.Request in case of success.
//
// The response body is a JSON buffer
type APIResponse struct {
	StatusCode int
	Body       []byte
	Headers    http.Header
}

func (e APIError) Error() string {
	s := e.Message
	if e.InnerError != nil {
		s += "\n" + e.InnerError.Error()
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
func (c *Client) Request(method string, path string, queryParams *url.Values, content *string) (*APIResponse, error) {
	client := http.Client{}
	if os.Getenv("DEBUG") == "true" {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	if c.Tags == "" {
		c.Tags = "go_sdk"
	}

	if queryParams == nil {
		queryParams = &url.Values{}
	}
	// queryParams.Set("apiTags", c.Tags)

	url := c.buildURL(path, queryParams)

	// fmt.Printf("url: %s\n", url)

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
		// fmt.Printf("Request body: %s\n", *content)
	}

	response, err := client.Do(request)

	if err != nil {
		// fmt.Println(err)
		return nil, APIError{InnerError: err}
	}

	body, err := ioutil.ReadAll(response.Body)
	if response.StatusCode/100 != 2 {
		if err != nil {
			return nil, APIError{
				Code:       "unknown error",
				Message:    "failed to parse http response body",
				InnerError: err,
			}
		}
		var ret APIError
		json.Unmarshal(body, &ret)
		return &APIResponse{StatusCode: response.StatusCode, Body: body, Headers: response.Header}, &ret
	}

	if err != nil {
		return nil, APIError{
			Message:    "Unable to read HTTP response body",
			InnerError: err,
		}
	}

	return &APIResponse{
		Body:       body,
		StatusCode: response.StatusCode,
		Headers:    response.Header,
	}, nil
}

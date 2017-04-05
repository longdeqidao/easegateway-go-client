package v1

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/go-resty/resty"
)

type APIClient struct {
	config *Configuration
}

func (c *APIClient) SelectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

func (c *APIClient) SelectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}
	if contains(accepts, "application/json") {
		return "application/json"
	}
	return strings.Join(accepts, ",")
}

func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

func (c *APIClient) CallAPI(path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values) (*resty.Response, error) {

	rClient := c.prepareClient()
	request := c.prepareRequest(rClient, postBody, headerParams, queryParams)

	switch strings.ToUpper(method) {
	case "GET":
		response, err := request.Get(path)
		return response, err
	case "POST":
		response, err := request.Post(path)
		return response, err
	case "PUT":
		response, err := request.Put(path)
		return response, err
	case "PATCH":
		response, err := request.Patch(path)
		return response, err
	case "DELETE":
		response, err := request.Delete(path)
		return response, err
	}

	return nil, fmt.Errorf("invalid method %v", method)
}

func (c *APIClient) prepareClient() *resty.Client {
	rClient := resty.New()

	rClient.SetDebug(c.config.Debug)
	if c.config.Transport != nil {
		rClient.SetTransport(c.config.Transport)
	}

	if c.config.Timeout != nil {
		rClient.SetTimeout(*c.config.Timeout)
	}
	rClient.SetLogger(ioutil.Discard)
	return rClient
}

func (c *APIClient) prepareRequest(
	rClient *resty.Client,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values) *resty.Request {

	request := rClient.R()
	request.SetBody(postBody)

	if c.config.UserAgent != "" {
		request.SetHeader("User-Agent", c.config.UserAgent)
	}

	// add header parameter, if any
	if len(headerParams) > 0 {
		request.SetHeaders(headerParams)
	}

	// add query parameter, if any
	if len(queryParams) > 0 {
		request.SetMultiValueQueryParams(queryParams)
	}

	return request
}

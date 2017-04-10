package v1

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/hexdecteam/easegateway-go-client/rest/1.0/admin/v1/pdu"
	"github.com/hexdecteam/easegateway-go-client/rest/1.0/common/v1"
)

type AdminApi struct {
	Configuration *v1.Configuration
}

func NewAdminApi() *AdminApi {
	configuration := v1.NewConfiguration("http://localhost:9090/admin/v1")
	return &AdminApi{
		Configuration: configuration,
	}
}

func NewAdminApiWithBasePath(basePath string) *AdminApi {
	configuration := v1.NewConfiguration(basePath)
	return &AdminApi{
		Configuration: configuration,
	}
}

func (a AdminApi) GetPipelineTypes() (*pdu.PipelineTypesRetrieveResponse, *v1.APIResponse, error) {
	method := strings.ToUpper("Get")
	path := a.Configuration.BasePath + "/pipeline-types"
	headers := make(map[string]string)
	queryParams := url.Values{}

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headers[key] = a.Configuration.DefaultHeader[key]
	}

	// set Content-Type header
	contentTypes := a.Configuration.APIClient.SelectHeaderContentType([]string{"application/json"})
	if contentTypes != "" {
		headers["Content-Type"] = contentTypes
	}

	// set Accept header
	accept := a.Configuration.APIClient.SelectHeaderAccept([]string{"application/json"})
	if accept != "" {
		headers["Accept"] = accept
	}

	pdu := new(pdu.PipelineTypesRetrieveResponse)
	response, err := a.Configuration.APIClient.CallAPI(path, method, nil, headers, queryParams)

	u, _ := url.Parse(path)
	u.RawQuery = queryParams.Encode()

	ret := &v1.APIResponse{Operation: "GetPipelineTypes", Method: method, RequestURL: u.String()}
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return pdu, ret, err
	}

	err = json.Unmarshal(response.Body(), &pdu)
	return pdu, ret, err
}

func (a AdminApi) GetPipelines(pipelinesRetrieveRequest pdu.PipelinesRetrieveRequest) (
	*pdu.PipelinesRetrieveResponse, *v1.APIResponse, error) {

	method := strings.ToUpper("Get")
	path := a.Configuration.BasePath + "/pipelines"
	headers := make(map[string]string)
	queryParams := url.Values{}

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headers[key] = a.Configuration.DefaultHeader[key]
	}

	// set Content-Type header
	contentType := a.Configuration.APIClient.SelectHeaderContentType([]string{"application/json"})
	if contentType != "" {
		headers["Content-Type"] = contentType
	}

	// set Accept header
	accept := a.Configuration.APIClient.SelectHeaderAccept([]string{"application/json"})
	if accept != "" {
		headers["Accept"] = accept
	}

	pdu := new(pdu.PipelinesRetrieveResponse)
	response, err := a.Configuration.APIClient.CallAPI(
		path, method, &pipelinesRetrieveRequest, headers, queryParams)

	u, _ := url.Parse(path)
	u.RawQuery = queryParams.Encode()
	ret := &v1.APIResponse{Operation: "GetPipelines", Method: method, RequestURL: u.String()}
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return pdu, ret, err
	}
	err = json.Unmarshal(response.Body(), &pdu)
	return pdu, ret, err
}

func (a AdminApi) DeletePipelineByName(pipelineName string) (*v1.APIResponse, error) {
	method := strings.ToUpper("Delete")
	path := fmt.Sprintf("%s/pipelines/%s", a.Configuration.BasePath, pipelineName)
	headers := make(map[string]string)
	queryParams := url.Values{}

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headers[key] = a.Configuration.DefaultHeader[key]
	}

	// set Content-Type header
	contentType := a.Configuration.APIClient.SelectHeaderContentType([]string{"application/json"})
	if contentType != "" {
		headers["Content-Type"] = contentType
	}

	// set Accept header
	accept := a.Configuration.APIClient.SelectHeaderAccept([]string{"application/json"})
	if accept != "" {
		headers["Accept"] = accept
	}

	response, err := a.Configuration.APIClient.CallAPI(path, method, nil, headers, queryParams)

	u, _ := url.Parse(path)
	u.RawQuery = queryParams.Encode()
	ret := &v1.APIResponse{Operation: "DeletePipelineByName", Method: method, RequestURL: u.String()}
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return ret, err
	}
	return ret, err
}

func (a AdminApi) GetPipelineByName(pipelineName string) (*pdu.PipelineSpec, *v1.APIResponse, error) {
	method := strings.ToUpper("Get")
	path := fmt.Sprintf("%s/pipelines/%s", a.Configuration.BasePath, pipelineName)
	headers := make(map[string]string)
	queryParams := url.Values{}

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headers[key] = a.Configuration.DefaultHeader[key]
	}

	// set Content-Type header
	contentType := a.Configuration.APIClient.SelectHeaderContentType([]string{"application/json"})
	if contentType != "" {
		headers["Content-Type"] = contentType
	}

	// set Accept header
	accept := a.Configuration.APIClient.SelectHeaderAccept([]string{"application/json"})
	if accept != "" {
		headers["Accept"] = accept
	}

	pdu := new(pdu.PipelineSpec)
	response, err := a.Configuration.APIClient.CallAPI(path, method, nil, headers, queryParams)

	u, _ := url.Parse(path)
	u.RawQuery = queryParams.Encode()
	ret := &v1.APIResponse{Operation: "GetPipelineByName", Method: method, RequestURL: u.String()}
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return pdu, ret, err
	}
	err = json.Unmarshal(response.Body(), &pdu)
	return pdu, ret, err
}

func (a AdminApi) CreatePipeline(pipelineCreationRequest pdu.PipelineCreationRequest) (*v1.APIResponse, error) {
	method := strings.ToUpper("Post")
	path := a.Configuration.BasePath + "/pipelines"
	headers := make(map[string]string)
	queryParams := url.Values{}

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headers[key] = a.Configuration.DefaultHeader[key]
	}

	// set Content-Type header
	contentType := a.Configuration.APIClient.SelectHeaderContentType([]string{"application/json"})
	if contentType != "" {
		headers["Content-Type"] = contentType
	}

	// set Accept header
	accept := a.Configuration.APIClient.SelectHeaderAccept([]string{"application/json"})
	if accept != "" {
		headers["Accept"] = accept
	}

	response, err := a.Configuration.APIClient.CallAPI(
		path, method, &pipelineCreationRequest, headers, queryParams)

	u, _ := url.Parse(path)
	u.RawQuery = queryParams.Encode()
	ret := &v1.APIResponse{Operation: "CreatePipeline", Method: method, RequestURL: u.String()}
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return ret, err
	}
	return ret, err
}

func (a AdminApi) UpdatePipeline(pipelineUpdateRequest pdu.PipelineUpdateRequest) (*v1.APIResponse, error) {
	method := strings.ToUpper("Put")
	path := a.Configuration.BasePath + "/pipelines"
	headers := make(map[string]string)
	queryParams := url.Values{}

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headers[key] = a.Configuration.DefaultHeader[key]
	}

	// set Content-Type header
	contentType := a.Configuration.APIClient.SelectHeaderContentType([]string{"application/json"})
	if contentType != "" {
		headers["Content-Type"] = contentType
	}

	// set Accept header
	accept := a.Configuration.APIClient.SelectHeaderAccept([]string{"application/json"})
	if accept != "" {
		headers["Accept"] = accept
	}

	response, err := a.Configuration.APIClient.CallAPI(
		path, method, &pipelineUpdateRequest, headers, queryParams)

	u, _ := url.Parse(path)
	u.RawQuery = queryParams.Encode()
	var ret = &v1.APIResponse{Operation: "UpdatePipeline", Method: method, RequestURL: u.String()}
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return ret, err
	}
	return ret, err
}

func (a AdminApi) GetPluginTypes() (*pdu.PluginTypesRetrieveResponse, *v1.APIResponse, error) {
	method := strings.ToUpper("Get")
	path := a.Configuration.BasePath + "/plugin-types"
	headers := make(map[string]string)
	queryParams := url.Values{}

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headers[key] = a.Configuration.DefaultHeader[key]
	}

	// set Content-Type header
	contentType := a.Configuration.APIClient.SelectHeaderContentType([]string{"application/json"})
	if contentType != "" {
		headers["Content-Type"] = contentType
	}

	// set Accept header
	accept := a.Configuration.APIClient.SelectHeaderAccept([]string{"application/json"})
	if accept != "" {
		headers["Accept"] = accept
	}

	pdu := new(pdu.PluginTypesRetrieveResponse)
	response, err := a.Configuration.APIClient.CallAPI(path, method, nil, headers, queryParams)

	u, _ := url.Parse(path)
	u.RawQuery = queryParams.Encode()
	ret := &v1.APIResponse{Operation: "GetPluginTypes", Method: method, RequestURL: u.String()}
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return pdu, ret, err
	}
	err = json.Unmarshal(response.Body(), &pdu)
	return pdu, ret, err
}

func (a AdminApi) GetPlugins(pluginsRetrieveRequest pdu.PluginsRetrieveRequest) (
	*pdu.PluginsRetrieveResponse, *v1.APIResponse, error) {

	var method = strings.ToUpper("Get")
	// create path and map variables
	path := a.Configuration.BasePath + "/plugins"
	headers := make(map[string]string)
	queryParams := url.Values{}

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headers[key] = a.Configuration.DefaultHeader[key]
	}

	// set Content-Type header
	contentType := a.Configuration.APIClient.SelectHeaderContentType([]string{"application/json"})
	if contentType != "" {
		headers["Content-Type"] = contentType
	}

	// set Accept header
	accept := a.Configuration.APIClient.SelectHeaderAccept([]string{"application/json"})
	if accept != "" {
		headers["Accept"] = accept
	}

	pdu := new(pdu.PluginsRetrieveResponse)
	response, err := a.Configuration.APIClient.CallAPI(path, method, &pluginsRetrieveRequest, headers, queryParams)

	u, _ := url.Parse(path)
	u.RawQuery = queryParams.Encode()
	ret := &v1.APIResponse{Operation: "GetPlugins", Method: method, RequestURL: u.String()}
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return pdu, ret, err
	}
	err = json.Unmarshal(response.Body(), &pdu)
	return pdu, ret, err
}

func (a AdminApi) DeletePluginByName(pluginName string) (*v1.APIResponse, error) {
	method := strings.ToUpper("Delete")
	// create path and map variables
	path := fmt.Sprintf("%s/plugins/%s", a.Configuration.BasePath, pluginName)
	headers := make(map[string]string)
	queryParams := url.Values{}

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headers[key] = a.Configuration.DefaultHeader[key]
	}

	// set Content-Type header
	contentType := a.Configuration.APIClient.SelectHeaderContentType([]string{"application/json"})
	if contentType != "" {
		headers["Content-Type"] = contentType
	}

	// set Accept header
	accept := a.Configuration.APIClient.SelectHeaderAccept([]string{"application/json"})
	if accept != "" {
		headers["Accept"] = accept
	}

	response, err := a.Configuration.APIClient.CallAPI(path, method, nil, headers, queryParams)

	u, _ := url.Parse(path)
	u.RawQuery = queryParams.Encode()
	ret := &v1.APIResponse{Operation: "DeletePluginByName", Method: method, RequestURL: u.String()}
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return ret, err
	}
	return ret, err
}

func (a AdminApi) GetPluginByName(pluginName string) (*pdu.PluginSpec, *v1.APIResponse, error) {
	method := strings.ToUpper("Get")
	path := fmt.Sprintf("%s/plugins/%s", a.Configuration.BasePath, pluginName)
	headers := make(map[string]string)
	queryParams := url.Values{}

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headers[key] = a.Configuration.DefaultHeader[key]
	}

	// set Content-Type header
	contentType := a.Configuration.APIClient.SelectHeaderContentType([]string{"application/json"})
	if contentType != "" {
		headers["Content-Type"] = contentType
	}

	// set Accept header
	accept := a.Configuration.APIClient.SelectHeaderAccept([]string{"application/json"})
	if accept != "" {
		headers["Accept"] = accept
	}

	pdu := new(pdu.PluginSpec)
	response, err := a.Configuration.APIClient.CallAPI(path, method, nil, headers, queryParams)

	u, _ := url.Parse(path)
	u.RawQuery = queryParams.Encode()
	ret := &v1.APIResponse{Operation: "GetPluginByName", Method: method, RequestURL: u.String()}
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return pdu, ret, err
	}
	err = json.Unmarshal(response.Body(), &pdu)
	return pdu, ret, err
}

func (a AdminApi) CreatePlugin(pluginCreationRequest pdu.PluginCreationRequest) (*v1.APIResponse, error) {
	method := strings.ToUpper("Post")
	path := a.Configuration.BasePath + "/plugins"
	headers := make(map[string]string)
	queryParams := url.Values{}

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headers[key] = a.Configuration.DefaultHeader[key]
	}

	// set Content-Type header
	contentType := a.Configuration.APIClient.SelectHeaderContentType([]string{"application/json"})
	if contentType != "" {
		headers["Content-Type"] = contentType
	}

	// set Accept header
	accept := a.Configuration.APIClient.SelectHeaderAccept([]string{"application/json"})
	if accept != "" {
		headers["Accept"] = accept
	}

	response, err := a.Configuration.APIClient.CallAPI(path, method, &pluginCreationRequest, headers, queryParams)

	u, _ := url.Parse(path)
	u.RawQuery = queryParams.Encode()
	ret := &v1.APIResponse{Operation: "CreatePlugin", Method: method, RequestURL: u.String()}
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return ret, err
	}
	return ret, err
}

func (a AdminApi) UpdatePlugin(pluginUpdateRequest pdu.PluginUpdateRequest) (*v1.APIResponse, error) {
	method := strings.ToUpper("Put")
	path := a.Configuration.BasePath + "/plugins"
	headers := make(map[string]string)
	queryParams := url.Values{}

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headers[key] = a.Configuration.DefaultHeader[key]
	}

	// set Content-Type header
	contentType := a.Configuration.APIClient.SelectHeaderContentType([]string{"application/json"})
	if contentType != "" {
		headers["Content-Type"] = contentType
	}

	// set Accept header
	accept := a.Configuration.APIClient.SelectHeaderAccept([]string{"application/json"})
	if accept != "" {
		headers["Accept"] = accept
	}

	response, err := a.Configuration.APIClient.CallAPI(path, method, &pluginUpdateRequest, headers, queryParams)

	u, _ := url.Parse(path)
	u.RawQuery = queryParams.Encode()
	ret := &v1.APIResponse{Operation: "UpdatePlugin", Method: method, RequestURL: u.String()}
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return ret, err
	}
	return ret, err
}

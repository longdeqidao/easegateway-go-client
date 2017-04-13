package v1

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/hexdecteam/easegateway-go-client/rest/1.0/admin/v1/pdu"
	"github.com/hexdecteam/easegateway-go-client/rest/1.0/common/v1"
	common_pdu "github.com/hexdecteam/easegateway-go-client/rest/1.0/common/v1/pdu"
)

type AdminApi struct {
	Configuration *v1.Configuration
}

func NewAdminApi(address string) *AdminApi {
	address = strings.TrimSpace(address)
	if len(address) == 0 {
		address = "localhost:9090"
	}

	configuration := v1.NewConfiguration(fmt.Sprintf("http://%s/admin/v1", address))
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

	ret := v1.NewAPIResponse("GetPipelineTypes", method, path, queryParams)
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return pdu, ret, err
	}

	if response.StatusCode() == 200 {
		err = json.Unmarshal(response.Body(), pdu)
	} else if len(response.Body()) > 0 {
		e := new(common_pdu.Error)
		err = json.Unmarshal(response.Body(), e)
		if err == nil {
			ret.Error = e
		}
	}

	return pdu, ret, err
}

func (a AdminApi) GetPipelines(pipelinesRetrieveRequest *pdu.PipelinesRetrieveRequest) (
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
	response, err := a.Configuration.APIClient.CallAPI(path, method, pipelinesRetrieveRequest, headers, queryParams)

	ret := v1.NewAPIResponse("GetPipelines", method, path, queryParams)
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return pdu, ret, err
	}

	if response.StatusCode() == 200 {
		err = json.Unmarshal(response.Body(), pdu)
	} else if len(response.Body()) > 0 {
		e := new(common_pdu.Error)
		err = json.Unmarshal(response.Body(), e)
		if err == nil {
			ret.Error = e
		}
	}

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

	ret := v1.NewAPIResponse("DeletePipelineByName", method, path, queryParams)
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return ret, err
	}

	if response.StatusCode() != 200 {
		e := new(common_pdu.Error)
		err = json.Unmarshal(response.Body(), e)
		if err == nil {
			ret.Error = e
		}
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

	ret := v1.NewAPIResponse("GetPipelineByName", method, path, queryParams)
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return pdu, ret, err
	}

	if response.StatusCode() == 200 {
		err = json.Unmarshal(response.Body(), pdu)
	} else if len(response.Body()) > 0 {
		e := new(common_pdu.Error)
		err = json.Unmarshal(response.Body(), e)
		if err == nil {
			ret.Error = e
		}
	}

	return pdu, ret, err
}

func (a AdminApi) CreatePipeline(pipelineCreationRequest *pdu.PipelineCreationRequest) (*v1.APIResponse, error) {
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

	response, err := a.Configuration.APIClient.CallAPI(path, method, pipelineCreationRequest, headers, queryParams)

	ret := v1.NewAPIResponse("CreatePipeline", method, path, queryParams)
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return ret, err
	}

	if response.StatusCode() != 200 {
		e := new(common_pdu.Error)
		err = json.Unmarshal(response.Body(), e)
		if err == nil {
			ret.Error = e
		}
	}

	return ret, err
}

func (a AdminApi) UpdatePipeline(pipelineUpdateRequest *pdu.PipelineUpdateRequest) (*v1.APIResponse, error) {
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

	response, err := a.Configuration.APIClient.CallAPI(path, method, pipelineUpdateRequest, headers, queryParams)

	ret := v1.NewAPIResponse("UpdatePipeline", method, path, queryParams)
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return ret, err
	}

	if response.StatusCode() != 200 {
		e := new(common_pdu.Error)
		err = json.Unmarshal(response.Body(), e)
		if err == nil {
			ret.Error = e
		}
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

	ret := v1.NewAPIResponse("GetPluginTypes", method, path, queryParams)
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return pdu, ret, err
	}

	if response.StatusCode() == 200 {
		err = json.Unmarshal(response.Body(), pdu)
	} else if len(response.Body()) > 0 {
		e := new(common_pdu.Error)
		err = json.Unmarshal(response.Body(), e)
		if err == nil {
			ret.Error = e
		}
	}

	return pdu, ret, err
}

func (a AdminApi) GetPlugins(pluginsRetrieveRequest *pdu.PluginsRetrieveRequest) (
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
	response, err := a.Configuration.APIClient.CallAPI(path, method, pluginsRetrieveRequest, headers, queryParams)

	ret := v1.NewAPIResponse("GetPlugins", method, path, queryParams)
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return pdu, ret, err
	}

	if response.StatusCode() == 200 {
		err = json.Unmarshal(response.Body(), pdu)
	} else if len(response.Body()) > 0 {
		e := new(common_pdu.Error)
		err = json.Unmarshal(response.Body(), e)
		if err == nil {
			ret.Error = e
		}
	}

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

	ret := v1.NewAPIResponse("DeletePluginByName", method, path, queryParams)
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return ret, err
	}

	if response.StatusCode() != 200 {
		e := new(common_pdu.Error)
		err = json.Unmarshal(response.Body(), e)
		if err == nil {
			ret.Error = e
		}
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

	ret := v1.NewAPIResponse("GetPluginByName", method, path, queryParams)
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return pdu, ret, err
	}

	if response.StatusCode() == 200 {
		err = json.Unmarshal(response.Body(), pdu)
	} else if len(response.Body()) > 0 {
		e := new(common_pdu.Error)
		err = json.Unmarshal(response.Body(), e)
		if err == nil {
			ret.Error = e
		}
	}

	return pdu, ret, err
}

func (a AdminApi) CreatePlugin(pluginCreationRequest *pdu.PluginCreationRequest) (*v1.APIResponse, error) {
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

	response, err := a.Configuration.APIClient.CallAPI(path, method, pluginCreationRequest, headers, queryParams)

	ret := v1.NewAPIResponse("CreatePlugin", method, path, queryParams)
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return ret, err
	}

	if response.StatusCode() != 200 {
		e := new(common_pdu.Error)
		err = json.Unmarshal(response.Body(), e)
		if err == nil {
			ret.Error = e
		}
	}

	return ret, err
}

func (a AdminApi) UpdatePlugin(pluginUpdateRequest *pdu.PluginUpdateRequest) (*v1.APIResponse, error) {
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

	response, err := a.Configuration.APIClient.CallAPI(path, method, pluginUpdateRequest, headers, queryParams)

	ret := v1.NewAPIResponse("UpdatePlugin", method, path, queryParams)
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return ret, err
	}

	if response.StatusCode() != 200 {
		e := new(common_pdu.Error)
		err = json.Unmarshal(response.Body(), e)
		if err == nil {
			ret.Error = e
		}
	}

	return ret, err
}

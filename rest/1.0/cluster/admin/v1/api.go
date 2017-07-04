package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/hexdecteam/easegateway-go-client/rest/1.0/cluster/admin/v1/pdu"
	"github.com/hexdecteam/easegateway-go-client/rest/1.0/common/v1"
	common_pdu "github.com/hexdecteam/easegateway-go-client/rest/1.0/common/v1/pdu"
)

type ClusterAdminApi struct {
	Configuration *v1.Configuration
}

func NewClusterAdminApi(addres) *ClusterAdminApi {
	address = strings.TrimSpace(address)
	if len(address) == 0 {
		address = "localhost:9090"
	}

	configuration := v1.NewConfiguration(fmt.Sprintf("http://%s/cluster/admin/v1", address))
	return &ClusterAdminApi{
		Configuration: configuration,
	}
}

func NewClusterAdminApiWithBasePath(basePath string) *ClusterAdminApi {
	configuration := v1.NewConfiguration(basePath)
	return &ClusterAdminApi{
		Configuration: configuration,
	}
}

func (a ClusterAdminApi) GetPipelineTypes(group string, clusterRetrieveRequest *pdu.ClusterRetrieveRequest) (
	*pdu.PipelineTypesRetrieveClusterResponse, *v1.APIResponse, error) {

	method := strings.ToUpper("Get")
	path := fmt.Sprintf("%s/%s/pipeline-types", a.Configuration.BasePath, group)
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

	pdu := new(pdu.PipelineTypesRetrieveClusterResponse)
	response, err := a.Configuration.APIClient.CallAPI(path, method, clusterRetrieveRequest, headers, queryParams)

	ret := v1.NewAPIResponse("GetPipelineTypes", method, path, queryParams)
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return pdu, ret, err
	}

	if response.StatusCode() == http.StatusOK {
		err = json.Unmarshal(response.Body(), pdu)
	} else if response.StatusCode() >= http.StatusBadRequest && len(response.Body()) > 0 {
		e := new(common_pdu.Error)
		err = json.Unmarshal(response.Body(), e)
		if err == nil {
			ret.Error = e
		}
	}

	return pdu, ret, err
}

func (a ClusterAdminApi) GetPipelines(group string,
	pipelinesRetrieveClusterRequest *pdu.PipelinesRetrieveClusterRequest) (
	*pdu.PipelinesRetrieveClusterResponse, *v1.APIResponse, error) {

	method := strings.ToUpper("Get")
	path := fmt.Sprintf("%s/%s/pipelines", a.Configuration.BasePath, group)
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

	pdu := new(pdu.PipelinesRetrieveClusterResponse)
	response, err := a.Configuration.APIClient.CallAPI(
		path, method, pipelinesRetrieveClusterRequest, headers, queryParams)

	ret := v1.NewAPIResponse("GetPipelines", method, path, queryParams)
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return pdu, ret, err
	}

	if response.StatusCode() == http.StatusOK {
		err = json.Unmarshal(response.Body(), pdu)
	} else if response.StatusCode() >= http.StatusBadRequest && len(response.Body()) > 0 {
		e := new(common_pdu.Error)
		err = json.Unmarshal(response.Body(), e)
		if err == nil {
			ret.Error = e
		}
	}

	return pdu, ret, err
}

func (a ClusterAdminApi) DeletePipelineByName(group string, pipelineName string,
	clusterOperationRequest *pdu.ClusterOperationRequest) (*v1.APIResponse, error) {

	method := strings.ToUpper("Delete")
	path := fmt.Sprintf("%s/%s/pipelines/%s", a.Configuration.BasePath, group, pipelineName)
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

	response, err := a.Configuration.APIClient.CallAPI(path, method, clusterOperationRequest, headers, queryParams)

	ret := v1.NewAPIResponse("DeletePipelineByName", method, path, queryParams)
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return ret, err
	}

	if response.StatusCode() >= http.StatusBadRequest {
		e := new(common_pdu.Error)
		err = json.Unmarshal(response.Body(), e)
		if err == nil {
			ret.Error = e
		}
	}

	return ret, err
}

func (a ClusterAdminApi) GetPipelineByName(group string, pipelineName string,
	clusterRetrieveRequest *pdu.ClusterRetrieveRequest) (*pdu.PipelineSpec, *v1.APIResponse, error) {

	method := strings.ToUpper("Get")
	path := fmt.Sprintf("%s/%s/pipelines/%s", a.Configuration.BasePath, group, pipelineName)
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
	response, err := a.Configuration.APIClient.CallAPI(path, method, clusterRetrieveRequest, headers, queryParams)

	ret := v1.NewAPIResponse("GetPipelineByName", method, path, queryParams)
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return pdu, ret, err
	}

	if response.StatusCode() == http.StatusOK {
		err = json.Unmarshal(response.Body(), pdu)
	} else if response.StatusCode() >= http.StatusBadRequest && len(response.Body()) > 0 {
		e := new(common_pdu.Error)
		err = json.Unmarshal(response.Body(), e)
		if err == nil {
			ret.Error = e
		}
	}

	return pdu, ret, err
}

func (a ClusterAdminApi) CreatePipeline(group string,
	pipelineCreationClusterRequest *pdu.PipelineCreationClusterRequest) (*v1.APIResponse, error) {

	method := strings.ToUpper("Post")
	path := fmt.Sprintf("%s/%s/pipelines", a.Configuration.BasePath, group)
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
		path, method, pipelineCreationClusterRequest, headers, queryParams)

	ret := v1.NewAPIResponse("CreatePipeline", method, path, queryParams)
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return ret, err
	}

	if response.StatusCode() >= http.StatusBadRequest {
		e := new(common_pdu.Error)
		err = json.Unmarshal(response.Body(), e)
		if err == nil {
			ret.Error = e
		}
	}

	return ret, err
}

func (a ClusterAdminApi) UpdatePipeline(group string,
	pipelineUpdateClusterRequest *pdu.PipelineUpdateClusterRequest) (*v1.APIResponse, error) {

	method := strings.ToUpper("Put")
	path := fmt.Sprintf("%s/%s/pipelines", a.Configuration.BasePath, group)
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
		path, method, pipelineUpdateClusterRequest, headers, queryParams)

	ret := v1.NewAPIResponse("UpdatePipeline", method, path, queryParams)
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return ret, err
	}

	if response.StatusCode() >= http.StatusBadRequest {
		e := new(common_pdu.Error)
		err = json.Unmarshal(response.Body(), e)
		if err == nil {
			ret.Error = e
		}
	}

	return ret, err
}

func (a ClusterAdminApi) GetPluginTypes(group string, clusterRetrieveRequest *pdu.ClusterRetrieveRequest) (
	*pdu.PluginTypesRetrieveClusterResponse, *v1.APIResponse, error) {

	method := strings.ToUpper("Get")
	path := fmt.Sprintf("%s/%s/plugin-types", a.Configuration.BasePath, group)
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

	pdu := new(pdu.PluginTypesRetrieveClusterResponse)
	response, err := a.Configuration.APIClient.CallAPI(path, method, clusterRetrieveRequest, headers, queryParams)

	ret := v1.NewAPIResponse("GetPluginTypes", method, path, queryParams)
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return pdu, ret, err
	}

	if response.StatusCode() == http.StatusOK {
		err = json.Unmarshal(response.Body(), pdu)
	} else if response.StatusCode() >= http.StatusBadRequest && len(response.Body()) > 0 {
		e := new(common_pdu.Error)
		err = json.Unmarshal(response.Body(), e)
		if err == nil {
			ret.Error = e
		}
	}

	return pdu, ret, err
}

func (a ClusterAdminApi) GetPlugins(group string, pluginsRetrieveClusterRequest *pdu.PluginsRetrieveClusterRequest) (
	*pdu.PluginsRetrieveClusterResponse, *v1.APIResponse, error) {

	method := strings.ToUpper("Get")
	path := fmt.Sprintf("%s/%s/plugins", a.Configuration.BasePath, group)
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

	pdu := new(pdu.PluginsRetrieveClusterResponse)
	response, err := a.Configuration.APIClient.CallAPI(
		path, method, pluginsRetrieveClusterRequest, headers, queryParams)

	ret := v1.NewAPIResponse("GetPlugins", method, path, queryParams)
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return pdu, ret, err
	}

	if response.StatusCode() == http.StatusOK {
		err = json.Unmarshal(response.Body(), pdu)
	} else if response.StatusCode() >= http.StatusBadRequest && len(response.Body()) > 0 {
		e := new(common_pdu.Error)
		err = json.Unmarshal(response.Body(), e)
		if err == nil {
			ret.Error = e
		}
	}

	return pdu, ret, err
}

func (a ClusterAdminApi) DeletePluginByName(group string, pluginName string,
	clusterOperationRequest *pdu.ClusterOperationRequest) (*v1.APIResponse, error) {

	method := strings.ToUpper("Delete")
	path := fmt.Sprintf("%s/%s/plugins/%s", a.Configuration.BasePath, group, pluginName)
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

	response, err := a.Configuration.APIClient.CallAPI(path, method, clusterOperationRequest, headers, queryParams)

	ret := v1.NewAPIResponse("DeletePluginByName", method, path, queryParams)
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return ret, err
	}

	if response.StatusCode() >= http.StatusBadRequest {
		e := new(common_pdu.Error)
		err = json.Unmarshal(response.Body(), e)
		if err == nil {
			ret.Error = e
		}
	}

	return ret, err
}

func (a ClusterAdminApi) GetPluginByName(group string, pluginName string,
	clusterRetrieveRequest *pdu.ClusterRetrieveRequest) (*pdu.PluginSpec, *v1.APIResponse, error) {

	method := strings.ToUpper("Get")
	path := fmt.Sprintf("%s/%s/plugins/%s", a.Configuration.BasePath, group, pluginName)
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
	response, err := a.Configuration.APIClient.CallAPI(path, method, clusterRetrieveRequest, headers, queryParams)

	ret := v1.NewAPIResponse("GetPluginByName", method, path, queryParams)
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return pdu, ret, err
	}

	if response.StatusCode() == http.StatusOK {
		err = json.Unmarshal(response.Body(), pdu)
	} else if response.StatusCode() >= http.StatusBadRequest && len(response.Body()) > 0 {
		e := new(common_pdu.Error)
		err = json.Unmarshal(response.Body(), e)
		if err == nil {
			ret.Error = e
		}
	}

	return pdu, ret, err
}

func (a ClusterAdminApi) CreatePlugin(group string,
	pluginCreationClusterRequest *pdu.PluginCreationClusterRequest) (*v1.APIResponse, error) {

	method := strings.ToUpper("Post")
	path := fmt.Sprintf("%s/%s/plugins", a.Configuration.BasePath, group)
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
		path, method, pluginCreationClusterRequest, headers, queryParams)

	ret := v1.NewAPIResponse("CreatePlugin", method, path, queryParams)
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return ret, err
	}

	if response.StatusCode() >= http.StatusBadRequest {
		e := new(common_pdu.Error)
		err = json.Unmarshal(response.Body(), e)
		if err == nil {
			ret.Error = e
		}
	}

	return ret, err
}

func (a ClusterAdminApi) GroupPluginsPut(group string,
	pluginUpdateClusterRequest *pdu.PluginUpdateClusterRequest) (*v1.APIResponse, error) {

	method := strings.ToUpper("Put")
	path := fmt.Sprintf("%s/%s/plugins", a.Configuration.BasePath, group)
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
		path, method, pluginUpdateClusterRequest, headers, queryParams)

	ret := v1.NewAPIResponse("UpdatePlugin", method, path, queryParams)
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return ret, err
	}

	if response.StatusCode() >= http.StatusBadRequest {
		e := new(common_pdu.Error)
		err = json.Unmarshal(response.Body(), e)
		if err == nil {
			ret.Error = e
		}
	}

	return ret, err
}

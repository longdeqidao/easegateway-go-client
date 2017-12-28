package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/hexdecteam/easegateway-go-client/rest/1.0/cluster/statistics/v1/pdu"
	"github.com/hexdecteam/easegateway-go-client/rest/1.0/common/v1"
	common_pdu "github.com/hexdecteam/easegateway-go-client/rest/1.0/common/v1/pdu"
)

type ClusterStatisticsApi struct {
	Configuration *v1.Configuration
}

func NewClusterStatisticsApi(address string) *ClusterStatisticsApi {
	address = strings.TrimSpace(address)
	if len(address) == 0 {
		address = "127.0.0.1:9090"
	}

	configuration := v1.NewConfiguration(fmt.Sprintf("http://%s/cluster/statistics/v1", address))
	return &ClusterStatisticsApi{
		Configuration: configuration,
	}
}

func NewClusterStatisticsApiWithBasePath(basePath string) *ClusterStatisticsApi {
	configuration := v1.NewConfiguration(basePath)
	return &ClusterStatisticsApi{
		Configuration: configuration,
	}
}

func (a *ClusterStatisticsApi) GetPipelineIndicatorNames(group string, pipelineName string,
	statisticsClusterRequest *pdu.StatisticsClusterRequest) (
	*pdu.ClusterPipelineIndicatorNames, *v1.APIResponse, error) {

	method := strings.ToUpper("Get")
	path := fmt.Sprintf("%s/%s/pipelines/%s/indicators", a.Configuration.BasePath, group, pipelineName)
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

	pdu := new(pdu.ClusterPipelineIndicatorNames)
	response, err := a.Configuration.APIClient.CallAPI(
		path, method, statisticsClusterRequest, headers, queryParams)

	ret := v1.NewAPIResponse("GetPipelineIndicatorNames", method, path, queryParams)
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

func (a *ClusterStatisticsApi) GetPipelineIndicatorDesc(group string, pipelineName string, indicatorName string,
	statisticsClusterRequest *pdu.StatisticsClusterRequest) (
	*pdu.ClusterPipelineIndicatorDescription, *v1.APIResponse, error) {

	method := strings.ToUpper("Get")
	path := fmt.Sprintf("%s/%s/pipelines/%s/indicators/%s/desc",
		a.Configuration.BasePath, group, pipelineName, indicatorName)
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

	pdu := new(pdu.ClusterPipelineIndicatorDescription)
	response, err := a.Configuration.APIClient.CallAPI(
		path, method, statisticsClusterRequest, headers, queryParams)

	ret := v1.NewAPIResponse("GetPipelineIndicatorDesc", method, path, queryParams)
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

func (a *ClusterStatisticsApi) GetPipelineIndicatorValue(group string, pipelineName string, indicatorName string,
	statisticsClusterRequest *pdu.StatisticsClusterRequest) (
	*pdu.ClusterPipelineIndicatorValue, *v1.APIResponse, error) {

	method := strings.ToUpper("Get")
	path := fmt.Sprintf("%s/%s/pipelines/%s/indicators/%s/value",
		a.Configuration.BasePath, group, pipelineName, indicatorName)
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

	pdu := new(pdu.ClusterPipelineIndicatorValue)
	response, err := a.Configuration.APIClient.CallAPI(
		path, method, statisticsClusterRequest, headers, queryParams)

	ret := v1.NewAPIResponse("GetPipelineIndicatorValue", method, path, queryParams)
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

func (a *ClusterStatisticsApi) GetPipelineIndicatorsValue(group string, pipelineName string,
	clusterPipelineIndicatorsValueRequest *pdu.ClusterPipelineIndicatorsValueRequest) (
	*pdu.ClusterPipelineIndicatorsValue, *v1.APIResponse, error) {

	method := strings.ToUpper("Get")
	path := fmt.Sprintf("%s/%s/pipelines/%s/indicators/value",
		a.Configuration.BasePath, group, pipelineName)
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

	pdu := new(pdu.ClusterPipelineIndicatorsValue)
	response, err := a.Configuration.APIClient.CallAPI(
		path, method, clusterPipelineIndicatorsValueRequest, headers, queryParams)

	ret := v1.NewAPIResponse("GetPipelineIndicatorValue", method, path, queryParams)
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

func (a *ClusterStatisticsApi) GetPluginIndicatorNames(group string, pipelineName string, pluginName string,
	statisticsClusterRequest *pdu.StatisticsClusterRequest) (
	*pdu.ClusterPluginIndicatorNames, *v1.APIResponse, error) {

	method := strings.ToUpper("Get")
	path := fmt.Sprintf("%s/%s/pipelines/%s/plugins/%s/indicators",
		a.Configuration.BasePath, group, pipelineName, pluginName)
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

	pdu := new(pdu.ClusterPluginIndicatorNames)
	response, err := a.Configuration.APIClient.CallAPI(
		path, method, statisticsClusterRequest, headers, queryParams)

	ret := v1.NewAPIResponse("GetPluginIndicatorNames", method, path, queryParams)
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

func (a *ClusterStatisticsApi) GetPluginIndicatorDesc(group string, pipelineName string, pluginName string,
	indicatorName string, statisticsClusterRequest *pdu.StatisticsClusterRequest) (
	*pdu.ClusterPluginIndicatorDescription, *v1.APIResponse, error) {

	method := strings.ToUpper("Get")
	path := fmt.Sprintf("%s/%s/pipelines/%s/plugins/%s/indicators/%s/desc",
		a.Configuration.BasePath, group, pipelineName, pluginName, indicatorName)

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

	pdu := new(pdu.ClusterPluginIndicatorDescription)
	response, err := a.Configuration.APIClient.CallAPI(
		path, method, statisticsClusterRequest, headers, queryParams)

	ret := v1.NewAPIResponse("GetPluginIndicatorDesc", method, path, queryParams)
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

func (a *ClusterStatisticsApi) GetPluginIndicatorValue(group string, pipelineName string, pluginName string,
	indicatorName string, statisticsClusterRequest *pdu.StatisticsClusterRequest) (
	*pdu.ClusterPluginIndicatorValue, *v1.APIResponse, error) {

	method := strings.ToUpper("Get")
	path := fmt.Sprintf("%s/%s/pipelines/%s/plugins/%s/indicators/%s/value",
		a.Configuration.BasePath, group, pipelineName, pluginName, indicatorName)
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

	pdu := new(pdu.ClusterPluginIndicatorValue)
	response, err := a.Configuration.APIClient.CallAPI(
		path, method, statisticsClusterRequest, headers, queryParams)

	ret := v1.NewAPIResponse("GetPluginIndicatorValue", method, path, queryParams)
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

func (a *ClusterStatisticsApi) GetTaskIndicatorNames(group string, pipelineName string,
	statisticsClusterRequest *pdu.StatisticsClusterRequest) (
	*pdu.ClusterTaskIndicatorNames, *v1.APIResponse, error) {

	method := strings.ToUpper("Get")
	path := fmt.Sprintf("%s/%s/pipelines/%s/task/indicators", a.Configuration.BasePath, group, pipelineName)
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

	pdu := new(pdu.ClusterTaskIndicatorNames)
	response, err := a.Configuration.APIClient.CallAPI(
		path, method, statisticsClusterRequest, headers, queryParams)

	ret := v1.NewAPIResponse("GetTaskIndicatorNames", method, path, queryParams)
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

func (a *ClusterStatisticsApi) GetTaskIndicatorDesc(group string, pipelineName string, indicatorName string,
	statisticsClusterRequest *pdu.StatisticsClusterRequest) (
	*pdu.ClusterTaskIndicatorDescription, *v1.APIResponse, error) {

	method := strings.ToUpper("Get")
	path := fmt.Sprintf("%s/%s/pipelines/%s/task/indicators/%s/desc",
		a.Configuration.BasePath, group, pipelineName, indicatorName)
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

	pdu := new(pdu.ClusterTaskIndicatorDescription)
	response, err := a.Configuration.APIClient.CallAPI(
		path, method, statisticsClusterRequest, headers, queryParams)

	ret := v1.NewAPIResponse("GetTaskIndicatorDesc", method, path, queryParams)
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

func (a *ClusterStatisticsApi) GetTaskIndicatorValue(group string, pipelineName string, indicatorName string,
	statisticsClusterRequest *pdu.StatisticsClusterRequest) (
	*pdu.ClusterTaskIndicatorValue, *v1.APIResponse, error) {

	method := strings.ToUpper("Get")
	path := fmt.Sprintf("%s/%s/pipelines/%s/task/indicators/%s/value",
		a.Configuration.BasePath, group, pipelineName, indicatorName)
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

	pdu := new(pdu.ClusterTaskIndicatorValue)
	response, err := a.Configuration.APIClient.CallAPI(
		path, method, statisticsClusterRequest, headers, queryParams)

	ret := v1.NewAPIResponse("GetTaskIndicatorValue", method, path, queryParams)
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

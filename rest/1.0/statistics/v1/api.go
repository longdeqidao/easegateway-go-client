package v1

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/hexdecteam/easegateway-go-client/rest/1.0/common/v1"
	common_pdu "github.com/hexdecteam/easegateway-go-client/rest/1.0/common/v1/pdu"
	"github.com/hexdecteam/easegateway-go-client/rest/1.0/statistics/v1/pdu"
)

type StatisticsApi struct {
	Configuration *v1.Configuration
}

func NewStatisticsApi(address string) *StatisticsApi {
	address = strings.TrimSpace(address)
	if len(address) == 0 {
		address = "localhost:9090"
	}

	configuration := v1.NewConfiguration(fmt.Sprintf("http://%s/statistics/v1", address))
	return &StatisticsApi{
		Configuration: configuration,
	}
}

func NewStatisticsApiWithBasePath(basePath string) *StatisticsApi {
	configuration := v1.NewConfiguration(basePath)
	return &StatisticsApi{
		Configuration: configuration,
	}
}

func (a StatisticsApi) GetGatewayAverageLoad() (*pdu.AvgLoad, *v1.APIResponse, error) {
	method := strings.ToUpper("Get")
	path := a.Configuration.BasePath + "/gateway/loadavg"
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

	pdu := new(pdu.AvgLoad)
	response, err := a.Configuration.APIClient.CallAPI(path, method, nil, headers, queryParams)

	ret := v1.NewAPIResponse("GetGatewayAverageLoad", method, path, queryParams)
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
		e := common_pdu.Error{}
		err = json.Unmarshal(response.Body(), &e)
		if err == nil {
			ret.Error = e
		}
	}

	return pdu, ret, err
}

func (a StatisticsApi) GetGatewayResourceUsage() (*pdu.ResourceUsage, *v1.APIResponse, error) {
	method := strings.ToUpper("Get")
	path := a.Configuration.BasePath + "/gateway/rusage"
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

	pdu := new(pdu.ResourceUsage)
	response, err := a.Configuration.APIClient.CallAPI(path, method, nil, headers, queryParams)

	ret := v1.NewAPIResponse("GetGatewayResourceUsage", method, path, queryParams)
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
		e := common_pdu.Error{}
		err = json.Unmarshal(response.Body(), &e)
		if err == nil {
			ret.Error = e
		}
	}

	return pdu, ret, err
}

func (a StatisticsApi) GetGatewayUpTime() (*pdu.UpTime, *v1.APIResponse, error) {
	method := strings.ToUpper("Get")
	path := a.Configuration.BasePath + "/gateway/uptime"
	headers := make(map[string]string)
	queryParams := url.Values{}

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headers[key] = a.Configuration.DefaultHeader[key]
	}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType([]string{"application/json"})
	if localVarHttpContentType != "" {
		headers["Content-Type"] = localVarHttpContentType
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept([]string{"application/json"})
	if localVarHttpHeaderAccept != "" {
		headers["Accept"] = localVarHttpHeaderAccept
	}

	pdu := new(pdu.UpTime)
	response, err := a.Configuration.APIClient.CallAPI(path, method, nil, headers, queryParams)

	ret := v1.NewAPIResponse("GetGatewayUpTime", method, path, queryParams)
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
		e := common_pdu.Error{}
		err = json.Unmarshal(response.Body(), &e)
		if err == nil {
			ret.Error = e
		}
	}

	return pdu, ret, err
}

func (a StatisticsApi) GetPipelineIndicatorNames(pipelineName string) (
	*pdu.PipelineIndicatorNames, *v1.APIResponse, error) {

	method := strings.ToUpper("Get")
	path := fmt.Sprintf("%s/pipelines/%s/indicators", a.Configuration.BasePath, pipelineName)
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

	pdu := new(pdu.PipelineIndicatorNames)
	response, err := a.Configuration.APIClient.CallAPI(path, method, nil, headers, queryParams)

	ret := v1.NewAPIResponse("GetPipelineIndicatorNames", method, path, queryParams)
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
		e := common_pdu.Error{}
		err = json.Unmarshal(response.Body(), &e)
		if err == nil {
			ret.Error = e
		}
	}

	return pdu, ret, err
}

func (a StatisticsApi) GetPipelineIndicatorDesc(pipelineName string, indicatorName string) (
	*pdu.PipelineIndicatorDescription, *v1.APIResponse, error) {

	method := strings.ToUpper("Get")
	path := fmt.Sprintf("%s/pipelines/%s/indicators/%s/desc", a.Configuration.BasePath, pipelineName, indicatorName)
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

	pdu := new(pdu.PipelineIndicatorDescription)
	response, err := a.Configuration.APIClient.CallAPI(path, method, nil, headers, queryParams)

	ret := v1.NewAPIResponse("GetPipelineIndicatorDesc", method, path, queryParams)
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
		e := common_pdu.Error{}
		err = json.Unmarshal(response.Body(), &e)
		if err == nil {
			ret.Error = e
		}
	}

	return pdu, ret, err
}

func (a StatisticsApi) GetPipelineIndicatorValue(pipelineName string, indicatorName string) (
	*pdu.PipelineIndicatorValue, *v1.APIResponse, error) {

	method := strings.ToUpper("Get")
	path := fmt.Sprintf("%s/pipelines/%s/indicators/%s/value", a.Configuration.BasePath, pipelineName, indicatorName)
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

	pdu := new(pdu.PipelineIndicatorValue)
	response, err := a.Configuration.APIClient.CallAPI(path, method, nil, headers, queryParams)

	ret := v1.NewAPIResponse("GetPipelineIndicatorValue", method, path, queryParams)
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
		e := common_pdu.Error{}
		err = json.Unmarshal(response.Body(), &e)
		if err == nil {
			ret.Error = e
		}
	}

	return pdu, ret, err
}

func (a StatisticsApi) GetPluginIndicatorNames(pipelineName string, pluginName string) (
	*pdu.PluginIndicatorNames, *v1.APIResponse, error) {

	method := strings.ToUpper("Get")
	path := fmt.Sprintf("%s/pipelines/%s/plugins/%s/indicators", a.Configuration.BasePath, pipelineName, pluginName)
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

	pdu := new(pdu.PluginIndicatorNames)
	response, err := a.Configuration.APIClient.CallAPI(path, method, nil, headers, queryParams)

	ret := v1.NewAPIResponse("GetPluginIndicatorNames", method, path, queryParams)
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
		e := common_pdu.Error{}
		err = json.Unmarshal(response.Body(), &e)
		if err == nil {
			ret.Error = e
		}
	}

	return pdu, ret, err
}

func (a StatisticsApi) GetPluginIndicatorDesc(pipelineName string, pluginName string, indicatorName string) (
	*pdu.PluginIndicatorDescription, *v1.APIResponse, error) {

	method := strings.ToUpper("Get")
	// create path and map variables
	path := fmt.Sprintf("%s/pipelines/%s/plugins/%s/indicators/%s/desc",
		a.Configuration.BasePath, pipelineName, pluginName, indicatorName)
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

	pdu := new(pdu.PluginIndicatorDescription)
	response, err := a.Configuration.APIClient.CallAPI(path, method, nil, headers, queryParams)

	ret := v1.NewAPIResponse("GetPluginIndicatorDesc", method, path, queryParams)
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
		e := common_pdu.Error{}
		err = json.Unmarshal(response.Body(), &e)
		if err == nil {
			ret.Error = e
		}
	}

	return pdu, ret, err
}

func (a StatisticsApi) GetPluginIndicatorValue(pipelineName string, pluginName string, indicatorName string) (
	*pdu.PluginIndicatorValue, *v1.APIResponse, error) {

	method := strings.ToUpper("Get")
	path := fmt.Sprintf("%s/pipelines/%s/plugins/%s/indicators/%s/value",
		a.Configuration.BasePath, pipelineName, pluginName, indicatorName)
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

	pdu := new(pdu.PluginIndicatorValue)
	response, err := a.Configuration.APIClient.CallAPI(path, method, nil, headers, queryParams)

	ret := v1.NewAPIResponse("GetPluginIndicatorValue", method, path, queryParams)
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
		e := common_pdu.Error{}
		err = json.Unmarshal(response.Body(), &e)
		if err == nil {
			ret.Error = e
		}
	}

	return pdu, ret, err
}

func (a StatisticsApi) GetTaskIndicatorNames(pipelineName string) (*pdu.TaskIndicatorNames, *v1.APIResponse, error) {

	method := strings.ToUpper("Get")
	// create path and map variables
	path := fmt.Sprintf("%s/pipelines/%s/task/indicators", a.Configuration.BasePath, pipelineName)
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

	pdu := new(pdu.TaskIndicatorNames)
	response, err := a.Configuration.APIClient.CallAPI(path, method, nil, headers, queryParams)

	ret := v1.NewAPIResponse("GetTaskIndicatorNames", method, path, queryParams)
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
		e := common_pdu.Error{}
		err = json.Unmarshal(response.Body(), &e)
		if err == nil {
			ret.Error = e
		}
	}

	return pdu, ret, err
}

func (a StatisticsApi) GetTaskIndicatorDesc(pipelineName string, indicatorName string) (
	*pdu.TaskIndicatorDescription, *v1.APIResponse, error) {

	method := strings.ToUpper("Get")
	path := fmt.Sprintf("%s/pipelines/%s/task/indicators/%s/desc",
		a.Configuration.BasePath, pipelineName, indicatorName)
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

	pdu := new(pdu.TaskIndicatorDescription)
	response, err := a.Configuration.APIClient.CallAPI(path, method, nil, headers, queryParams)

	ret := v1.NewAPIResponse("GetTaskIndicatorDesc", method, path, queryParams)
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
		e := common_pdu.Error{}
		err = json.Unmarshal(response.Body(), &e)
		if err == nil {
			ret.Error = e
		}
	}

	return pdu, ret, err
}

func (a StatisticsApi) GetTaskIndicatorValue(pipelineName string, indicatorName string) (
	*pdu.TaskIndicatorValue, *v1.APIResponse, error) {

	method := strings.ToUpper("Get")
	path := fmt.Sprintf("%s/pipelines/%s/task/indicators/%s/value",
		a.Configuration.BasePath, pipelineName, indicatorName)
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

	pdu := new(pdu.TaskIndicatorValue)
	response, err := a.Configuration.APIClient.CallAPI(path, method, nil, headers, queryParams)

	ret := v1.NewAPIResponse("GetTaskIndicatorValue", method, path, queryParams)
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
		e := common_pdu.Error{}
		err = json.Unmarshal(response.Body(), &e)
		if err == nil {
			ret.Error = e
		}
	}

	return pdu, ret, err
}

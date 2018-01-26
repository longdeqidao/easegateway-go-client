package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/hexdecteam/easegateway-go-client/rest/1.0/cluster/health/v1/pdu"
	"github.com/hexdecteam/easegateway-go-client/rest/1.0/common/v1"
	common_pdu "github.com/hexdecteam/easegateway-go-client/rest/1.0/common/v1/pdu"
)

type ClusterHealthApi struct {
	Configuration *v1.Configuration
}

func NewClusterHealthApi(address string) *ClusterHealthApi {
	address = strings.TrimSpace(address)
	if len(address) == 0 {
		address = "localhost:9090"
	}

	configuration := v1.NewConfiguration(fmt.Sprintf("http://%s/cluster/health/v1", address))
	return &ClusterHealthApi{
		Configuration: configuration,
	}
}

func NewClusterHealthApiWithBasePath(basePath string) *ClusterHealthApi {
	configuration := v1.NewConfiguration(basePath)
	return &ClusterHealthApi{
		Configuration: configuration,
	}
}

func (a *ClusterHealthApi) GetGroupHealth(group string, clusterRequest *common_pdu.ClusterRequest) (
	*pdu.ClusterQueryGroupHealthResponse, *v1.APIResponse, error) {
	path := fmt.Sprintf("%s/check/groups/%s", a.Configuration.BasePath, group)
	pdu := new(pdu.ClusterQueryGroupHealthResponse)
	body, ret, err := a.getResponse(clusterRequest, "GetGroupHealth", path)
	if err != nil || ret.Error != nil {
		return nil, ret, err
	}
	err = json.Unmarshal(body, &pdu)

	return pdu, ret, err
}

func (a *ClusterHealthApi) GetGroupsHealth(clusterRequest *common_pdu.ClusterRequest) (
	*pdu.ClusterQueryGroupsHealthResponse, *v1.APIResponse, error) {
	path := fmt.Sprintf("%s/check/groups", a.Configuration.BasePath)
	pdu := new(pdu.ClusterQueryGroupsHealthResponse)
	body, ret, err := a.getResponse(clusterRequest, "GetGroupsHealth", path)
	if err != nil || ret.Error != nil {
		return nil, ret, err
	}
	err = json.Unmarshal(body, &pdu)

	return pdu, ret, err
}

func (a *ClusterHealthApi) GetGroups(clusterRequest *common_pdu.ClusterRequest) (
	*pdu.ClusterQueryGroupsResponse, *v1.APIResponse, error) {

	path := fmt.Sprintf("%s/info/groups", a.Configuration.BasePath)
	pdu := new(pdu.ClusterQueryGroupsResponse)
	body, ret, err := a.getResponse(clusterRequest, "GetGroups", path)
	if err != nil || ret.Error != nil {
		return nil, ret, err
	}
	err = json.Unmarshal(body, &pdu)

	return pdu, ret, err
}

func (a *ClusterHealthApi) GetMembers(group string, clusterRequest *common_pdu.ClusterRequest) (
	*pdu.ClusterQueryMembersResponse, *v1.APIResponse, error) {

	path := fmt.Sprintf("%s/info/groups/%s/members", a.Configuration.BasePath, group)
	pdu := new(pdu.ClusterQueryMembersResponse)
	body, ret, err := a.getResponse(clusterRequest, "GetMembers", path)
	if err != nil || ret.Error != nil {
		return nil, ret, err
	}
	err = json.Unmarshal(body, &pdu)

	return pdu, ret, err
}

func (a *ClusterHealthApi) GetGroup(group string, clusterRequest *common_pdu.ClusterRequest) (
	*pdu.ClusterQueryGroupResponse, *v1.APIResponse, error) {

	path := fmt.Sprintf("%s/info/groups/%s", a.Configuration.BasePath, group)
	pdu := new(pdu.ClusterQueryGroupResponse)
	body, ret, err := a.getResponse(clusterRequest, "GetGroup", path)
	if err != nil || ret.Error != nil {
		return nil, ret, err
	}
	err = json.Unmarshal(body, &pdu)

	return pdu, ret, err
}

func (a *ClusterHealthApi) GetMember(group, member string, clusterRequest *common_pdu.ClusterRequest) (
	*pdu.ClusterQueryMemberResponse, *v1.APIResponse, error) {

	path := fmt.Sprintf("%s/info/groups/%s/members/%s", a.Configuration.BasePath, group, member)
	pdu := new(pdu.ClusterQueryMemberResponse)
	body, ret, err := a.getResponse(clusterRequest, "GetMember", path)
	if err != nil || ret.Error != nil {
		return nil, ret, err
	}
	err = json.Unmarshal(body, &pdu)

	return pdu, ret, err
}

func (a *ClusterHealthApi) getResponse(clusterRequest *common_pdu.ClusterRequest, operationName, path string) ([]byte, *v1.APIResponse, error) {
	method := strings.ToUpper("Get")
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
		path, method, clusterRequest, headers, queryParams)

	ret := v1.NewAPIResponse(operationName, method, path, queryParams)
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return nil, ret, err
	}

	var body []byte
	if response.StatusCode() == http.StatusOK {
		body = response.Body()
	} else if response.StatusCode() >= http.StatusBadRequest && len(response.Body()) > 0 {
		e := new(common_pdu.Error)
		err = json.Unmarshal(response.Body(), e)
		if err == nil {
			ret.Error = e
		} else {
			ret.Error = &common_pdu.Error{
				Error: "unmarshal error response failed",
			}
		}
	}

	return body, ret, err
}

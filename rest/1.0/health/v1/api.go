package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/hexdecteam/easegateway-go-client/rest/1.0/common/v1"
	common_pdu "github.com/hexdecteam/easegateway-go-client/rest/1.0/common/v1/pdu"
)

type HealthApi struct {
	Configuration *v1.Configuration
}

func NewHealthApi(address string) *HealthApi {
	address = strings.TrimSpace(address)
	if len(address) == 0 {
		address = "localhost:9090"
	}

	configuration := v1.NewConfiguration(fmt.Sprintf("http://%s/health/v1", address))
	return &HealthApi{
		Configuration: configuration,
	}
}

func NewHealthApiWithBasePath(basePath string) *HealthApi {
	configuration := v1.NewConfiguration(basePath)
	return &HealthApi{
		Configuration: configuration,
	}
}

func (a HealthApi) Check() (*v1.APIResponse, error) {
	method := strings.ToUpper("Get")
	path := a.Configuration.BasePath + "/check"
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
	accepts := a.Configuration.APIClient.SelectHeaderAccept([]string{"application/json"})
	if accepts != "" {
		headers["Accept"] = accepts
	}
	response, err := a.Configuration.APIClient.CallAPI(path, method, nil, headers, queryParams)

	ret := v1.NewAPIResponse("Check", method, path, queryParams)
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return ret, err
	}

	if response.StatusCode() != http.StatusOK {
		e := new(common_pdu.Error)
		err = json.Unmarshal(response.Body(), e)
		if err == nil {
			ret.Error = e
		}
	}

	return ret, err
}

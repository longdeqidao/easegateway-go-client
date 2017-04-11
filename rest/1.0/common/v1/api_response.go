package v1

import (
	"net/http"
	"net/url"

	"github.com/hexdecteam/easegateway-go-client/rest/1.0/common/v1/pdu"
)

type APIResponse struct {
	*http.Response `json:"-"`
	Error          *pdu.Error `json:"message,omitempty"`
	Operation      string    `json:"operation,omitempty"`
	RequestURL     *url.URL  `json:"url,omitempty"`
	Method         string    `json:"method,omitempty"`
	Payload        []byte    `json:"-"`
}

func NewAPIResponse(operation, method, path string, queryParams url.Values) *APIResponse {
	u, _ := url.Parse(path)
	u.RawQuery = queryParams.Encode()
	response := &APIResponse{Operation: operation, Method: method, RequestURL: u}
	return response
}

func NewAPIResponseWithError(message string) *APIResponse {
	response := &APIResponse{Error: &pdu.Error{Error: message}}
	return response
}

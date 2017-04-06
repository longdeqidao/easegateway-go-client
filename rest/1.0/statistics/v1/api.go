package v1

import (
	"github.com/hexdecteam/easegateway-go-client/rest/1.0/common/v1"
)

type DefaultApi struct {
	Configuration *v1.Configuration
}

func NewStatisticsApi() *DefaultApi {
	configuration := v1.NewConfiguration("http://localhost:9090/statistics/v1")
	return &DefaultApi{
		Configuration: configuration,
	}
}

func NewStatisticsApiWithBasePath(basePath string) *DefaultApi {
	configuration := v1.NewConfiguration(basePath)
	return &DefaultApi{
		Configuration: configuration,
	}
}

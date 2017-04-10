# Go API client for Ease Gateway

## Overview

- API version: 1.0
- Package version: 1.0

## Installation
1. Get the package under your project folder

```
go get -d -v github.com/hexdecteam/easegateway-go-client
```

2. Add the following in import:

```
import "github.com/hexdecteam/easegateway-go-client"
```

## Documentation for Administration API Endpoints

All URIs are relative to *http://localhost:9090/admin/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AdminApi* | **GetPipelineTypes** | **Get** /pipeline-types | Retrieves pipeline type
*AdminApi* | **GetPipelines** | **Get** /pipelines | Retrieves pipeline instances
*AdminApi* | **DeletePipelineByName** | **Delete** /pipelines/{pipelineName} | Deletes pipeline instance
*AdminApi* | **GetPipelineByName** | **Get** /pipelines/{pipelineName} | Retrieves pipeline instance
*AdminApi* | **CreatePipeline** | **Post** /pipelines | Creates pipeline instance
*AdminApi* | **UpdatePipeline** | **Put** /pipelines | Updates pipeline instance
*AdminApi* | **GetPluginTypes** | **Get** /plugin-types | Retrieves plugin type
*AdminApi* | **GetPlugins** | **Get** /plugins | Retrieves plugin instances
*AdminApi* | **DeletePluginByName** | **Delete** /plugins/{pluginName} | Deletes plugin instance
*AdminApi* | **GetPluginByName** | **Get** /plugins/{pluginName} | Retrieves plugin instance
*AdminApi* | **CreatePlugin** | **Post** /plugins | Creates plugin instance
*AdminApi* | **UpdatePlugin** | **Put** /plugins | Updates plugin instance

## Documentation for Statistics API Endpoints

All URIs are relative to *http://localhost:9090/statistics/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*StatisticsApi* | **GetGatewayAverageLoad** | **Get** /gateway/loadavg | Retrieves gateway host average load
*StatisticsApi* | **GetGatewayResourceUsage** | **Get** /gateway/rusage | Retrieves resource usage of gateway service instance
*StatisticsApi* | **GetGatewayUpTime** | **Get** /gateway/uptime | Retrieves up time of gateway service instance
*StatisticsApi* | **GetPipelineIndicatorNames** | **Get** /pipelines/{pipelineName}/indicators | Retrieves pipeline statistics indicator names
*StatisticsApi* | **GetPipelineIndicatorValue** | **Get** /pipelines/{pipelineName}/indicators/{indicatorName}/value | Retrieves pipeline statistics indicator value
*StatisticsApi* | **GetPipelineIndicatorDesc** | **Get** /pipelines/{pipelineName}/indicators/{indicatorName}/desc | Retrieves pipeline statistics indicator description
*StatisticsApi* | **GetPluginIndicatorNames** | **Get** /pipelines/{pipelineName}/plugins/{pluginName}/indicators | Retrieves plugin statistics indicator names
*StatisticsApi* | **GetPluginIndicatorValue** | **Get** /pipelines/{pipelineName}/plugins/{pluginName}/indicators/{indicatorName}/value | Retrieves plugin statistics indicator value
*StatisticsApi* | **GetPluginIndicatorDesc** | **Get** /pipelines/{pipelineName}/plugins/{pluginName}/indicators/{indicatorName}/desc | Retrieves plugin statistics indicator description
*StatisticsApi* | **GetTaskIndicatorNames** | **Get** /pipelines/{pipelineName}/task/indicators | Retrieves task statistics indicator names
*StatisticsApi* | **GetTaskIndicatorValue** | **Get** /pipelines/{pipelineName}/task/indicators/{indicatorName}/value | Retrieves task statistics indicator value
*StatisticsApi* | **GetTaskIndicatorDesc** | **Get** /pipelines/{pipelineName}/task/indicators/{indicatorName}/desc | Retrieves task statistics indicator description

## Documentation for Health API Endpoints

All URIs are relative to *http://localhost:9090/health/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*HealthApi* | **Check** | **Get** /check | Checks gateway service instance existing

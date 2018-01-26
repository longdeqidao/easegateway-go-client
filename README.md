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

## Administration API Endpoints

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

## Statistics API Endpoints

All URIs are relative to *http://localhost:9090/statistics/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*StatisticsApi* | **GetGatewayAverageLoad** | **Get** /gateway/loadavg | Retrieves gateway host average load
*StatisticsApi* | **GetGatewayResourceUsage** | **Get** /gateway/rusage | Retrieves resource usage of gateway service instance
*StatisticsApi* | **GetGatewayUpTime** | **Get** /gateway/uptime | Retrieves up time of gateway service instance
*StatisticsApi* | **GetPipelineIndicatorNames** | **Get** /pipelines/{pipelineName}/indicators | Retrieves pipeline statistics indicator names
*StatisticsApi* | **GetPipelineIndicatorValue** | **Get** /pipelines/{pipelineName}/indicators/{indicatorName}/value | Retrieves pipeline statistics indicator value
*StatisticsApi* | **GetPipelineIndicatorsValue** | **Get** /pipelines/{pipelineName}/indicators/value | Retrieves pipeline statistics values from multiple indicators
*StatisticsApi* | **GetPipelineIndicatorDesc** | **Get** /pipelines/{pipelineName}/indicators/{indicatorName}/desc | Retrieves pipeline statistics indicator description
*StatisticsApi* | **GetPluginIndicatorNames** | **Get** /pipelines/{pipelineName}/plugins/{pluginName}/indicators | Retrieves plugin statistics indicator names
*StatisticsApi* | **GetPluginIndicatorValue** | **Get** /pipelines/{pipelineName}/plugins/{pluginName}/indicators/{indicatorName}/value | Retrieves plugin statistics indicator value
*StatisticsApi* | **GetPluginIndicatorDesc** | **Get** /pipelines/{pipelineName}/plugins/{pluginName}/indicators/{indicatorName}/desc | Retrieves plugin statistics indicator description
*StatisticsApi* | **GetTaskIndicatorNames** | **Get** /pipelines/{pipelineName}/task/indicators | Retrieves task statistics indicator names
*StatisticsApi* | **GetTaskIndicatorValue** | **Get** /pipelines/{pipelineName}/task/indicators/{indicatorName}/value | Retrieves task statistics indicator value
*StatisticsApi* | **GetTaskIndicatorDesc** | **Get** /pipelines/{pipelineName}/task/indicators/{indicatorName}/desc | Retrieves task statistics indicator description

## Health API Endpoints

All URIs are relative to *http://localhost:9090/health/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*HealthApi* | **Check** | **Get** /check | Checks gateway service instance existing
*HealthApi* | **GetInfo** | **Get** /info | Retrieves gateway service instance information

## Cluster Administration API Endpoints

All URIs are relative to *http://localhost:9090/cluster/admin/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ClusterAdminApi* | **GetMaxOperationSequence** | **Get** /{group}/sequence | Retrieves maximum operation sequence in a group
*ClusterAdminApi* | **GetPipelineTypes** | **Get** /{group}/pipeline-types | Retrieves pipeline type in a group
*ClusterAdminApi* | **GetPipelines** | **Get** /{group}/pipelines | Retrieves pipeline instances in a group
*ClusterAdminApi* | **DeletePipelineByName** | **Delete** /{group}/pipelines/{pipelineName} | Deletes pipeline instance in a group
*ClusterAdminApi* | **GetPipelineByName** | **Get** /{group}/pipelines/{pipelineName} | Retrieves pipeline instance in a group
*ClusterAdminApi* | **CreatePipeline** | **Post** /{group}/pipelines | Creates pipeline instance in a group
*ClusterAdminApi* | **UpdatePipeline** | **Put** /{group}/pipelines | Updates pipeline instance in a group
*ClusterAdminApi* | **GetPluginTypes** | **Get** /{group}/plugin-types | Retrieves plugin type in a group
*ClusterAdminApi* | **GetPlugins** | **Get** /{group}/plugins | Retrieves plugin instances in a group
*ClusterAdminApi* | **DeletePluginByName** | **Delete** /{group}/plugins/{pluginName} | Deletes plugin instance in a group
*ClusterAdminApi* | **GetPluginByName** | **Get** /{group}/plugins/{pluginName} | Retrieves plugin instance in a group
*ClusterAdminApi* | **CreatePlugin** | **Post** /{group}/plugins | Creates plugin instance in a group
*ClusterAdminApi* | **UpdatePlugin** | **Put** /{group}/plugins | Updates plugin instance in a group

## Cluster Statistics API Endpoints

All URIs are relative to *http://localhost:9090/cluster/statistics/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ClusterStatisticsApi* | **GetPipelineIndicatorNames** | **Get** /{group}/pipelines/{pipelineName}/indicators | Retrieves pipeline statistics indicator names in a group
*ClusterStatisticsApi* | **GetPipelineIndicatorValue** | **Get** /{group}/pipelines/{pipelineName}/indicators/{indicatorName}/value | Retrieves pipeline statistics indicator value in a group
*ClusterStatisticsApi* | **GetPipelineIndicatorsValue** | **Get** /{group}/pipelines/{pipelineName}/indicators/value | Retrieves pipeline statistics values from multiple indicators in a group
*ClusterStatisticsApi* | **GetPipelineIndicatorDesc** | **Get** /{group}/pipelines/{pipelineName}/indicators/{indicatorName}/desc | Retrieves pipeline statistics indicator description in a group
*ClusterStatisticsApi* | **GetPluginIndicatorNames** | **Get** /{group}/pipelines/{pipelineName}/plugins/{pluginName}/indicators | Retrieves plugin statistics indicator names in a group
*ClusterStatisticsApi* | **GetPluginIndicatorValue** | **Get** /{group}/pipelines/{pipelineName}/plugins/{pluginName}/indicators/{indicatorName}/value | Retrieves plugin statistics indicator value in a group
*ClusterStatisticsApi* | **GetPluginIndicatorDesc** | **Get** /{group}/pipelines/{pipelineName}/plugins/{pluginName}/indicators/{indicatorName}/desc | Retrieves plugin statistics indicator description in a group
*ClusterStatisticsApi* | **GetTaskIndicatorNames** | **Get** /{group}/pipelines/{pipelineName}/task/indicators | Retrieves task statistics indicator names in a group
*ClusterStatisticsApi* | **GetTaskIndicatorValue** | **Get** /{group}/pipelines/{pipelineName}/task/indicators/{indicatorName}/value | Retrieves task statistics indicator value in a group
*ClusterStatisticsApi* | **GetTaskIndicatorDesc** | **Get** /{group}/pipelines/{pipelineName}/task/indicators/{indicatorName}/desc | Retrieves task statistics indicator description in a group

## Cluster Administration API Endpoints

All URIs are relative to *http://localhost:9090/cluster/health/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ClusterHealthApi* | **GetGroupHealth** | **Get** /check/groups/{group} | Retrieves health check status in a group
*ClusterHealthApi* | **GetGroupsHealth** | **Get** /check/groups | Retrieves health check status in a cluster
*ClusterHealthApi* | **GetGroups** | **Get** /info/groups | Retrieves groups in a cluster
*ClusterHealthApi* | **GetMembers** | **Get** /info/groups/{group}/members | Retrieves members in a group
*ClusterHealthApi* | **GetGroup** | **Get** /info/groups/{group} | Retrieves group information in a cluster
*ClusterHealthApi* | **GetMember** | **Get** /info/groups/{group}/members/{member} | Retrieves member information in a group

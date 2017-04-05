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

All URIs are relative to *http://localhost/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AdminApi* | **GetPipelineTypes** | **Get** /pipeline-types | Retrieves Pipeline Type
*AdminApi* | **GetPipelines** | **Get** /pipelines | Retrieves Pipeline Instances
*AdminApi* | **DeletePipelineByName** | **Delete** /pipelines/{pipelineName} | Deletes Pipeline Instance
*AdminApi* | **GetPipelineByName** | **Get** /pipelines/{pipelineName} | Retrieves Pipeline Instance
*AdminApi* | **CreatePipeline** | **Post** /pipelines | Creates Pipwline Instance
*AdminApi* | **UpdatePipeline** | **Put** /pipelines | Updates Pipeline Instance
*AdminApi* | **GetPluginTypes** | **Get** /plugin-types | Retrieves Plugin Type
*AdminApi* | **GetPlugins** | **Get** /plugins | Retrieves Plugin Instances
*AdminApi* | **DeletePluginByName** | **Delete** /plugins/{pluginName} | Deletes Plugin Instance
*AdminApi* | **GetPluginByName** | **Get** /plugins/{pluginName} | Retrieves Plugin Instance
*AdminApi* | **CreatePlugin** | **Post** /plugins | Creates Plugin Instance
*AdminApi* | **UpdatePlugin** | **Put** /plugins | Updates Plugin Instance

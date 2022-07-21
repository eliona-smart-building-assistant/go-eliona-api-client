# Go API client for api

API to access Eliona Smart Building Assistant

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.0.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import api "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), api.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), api.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), api.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), api.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://api.eliona.io/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AppsApi* | [**GetAppByName**](docs/AppsApi.md#getappbyname) | **Get** /apps/{app-name} | Information about an app
*AppsApi* | [**GetPatchByName**](docs/AppsApi.md#getpatchbyname) | **Get** /apps/{app-name}/patches/{patch-name} | Information about a patch for an app
*AppsApi* | [**PatchAppByName**](docs/AppsApi.md#patchappbyname) | **Patch** /apps/{app-name} | Update an app
*AppsApi* | [**PatchPatchByName**](docs/AppsApi.md#patchpatchbyname) | **Patch** /apps/{app-name}/patches/{patch-name} | Updates a patch
*AssetTypesApi* | [**GetAssetTypes**](docs/AssetTypesApi.md#getassettypes) | **Get** /asset-types | List of asset types
*AssetTypesApi* | [**PutAssetType**](docs/AssetTypesApi.md#putassettype) | **Put** /asset-types | Create or update an asset type
*AssetTypesApi* | [**PutAssetTypeAttribute**](docs/AssetTypesApi.md#putassettypeattribute) | **Put** /asset-types/{asset-type-name}/attributes | Create or update an asset type attribute
*AssetsApi* | [**GetAssetById**](docs/AssetsApi.md#getassetbyid) | **Get** /assets/{asset-id} | Information about an asset
*AssetsApi* | [**GetAssets**](docs/AssetsApi.md#getassets) | **Get** /assets | Information about assets
*AssetsApi* | [**PutAsset**](docs/AssetsApi.md#putasset) | **Put** /assets | Create or update an asset
*DashboardsApi* | [**PostDashboard**](docs/DashboardsApi.md#postdashboard) | **Post** /dashboards | Creates a new dashboard
*DashboardsApi* | [**PostDashboardWidget**](docs/DashboardsApi.md#postdashboardwidget) | **Post** /dashboards/{dashboard-id}/widgets | Adds widget to dashboard
*DashboardsApi* | [**PutWidgetType**](docs/DashboardsApi.md#putwidgettype) | **Put** /widget-types | Create or update a widget type
*HeapsApi* | [**PutHeap**](docs/HeapsApi.md#putheap) | **Put** /heaps | Create or update heap data


## Documentation For Models

 - [App](docs/App.md)
 - [Asset](docs/Asset.md)
 - [AssetType](docs/AssetType.md)
 - [AssetTypeAttribute](docs/AssetTypeAttribute.md)
 - [Dashboard](docs/Dashboard.md)
 - [Heap](docs/Heap.md)
 - [HeapSubtype](docs/HeapSubtype.md)
 - [Patch](docs/Patch.md)
 - [Pipeline](docs/Pipeline.md)
 - [Translation](docs/Translation.md)
 - [Widget](docs/Widget.md)
 - [WidgetData](docs/WidgetData.md)
 - [WidgetType](docs/WidgetType.md)
 - [WidgetTypeElement](docs/WidgetTypeElement.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author




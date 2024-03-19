# Go API client for api

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.6.8
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://eliona.io](https://eliona.io)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import api "github.com/eliona-smart-building-assistant/go-eliona-api-client/v2"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `api.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), api.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `api.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), api.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `api.ContextOperationServerIndices` and `api.ContextOperationServerVariables` context maps.

```go
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

All URIs are relative to *https://name.eliona.io/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AgentsAPI* | [**GetAgentByClassAndId**](docs/AgentsAPI.md#getagentbyclassandid) | **Get** /agents/{agent-class}/{agent-id} | Information about an agent
*AgentsAPI* | [**GetAgentDeviceById**](docs/AgentsAPI.md#getagentdevicebyid) | **Get** /agent-devices/{agent-class}/{agent-device-id} | Information about agent device
*AgentsAPI* | [**GetAgentDeviceMappingById**](docs/AgentsAPI.md#getagentdevicemappingbyid) | **Get** /agent-device-mappings/{agent-class}/{agent-device-mapping-id} | Information about agent device mapping
*AgentsAPI* | [**GetAgentDeviceMappingsByDeviceId**](docs/AgentsAPI.md#getagentdevicemappingsbydeviceid) | **Get** /agent-devices/{agent-class}/{agent-device-id}/mappings | Information about agent device mappings
*AgentsAPI* | [**GetAgentDevicesByAgentId**](docs/AgentsAPI.md#getagentdevicesbyagentid) | **Get** /agents/{agent-class}/{agent-id}/devices | Information about agent devices
*AgentsAPI* | [**GetAgents**](docs/AgentsAPI.md#getagents) | **Get** /agents | Information about agents
*AgentsAPI* | [**GetAgentsByClass**](docs/AgentsAPI.md#getagentsbyclass) | **Get** /agents/{agent-class} | Information about agents for a specific class
*AgentsAPI* | [**PostAgentByClass**](docs/AgentsAPI.md#postagentbyclass) | **Post** /agents/{agent-class} | Create an agent
*AgentsAPI* | [**PostAgentDeviceByAgentId**](docs/AgentsAPI.md#postagentdevicebyagentid) | **Post** /agents/{agent-class}/{agent-id}/devices | Create an agent device
*AgentsAPI* | [**PostAgentDeviceMappingByDeviceId**](docs/AgentsAPI.md#postagentdevicemappingbydeviceid) | **Post** /agent-devices/{agent-class}/{agent-device-id}/mappings | Create an agent device mapping
*AgentsAPI* | [**PutAgentByClass**](docs/AgentsAPI.md#putagentbyclass) | **Put** /agents/{agent-class} | Create or update an agent
*AgentsAPI* | [**PutAgentByClassAndId**](docs/AgentsAPI.md#putagentbyclassandid) | **Put** /agents/{agent-class}/{agent-id} | Update an agent
*AgentsAPI* | [**PutAgentDeviceByAgentId**](docs/AgentsAPI.md#putagentdevicebyagentid) | **Put** /agents/{agent-class}/{agent-id}/devices | Create or update an agent device
*AgentsAPI* | [**PutAgentDeviceById**](docs/AgentsAPI.md#putagentdevicebyid) | **Put** /agent-devices/{agent-class}/{agent-device-id} | Update an agent device
*AgentsAPI* | [**PutAgentDeviceMappingByDeviceId**](docs/AgentsAPI.md#putagentdevicemappingbydeviceid) | **Put** /agent-devices/{agent-class}/{agent-device-id}/mappings | Create or update an agent device mapping
*AgentsAPI* | [**PutAgentDeviceMappingById**](docs/AgentsAPI.md#putagentdevicemappingbyid) | **Put** /agent-device-mappings/{agent-class}/{agent-device-mapping-id} | Update an agent device mapping
*AggregationsAPI* | [**DeleteAggregationById**](docs/AggregationsAPI.md#deleteaggregationbyid) | **Delete** /aggregations/{aggregation-id} | Delete an aggregation
*AggregationsAPI* | [**GetAggregationById**](docs/AggregationsAPI.md#getaggregationbyid) | **Get** /aggregations/{aggregation-id} | Information about an aggregation
*AggregationsAPI* | [**GetAggregations**](docs/AggregationsAPI.md#getaggregations) | **Get** /aggregations | Information about aggregations
*AggregationsAPI* | [**PostAggregation**](docs/AggregationsAPI.md#postaggregation) | **Post** /aggregations | Creates an aggregation
*AggregationsAPI* | [**PutAggregation**](docs/AggregationsAPI.md#putaggregation) | **Put** /aggregations | Creates or updates an aggregation
*AggregationsAPI* | [**PutAggregationById**](docs/AggregationsAPI.md#putaggregationbyid) | **Put** /aggregations/{aggregation-id} | Updates an aggregation
*AlarmRulesAPI* | [**DeleteAlarmRuleById**](docs/AlarmRulesAPI.md#deletealarmrulebyid) | **Delete** /alarm-rules/{alarm-rule-id} | Delete an alarm rule
*AlarmRulesAPI* | [**GetAlarmRuleById**](docs/AlarmRulesAPI.md#getalarmrulebyid) | **Get** /alarm-rules/{alarm-rule-id} | Information about an alarm rule
*AlarmRulesAPI* | [**GetAlarmRules**](docs/AlarmRulesAPI.md#getalarmrules) | **Get** /alarm-rules | Information about alarm rules
*AlarmRulesAPI* | [**PostAlarmRule**](docs/AlarmRulesAPI.md#postalarmrule) | **Post** /alarm-rules | Create an alarm rule
*AlarmRulesAPI* | [**PutAlarmRule**](docs/AlarmRulesAPI.md#putalarmrule) | **Put** /alarm-rules | Create or update an alarm rule
*AlarmRulesAPI* | [**PutAlarmRuleById**](docs/AlarmRulesAPI.md#putalarmrulebyid) | **Put** /alarm-rules/{alarm-rule-id} | Update an alarm rule
*AlarmsAPI* | [**GetAlarmById**](docs/AlarmsAPI.md#getalarmbyid) | **Get** /alarms/{alarm-rule-id} | Information about alarm
*AlarmsAPI* | [**GetAlarmHistoryById**](docs/AlarmsAPI.md#getalarmhistorybyid) | **Get** /alarms-history/{alarm-rule-id} | Information about alarm history
*AlarmsAPI* | [**GetAlarms**](docs/AlarmsAPI.md#getalarms) | **Get** /alarms | Information about alarms
*AlarmsAPI* | [**GetAlarmsHistory**](docs/AlarmsAPI.md#getalarmshistory) | **Get** /alarms-history | Information about alarms history
*AlarmsAPI* | [**GetHighestAlarms**](docs/AlarmsAPI.md#gethighestalarms) | **Get** /alarms-highest | Information about most prioritized alarms
*AlarmsAPI* | [**ListenAlarm**](docs/AlarmsAPI.md#listenalarm) | **Get** /alarm-listener | WebSocket connection for alarm changes
*AlarmsAPI* | [**PatchAlarmById**](docs/AlarmsAPI.md#patchalarmbyid) | **Patch** /alarms/{alarm-rule-id} | Update alarm
*AppsAPI* | [**GetAppByName**](docs/AppsAPI.md#getappbyname) | **Get** /apps/{app-name} | Information about an app
*AppsAPI* | [**GetPatchByName**](docs/AppsAPI.md#getpatchbyname) | **Get** /apps/{app-name}/patches/{patch-name} | Information about a patch for an app
*AppsAPI* | [**PatchAppByName**](docs/AppsAPI.md#patchappbyname) | **Patch** /apps/{app-name} | Update an app
*AppsAPI* | [**PatchPatchByName**](docs/AppsAPI.md#patchpatchbyname) | **Patch** /apps/{app-name}/patches/{patch-name} | Updates a patch
*AssetTypesAPI* | [**DeleteAssetTypeByName**](docs/AssetTypesAPI.md#deleteassettypebyname) | **Delete** /asset-types/{asset-type-name} | Delete an asset type
*AssetTypesAPI* | [**GetAssetTypeByName**](docs/AssetTypesAPI.md#getassettypebyname) | **Get** /asset-types/{asset-type-name} | Information about an asset type
*AssetTypesAPI* | [**GetAssetTypes**](docs/AssetTypesAPI.md#getassettypes) | **Get** /asset-types | List of asset types
*AssetTypesAPI* | [**PostAssetType**](docs/AssetTypesAPI.md#postassettype) | **Post** /asset-types | Create an asset type
*AssetTypesAPI* | [**PostAssetTypeAttribute**](docs/AssetTypesAPI.md#postassettypeattribute) | **Post** /asset-types/{asset-type-name}/attributes | Create asset type attribute
*AssetTypesAPI* | [**PutAssetType**](docs/AssetTypesAPI.md#putassettype) | **Put** /asset-types | Create or update an asset type
*AssetTypesAPI* | [**PutAssetTypeAttribute**](docs/AssetTypesAPI.md#putassettypeattribute) | **Put** /asset-types/{asset-type-name}/attributes | Create or update an asset type attribute
*AssetTypesAPI* | [**PutAssetTypeByName**](docs/AssetTypesAPI.md#putassettypebyname) | **Put** /asset-types/{asset-type-name} | Update an asset type
*AssetsAPI* | [**DeleteAssetById**](docs/AssetsAPI.md#deleteassetbyid) | **Delete** /assets/{asset-id} | Delete an asset
*AssetsAPI* | [**DeleteBulkAssets**](docs/AssetsAPI.md#deletebulkassets) | **Delete** /assets-bulk | Delete a list of assets
*AssetsAPI* | [**DryRunDeleteBulkAssets**](docs/AssetsAPI.md#dryrundeletebulkassets) | **Delete** /assets-bulk/dry-run | Dry-run for deleting a list of assets
*AssetsAPI* | [**DryRunPostBulkAssets**](docs/AssetsAPI.md#dryrunpostbulkassets) | **Post** /assets-bulk/dry-run | Dry-run for creating a list of assets
*AssetsAPI* | [**DryRunPutBulkAssets**](docs/AssetsAPI.md#dryrunputbulkassets) | **Put** /assets-bulk/dry-run | Dry-run for creating or updating a list of assets
*AssetsAPI* | [**GetAssetById**](docs/AssetsAPI.md#getassetbyid) | **Get** /assets/{asset-id} | Information about an asset
*AssetsAPI* | [**GetAssets**](docs/AssetsAPI.md#getassets) | **Get** /assets | Information about assets
*AssetsAPI* | [**GetAttributeDisplay**](docs/AssetsAPI.md#getattributedisplay) | **Get** /attribute-display | How attributes are displayed
*AssetsAPI* | [**ListenAsset**](docs/AssetsAPI.md#listenasset) | **Get** /asset-listener | WebSocket connection for asset changes
*AssetsAPI* | [**PostAsset**](docs/AssetsAPI.md#postasset) | **Post** /assets | Create an asset
*AssetsAPI* | [**PostBulkAssets**](docs/AssetsAPI.md#postbulkassets) | **Post** /assets-bulk | Create a list of assets
*AssetsAPI* | [**PutAsset**](docs/AssetsAPI.md#putasset) | **Put** /assets | Create or update an asset
*AssetsAPI* | [**PutAssetById**](docs/AssetsAPI.md#putassetbyid) | **Put** /assets/{asset-id} | Update an asset
*AssetsAPI* | [**PutAttributeDisplay**](docs/AssetsAPI.md#putattributedisplay) | **Put** /attribute-display | Create or update how attributes are displayed
*AssetsAPI* | [**PutBulkAssets**](docs/AssetsAPI.md#putbulkassets) | **Put** /assets-bulk | Create or update a list of assets
*CommunicationAPI* | [**GetMessageReceiptById**](docs/CommunicationAPI.md#getmessagereceiptbyid) | **Get** /message-receipts/{message-id} | Information about a message
*CommunicationAPI* | [**PostMail**](docs/CommunicationAPI.md#postmail) | **Post** /send-mail | Send e-mail
*CommunicationAPI* | [**PostNotification**](docs/CommunicationAPI.md#postnotification) | **Post** /send-notification | Send notification
*DashboardsAPI* | [**GetDashboardById**](docs/DashboardsAPI.md#getdashboardbyid) | **Get** /dashboards/{dashboard-id} | Information about a dashboard
*DashboardsAPI* | [**GetDashboards**](docs/DashboardsAPI.md#getdashboards) | **Get** /dashboards | Information about dashboards
*DashboardsAPI* | [**PostDashboard**](docs/DashboardsAPI.md#postdashboard) | **Post** /dashboards | Creates a new dashboard
*DataAPI* | [**GetData**](docs/DataAPI.md#getdata) | **Get** /data | Gets all data
*DataAPI* | [**GetDataAggregated**](docs/DataAPI.md#getdataaggregated) | **Get** /data-aggregated | Get aggregated data
*DataAPI* | [**GetDataTrends**](docs/DataAPI.md#getdatatrends) | **Get** /data-trends | Get trend of historical data
*DataAPI* | [**ListenData**](docs/DataAPI.md#listendata) | **Get** /data-listener | WebSocket connection for asset data changes
*DataAPI* | [**PutBulkData**](docs/DataAPI.md#putbulkdata) | **Put** /data-bulk | Create or update multiple asset data
*DataAPI* | [**PutData**](docs/DataAPI.md#putdata) | **Put** /data | Create or update asset data
*NodesAPI* | [**GetNodeByIdent**](docs/NodesAPI.md#getnodebyident) | **Get** /nodes/{node-ident} | Information about a node
*NodesAPI* | [**GetNodes**](docs/NodesAPI.md#getnodes) | **Get** /nodes | Information about nodes
*NodesAPI* | [**PostNode**](docs/NodesAPI.md#postnode) | **Post** /nodes | Create a node
*NodesAPI* | [**PutNode**](docs/NodesAPI.md#putnode) | **Put** /nodes | Create or update a node
*NodesAPI* | [**PutNodeByIdent**](docs/NodesAPI.md#putnodebyident) | **Put** /nodes/{node-ident} | Update a node
*ProjectsAPI* | [**GetProjectById**](docs/ProjectsAPI.md#getprojectbyid) | **Get** /projects/{project-id} | Information about a project
*ProjectsAPI* | [**GetProjects**](docs/ProjectsAPI.md#getprojects) | **Get** /projects | Information about projects
*ProjectsAPI* | [**PutProject**](docs/ProjectsAPI.md#putproject) | **Put** /projects | Create or update a project
*QRCodesAPI* | [**GetQrCodeByAssetId**](docs/QRCodesAPI.md#getqrcodebyassetid) | **Get** /qr-codes/assets/{asset-id} | QR code for assets
*TagsAPI* | [**GetTagByName**](docs/TagsAPI.md#gettagbyname) | **Get** /tags/{tag-name} | Information about a tag
*TagsAPI* | [**GetTags**](docs/TagsAPI.md#gettags) | **Get** /tags | Information about tags
*TagsAPI* | [**PutTag**](docs/TagsAPI.md#puttag) | **Put** /tags | Create or update a tag
*UsersAPI* | [**GetUserById**](docs/UsersAPI.md#getuserbyid) | **Get** /users/{user-id} | Information about an user
*UsersAPI* | [**GetUsers**](docs/UsersAPI.md#getusers) | **Get** /users | Information about users
*UsersAPI* | [**PutUser**](docs/UsersAPI.md#putuser) | **Put** /users | Create or update an user
*VersionAPI* | [**GetOpenAPI**](docs/VersionAPI.md#getopenapi) | **Get** /version/openapi.json | OpenAPI specification for this API version
*VersionAPI* | [**GetVersion**](docs/VersionAPI.md#getversion) | **Get** /version | Version of the API
*WidgetsAPI* | [**GetDashboardWidgets**](docs/WidgetsAPI.md#getdashboardwidgets) | **Get** /dashboards/{dashboard-id}/widgets | Information about widgets on dashboard
*WidgetsAPI* | [**PostDashboardWidget**](docs/WidgetsAPI.md#postdashboardwidget) | **Post** /dashboards/{dashboard-id}/widgets | Adds widget to dashboard
*WidgetsTypesAPI* | [**DeleteWidgetTypeByName**](docs/WidgetsTypesAPI.md#deletewidgettypebyname) | **Delete** /widget-types/{widget-type-name} | Delete a widget type
*WidgetsTypesAPI* | [**GetWidgetTypeByName**](docs/WidgetsTypesAPI.md#getwidgettypebyname) | **Get** /widget-types/{widget-type-name} | Information about a widget type
*WidgetsTypesAPI* | [**GetWidgetTypes**](docs/WidgetsTypesAPI.md#getwidgettypes) | **Get** /widget-types | List of widget types
*WidgetsTypesAPI* | [**PostWidgetType**](docs/WidgetsTypesAPI.md#postwidgettype) | **Post** /widget-types | Create a widget type
*WidgetsTypesAPI* | [**PutWidgetType**](docs/WidgetsTypesAPI.md#putwidgettype) | **Put** /widget-types | Create or update a widget type
*WidgetsTypesAPI* | [**PutWidgetTypeByName**](docs/WidgetsTypesAPI.md#putwidgettypebyname) | **Put** /widget-types/{widget-type-name} | Update a widget type


## Documentation For Models

 - [Agent](docs/Agent.md)
 - [AgentClass](docs/AgentClass.md)
 - [AgentDevice](docs/AgentDevice.md)
 - [AgentDeviceGeneral](docs/AgentDeviceGeneral.md)
 - [AgentDeviceMapping](docs/AgentDeviceMapping.md)
 - [AgentDeviceMappingGeneral](docs/AgentDeviceMappingGeneral.md)
 - [Aggregation](docs/Aggregation.md)
 - [Alarm](docs/Alarm.md)
 - [AlarmListen](docs/AlarmListen.md)
 - [AlarmPriority](docs/AlarmPriority.md)
 - [AlarmRule](docs/AlarmRule.md)
 - [App](docs/App.md)
 - [Asset](docs/Asset.md)
 - [AssetDryRun](docs/AssetDryRun.md)
 - [AssetIdentifyBy](docs/AssetIdentifyBy.md)
 - [AssetListen](docs/AssetListen.md)
 - [AssetType](docs/AssetType.md)
 - [AssetTypeAttribute](docs/AssetTypeAttribute.md)
 - [Attachment](docs/Attachment.md)
 - [AttributeDisplay](docs/AttributeDisplay.md)
 - [Dashboard](docs/Dashboard.md)
 - [Data](docs/Data.md)
 - [DataAggregated](docs/DataAggregated.md)
 - [DataListen](docs/DataListen.md)
 - [DataSubtype](docs/DataSubtype.md)
 - [DryRunGeneral](docs/DryRunGeneral.md)
 - [IosysAgentDevice](docs/IosysAgentDevice.md)
 - [IosysAgentDeviceMapping](docs/IosysAgentDeviceMapping.md)
 - [MbusAgentDevice](docs/MbusAgentDevice.md)
 - [MbusAgentDeviceMapping](docs/MbusAgentDeviceMapping.md)
 - [Message](docs/Message.md)
 - [MessageReceipt](docs/MessageReceipt.md)
 - [Node](docs/Node.md)
 - [Notification](docs/Notification.md)
 - [Patch](docs/Patch.md)
 - [Project](docs/Project.md)
 - [Tag](docs/Tag.md)
 - [Translation](docs/Translation.md)
 - [User](docs/User.md)
 - [Widget](docs/Widget.md)
 - [WidgetData](docs/WidgetData.md)
 - [WidgetType](docs/WidgetType.md)
 - [WidgetTypeElement](docs/WidgetTypeElement.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### ApiKeyAuth

- **Type**: API key
- **API key parameter name**: X-API-Key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-API-Key and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		api.ContextAPIKeys,
		map[string]api.APIKey{
			"X-API-Key": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### BearerAuth

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), api.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


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

hello@eliona.io


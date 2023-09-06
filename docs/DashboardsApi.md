# \DashboardsAPI

All URIs are relative to *https://api.eliona.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDashboardById**](DashboardsAPI.md#GetDashboardById) | **Get** /dashboards/{dashboard-id} | Information about a dashboard
[**GetDashboards**](DashboardsAPI.md#GetDashboards) | **Get** /dashboards | Information about dashboards
[**PostDashboard**](DashboardsAPI.md#PostDashboard) | **Post** /dashboards | Creates a new dashboard



## GetDashboardById

> Dashboard GetDashboardById(ctx, dashboardId).Expansions(expansions).Execute()

Information about a dashboard



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/eliona-smart-building-assistant/go-eliona-api-client/v2"
)

func main() {
    dashboardId := int32(4711) // int32 | The id of the dashboard
    expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardsAPI.GetDashboardById(context.Background(), dashboardId).Expansions(expansions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.GetDashboardById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboardById`: Dashboard
    fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.GetDashboardById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **int32** | The id of the dashboard | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboards

> []Dashboard GetDashboards(ctx).Expansions(expansions).Execute()

Information about dashboards



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/eliona-smart-building-assistant/go-eliona-api-client/v2"
)

func main() {
    expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardsAPI.GetDashboards(context.Background()).Expansions(expansions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.GetDashboards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboards`: []Dashboard
    fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.GetDashboards`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**[]Dashboard**](Dashboard.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDashboard

> Dashboard PostDashboard(ctx).Dashboard(dashboard).Expansions(expansions).Execute()

Creates a new dashboard



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/eliona-smart-building-assistant/go-eliona-api-client/v2"
)

func main() {
    dashboard := *openapiclient.NewDashboard("Weather info", "99", "42") // Dashboard | 
    expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardsAPI.PostDashboard(context.Background()).Dashboard(dashboard).Expansions(expansions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.PostDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostDashboard`: Dashboard
    fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.PostDashboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboard** | [**Dashboard**](Dashboard.md) |  | 
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


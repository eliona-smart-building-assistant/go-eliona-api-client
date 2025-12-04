# \WidgetsTypesAPI

All URIs are relative to *https://name.eliona.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWidgetTypeByName**](WidgetsTypesAPI.md#DeleteWidgetTypeByName) | **Delete** /widget-types/{widget-type-name} | Delete a widget type
[**GetWidgetTypeByName**](WidgetsTypesAPI.md#GetWidgetTypeByName) | **Get** /widget-types/{widget-type-name} | Information about a widget type
[**GetWidgetTypes**](WidgetsTypesAPI.md#GetWidgetTypes) | **Get** /widget-types | List of widget types
[**PostWidgetType**](WidgetsTypesAPI.md#PostWidgetType) | **Post** /widget-types | Create a widget type
[**PutWidgetType**](WidgetsTypesAPI.md#PutWidgetType) | **Put** /widget-types | Create or update a widget type



## DeleteWidgetTypeByName

> DeleteWidgetTypeByName(ctx, widgetTypeName).Execute()

Delete a widget type



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eliona-smart-building-assistant/go-eliona-api-client/v3"
)

func main() {
	widgetTypeName := "weather" // string | The name of the widget type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WidgetsTypesAPI.DeleteWidgetTypeByName(context.Background(), widgetTypeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WidgetsTypesAPI.DeleteWidgetTypeByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**widgetTypeName** | **string** | The name of the widget type | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWidgetTypeByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWidgetTypeByName

> WidgetType GetWidgetTypeByName(ctx, widgetTypeName).Expansions(expansions).Execute()

Information about a widget type



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eliona-smart-building-assistant/go-eliona-api-client/v3"
)

func main() {
	widgetTypeName := "weather" // string | The name of the widget type
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WidgetsTypesAPI.GetWidgetTypeByName(context.Background(), widgetTypeName).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WidgetsTypesAPI.GetWidgetTypeByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWidgetTypeByName`: WidgetType
	fmt.Fprintf(os.Stdout, "Response from `WidgetsTypesAPI.GetWidgetTypeByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**widgetTypeName** | **string** | The name of the widget type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWidgetTypeByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**WidgetType**](WidgetType.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWidgetTypes

> []WidgetType GetWidgetTypes(ctx).Offset(offset).Size(size).Expansions(expansions).Execute()

List of widget types



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eliona-smart-building-assistant/go-eliona-api-client/v3"
)

func main() {
	offset := int64(3) // int64 | Specifies the starting point for pagination by indicating the number of items to skip.  (optional)
	size := int64(10) // int64 | Specifies the number of items per page for pagination.  (optional)
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WidgetsTypesAPI.GetWidgetTypes(context.Background()).Offset(offset).Size(size).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WidgetsTypesAPI.GetWidgetTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWidgetTypes`: []WidgetType
	fmt.Fprintf(os.Stdout, "Response from `WidgetsTypesAPI.GetWidgetTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWidgetTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int64** | Specifies the starting point for pagination by indicating the number of items to skip.  | 
 **size** | **int64** | Specifies the number of items per page for pagination.  | 
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**[]WidgetType**](WidgetType.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWidgetType

> WidgetType PostWidgetType(ctx).WidgetType(widgetType).Expansions(expansions).Execute()

Create a widget type



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eliona-smart-building-assistant/go-eliona-api-client/v3"
)

func main() {
	widgetType := *openapiclient.NewWidgetType("weather", "TODO", []openapiclient.WidgetTypeElement{*openapiclient.NewWidgetTypeElement("weather")}) // WidgetType | 
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WidgetsTypesAPI.PostWidgetType(context.Background()).WidgetType(widgetType).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WidgetsTypesAPI.PostWidgetType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWidgetType`: WidgetType
	fmt.Fprintf(os.Stdout, "Response from `WidgetsTypesAPI.PostWidgetType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostWidgetTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **widgetType** | [**WidgetType**](WidgetType.md) |  | 
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**WidgetType**](WidgetType.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutWidgetType

> WidgetType PutWidgetType(ctx).WidgetType(widgetType).Expansions(expansions).Execute()

Create or update a widget type



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eliona-smart-building-assistant/go-eliona-api-client/v3"
)

func main() {
	widgetType := *openapiclient.NewWidgetType("weather", "TODO", []openapiclient.WidgetTypeElement{*openapiclient.NewWidgetTypeElement("weather")}) // WidgetType | 
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WidgetsTypesAPI.PutWidgetType(context.Background()).WidgetType(widgetType).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WidgetsTypesAPI.PutWidgetType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutWidgetType`: WidgetType
	fmt.Fprintf(os.Stdout, "Response from `WidgetsTypesAPI.PutWidgetType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutWidgetTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **widgetType** | [**WidgetType**](WidgetType.md) |  | 
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**WidgetType**](WidgetType.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


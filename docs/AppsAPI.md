# \AppsAPI

All URIs are relative to *https://name.eliona.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAppByName**](AppsAPI.md#GetAppByName) | **Get** /apps/{app-name} | Information about an app
[**GetPatchByName**](AppsAPI.md#GetPatchByName) | **Get** /apps/{app-name}/patches/{patch-name} | Information about a patch for an app
[**PatchAppByName**](AppsAPI.md#PatchAppByName) | **Patch** /apps/{app-name} | Update an app
[**PatchPatchByName**](AppsAPI.md#PatchPatchByName) | **Patch** /apps/{app-name}/patches/{patch-name} | Updates a patch



## GetAppByName

> App GetAppByName(ctx, appName).Execute()

Information about an app



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
	appName := "weather" // string | The name of the app

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.GetAppByName(context.Background(), appName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetAppByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppByName`: App
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetAppByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | The name of the app | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**App**](App.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatchByName

> Patch GetPatchByName(ctx, appName, patchName).Execute()

Information about a patch for an app



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
	appName := "weather" // string | The name of the app
	patchName := "2.0.0" // string | The name of the patch

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.GetPatchByName(context.Background(), appName, patchName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetPatchByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatchByName`: Patch
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetPatchByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | The name of the app | 
**patchName** | **string** | The name of the patch | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPatchByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Patch**](Patch.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAppByName

> PatchAppByName(ctx, appName).Registered(registered).Version(version).Execute()

Update an app



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
	appName := "weather" // string | The name of the app
	registered := true // bool | Marks that the app is now initialized and installed. Further request to get app information returns { \"registered\": true } (optional)
	version := "version_example" // string | Sets the current version that is running and initialized (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppsAPI.PatchAppByName(context.Background(), appName).Registered(registered).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.PatchAppByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | The name of the app | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAppByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registered** | **bool** | Marks that the app is now initialized and installed. Further request to get app information returns { \&quot;registered\&quot;: true } | 
 **version** | **string** | Sets the current version that is running and initialized | 

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


## PatchPatchByName

> PatchPatchByName(ctx, appName, patchName).Apply(apply).Execute()

Updates a patch



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
	appName := "weather" // string | The name of the app
	patchName := "2.0.0" // string | The name of the patch
	apply := true // bool | Marks that the patch is now applied. Further request to get patch information returns { \"applied\": true } (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppsAPI.PatchPatchByName(context.Background(), appName, patchName).Apply(apply).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.PatchPatchByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | The name of the app | 
**patchName** | **string** | The name of the patch | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPatchByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apply** | **bool** | Marks that the patch is now applied. Further request to get patch information returns { \&quot;applied\&quot;: true } | 

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


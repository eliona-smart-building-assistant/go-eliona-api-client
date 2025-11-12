# \AggregationsAPI

All URIs are relative to *https://name.eliona.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAggregationById**](AggregationsAPI.md#DeleteAggregationById) | **Delete** /aggregations/{aggregation-id} | Delete an aggregation
[**GetAggregationById**](AggregationsAPI.md#GetAggregationById) | **Get** /aggregations/{aggregation-id} | Information about an aggregation
[**GetAggregations**](AggregationsAPI.md#GetAggregations) | **Get** /aggregations | Information about aggregations
[**PostAggregation**](AggregationsAPI.md#PostAggregation) | **Post** /aggregations | Creates an aggregation
[**PutAggregation**](AggregationsAPI.md#PutAggregation) | **Put** /aggregations | Creates or updates an aggregation
[**PutAggregationById**](AggregationsAPI.md#PutAggregationById) | **Put** /aggregations/{aggregation-id} | Updates an aggregation



## DeleteAggregationById

> DeleteAggregationById(ctx, aggregationId).Execute()

Delete an aggregation



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
	aggregationId := int32(4711) // int32 | The id of the aggregation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AggregationsAPI.DeleteAggregationById(context.Background(), aggregationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AggregationsAPI.DeleteAggregationById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aggregationId** | **int32** | The id of the aggregation | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAggregationByIdRequest struct via the builder pattern


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


## GetAggregationById

> Aggregation GetAggregationById(ctx, aggregationId).Execute()

Information about an aggregation



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
	aggregationId := int32(4711) // int32 | The id of the aggregation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AggregationsAPI.GetAggregationById(context.Background(), aggregationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AggregationsAPI.GetAggregationById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAggregationById`: Aggregation
	fmt.Fprintf(os.Stdout, "Response from `AggregationsAPI.GetAggregationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aggregationId** | **int32** | The id of the aggregation | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAggregationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Aggregation**](Aggregation.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAggregations

> []Aggregation GetAggregations(ctx).Execute()

Information about aggregations



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AggregationsAPI.GetAggregations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AggregationsAPI.GetAggregations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAggregations`: []Aggregation
	fmt.Fprintf(os.Stdout, "Response from `AggregationsAPI.GetAggregations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAggregationsRequest struct via the builder pattern


### Return type

[**[]Aggregation**](Aggregation.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAggregation

> Aggregation PostAggregation(ctx).Aggregation(aggregation).Execute()

Creates an aggregation



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
	aggregation := *openapiclient.NewAggregation(int32(4711), openapiclient.DataSubtype("input"), "avg") // Aggregation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AggregationsAPI.PostAggregation(context.Background()).Aggregation(aggregation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AggregationsAPI.PostAggregation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAggregation`: Aggregation
	fmt.Fprintf(os.Stdout, "Response from `AggregationsAPI.PostAggregation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAggregationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aggregation** | [**Aggregation**](Aggregation.md) |  | 

### Return type

[**Aggregation**](Aggregation.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAggregation

> Aggregation PutAggregation(ctx).Aggregation(aggregation).Execute()

Creates or updates an aggregation



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
	aggregation := *openapiclient.NewAggregation(int32(4711), openapiclient.DataSubtype("input"), "avg") // Aggregation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AggregationsAPI.PutAggregation(context.Background()).Aggregation(aggregation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AggregationsAPI.PutAggregation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAggregation`: Aggregation
	fmt.Fprintf(os.Stdout, "Response from `AggregationsAPI.PutAggregation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutAggregationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aggregation** | [**Aggregation**](Aggregation.md) |  | 

### Return type

[**Aggregation**](Aggregation.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAggregationById

> Aggregation PutAggregationById(ctx, aggregationId).Aggregation(aggregation).Execute()

Updates an aggregation



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
	aggregationId := int32(4711) // int32 | The id of the aggregation
	aggregation := *openapiclient.NewAggregation(int32(4711), openapiclient.DataSubtype("input"), "avg") // Aggregation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AggregationsAPI.PutAggregationById(context.Background(), aggregationId).Aggregation(aggregation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AggregationsAPI.PutAggregationById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAggregationById`: Aggregation
	fmt.Fprintf(os.Stdout, "Response from `AggregationsAPI.PutAggregationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aggregationId** | **int32** | The id of the aggregation | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAggregationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aggregation** | [**Aggregation**](Aggregation.md) |  | 

### Return type

[**Aggregation**](Aggregation.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \DataAPI

All URIs are relative to *https://name.eliona.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetData**](DataAPI.md#GetData) | **Get** /data | Gets all data
[**GetDataAggregated**](DataAPI.md#GetDataAggregated) | **Get** /data-aggregated | Get aggregated data
[**GetDataTrendAggregatedById**](DataAPI.md#GetDataTrendAggregatedById) | **Get** /data-trend-aggregated/{asset-id} | Get aggregated trend of historical data
[**GetDataTrendById**](DataAPI.md#GetDataTrendById) | **Get** /data-trend/{asset-id} | Get trend of historical data
[**GetDataTrends**](DataAPI.md#GetDataTrends) | **Get** /data-trends | Get trend of historical data
[**ListenData**](DataAPI.md#ListenData) | **Get** /data-listener | WebSocket connection for asset data changes
[**PutBulkData**](DataAPI.md#PutBulkData) | **Put** /data-bulk | Create or update multiple asset data
[**PutData**](DataAPI.md#PutData) | **Put** /data | Create or update asset data
[**PutDataTrend**](DataAPI.md#PutDataTrend) | **Put** /data-trend | Create or update historical data



## GetData

> []Data GetData(ctx).AssetId(assetId).ParentAssetId(parentAssetId).DataSubtype(dataSubtype).AssetTypeName(assetTypeName).Execute()

Gets all data



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
	assetId := int32(4711) // int32 | Filter for a specific asset id (optional)
	parentAssetId := int32(4711) // int32 | Filter for a specific parent asset id (optional)
	dataSubtype := "input" // string | Filter for a specific type of asset data (optional)
	assetTypeName := "weather_location" // string | Filter the name of the asset type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAPI.GetData(context.Background()).AssetId(assetId).ParentAssetId(parentAssetId).DataSubtype(dataSubtype).AssetTypeName(assetTypeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAPI.GetData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetData`: []Data
	fmt.Fprintf(os.Stdout, "Response from `DataAPI.GetData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetId** | **int32** | Filter for a specific asset id | 
 **parentAssetId** | **int32** | Filter for a specific parent asset id | 
 **dataSubtype** | **string** | Filter for a specific type of asset data | 
 **assetTypeName** | **string** | Filter the name of the asset type | 

### Return type

[**[]Data**](Data.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataAggregated

> []DataAggregated GetDataAggregated(ctx).FromDate(fromDate).ToDate(toDate).AssetId(assetId).DataSubtype(dataSubtype).AssetTypeName(assetTypeName).AggregationId(aggregationId).AggregationRaster(aggregationRaster).Execute()

Get aggregated data



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
	fromDate := "2020-01-01T09:00:00.000Z" // string | Filter by lower date time (RFC3339) limit inclusive (optional)
	toDate := "2021-12-31T23:00:00.000Z" // string | Filter by upper date time (RFC3339) limit exclusive (optional)
	assetId := int32(4711) // int32 | Filter for a specific asset id (optional)
	dataSubtype := "input" // string | Filter for a specific type of asset data (optional)
	assetTypeName := "weather_location" // string | Filter the name of the asset type (optional)
	aggregationId := int32(0815) // int32 | Filter for a specific aggregation id (optional)
	aggregationRaster := "DAY" // string | Aggregation calculation interval (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAPI.GetDataAggregated(context.Background()).FromDate(fromDate).ToDate(toDate).AssetId(assetId).DataSubtype(dataSubtype).AssetTypeName(assetTypeName).AggregationId(aggregationId).AggregationRaster(aggregationRaster).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAPI.GetDataAggregated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataAggregated`: []DataAggregated
	fmt.Fprintf(os.Stdout, "Response from `DataAPI.GetDataAggregated`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDataAggregatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromDate** | **string** | Filter by lower date time (RFC3339) limit inclusive | 
 **toDate** | **string** | Filter by upper date time (RFC3339) limit exclusive | 
 **assetId** | **int32** | Filter for a specific asset id | 
 **dataSubtype** | **string** | Filter for a specific type of asset data | 
 **assetTypeName** | **string** | Filter the name of the asset type | 
 **aggregationId** | **int32** | Filter for a specific aggregation id | 
 **aggregationRaster** | **string** | Aggregation calculation interval | 

### Return type

[**[]DataAggregated**](DataAggregated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataTrendAggregatedById

> []DataTrendAggregated GetDataTrendAggregatedById(ctx, assetId).DataSubtype(dataSubtype).AttributeName(attributeName).AggregationRaster(aggregationRaster).FromDate(fromDate).ToDate(toDate).SumMethod(sumMethod).Offset(offset).Size(size).Execute()

Get aggregated trend of historical data



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
	assetId := int32(4711) // int32 | The id of the asset
	dataSubtype := "input" // string | Type of asset data
	attributeName := "temperature" // string | Data attribute name
	aggregationRaster := "1d" // string | Aggregation calculation period
	fromDate := "2020-01-01T09:00:00.000Z" // string | Lower date time (RFC3339) limit inclusive
	toDate := "2021-12-31T23:00:00.000Z" // string | Upper date time (RFC3339) limit inclusive
	sumMethod := "sum" // string | Method for summarize data (optional)
	offset := int32(3) // int32 | Specifies the starting point for pagination by indicating the number of items to skip.  (optional)
	size := int32(10) // int32 | Specifies the number of items per page for pagination.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAPI.GetDataTrendAggregatedById(context.Background(), assetId).DataSubtype(dataSubtype).AttributeName(attributeName).AggregationRaster(aggregationRaster).FromDate(fromDate).ToDate(toDate).SumMethod(sumMethod).Offset(offset).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAPI.GetDataTrendAggregatedById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataTrendAggregatedById`: []DataTrendAggregated
	fmt.Fprintf(os.Stdout, "Response from `DataAPI.GetDataTrendAggregatedById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **int32** | The id of the asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataTrendAggregatedByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataSubtype** | **string** | Type of asset data | 
 **attributeName** | **string** | Data attribute name | 
 **aggregationRaster** | **string** | Aggregation calculation period | 
 **fromDate** | **string** | Lower date time (RFC3339) limit inclusive | 
 **toDate** | **string** | Upper date time (RFC3339) limit inclusive | 
 **sumMethod** | **string** | Method for summarize data | 
 **offset** | **int32** | Specifies the starting point for pagination by indicating the number of items to skip.  | 
 **size** | **int32** | Specifies the number of items per page for pagination.  | 

### Return type

[**[]DataTrendAggregated**](DataTrendAggregated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataTrendById

> []Data GetDataTrendById(ctx, assetId).DataSubtype(dataSubtype).AttributeName(attributeName).FromDate(fromDate).ToDate(toDate).Offset(offset).Size(size).Execute()

Get trend of historical data



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
	assetId := int32(4711) // int32 | The id of the asset
	dataSubtype := "input" // string | Type of asset data
	attributeName := "temperature" // string | Data attribute name
	fromDate := "2020-01-01T09:00:00.000Z" // string | Lower date time (RFC3339) limit inclusive
	toDate := "2021-12-31T23:00:00.000Z" // string | Upper date time (RFC3339) limit inclusive
	offset := int32(3) // int32 | Specifies the starting point for pagination by indicating the number of items to skip.  (optional)
	size := int32(10) // int32 | Specifies the number of items per page for pagination.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAPI.GetDataTrendById(context.Background(), assetId).DataSubtype(dataSubtype).AttributeName(attributeName).FromDate(fromDate).ToDate(toDate).Offset(offset).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAPI.GetDataTrendById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataTrendById`: []Data
	fmt.Fprintf(os.Stdout, "Response from `DataAPI.GetDataTrendById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **int32** | The id of the asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataTrendByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataSubtype** | **string** | Type of asset data | 
 **attributeName** | **string** | Data attribute name | 
 **fromDate** | **string** | Lower date time (RFC3339) limit inclusive | 
 **toDate** | **string** | Upper date time (RFC3339) limit inclusive | 
 **offset** | **int32** | Specifies the starting point for pagination by indicating the number of items to skip.  | 
 **size** | **int32** | Specifies the number of items per page for pagination.  | 

### Return type

[**[]Data**](Data.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataTrends

> []Data GetDataTrends(ctx).FromDate(fromDate).ToDate(toDate).AssetId(assetId).DataSubtype(dataSubtype).AssetTypeName(assetTypeName).Execute()

Get trend of historical data



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
	fromDate := "2020-01-01T09:00:00.000Z" // string | Filter by lower date time (RFC3339) limit inclusive (optional)
	toDate := "2021-12-31T23:00:00.000Z" // string | Filter by upper date time (RFC3339) limit exclusive (optional)
	assetId := int32(4711) // int32 | Filter for a specific asset id (optional)
	dataSubtype := "input" // string | Filter for a specific type of asset data (optional)
	assetTypeName := "weather_location" // string | Filter the name of the asset type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAPI.GetDataTrends(context.Background()).FromDate(fromDate).ToDate(toDate).AssetId(assetId).DataSubtype(dataSubtype).AssetTypeName(assetTypeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAPI.GetDataTrends``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataTrends`: []Data
	fmt.Fprintf(os.Stdout, "Response from `DataAPI.GetDataTrends`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDataTrendsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromDate** | **string** | Filter by lower date time (RFC3339) limit inclusive | 
 **toDate** | **string** | Filter by upper date time (RFC3339) limit exclusive | 
 **assetId** | **int32** | Filter for a specific asset id | 
 **dataSubtype** | **string** | Filter for a specific type of asset data | 
 **assetTypeName** | **string** | Filter the name of the asset type | 

### Return type

[**[]Data**](Data.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListenData

> DataListen ListenData(ctx).AssetId(assetId).DataSubtype(dataSubtype).Execute()

WebSocket connection for asset data changes



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
	assetId := int32(4711) // int32 | Filter for a specific asset id (optional)
	dataSubtype := "input" // string | Filter for a specific type of asset data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAPI.ListenData(context.Background()).AssetId(assetId).DataSubtype(dataSubtype).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAPI.ListenData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListenData`: DataListen
	fmt.Fprintf(os.Stdout, "Response from `DataAPI.ListenData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListenDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetId** | **int32** | Filter for a specific asset id | 
 **dataSubtype** | **string** | Filter for a specific type of asset data | 

### Return type

[**DataListen**](DataListen.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutBulkData

> PutBulkData(ctx).Data(data).DirectMode(directMode).Execute()

Create or update multiple asset data



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
	data := []openapiclient.Data{*openapiclient.NewData(int32(4711), openapiclient.DataSubtype("input"), map[string]interface{}({"temperature":42}))} // []Data | 
	directMode := "true" // string | Executes the operation directly without using other services. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataAPI.PutBulkData(context.Background()).Data(data).DirectMode(directMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAPI.PutBulkData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutBulkDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**[]Data**](Data.md) |  | 
 **directMode** | **string** | Executes the operation directly without using other services. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutData

> PutData(ctx).Data(data).DirectMode(directMode).Execute()

Create or update asset data



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
	data := *openapiclient.NewData(int32(4711), openapiclient.DataSubtype("input"), map[string]interface{}({"temperature":42})) // Data | 
	directMode := "true" // string | Executes the operation directly without using other services. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataAPI.PutData(context.Background()).Data(data).DirectMode(directMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAPI.PutData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**Data**](Data.md) |  | 
 **directMode** | **string** | Executes the operation directly without using other services. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDataTrend

> Data PutDataTrend(ctx).Data(data).Execute()

Create or update historical data



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
	data := *openapiclient.NewData(int32(4711), openapiclient.DataSubtype("input"), map[string]interface{}({"temperature":42})) // Data | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAPI.PutDataTrend(context.Background()).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAPI.PutDataTrend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutDataTrend`: Data
	fmt.Fprintf(os.Stdout, "Response from `DataAPI.PutDataTrend`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutDataTrendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**Data**](Data.md) |  | 

### Return type

[**Data**](Data.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


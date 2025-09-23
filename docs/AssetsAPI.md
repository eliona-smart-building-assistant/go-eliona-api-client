# \AssetsAPI

All URIs are relative to *https://name.eliona.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAssetById**](AssetsAPI.md#DeleteAssetById) | **Delete** /assets/{asset-id} | Delete an asset
[**DeleteBulkAssets**](AssetsAPI.md#DeleteBulkAssets) | **Delete** /assets-bulk | Delete a list of assets
[**DryRunDeleteBulkAssets**](AssetsAPI.md#DryRunDeleteBulkAssets) | **Delete** /assets-bulk/dry-run | Dry-run for deleting a list of assets
[**DryRunPostBulkAssets**](AssetsAPI.md#DryRunPostBulkAssets) | **Post** /assets-bulk/dry-run | Dry-run for creating a list of assets
[**DryRunPutBulkAssets**](AssetsAPI.md#DryRunPutBulkAssets) | **Put** /assets-bulk/dry-run | Dry-run for creating or updating a list of assets
[**GetAssetById**](AssetsAPI.md#GetAssetById) | **Get** /assets/{asset-id} | Information about an asset
[**GetAssets**](AssetsAPI.md#GetAssets) | **Get** /assets | Information about assets
[**GetAttributeDisplay**](AssetsAPI.md#GetAttributeDisplay) | **Get** /attribute-display | How attributes are displayed
[**ListenAsset**](AssetsAPI.md#ListenAsset) | **Get** /asset-listener | WebSocket connection for asset changes
[**PostAsset**](AssetsAPI.md#PostAsset) | **Post** /assets | Create an asset
[**PostBulkAssets**](AssetsAPI.md#PostBulkAssets) | **Post** /assets-bulk | Create a list of assets
[**PutAsset**](AssetsAPI.md#PutAsset) | **Put** /assets | Create or update an asset
[**PutAssetById**](AssetsAPI.md#PutAssetById) | **Put** /assets/{asset-id} | Update an asset
[**PutAttributeDisplay**](AssetsAPI.md#PutAttributeDisplay) | **Put** /attribute-display | Create or update how attributes are displayed
[**PutBulkAssets**](AssetsAPI.md#PutBulkAssets) | **Put** /assets-bulk | Create or update a list of assets



## DeleteAssetById

> DeleteAssetById(ctx, assetId).Execute()

Delete an asset



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
	assetId := int32(4711) // int32 | The id of the asset

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AssetsAPI.DeleteAssetById(context.Background(), assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.DeleteAssetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **int32** | The id of the asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAssetByIdRequest struct via the builder pattern


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


## DeleteBulkAssets

> DeleteBulkAssets(ctx).RequestBody(requestBody).IdentifyBy(identifyBy).Expansions(expansions).Execute()

Delete a list of assets



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
	requestBody := []string{"1234"} // []string | 
	identifyBy := "resourceId" // string | Serves the field name send in the request body as a unique identifier for the asset, essential for operations like updates or deletions. Please refer to the Asset schema definition for further information about this fields.  In cases where this parameter isn't defined, all field names are used in the order defined. So if there is no 'resourceId' present in the request body, the 'deviceId' is used and when there is also no deviceId present the 'id' field (assetId) is used.  (optional)
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AssetsAPI.DeleteBulkAssets(context.Background()).RequestBody(requestBody).IdentifyBy(identifyBy).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.DeleteBulkAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBulkAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 
 **identifyBy** | **string** | Serves the field name send in the request body as a unique identifier for the asset, essential for operations like updates or deletions. Please refer to the Asset schema definition for further information about this fields.  In cases where this parameter isn&#39;t defined, all field names are used in the order defined. So if there is no &#39;resourceId&#39; present in the request body, the &#39;deviceId&#39; is used and when there is also no deviceId present the &#39;id&#39; field (assetId) is used.  | 
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DryRunDeleteBulkAssets

> []AssetDryRun DryRunDeleteBulkAssets(ctx).RequestBody(requestBody).IdentifyBy(identifyBy).Expansions(expansions).Execute()

Dry-run for deleting a list of assets



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
	requestBody := []string{"1234"} // []string | 
	identifyBy := "resourceId" // string | Serves the field name send in the request body as a unique identifier for the asset, essential for operations like updates or deletions. Please refer to the Asset schema definition for further information about this fields.  In cases where this parameter isn't defined, all field names are used in the order defined. So if there is no 'resourceId' present in the request body, the 'deviceId' is used and when there is also no deviceId present the 'id' field (assetId) is used.  (optional)
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.DryRunDeleteBulkAssets(context.Background()).RequestBody(requestBody).IdentifyBy(identifyBy).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.DryRunDeleteBulkAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DryRunDeleteBulkAssets`: []AssetDryRun
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.DryRunDeleteBulkAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDryRunDeleteBulkAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 
 **identifyBy** | **string** | Serves the field name send in the request body as a unique identifier for the asset, essential for operations like updates or deletions. Please refer to the Asset schema definition for further information about this fields.  In cases where this parameter isn&#39;t defined, all field names are used in the order defined. So if there is no &#39;resourceId&#39; present in the request body, the &#39;deviceId&#39; is used and when there is also no deviceId present the &#39;id&#39; field (assetId) is used.  | 
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**[]AssetDryRun**](AssetDryRun.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DryRunPostBulkAssets

> []AssetDryRun DryRunPostBulkAssets(ctx).Asset(asset).IdentifyBy(identifyBy).Expansions(expansions).Execute()

Dry-run for creating a list of assets



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
	asset := []openapiclient.Asset{*openapiclient.NewAsset("zurich_swiss", "AssetType_example")} // []Asset | 
	identifyBy := "resourceId" // string | Serves the field name send in the request body as a unique identifier for the asset, essential for operations like updates or deletions. Please refer to the Asset schema definition for further information about this fields.  In cases where this parameter isn't defined, all field names are used in the order defined. So if there is no 'resourceId' present in the request body, the 'deviceId' is used and when there is also no deviceId present the 'id' field (assetId) is used.  (optional)
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.DryRunPostBulkAssets(context.Background()).Asset(asset).IdentifyBy(identifyBy).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.DryRunPostBulkAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DryRunPostBulkAssets`: []AssetDryRun
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.DryRunPostBulkAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDryRunPostBulkAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | [**[]Asset**](Asset.md) |  | 
 **identifyBy** | **string** | Serves the field name send in the request body as a unique identifier for the asset, essential for operations like updates or deletions. Please refer to the Asset schema definition for further information about this fields.  In cases where this parameter isn&#39;t defined, all field names are used in the order defined. So if there is no &#39;resourceId&#39; present in the request body, the &#39;deviceId&#39; is used and when there is also no deviceId present the &#39;id&#39; field (assetId) is used.  | 
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**[]AssetDryRun**](AssetDryRun.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DryRunPutBulkAssets

> []AssetDryRun DryRunPutBulkAssets(ctx).Asset(asset).IdentifyBy(identifyBy).Expansions(expansions).Execute()

Dry-run for creating or updating a list of assets



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
	asset := []openapiclient.Asset{*openapiclient.NewAsset("zurich_swiss", "AssetType_example")} // []Asset | 
	identifyBy := "resourceId" // string | Serves the field name send in the request body as a unique identifier for the asset, essential for operations like updates or deletions. Please refer to the Asset schema definition for further information about this fields.  In cases where this parameter isn't defined, all field names are used in the order defined. So if there is no 'resourceId' present in the request body, the 'deviceId' is used and when there is also no deviceId present the 'id' field (assetId) is used.  (optional)
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.DryRunPutBulkAssets(context.Background()).Asset(asset).IdentifyBy(identifyBy).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.DryRunPutBulkAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DryRunPutBulkAssets`: []AssetDryRun
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.DryRunPutBulkAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDryRunPutBulkAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | [**[]Asset**](Asset.md) |  | 
 **identifyBy** | **string** | Serves the field name send in the request body as a unique identifier for the asset, essential for operations like updates or deletions. Please refer to the Asset schema definition for further information about this fields.  In cases where this parameter isn&#39;t defined, all field names are used in the order defined. So if there is no &#39;resourceId&#39; present in the request body, the &#39;deviceId&#39; is used and when there is also no deviceId present the &#39;id&#39; field (assetId) is used.  | 
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**[]AssetDryRun**](AssetDryRun.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetById

> Asset GetAssetById(ctx, assetId).Expansions(expansions).Execute()

Information about an asset



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
	assetId := int32(4711) // int32 | The id of the asset
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.GetAssetById(context.Background(), assetId).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.GetAssetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetById`: Asset
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.GetAssetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **int32** | The id of the asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**Asset**](Asset.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssets

> []Asset GetAssets(ctx).AssetTypeName(assetTypeName).SiteId(siteId).ProjectId(projectId).Offset(offset).Size(size).Expansions(expansions).Execute()

Information about assets



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
	assetTypeName := "weather_location" // string | Filter the name of the asset type (optional)
	siteId := "69133767-7de1-4134-9130-5f5403e1ee98" // string | Filter for a specific site (optional)
	projectId := "79d7f1a0-0ec1-4211-8626-21f79491650a" // string | Filter for a specific project. The project ID is deprecated due to the removal of projects and is broadly replaced by sites. This filter now uses siteId, which is internally represented as a UUID. (optional)
	offset := int64(3) // int64 | Specifies the starting point for pagination by indicating the number of items to skip.  (optional)
	size := int64(10) // int64 | Specifies the number of items per page for pagination.  (optional)
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.GetAssets(context.Background()).AssetTypeName(assetTypeName).SiteId(siteId).ProjectId(projectId).Offset(offset).Size(size).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.GetAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssets`: []Asset
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.GetAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetTypeName** | **string** | Filter the name of the asset type | 
 **siteId** | **string** | Filter for a specific site | 
 **projectId** | **string** | Filter for a specific project. The project ID is deprecated due to the removal of projects and is broadly replaced by sites. This filter now uses siteId, which is internally represented as a UUID. | 
 **offset** | **int64** | Specifies the starting point for pagination by indicating the number of items to skip.  | 
 **size** | **int64** | Specifies the number of items per page for pagination.  | 
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**[]Asset**](Asset.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttributeDisplay

> []AttributeDisplay GetAttributeDisplay(ctx).Offset(offset).Size(size).Execute()

How attributes are displayed



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.GetAttributeDisplay(context.Background()).Offset(offset).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.GetAttributeDisplay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttributeDisplay`: []AttributeDisplay
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.GetAttributeDisplay`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAttributeDisplayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int64** | Specifies the starting point for pagination by indicating the number of items to skip.  | 
 **size** | **int64** | Specifies the number of items per page for pagination.  | 

### Return type

[**[]AttributeDisplay**](AttributeDisplay.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListenAsset

> AssetListen ListenAsset(ctx).Expansions(expansions).Execute()

WebSocket connection for asset changes



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
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.ListenAsset(context.Background()).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.ListenAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListenAsset`: AssetListen
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.ListenAsset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListenAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**AssetListen**](AssetListen.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAsset

> Asset PostAsset(ctx).Asset(asset).IdentifyBy(identifyBy).Expansions(expansions).Execute()

Create an asset



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
	asset := *openapiclient.NewAsset("zurich_swiss", "AssetType_example") // Asset | 
	identifyBy := "resourceId" // string | Serves the field name send in the request body as a unique identifier for the asset, essential for operations like updates or deletions. Please refer to the Asset schema definition for further information about this fields.  In cases where this parameter isn't defined, all field names are used in the order defined. So if there is no 'resourceId' present in the request body, the 'deviceId' is used and when there is also no deviceId present the 'id' field (assetId) is used.  (optional)
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.PostAsset(context.Background()).Asset(asset).IdentifyBy(identifyBy).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.PostAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAsset`: Asset
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.PostAsset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | [**Asset**](Asset.md) |  | 
 **identifyBy** | **string** | Serves the field name send in the request body as a unique identifier for the asset, essential for operations like updates or deletions. Please refer to the Asset schema definition for further information about this fields.  In cases where this parameter isn&#39;t defined, all field names are used in the order defined. So if there is no &#39;resourceId&#39; present in the request body, the &#39;deviceId&#39; is used and when there is also no deviceId present the &#39;id&#39; field (assetId) is used.  | 
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**Asset**](Asset.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBulkAssets

> []Asset PostBulkAssets(ctx).Asset(asset).IdentifyBy(identifyBy).Expansions(expansions).Execute()

Create a list of assets



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
	asset := []openapiclient.Asset{*openapiclient.NewAsset("zurich_swiss", "AssetType_example")} // []Asset | 
	identifyBy := "resourceId" // string | Serves the field name send in the request body as a unique identifier for the asset, essential for operations like updates or deletions. Please refer to the Asset schema definition for further information about this fields.  In cases where this parameter isn't defined, all field names are used in the order defined. So if there is no 'resourceId' present in the request body, the 'deviceId' is used and when there is also no deviceId present the 'id' field (assetId) is used.  (optional)
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.PostBulkAssets(context.Background()).Asset(asset).IdentifyBy(identifyBy).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.PostBulkAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBulkAssets`: []Asset
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.PostBulkAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBulkAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | [**[]Asset**](Asset.md) |  | 
 **identifyBy** | **string** | Serves the field name send in the request body as a unique identifier for the asset, essential for operations like updates or deletions. Please refer to the Asset schema definition for further information about this fields.  In cases where this parameter isn&#39;t defined, all field names are used in the order defined. So if there is no &#39;resourceId&#39; present in the request body, the &#39;deviceId&#39; is used and when there is also no deviceId present the &#39;id&#39; field (assetId) is used.  | 
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**[]Asset**](Asset.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAsset

> Asset PutAsset(ctx).Asset(asset).IdentifyBy(identifyBy).Expansions(expansions).Execute()

Create or update an asset



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
	asset := *openapiclient.NewAsset("zurich_swiss", "AssetType_example") // Asset | 
	identifyBy := "resourceId" // string | Serves the field name send in the request body as a unique identifier for the asset, essential for operations like updates or deletions. Please refer to the Asset schema definition for further information about this fields.  In cases where this parameter isn't defined, all field names are used in the order defined. So if there is no 'resourceId' present in the request body, the 'deviceId' is used and when there is also no deviceId present the 'id' field (assetId) is used.  (optional)
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.PutAsset(context.Background()).Asset(asset).IdentifyBy(identifyBy).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.PutAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAsset`: Asset
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.PutAsset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | [**Asset**](Asset.md) |  | 
 **identifyBy** | **string** | Serves the field name send in the request body as a unique identifier for the asset, essential for operations like updates or deletions. Please refer to the Asset schema definition for further information about this fields.  In cases where this parameter isn&#39;t defined, all field names are used in the order defined. So if there is no &#39;resourceId&#39; present in the request body, the &#39;deviceId&#39; is used and when there is also no deviceId present the &#39;id&#39; field (assetId) is used.  | 
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**Asset**](Asset.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAssetById

> Asset PutAssetById(ctx, assetId).Asset(asset).Expansions(expansions).Execute()

Update an asset



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
	assetId := int32(4711) // int32 | The id of the asset
	asset := *openapiclient.NewAsset("zurich_swiss", "AssetType_example") // Asset | 
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.PutAssetById(context.Background(), assetId).Asset(asset).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.PutAssetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAssetById`: Asset
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.PutAssetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **int32** | The id of the asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAssetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **asset** | [**Asset**](Asset.md) |  | 
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**Asset**](Asset.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAttributeDisplay

> AttributeDisplay PutAttributeDisplay(ctx).AttributeDisplay(attributeDisplay).Execute()

Create or update how attributes are displayed



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
	attributeDisplay := *openapiclient.NewAttributeDisplay(int32(4711), openapiclient.DataSubtype("input"), "temperature") // AttributeDisplay | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.PutAttributeDisplay(context.Background()).AttributeDisplay(attributeDisplay).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.PutAttributeDisplay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAttributeDisplay`: AttributeDisplay
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.PutAttributeDisplay`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutAttributeDisplayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attributeDisplay** | [**AttributeDisplay**](AttributeDisplay.md) |  | 

### Return type

[**AttributeDisplay**](AttributeDisplay.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutBulkAssets

> []Asset PutBulkAssets(ctx).Asset(asset).IdentifyBy(identifyBy).Expansions(expansions).Execute()

Create or update a list of assets



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
	asset := []openapiclient.Asset{*openapiclient.NewAsset("zurich_swiss", "AssetType_example")} // []Asset | 
	identifyBy := "resourceId" // string | Serves the field name send in the request body as a unique identifier for the asset, essential for operations like updates or deletions. Please refer to the Asset schema definition for further information about this fields.  In cases where this parameter isn't defined, all field names are used in the order defined. So if there is no 'resourceId' present in the request body, the 'deviceId' is used and when there is also no deviceId present the 'id' field (assetId) is used.  (optional)
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.PutBulkAssets(context.Background()).Asset(asset).IdentifyBy(identifyBy).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.PutBulkAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutBulkAssets`: []Asset
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.PutBulkAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutBulkAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | [**[]Asset**](Asset.md) |  | 
 **identifyBy** | **string** | Serves the field name send in the request body as a unique identifier for the asset, essential for operations like updates or deletions. Please refer to the Asset schema definition for further information about this fields.  In cases where this parameter isn&#39;t defined, all field names are used in the order defined. So if there is no &#39;resourceId&#39; present in the request body, the &#39;deviceId&#39; is used and when there is also no deviceId present the &#39;id&#39; field (assetId) is used.  | 
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**[]Asset**](Asset.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


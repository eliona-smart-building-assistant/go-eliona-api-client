# \AssetTypesAPI

All URIs are relative to *https://name.eliona.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAssetTypeByName**](AssetTypesAPI.md#DeleteAssetTypeByName) | **Delete** /asset-types/{asset-type-name} | Delete an asset type
[**DeleteAssetTypeCategoryByName**](AssetTypesAPI.md#DeleteAssetTypeCategoryByName) | **Delete** /asset-type-categories/{asset-type-category-name} | Delete an asset type category
[**GetAssetTypeByName**](AssetTypesAPI.md#GetAssetTypeByName) | **Get** /asset-types/{asset-type-name} | Information about an asset type
[**GetAssetTypeCategories**](AssetTypesAPI.md#GetAssetTypeCategories) | **Get** /asset-type-categories | List of asset type categories
[**GetAssetTypes**](AssetTypesAPI.md#GetAssetTypes) | **Get** /asset-types | List of asset types
[**PostAssetType**](AssetTypesAPI.md#PostAssetType) | **Post** /asset-types | Create an asset type
[**PostAssetTypeAttribute**](AssetTypesAPI.md#PostAssetTypeAttribute) | **Post** /asset-types/{asset-type-name}/attributes | Create asset type attribute
[**PutAssetType**](AssetTypesAPI.md#PutAssetType) | **Put** /asset-types | Create or update an asset type
[**PutAssetTypeAttribute**](AssetTypesAPI.md#PutAssetTypeAttribute) | **Put** /asset-types/{asset-type-name}/attributes | Create or update an asset type attribute
[**PutAssetTypeByName**](AssetTypesAPI.md#PutAssetTypeByName) | **Put** /asset-types/{asset-type-name} | Update an asset type
[**PutAssetTypeCategory**](AssetTypesAPI.md#PutAssetTypeCategory) | **Put** /asset-type-categories | Create or update an asset type categories



## DeleteAssetTypeByName

> DeleteAssetTypeByName(ctx, assetTypeName).Execute()

Delete an asset type



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
	assetTypeName := "weather_location" // string | The name of the asset type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AssetTypesAPI.DeleteAssetTypeByName(context.Background(), assetTypeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetTypesAPI.DeleteAssetTypeByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetTypeName** | **string** | The name of the asset type | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAssetTypeByNameRequest struct via the builder pattern


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


## DeleteAssetTypeCategoryByName

> DeleteAssetTypeCategoryByName(ctx, assetTypeCategoryName).Execute()

Delete an asset type category



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
	assetTypeCategoryName := "location" // string | The name of the asset type category

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AssetTypesAPI.DeleteAssetTypeCategoryByName(context.Background(), assetTypeCategoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetTypesAPI.DeleteAssetTypeCategoryByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetTypeCategoryName** | **string** | The name of the asset type category | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAssetTypeCategoryByNameRequest struct via the builder pattern


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


## GetAssetTypeByName

> AssetType GetAssetTypeByName(ctx, assetTypeName).Expansions(expansions).Execute()

Information about an asset type



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
	assetTypeName := "weather_location" // string | The name of the asset type
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetTypesAPI.GetAssetTypeByName(context.Background(), assetTypeName).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetTypesAPI.GetAssetTypeByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetTypeByName`: AssetType
	fmt.Fprintf(os.Stdout, "Response from `AssetTypesAPI.GetAssetTypeByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetTypeName** | **string** | The name of the asset type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetTypeByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**AssetType**](AssetType.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetTypeCategories

> []AssetTypeCategory GetAssetTypeCategories(ctx).Execute()

List of asset type categories



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
	resp, r, err := apiClient.AssetTypesAPI.GetAssetTypeCategories(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetTypesAPI.GetAssetTypeCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetTypeCategories`: []AssetTypeCategory
	fmt.Fprintf(os.Stdout, "Response from `AssetTypesAPI.GetAssetTypeCategories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetTypeCategoriesRequest struct via the builder pattern


### Return type

[**[]AssetTypeCategory**](AssetTypeCategory.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetTypes

> []AssetType GetAssetTypes(ctx).Expansions(expansions).Execute()

List of asset types



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
	resp, r, err := apiClient.AssetTypesAPI.GetAssetTypes(context.Background()).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetTypesAPI.GetAssetTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetTypes`: []AssetType
	fmt.Fprintf(os.Stdout, "Response from `AssetTypesAPI.GetAssetTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**[]AssetType**](AssetType.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAssetType

> AssetType PostAssetType(ctx).AssetType(assetType).Expansions(expansions).Execute()

Create an asset type



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
	assetType := *openapiclient.NewAssetType("weather_location") // AssetType | 
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetTypesAPI.PostAssetType(context.Background()).AssetType(assetType).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetTypesAPI.PostAssetType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAssetType`: AssetType
	fmt.Fprintf(os.Stdout, "Response from `AssetTypesAPI.PostAssetType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAssetTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetType** | [**AssetType**](AssetType.md) |  | 
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**AssetType**](AssetType.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAssetTypeAttribute

> AssetTypeAttribute PostAssetTypeAttribute(ctx, assetTypeName).AssetTypeAttribute(assetTypeAttribute).Execute()

Create asset type attribute



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
	assetTypeName := "weather_location" // string | The name of the asset type
	assetTypeAttribute := *openapiclient.NewAssetTypeAttribute("temperature", openapiclient.DataSubtype("input")) // AssetTypeAttribute | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetTypesAPI.PostAssetTypeAttribute(context.Background(), assetTypeName).AssetTypeAttribute(assetTypeAttribute).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetTypesAPI.PostAssetTypeAttribute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAssetTypeAttribute`: AssetTypeAttribute
	fmt.Fprintf(os.Stdout, "Response from `AssetTypesAPI.PostAssetTypeAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetTypeName** | **string** | The name of the asset type | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAssetTypeAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assetTypeAttribute** | [**AssetTypeAttribute**](AssetTypeAttribute.md) |  | 

### Return type

[**AssetTypeAttribute**](AssetTypeAttribute.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAssetType

> AssetType PutAssetType(ctx).AssetType(assetType).Expansions(expansions).Execute()

Create or update an asset type



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
	assetType := *openapiclient.NewAssetType("weather_location") // AssetType | 
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetTypesAPI.PutAssetType(context.Background()).AssetType(assetType).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetTypesAPI.PutAssetType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAssetType`: AssetType
	fmt.Fprintf(os.Stdout, "Response from `AssetTypesAPI.PutAssetType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutAssetTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetType** | [**AssetType**](AssetType.md) |  | 
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**AssetType**](AssetType.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAssetTypeAttribute

> AssetTypeAttribute PutAssetTypeAttribute(ctx, assetTypeName).AssetTypeAttribute(assetTypeAttribute).Execute()

Create or update an asset type attribute



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
	assetTypeName := "weather_location" // string | The name of the asset type
	assetTypeAttribute := *openapiclient.NewAssetTypeAttribute("temperature", openapiclient.DataSubtype("input")) // AssetTypeAttribute | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetTypesAPI.PutAssetTypeAttribute(context.Background(), assetTypeName).AssetTypeAttribute(assetTypeAttribute).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetTypesAPI.PutAssetTypeAttribute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAssetTypeAttribute`: AssetTypeAttribute
	fmt.Fprintf(os.Stdout, "Response from `AssetTypesAPI.PutAssetTypeAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetTypeName** | **string** | The name of the asset type | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAssetTypeAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assetTypeAttribute** | [**AssetTypeAttribute**](AssetTypeAttribute.md) |  | 

### Return type

[**AssetTypeAttribute**](AssetTypeAttribute.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAssetTypeByName

> AssetType PutAssetTypeByName(ctx, assetTypeName).AssetType(assetType).Expansions(expansions).Execute()

Update an asset type



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
	assetTypeName := "weather_location" // string | The name of the asset type
	assetType := *openapiclient.NewAssetType("weather_location") // AssetType | 
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetTypesAPI.PutAssetTypeByName(context.Background(), assetTypeName).AssetType(assetType).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetTypesAPI.PutAssetTypeByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAssetTypeByName`: AssetType
	fmt.Fprintf(os.Stdout, "Response from `AssetTypesAPI.PutAssetTypeByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetTypeName** | **string** | The name of the asset type | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAssetTypeByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assetType** | [**AssetType**](AssetType.md) |  | 
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**AssetType**](AssetType.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAssetTypeCategory

> AssetTypeCategory PutAssetTypeCategory(ctx).AssetTypeCategory(assetTypeCategory).Execute()

Create or update an asset type categories



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
	assetTypeCategory := *openapiclient.NewAssetTypeCategory("location") // AssetTypeCategory | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetTypesAPI.PutAssetTypeCategory(context.Background()).AssetTypeCategory(assetTypeCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetTypesAPI.PutAssetTypeCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAssetTypeCategory`: AssetTypeCategory
	fmt.Fprintf(os.Stdout, "Response from `AssetTypesAPI.PutAssetTypeCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutAssetTypeCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetTypeCategory** | [**AssetTypeCategory**](AssetTypeCategory.md) |  | 

### Return type

[**AssetTypeCategory**](AssetTypeCategory.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \ElionaAPI

All URIs are relative to *https://name.eliona.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTenant**](ElionaAPI.md#DeleteTenant) | **Delete** /tenants | Delete the current tenant
[**GetApiKeys**](ElionaAPI.md#GetApiKeys) | **Get** /api-keys | Information about API keys
[**PutTenant**](ElionaAPI.md#PutTenant) | **Put** /tenants | Setup or update a tenant



## DeleteTenant

> DeleteTenant(ctx).Execute()

Delete the current tenant



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ElionaAPI.DeleteTenant(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElionaAPI.DeleteTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantRequest struct via the builder pattern


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


## GetApiKeys

> []ApiKey GetApiKeys(ctx).Offset(offset).Size(size).Execute()

Information about API keys



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
	resp, r, err := apiClient.ElionaAPI.GetApiKeys(context.Background()).Offset(offset).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElionaAPI.GetApiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiKeys`: []ApiKey
	fmt.Fprintf(os.Stdout, "Response from `ElionaAPI.GetApiKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int64** | Specifies the starting point for pagination by indicating the number of items to skip.  | 
 **size** | **int64** | Specifies the number of items per page for pagination.  | 

### Return type

[**[]ApiKey**](ApiKey.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTenant

> Tenant PutTenant(ctx).SignatureInput(signatureInput).ContentDigest(contentDigest).Tenant(tenant).Execute()

Setup or update a tenant



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
	signatureInput := "sig1=("@method" "@path" "content-digest");created=1618884473;keyid="client-key-1"" // string | Defines what was signed (e.g. method, path, `content-digest`) and metadata like creation time and key ID; must match the `Signature`.
	contentDigest := "sha-256=:47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=:" // string | Hash of the request body; server recomputes and compares. Must be included in `Signature-Input`.
	tenant := *openapiclient.NewTenant("Main") // Tenant | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElionaAPI.PutTenant(context.Background()).SignatureInput(signatureInput).ContentDigest(contentDigest).Tenant(tenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElionaAPI.PutTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTenant`: Tenant
	fmt.Fprintf(os.Stdout, "Response from `ElionaAPI.PutTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signatureInput** | **string** | Defines what was signed (e.g. method, path, &#x60;content-digest&#x60;) and metadata like creation time and key ID; must match the &#x60;Signature&#x60;. | 
 **contentDigest** | **string** | Hash of the request body; server recomputes and compares. Must be included in &#x60;Signature-Input&#x60;. | 
 **tenant** | [**Tenant**](Tenant.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[HttpSig](../README.md#HttpSig)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


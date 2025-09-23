# \TagsAPI

All URIs are relative to *https://name.eliona.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTagByName**](TagsAPI.md#GetTagByName) | **Get** /tags/{tag-name} | Information about a tag
[**GetTags**](TagsAPI.md#GetTags) | **Get** /tags | Information about tags
[**PutTag**](TagsAPI.md#PutTag) | **Put** /tags | Create or update a tag



## GetTagByName

> Tag GetTagByName(ctx, tagName).Execute()

Information about a tag



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
	tagName := "Support" // string | The name of the tag

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.GetTagByName(context.Background(), tagName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.GetTagByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTagByName`: Tag
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.GetTagByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagName** | **string** | The name of the tag | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Tag**](Tag.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTags

> []Tag GetTags(ctx).Offset(offset).Size(size).Execute()

Information about tags



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
	resp, r, err := apiClient.TagsAPI.GetTags(context.Background()).Offset(offset).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.GetTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTags`: []Tag
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.GetTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int64** | Specifies the starting point for pagination by indicating the number of items to skip.  | 
 **size** | **int64** | Specifies the number of items per page for pagination.  | 

### Return type

[**[]Tag**](Tag.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTag

> Tag PutTag(ctx).Tag(tag).Execute()

Create or update a tag



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
	tag := *openapiclient.NewTag("Support") // Tag | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.PutTag(context.Background()).Tag(tag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.PutTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTag`: Tag
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.PutTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | [**Tag**](Tag.md) |  | 

### Return type

[**Tag**](Tag.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


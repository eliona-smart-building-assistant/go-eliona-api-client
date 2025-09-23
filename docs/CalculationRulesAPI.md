# \CalculationRulesAPI

All URIs are relative to *https://name.eliona.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCalculationRuleById**](CalculationRulesAPI.md#DeleteCalculationRuleById) | **Delete** /calculation-rules/{calculation-rule-id} | Delete a calculation rule
[**GetCalculationRuleById**](CalculationRulesAPI.md#GetCalculationRuleById) | **Get** /calculation-rules/{calculation-rule-id} | Information about a calculation rules rule
[**GetCalculationRules**](CalculationRulesAPI.md#GetCalculationRules) | **Get** /calculation-rules | Information about calculation rules
[**PutCalculationRule**](CalculationRulesAPI.md#PutCalculationRule) | **Put** /calculation-rules | Creates or updates a calculation rule
[**PutCalculationRuleById**](CalculationRulesAPI.md#PutCalculationRuleById) | **Put** /calculation-rules/{calculation-rule-id} | Update a calculation rule



## DeleteCalculationRuleById

> DeleteCalculationRuleById(ctx, calculationRuleId).Execute()

Delete a calculation rule



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
	calculationRuleId := int32(4711) // int32 | The id of the calculation rule

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CalculationRulesAPI.DeleteCalculationRuleById(context.Background(), calculationRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalculationRulesAPI.DeleteCalculationRuleById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**calculationRuleId** | **int32** | The id of the calculation rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCalculationRuleByIdRequest struct via the builder pattern


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


## GetCalculationRuleById

> CalculationRule GetCalculationRuleById(ctx, calculationRuleId).Execute()

Information about a calculation rules rule



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
	calculationRuleId := int32(4711) // int32 | The id of the calculation rule

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalculationRulesAPI.GetCalculationRuleById(context.Background(), calculationRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalculationRulesAPI.GetCalculationRuleById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCalculationRuleById`: CalculationRule
	fmt.Fprintf(os.Stdout, "Response from `CalculationRulesAPI.GetCalculationRuleById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**calculationRuleId** | **int32** | The id of the calculation rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCalculationRuleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CalculationRule**](CalculationRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCalculationRules

> []CalculationRule GetCalculationRules(ctx).CalculationRuleIds(calculationRuleIds).Offset(offset).Size(size).Execute()

Information about calculation rules



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
	calculationRuleIds := []int32{int32(123)} // []int32 | List of calculation rule ids for filtering (optional)
	offset := int64(3) // int64 | Specifies the starting point for pagination by indicating the number of items to skip.  (optional)
	size := int64(10) // int64 | Specifies the number of items per page for pagination.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalculationRulesAPI.GetCalculationRules(context.Background()).CalculationRuleIds(calculationRuleIds).Offset(offset).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalculationRulesAPI.GetCalculationRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCalculationRules`: []CalculationRule
	fmt.Fprintf(os.Stdout, "Response from `CalculationRulesAPI.GetCalculationRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCalculationRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **calculationRuleIds** | **[]int32** | List of calculation rule ids for filtering | 
 **offset** | **int64** | Specifies the starting point for pagination by indicating the number of items to skip.  | 
 **size** | **int64** | Specifies the number of items per page for pagination.  | 

### Return type

[**[]CalculationRule**](CalculationRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCalculationRule

> CalculationRule PutCalculationRule(ctx).CalculationRule(calculationRule).Execute()

Creates or updates a calculation rule



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
	calculationRule := *openapiclient.NewCalculationRule(int32(4711), openapiclient.DataSubtype("input"), "temperature") // CalculationRule | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalculationRulesAPI.PutCalculationRule(context.Background()).CalculationRule(calculationRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalculationRulesAPI.PutCalculationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCalculationRule`: CalculationRule
	fmt.Fprintf(os.Stdout, "Response from `CalculationRulesAPI.PutCalculationRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutCalculationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **calculationRule** | [**CalculationRule**](CalculationRule.md) |  | 

### Return type

[**CalculationRule**](CalculationRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCalculationRuleById

> CalculationRule PutCalculationRuleById(ctx, calculationRuleId).CalculationRule(calculationRule).Execute()

Update a calculation rule



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
	calculationRuleId := int32(4711) // int32 | The id of the calculation rule
	calculationRule := *openapiclient.NewCalculationRule(int32(4711), openapiclient.DataSubtype("input"), "temperature") // CalculationRule | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalculationRulesAPI.PutCalculationRuleById(context.Background(), calculationRuleId).CalculationRule(calculationRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalculationRulesAPI.PutCalculationRuleById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCalculationRuleById`: CalculationRule
	fmt.Fprintf(os.Stdout, "Response from `CalculationRulesAPI.PutCalculationRuleById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**calculationRuleId** | **int32** | The id of the calculation rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCalculationRuleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **calculationRule** | [**CalculationRule**](CalculationRule.md) |  | 

### Return type

[**CalculationRule**](CalculationRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


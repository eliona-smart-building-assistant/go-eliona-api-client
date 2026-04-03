# \AlarmRulesAPI

All URIs are relative to *https://name.eliona.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAlarmRuleById**](AlarmRulesAPI.md#DeleteAlarmRuleById) | **Delete** /alarm-rules/{alarm-rule-id} | Delete an alarm rule
[**GetAlarmRuleById**](AlarmRulesAPI.md#GetAlarmRuleById) | **Get** /alarm-rules/{alarm-rule-id} | Information about an alarm rule
[**GetAlarmRules**](AlarmRulesAPI.md#GetAlarmRules) | **Get** /alarm-rules | Information about alarm rules
[**PostAlarmRule**](AlarmRulesAPI.md#PostAlarmRule) | **Post** /alarm-rules | Create an alarm rule
[**PutAlarmRule**](AlarmRulesAPI.md#PutAlarmRule) | **Put** /alarm-rules | Create or update an alarm rule
[**PutAlarmRuleById**](AlarmRulesAPI.md#PutAlarmRuleById) | **Put** /alarm-rules/{alarm-rule-id} | Update an alarm rule



## DeleteAlarmRuleById

> DeleteAlarmRuleById(ctx, alarmRuleId).Execute()

Delete an alarm rule



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
	alarmRuleId := int32(4711) // int32 | The id of the alarm rule

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlarmRulesAPI.DeleteAlarmRuleById(context.Background(), alarmRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlarmRulesAPI.DeleteAlarmRuleById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alarmRuleId** | **int32** | The id of the alarm rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlarmRuleByIdRequest struct via the builder pattern


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


## GetAlarmRuleById

> AlarmRule GetAlarmRuleById(ctx, alarmRuleId).Expansions(expansions).Execute()

Information about an alarm rule



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
	alarmRuleId := int32(4711) // int32 | The id of the alarm rule
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlarmRulesAPI.GetAlarmRuleById(context.Background(), alarmRuleId).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlarmRulesAPI.GetAlarmRuleById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlarmRuleById`: AlarmRule
	fmt.Fprintf(os.Stdout, "Response from `AlarmRulesAPI.GetAlarmRuleById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alarmRuleId** | **int32** | The id of the alarm rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlarmRuleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**AlarmRule**](AlarmRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlarmRules

> []AlarmRule GetAlarmRules(ctx).AlarmRuleIds(alarmRuleIds).AssetId(assetId).Offset(offset).Size(size).Expansions(expansions).Execute()

Information about alarm rules



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
	alarmRuleIds := []int32{int32(123)} // []int32 | List of alarm rule ids for filtering (optional)
	assetId := int32(4711) // int32 | Filter for a specific asset id (optional)
	offset := int64(3) // int64 | Specifies the starting point for pagination by indicating the number of items to skip.  (optional)
	size := int64(10) // int64 | Specifies the number of items per page for pagination.  (optional)
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlarmRulesAPI.GetAlarmRules(context.Background()).AlarmRuleIds(alarmRuleIds).AssetId(assetId).Offset(offset).Size(size).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlarmRulesAPI.GetAlarmRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlarmRules`: []AlarmRule
	fmt.Fprintf(os.Stdout, "Response from `AlarmRulesAPI.GetAlarmRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlarmRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alarmRuleIds** | **[]int32** | List of alarm rule ids for filtering | 
 **assetId** | **int32** | Filter for a specific asset id | 
 **offset** | **int64** | Specifies the starting point for pagination by indicating the number of items to skip.  | 
 **size** | **int64** | Specifies the number of items per page for pagination.  | 
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**[]AlarmRule**](AlarmRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAlarmRule

> AlarmRule PostAlarmRule(ctx).AlarmRule(alarmRule).Execute()

Create an alarm rule



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
	alarmRule := *openapiclient.NewAlarmRule(int32(4711), openapiclient.DataSubtype("input"), "temperature", openapiclient.AlarmPriority(1)) // AlarmRule | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlarmRulesAPI.PostAlarmRule(context.Background()).AlarmRule(alarmRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlarmRulesAPI.PostAlarmRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAlarmRule`: AlarmRule
	fmt.Fprintf(os.Stdout, "Response from `AlarmRulesAPI.PostAlarmRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAlarmRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alarmRule** | [**AlarmRule**](AlarmRule.md) |  | 

### Return type

[**AlarmRule**](AlarmRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAlarmRule

> AlarmRule PutAlarmRule(ctx).AlarmRule(alarmRule).Execute()

Create or update an alarm rule



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
	alarmRule := *openapiclient.NewAlarmRule(int32(4711), openapiclient.DataSubtype("input"), "temperature", openapiclient.AlarmPriority(1)) // AlarmRule | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlarmRulesAPI.PutAlarmRule(context.Background()).AlarmRule(alarmRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlarmRulesAPI.PutAlarmRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAlarmRule`: AlarmRule
	fmt.Fprintf(os.Stdout, "Response from `AlarmRulesAPI.PutAlarmRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutAlarmRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alarmRule** | [**AlarmRule**](AlarmRule.md) |  | 

### Return type

[**AlarmRule**](AlarmRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAlarmRuleById

> AlarmRule PutAlarmRuleById(ctx, alarmRuleId).AlarmRule(alarmRule).Execute()

Update an alarm rule



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
	alarmRuleId := int32(4711) // int32 | The id of the alarm rule
	alarmRule := *openapiclient.NewAlarmRule(int32(4711), openapiclient.DataSubtype("input"), "temperature", openapiclient.AlarmPriority(1)) // AlarmRule | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlarmRulesAPI.PutAlarmRuleById(context.Background(), alarmRuleId).AlarmRule(alarmRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlarmRulesAPI.PutAlarmRuleById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAlarmRuleById`: AlarmRule
	fmt.Fprintf(os.Stdout, "Response from `AlarmRulesAPI.PutAlarmRuleById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alarmRuleId** | **int32** | The id of the alarm rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAlarmRuleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alarmRule** | [**AlarmRule**](AlarmRule.md) |  | 

### Return type

[**AlarmRule**](AlarmRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


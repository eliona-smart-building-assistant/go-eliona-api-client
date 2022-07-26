# \AlarmRulesApi

All URIs are relative to *http://api.eliona.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAlarmRuleById**](AlarmRulesApi.md#GetAlarmRuleById) | **Get** /alarm-rules/{alarm-rule-id} | Information about an alarm rule
[**GetAlarmRules**](AlarmRulesApi.md#GetAlarmRules) | **Get** /alarm-rules | Information about alarm rules
[**PostAlarmRule**](AlarmRulesApi.md#PostAlarmRule) | **Post** /alarm-rules | Create an alarm rule
[**PutAlarmRule**](AlarmRulesApi.md#PutAlarmRule) | **Put** /alarm-rules | Create or update an alarm rule
[**PutAlarmRuleById**](AlarmRulesApi.md#PutAlarmRuleById) | **Put** /alarm-rules/{alarm-rule-id} | Update an alarm rule



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
    openapiclient "./openapi"
)

func main() {
    alarmRuleId := int32(4711) // int32 | The id of the alarm rule
    expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlarmRulesApi.GetAlarmRuleById(context.Background(), alarmRuleId).Expansions(expansions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlarmRulesApi.GetAlarmRuleById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlarmRuleById`: AlarmRule
    fmt.Fprintf(os.Stdout, "Response from `AlarmRulesApi.GetAlarmRuleById`: %v\n", resp)
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

> []AlarmRule GetAlarmRules(ctx).Expansions(expansions).Execute()

Information about alarm rules



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlarmRulesApi.GetAlarmRules(context.Background()).Expansions(expansions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlarmRulesApi.GetAlarmRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlarmRules`: []AlarmRule
    fmt.Fprintf(os.Stdout, "Response from `AlarmRulesApi.GetAlarmRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlarmRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**[]AlarmRule**](AlarmRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
    openapiclient "./openapi"
)

func main() {
    alarmRule := *openapiclient.NewAlarmRule(int32(4711), openapiclient.DataSubtype("input"), "temperature", openapiclient.AlarmPriority(1)) // AlarmRule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlarmRulesApi.PostAlarmRule(context.Background()).AlarmRule(alarmRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlarmRulesApi.PostAlarmRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostAlarmRule`: AlarmRule
    fmt.Fprintf(os.Stdout, "Response from `AlarmRulesApi.PostAlarmRule`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    alarmRule := *openapiclient.NewAlarmRule(int32(4711), openapiclient.DataSubtype("input"), "temperature", openapiclient.AlarmPriority(1)) // AlarmRule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlarmRulesApi.PutAlarmRule(context.Background()).AlarmRule(alarmRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlarmRulesApi.PutAlarmRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutAlarmRule`: AlarmRule
    fmt.Fprintf(os.Stdout, "Response from `AlarmRulesApi.PutAlarmRule`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    alarmRuleId := int32(4711) // int32 | The id of the alarm rule
    alarmRule := *openapiclient.NewAlarmRule(int32(4711), openapiclient.DataSubtype("input"), "temperature", openapiclient.AlarmPriority(1)) // AlarmRule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlarmRulesApi.PutAlarmRuleById(context.Background(), alarmRuleId).AlarmRule(alarmRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlarmRulesApi.PutAlarmRuleById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutAlarmRuleById`: AlarmRule
    fmt.Fprintf(os.Stdout, "Response from `AlarmRulesApi.PutAlarmRuleById`: %v\n", resp)
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


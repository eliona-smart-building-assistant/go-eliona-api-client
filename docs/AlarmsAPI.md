# \AlarmsAPI

All URIs are relative to *https://name.eliona.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAlarmById**](AlarmsAPI.md#DeleteAlarmById) | **Delete** /alarms/{alarm-rule-id} | Removes an alarm
[**GetAlarmById**](AlarmsAPI.md#GetAlarmById) | **Get** /alarms/{alarm-rule-id} | Information about alarm
[**GetAlarmHistoryById**](AlarmsAPI.md#GetAlarmHistoryById) | **Get** /alarms-history/{alarm-rule-id} | Information about alarm history
[**GetAlarms**](AlarmsAPI.md#GetAlarms) | **Get** /alarms | Information about alarms
[**GetAlarmsHistory**](AlarmsAPI.md#GetAlarmsHistory) | **Get** /alarms-history | Information about alarms history
[**GetHighestAlarms**](AlarmsAPI.md#GetHighestAlarms) | **Get** /alarms-highest | Information about most prioritized alarms
[**ListenAlarm**](AlarmsAPI.md#ListenAlarm) | **Get** /alarm-listener | WebSocket connection for alarm changes
[**PatchAlarmById**](AlarmsAPI.md#PatchAlarmById) | **Patch** /alarms/{alarm-rule-id} | Update alarm
[**PutAlarm**](AlarmsAPI.md#PutAlarm) | **Put** /alarms | Create or update an alarm



## DeleteAlarmById

> DeleteAlarmById(ctx, alarmRuleId).Execute()

Removes an alarm



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
	alarmRuleId := int32(4711) // int32 | The id of the alarm rule

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlarmsAPI.DeleteAlarmById(context.Background(), alarmRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlarmsAPI.DeleteAlarmById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAlarmByIdRequest struct via the builder pattern


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


## GetAlarmById

> Alarm GetAlarmById(ctx, alarmRuleId).Expansions(expansions).Execute()

Information about alarm



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
	alarmRuleId := int32(4711) // int32 | The id of the alarm rule
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlarmsAPI.GetAlarmById(context.Background(), alarmRuleId).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlarmsAPI.GetAlarmById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlarmById`: Alarm
	fmt.Fprintf(os.Stdout, "Response from `AlarmsAPI.GetAlarmById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alarmRuleId** | **int32** | The id of the alarm rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlarmByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**Alarm**](Alarm.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlarmHistoryById

> []Alarm GetAlarmHistoryById(ctx, alarmRuleId).FromDate(fromDate).ToDate(toDate).Tags(tags).Expansions(expansions).Execute()

Information about alarm history



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
	alarmRuleId := int32(4711) // int32 | The id of the alarm rule
	fromDate := "2020-01-01T09:00:00.000Z" // string | Filter by lower date time (RFC3339) limit inclusive (optional)
	toDate := "2021-12-31T23:00:00.000Z" // string | Filter by upper date time (RFC3339) limit exclusive (optional)
	tags := []string{"Inner_example"} // []string | A list of defined tags. Result must include all of these tags, not just some. (optional)
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlarmsAPI.GetAlarmHistoryById(context.Background(), alarmRuleId).FromDate(fromDate).ToDate(toDate).Tags(tags).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlarmsAPI.GetAlarmHistoryById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlarmHistoryById`: []Alarm
	fmt.Fprintf(os.Stdout, "Response from `AlarmsAPI.GetAlarmHistoryById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alarmRuleId** | **int32** | The id of the alarm rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlarmHistoryByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | **string** | Filter by lower date time (RFC3339) limit inclusive | 
 **toDate** | **string** | Filter by upper date time (RFC3339) limit exclusive | 
 **tags** | **[]string** | A list of defined tags. Result must include all of these tags, not just some. | 
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**[]Alarm**](Alarm.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlarms

> []Alarm GetAlarms(ctx).ProjectId(projectId).FromDate(fromDate).ToDate(toDate).Tags(tags).Expansions(expansions).Execute()

Information about alarms



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
	projectId := "projectId_example" // string | Filter for a specific project (optional)
	fromDate := "2020-01-01T09:00:00.000Z" // string | Filter by lower date time (RFC3339) limit inclusive (optional)
	toDate := "2021-12-31T23:00:00.000Z" // string | Filter by upper date time (RFC3339) limit exclusive (optional)
	tags := []string{"Inner_example"} // []string | A list of defined tags. Result must include all of these tags, not just some. (optional)
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlarmsAPI.GetAlarms(context.Background()).ProjectId(projectId).FromDate(fromDate).ToDate(toDate).Tags(tags).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlarmsAPI.GetAlarms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlarms`: []Alarm
	fmt.Fprintf(os.Stdout, "Response from `AlarmsAPI.GetAlarms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlarmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** | Filter for a specific project | 
 **fromDate** | **string** | Filter by lower date time (RFC3339) limit inclusive | 
 **toDate** | **string** | Filter by upper date time (RFC3339) limit exclusive | 
 **tags** | **[]string** | A list of defined tags. Result must include all of these tags, not just some. | 
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**[]Alarm**](Alarm.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlarmsHistory

> []Alarm GetAlarmsHistory(ctx).ProjectId(projectId).FromDate(fromDate).ToDate(toDate).Tags(tags).Expansions(expansions).Execute()

Information about alarms history



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
	projectId := "projectId_example" // string | Filter for a specific project (optional)
	fromDate := "2020-01-01T09:00:00.000Z" // string | Filter by lower date time (RFC3339) limit inclusive (optional)
	toDate := "2021-12-31T23:00:00.000Z" // string | Filter by upper date time (RFC3339) limit exclusive (optional)
	tags := []string{"Inner_example"} // []string | A list of defined tags. Result must include all of these tags, not just some. (optional)
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlarmsAPI.GetAlarmsHistory(context.Background()).ProjectId(projectId).FromDate(fromDate).ToDate(toDate).Tags(tags).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlarmsAPI.GetAlarmsHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlarmsHistory`: []Alarm
	fmt.Fprintf(os.Stdout, "Response from `AlarmsAPI.GetAlarmsHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlarmsHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** | Filter for a specific project | 
 **fromDate** | **string** | Filter by lower date time (RFC3339) limit inclusive | 
 **toDate** | **string** | Filter by upper date time (RFC3339) limit exclusive | 
 **tags** | **[]string** | A list of defined tags. Result must include all of these tags, not just some. | 
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**[]Alarm**](Alarm.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHighestAlarms

> []Alarm GetHighestAlarms(ctx).ProjectId(projectId).FromDate(fromDate).ToDate(toDate).Tags(tags).Expansions(expansions).Execute()

Information about most prioritized alarms



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
	projectId := "projectId_example" // string | Filter for a specific project (optional)
	fromDate := "2020-01-01T09:00:00.000Z" // string | Filter by lower date time (RFC3339) limit inclusive (optional)
	toDate := "2021-12-31T23:00:00.000Z" // string | Filter by upper date time (RFC3339) limit exclusive (optional)
	tags := []string{"Inner_example"} // []string | A list of defined tags. Result must include all of these tags, not just some. (optional)
	expansions := []string{"Inner_example"} // []string | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows 'ObjectName.fieldName'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlarmsAPI.GetHighestAlarms(context.Background()).ProjectId(projectId).FromDate(fromDate).ToDate(toDate).Tags(tags).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlarmsAPI.GetHighestAlarms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHighestAlarms`: []Alarm
	fmt.Fprintf(os.Stdout, "Response from `AlarmsAPI.GetHighestAlarms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHighestAlarmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** | Filter for a specific project | 
 **fromDate** | **string** | Filter by lower date time (RFC3339) limit inclusive | 
 **toDate** | **string** | Filter by upper date time (RFC3339) limit exclusive | 
 **tags** | **[]string** | A list of defined tags. Result must include all of these tags, not just some. | 
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**[]Alarm**](Alarm.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListenAlarm

> AlarmListen ListenAlarm(ctx).Expansions(expansions).Execute()

WebSocket connection for alarm changes



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
	resp, r, err := apiClient.AlarmsAPI.ListenAlarm(context.Background()).Expansions(expansions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlarmsAPI.ListenAlarm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListenAlarm`: AlarmListen
	fmt.Fprintf(os.Stdout, "Response from `AlarmsAPI.ListenAlarm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListenAlarmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expansions** | **[]string** | List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;. | 

### Return type

[**AlarmListen**](AlarmListen.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAlarmById

> PatchAlarmById(ctx, alarmRuleId).Acknowledged(acknowledged).AcknowledgeText(acknowledgeText).Execute()

Update alarm



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
	alarmRuleId := int32(4711) // int32 | The id of the alarm rule
	acknowledged := true // bool | Marks the alarm as acknowledged or not acknowledged by setting the acknowledge timestamp to now or to null.
	acknowledgeText := "acknowledgeText_example" // string | Sets the text for acknowledgement if acknowledged is set to true (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlarmsAPI.PatchAlarmById(context.Background(), alarmRuleId).Acknowledged(acknowledged).AcknowledgeText(acknowledgeText).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlarmsAPI.PatchAlarmById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPatchAlarmByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acknowledged** | **bool** | Marks the alarm as acknowledged or not acknowledged by setting the acknowledge timestamp to now or to null. | 
 **acknowledgeText** | **string** | Sets the text for acknowledgement if acknowledged is set to true | 

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


## PutAlarm

> Alarm PutAlarm(ctx).Alarm(alarm).Execute()

Create or update an alarm



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
	alarm := *openapiclient.NewAlarm(int32(123), map[string]interface{}({})) // Alarm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlarmsAPI.PutAlarm(context.Background()).Alarm(alarm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlarmsAPI.PutAlarm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAlarm`: Alarm
	fmt.Fprintf(os.Stdout, "Response from `AlarmsAPI.PutAlarm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutAlarmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alarm** | [**Alarm**](Alarm.md) |  | 

### Return type

[**Alarm**](Alarm.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \CommunicationAPI

All URIs are relative to *https://name.eliona.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMessageReceiptById**](CommunicationAPI.md#GetMessageReceiptById) | **Get** /message-receipts/{message-id} | Information about a message
[**PostMail**](CommunicationAPI.md#PostMail) | **Post** /send-mail | Send e-mail
[**PostNotification**](CommunicationAPI.md#PostNotification) | **Post** /send-notification | Send notification



## GetMessageReceiptById

> MessageReceipt GetMessageReceiptById(ctx, messageId).Execute()

Information about a message



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
	messageId := "AB0815" // string | The id of the message

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommunicationAPI.GetMessageReceiptById(context.Background(), messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommunicationAPI.GetMessageReceiptById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessageReceiptById`: MessageReceipt
	fmt.Fprintf(os.Stdout, "Response from `CommunicationAPI.GetMessageReceiptById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | The id of the message | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageReceiptByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageReceipt**](MessageReceipt.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMail

> MessageReceipt PostMail(ctx).Message(message).Execute()

Send e-mail



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
	message := *openapiclient.NewMessage([]string{"recipient@example.com"}, "<h1>Example</h1>") // Message | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommunicationAPI.PostMail(context.Background()).Message(message).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommunicationAPI.PostMail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMail`: MessageReceipt
	fmt.Fprintf(os.Stdout, "Response from `CommunicationAPI.PostMail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **message** | [**Message**](Message.md) |  | 

### Return type

[**MessageReceipt**](MessageReceipt.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNotification

> MessageReceipt PostNotification(ctx).Notification(notification).Execute()

Send notification



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
	notification := *openapiclient.NewNotification("peter.fox@example.com", "TODO") // Notification | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommunicationAPI.PostNotification(context.Background()).Notification(notification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommunicationAPI.PostNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostNotification`: MessageReceipt
	fmt.Fprintf(os.Stdout, "Response from `CommunicationAPI.PostNotification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notification** | [**Notification**](Notification.md) |  | 

### Return type

[**MessageReceipt**](MessageReceipt.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


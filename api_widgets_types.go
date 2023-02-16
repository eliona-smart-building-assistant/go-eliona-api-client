/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.4.9
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// WidgetsTypesApiService WidgetsTypesApi service
type WidgetsTypesApiService service

type ApiDeleteWidgetTypeByNameRequest struct {
	ctx            context.Context
	ApiService     *WidgetsTypesApiService
	widgetTypeName string
}

func (r ApiDeleteWidgetTypeByNameRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteWidgetTypeByNameExecute(r)
}

/*
DeleteWidgetTypeByName Delete a widget type

Deletes a widget type and the elements for this widget type.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param widgetTypeName The name of the widget type
 @return ApiDeleteWidgetTypeByNameRequest
*/
func (a *WidgetsTypesApiService) DeleteWidgetTypeByName(ctx context.Context, widgetTypeName string) ApiDeleteWidgetTypeByNameRequest {
	return ApiDeleteWidgetTypeByNameRequest{
		ApiService:     a,
		ctx:            ctx,
		widgetTypeName: widgetTypeName,
	}
}

// Execute executes the request
func (a *WidgetsTypesApiService) DeleteWidgetTypeByNameExecute(r ApiDeleteWidgetTypeByNameRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WidgetsTypesApiService.DeleteWidgetTypeByName")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/widget-types/{widget-type-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"widget-type-name"+"}", url.PathEscape(parameterValueToString(r.widgetTypeName, "widgetTypeName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetWidgetTypeByNameRequest struct {
	ctx            context.Context
	ApiService     *WidgetsTypesApiService
	widgetTypeName string
	expansions     *[]string
}

// List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;.
func (r ApiGetWidgetTypeByNameRequest) Expansions(expansions []string) ApiGetWidgetTypeByNameRequest {
	r.expansions = &expansions
	return r
}

func (r ApiGetWidgetTypeByNameRequest) Execute() (*WidgetType, *http.Response, error) {
	return r.ApiService.GetWidgetTypeByNameExecute(r)
}

/*
GetWidgetTypeByName Information about a widget type

Gets information about a widget type.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param widgetTypeName The name of the widget type
 @return ApiGetWidgetTypeByNameRequest
*/
func (a *WidgetsTypesApiService) GetWidgetTypeByName(ctx context.Context, widgetTypeName string) ApiGetWidgetTypeByNameRequest {
	return ApiGetWidgetTypeByNameRequest{
		ApiService:     a,
		ctx:            ctx,
		widgetTypeName: widgetTypeName,
	}
}

// Execute executes the request
//  @return WidgetType
func (a *WidgetsTypesApiService) GetWidgetTypeByNameExecute(r ApiGetWidgetTypeByNameRequest) (*WidgetType, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *WidgetType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WidgetsTypesApiService.GetWidgetTypeByName")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/widget-types/{widget-type-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"widget-type-name"+"}", url.PathEscape(parameterValueToString(r.widgetTypeName, "widgetTypeName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.expansions != nil {
		parameterAddToQuery(localVarQueryParams, "expansions", r.expansions, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetWidgetTypesRequest struct {
	ctx        context.Context
	ApiService *WidgetsTypesApiService
	expansions *[]string
}

// List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;.
func (r ApiGetWidgetTypesRequest) Expansions(expansions []string) ApiGetWidgetTypesRequest {
	r.expansions = &expansions
	return r
}

func (r ApiGetWidgetTypesRequest) Execute() ([]WidgetType, *http.Response, error) {
	return r.ApiService.GetWidgetTypesExecute(r)
}

/*
GetWidgetTypes List of widget types

Returns a list of widget types

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetWidgetTypesRequest
*/
func (a *WidgetsTypesApiService) GetWidgetTypes(ctx context.Context) ApiGetWidgetTypesRequest {
	return ApiGetWidgetTypesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return []WidgetType
func (a *WidgetsTypesApiService) GetWidgetTypesExecute(r ApiGetWidgetTypesRequest) ([]WidgetType, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []WidgetType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WidgetsTypesApiService.GetWidgetTypes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/widget-types"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.expansions != nil {
		parameterAddToQuery(localVarQueryParams, "expansions", r.expansions, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostWidgetTypeRequest struct {
	ctx        context.Context
	ApiService *WidgetsTypesApiService
	widgetType *WidgetType
	expansions *[]string
}

func (r ApiPostWidgetTypeRequest) WidgetType(widgetType WidgetType) ApiPostWidgetTypeRequest {
	r.widgetType = &widgetType
	return r
}

// List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;.
func (r ApiPostWidgetTypeRequest) Expansions(expansions []string) ApiPostWidgetTypeRequest {
	r.expansions = &expansions
	return r
}

func (r ApiPostWidgetTypeRequest) Execute() (*WidgetType, *http.Response, error) {
	return r.ApiService.PostWidgetTypeExecute(r)
}

/*
PostWidgetType Create a widget type

Create a new widget type.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostWidgetTypeRequest
*/
func (a *WidgetsTypesApiService) PostWidgetType(ctx context.Context) ApiPostWidgetTypeRequest {
	return ApiPostWidgetTypeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return WidgetType
func (a *WidgetsTypesApiService) PostWidgetTypeExecute(r ApiPostWidgetTypeRequest) (*WidgetType, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *WidgetType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WidgetsTypesApiService.PostWidgetType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/widget-types"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.widgetType == nil {
		return localVarReturnValue, nil, reportError("widgetType is required and must be specified")
	}

	if r.expansions != nil {
		parameterAddToQuery(localVarQueryParams, "expansions", r.expansions, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.widgetType
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPutWidgetTypeRequest struct {
	ctx        context.Context
	ApiService *WidgetsTypesApiService
	widgetType *WidgetType
	expansions *[]string
}

func (r ApiPutWidgetTypeRequest) WidgetType(widgetType WidgetType) ApiPutWidgetTypeRequest {
	r.widgetType = &widgetType
	return r
}

// List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;.
func (r ApiPutWidgetTypeRequest) Expansions(expansions []string) ApiPutWidgetTypeRequest {
	r.expansions = &expansions
	return r
}

func (r ApiPutWidgetTypeRequest) Execute() (*WidgetType, *http.Response, error) {
	return r.ApiService.PutWidgetTypeExecute(r)
}

/*
PutWidgetType Create or update a widget type

Create a new widget type or update it if already exists. Uses the unique widget type name for updating.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPutWidgetTypeRequest
*/
func (a *WidgetsTypesApiService) PutWidgetType(ctx context.Context) ApiPutWidgetTypeRequest {
	return ApiPutWidgetTypeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return WidgetType
func (a *WidgetsTypesApiService) PutWidgetTypeExecute(r ApiPutWidgetTypeRequest) (*WidgetType, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *WidgetType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WidgetsTypesApiService.PutWidgetType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/widget-types"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.widgetType == nil {
		return localVarReturnValue, nil, reportError("widgetType is required and must be specified")
	}

	if r.expansions != nil {
		parameterAddToQuery(localVarQueryParams, "expansions", r.expansions, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.widgetType
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPutWidgetTypeByNameRequest struct {
	ctx            context.Context
	ApiService     *WidgetsTypesApiService
	widgetTypeName string
	widgetType     *WidgetType
	expansions     *[]string
}

func (r ApiPutWidgetTypeByNameRequest) WidgetType(widgetType WidgetType) ApiPutWidgetTypeByNameRequest {
	r.widgetType = &widgetType
	return r
}

// List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;.
func (r ApiPutWidgetTypeByNameRequest) Expansions(expansions []string) ApiPutWidgetTypeByNameRequest {
	r.expansions = &expansions
	return r
}

func (r ApiPutWidgetTypeByNameRequest) Execute() (*WidgetType, *http.Response, error) {
	return r.ApiService.PutWidgetTypeByNameExecute(r)
}

/*
PutWidgetTypeByName Update a widget type

Update a widget type.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param widgetTypeName The name of the widget type
 @return ApiPutWidgetTypeByNameRequest
*/
func (a *WidgetsTypesApiService) PutWidgetTypeByName(ctx context.Context, widgetTypeName string) ApiPutWidgetTypeByNameRequest {
	return ApiPutWidgetTypeByNameRequest{
		ApiService:     a,
		ctx:            ctx,
		widgetTypeName: widgetTypeName,
	}
}

// Execute executes the request
//  @return WidgetType
func (a *WidgetsTypesApiService) PutWidgetTypeByNameExecute(r ApiPutWidgetTypeByNameRequest) (*WidgetType, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *WidgetType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WidgetsTypesApiService.PutWidgetTypeByName")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/widget-types/{widget-type-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"widget-type-name"+"}", url.PathEscape(parameterValueToString(r.widgetTypeName, "widgetTypeName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.widgetType == nil {
		return localVarReturnValue, nil, reportError("widgetType is required and must be specified")
	}

	if r.expansions != nil {
		parameterAddToQuery(localVarQueryParams, "expansions", r.expansions, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.widgetType
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

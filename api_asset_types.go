/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.5.5
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// AssetTypesAPIService AssetTypesAPI service
type AssetTypesAPIService service

type ApiDeleteAssetTypeByNameRequest struct {
	ctx           context.Context
	ApiService    *AssetTypesAPIService
	assetTypeName string
}

func (r ApiDeleteAssetTypeByNameRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAssetTypeByNameExecute(r)
}

/*
DeleteAssetTypeByName Delete an asset type

Deletes an asset type and the attributes for this asset type.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param assetTypeName The name of the asset type
 @return ApiDeleteAssetTypeByNameRequest
*/
func (a *AssetTypesAPIService) DeleteAssetTypeByName(ctx context.Context, assetTypeName string) ApiDeleteAssetTypeByNameRequest {
	return ApiDeleteAssetTypeByNameRequest{
		ApiService:    a,
		ctx:           ctx,
		assetTypeName: assetTypeName,
	}
}

// Execute executes the request
func (a *AssetTypesAPIService) DeleteAssetTypeByNameExecute(r ApiDeleteAssetTypeByNameRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssetTypesAPIService.DeleteAssetTypeByName")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/asset-types/{asset-type-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"asset-type-name"+"}", url.PathEscape(parameterValueToString(r.assetTypeName, "assetTypeName")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetAssetTypeByNameRequest struct {
	ctx           context.Context
	ApiService    *AssetTypesAPIService
	assetTypeName string
	expansions    *[]string
}

// List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;.
func (r ApiGetAssetTypeByNameRequest) Expansions(expansions []string) ApiGetAssetTypeByNameRequest {
	r.expansions = &expansions
	return r
}

func (r ApiGetAssetTypeByNameRequest) Execute() (*AssetType, *http.Response, error) {
	return r.ApiService.GetAssetTypeByNameExecute(r)
}

/*
GetAssetTypeByName Information about an asset type

Gets information about an asset type.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param assetTypeName The name of the asset type
 @return ApiGetAssetTypeByNameRequest
*/
func (a *AssetTypesAPIService) GetAssetTypeByName(ctx context.Context, assetTypeName string) ApiGetAssetTypeByNameRequest {
	return ApiGetAssetTypeByNameRequest{
		ApiService:    a,
		ctx:           ctx,
		assetTypeName: assetTypeName,
	}
}

// Execute executes the request
//  @return AssetType
func (a *AssetTypesAPIService) GetAssetTypeByNameExecute(r ApiGetAssetTypeByNameRequest) (*AssetType, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AssetType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssetTypesAPIService.GetAssetTypeByName")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/asset-types/{asset-type-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"asset-type-name"+"}", url.PathEscape(parameterValueToString(r.assetTypeName, "assetTypeName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.expansions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "expansions", r.expansions, "csv")
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetAssetTypesRequest struct {
	ctx        context.Context
	ApiService *AssetTypesAPIService
	expansions *[]string
}

// List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;.
func (r ApiGetAssetTypesRequest) Expansions(expansions []string) ApiGetAssetTypesRequest {
	r.expansions = &expansions
	return r
}

func (r ApiGetAssetTypesRequest) Execute() ([]AssetType, *http.Response, error) {
	return r.ApiService.GetAssetTypesExecute(r)
}

/*
GetAssetTypes List of asset types

Returns a list of asset types

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAssetTypesRequest
*/
func (a *AssetTypesAPIService) GetAssetTypes(ctx context.Context) ApiGetAssetTypesRequest {
	return ApiGetAssetTypesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return []AssetType
func (a *AssetTypesAPIService) GetAssetTypesExecute(r ApiGetAssetTypesRequest) ([]AssetType, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []AssetType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssetTypesAPIService.GetAssetTypes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/asset-types"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.expansions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "expansions", r.expansions, "csv")
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiPostAssetTypeRequest struct {
	ctx        context.Context
	ApiService *AssetTypesAPIService
	assetType  *AssetType
	expansions *[]string
}

func (r ApiPostAssetTypeRequest) AssetType(assetType AssetType) ApiPostAssetTypeRequest {
	r.assetType = &assetType
	return r
}

// List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;.
func (r ApiPostAssetTypeRequest) Expansions(expansions []string) ApiPostAssetTypeRequest {
	r.expansions = &expansions
	return r
}

func (r ApiPostAssetTypeRequest) Execute() (*AssetType, *http.Response, error) {
	return r.ApiService.PostAssetTypeExecute(r)
}

/*
PostAssetType Create an asset type

Create a new asset type.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostAssetTypeRequest
*/
func (a *AssetTypesAPIService) PostAssetType(ctx context.Context) ApiPostAssetTypeRequest {
	return ApiPostAssetTypeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return AssetType
func (a *AssetTypesAPIService) PostAssetTypeExecute(r ApiPostAssetTypeRequest) (*AssetType, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AssetType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssetTypesAPIService.PostAssetType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/asset-types"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.assetType == nil {
		return localVarReturnValue, nil, reportError("assetType is required and must be specified")
	}

	if r.expansions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "expansions", r.expansions, "csv")
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
	localVarPostBody = r.assetType
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiPostAssetTypeAttributeRequest struct {
	ctx                context.Context
	ApiService         *AssetTypesAPIService
	assetTypeName      string
	assetTypeAttribute *AssetTypeAttribute
}

func (r ApiPostAssetTypeAttributeRequest) AssetTypeAttribute(assetTypeAttribute AssetTypeAttribute) ApiPostAssetTypeAttributeRequest {
	r.assetTypeAttribute = &assetTypeAttribute
	return r
}

func (r ApiPostAssetTypeAttributeRequest) Execute() (*AssetTypeAttribute, *http.Response, error) {
	return r.ApiService.PostAssetTypeAttributeExecute(r)
}

/*
PostAssetTypeAttribute Create asset type attribute

Create a new asset type attribute.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param assetTypeName The name of the asset type
 @return ApiPostAssetTypeAttributeRequest
*/
func (a *AssetTypesAPIService) PostAssetTypeAttribute(ctx context.Context, assetTypeName string) ApiPostAssetTypeAttributeRequest {
	return ApiPostAssetTypeAttributeRequest{
		ApiService:    a,
		ctx:           ctx,
		assetTypeName: assetTypeName,
	}
}

// Execute executes the request
//  @return AssetTypeAttribute
func (a *AssetTypesAPIService) PostAssetTypeAttributeExecute(r ApiPostAssetTypeAttributeRequest) (*AssetTypeAttribute, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AssetTypeAttribute
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssetTypesAPIService.PostAssetTypeAttribute")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/asset-types/{asset-type-name}/attributes"
	localVarPath = strings.Replace(localVarPath, "{"+"asset-type-name"+"}", url.PathEscape(parameterValueToString(r.assetTypeName, "assetTypeName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.assetTypeAttribute == nil {
		return localVarReturnValue, nil, reportError("assetTypeAttribute is required and must be specified")
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
	localVarPostBody = r.assetTypeAttribute
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiPutAssetTypeRequest struct {
	ctx        context.Context
	ApiService *AssetTypesAPIService
	assetType  *AssetType
	expansions *[]string
}

func (r ApiPutAssetTypeRequest) AssetType(assetType AssetType) ApiPutAssetTypeRequest {
	r.assetType = &assetType
	return r
}

// List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;.
func (r ApiPutAssetTypeRequest) Expansions(expansions []string) ApiPutAssetTypeRequest {
	r.expansions = &expansions
	return r
}

func (r ApiPutAssetTypeRequest) Execute() (*AssetType, *http.Response, error) {
	return r.ApiService.PutAssetTypeExecute(r)
}

/*
PutAssetType Create or update an asset type

Create a new asset type or update an asset type if already exists. Uses the unique asset type name for updating.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPutAssetTypeRequest
*/
func (a *AssetTypesAPIService) PutAssetType(ctx context.Context) ApiPutAssetTypeRequest {
	return ApiPutAssetTypeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return AssetType
func (a *AssetTypesAPIService) PutAssetTypeExecute(r ApiPutAssetTypeRequest) (*AssetType, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AssetType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssetTypesAPIService.PutAssetType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/asset-types"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.assetType == nil {
		return localVarReturnValue, nil, reportError("assetType is required and must be specified")
	}

	if r.expansions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "expansions", r.expansions, "csv")
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
	localVarPostBody = r.assetType
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiPutAssetTypeAttributeRequest struct {
	ctx                context.Context
	ApiService         *AssetTypesAPIService
	assetTypeName      string
	assetTypeAttribute *AssetTypeAttribute
}

func (r ApiPutAssetTypeAttributeRequest) AssetTypeAttribute(assetTypeAttribute AssetTypeAttribute) ApiPutAssetTypeAttributeRequest {
	r.assetTypeAttribute = &assetTypeAttribute
	return r
}

func (r ApiPutAssetTypeAttributeRequest) Execute() (*AssetTypeAttribute, *http.Response, error) {
	return r.ApiService.PutAssetTypeAttributeExecute(r)
}

/*
PutAssetTypeAttribute Create or update an asset type attribute

Create a new asset type attribute or update an asset type attribute if already exists. Uses the unique combination of asset type name, subtype and attribute name for updating.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param assetTypeName The name of the asset type
 @return ApiPutAssetTypeAttributeRequest
*/
func (a *AssetTypesAPIService) PutAssetTypeAttribute(ctx context.Context, assetTypeName string) ApiPutAssetTypeAttributeRequest {
	return ApiPutAssetTypeAttributeRequest{
		ApiService:    a,
		ctx:           ctx,
		assetTypeName: assetTypeName,
	}
}

// Execute executes the request
//  @return AssetTypeAttribute
func (a *AssetTypesAPIService) PutAssetTypeAttributeExecute(r ApiPutAssetTypeAttributeRequest) (*AssetTypeAttribute, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AssetTypeAttribute
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssetTypesAPIService.PutAssetTypeAttribute")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/asset-types/{asset-type-name}/attributes"
	localVarPath = strings.Replace(localVarPath, "{"+"asset-type-name"+"}", url.PathEscape(parameterValueToString(r.assetTypeName, "assetTypeName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.assetTypeAttribute == nil {
		return localVarReturnValue, nil, reportError("assetTypeAttribute is required and must be specified")
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
	localVarPostBody = r.assetTypeAttribute
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiPutAssetTypeByNameRequest struct {
	ctx           context.Context
	ApiService    *AssetTypesAPIService
	assetTypeName string
	assetType     *AssetType
	expansions    *[]string
}

func (r ApiPutAssetTypeByNameRequest) AssetType(assetType AssetType) ApiPutAssetTypeByNameRequest {
	r.assetType = &assetType
	return r
}

// List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;.
func (r ApiPutAssetTypeByNameRequest) Expansions(expansions []string) ApiPutAssetTypeByNameRequest {
	r.expansions = &expansions
	return r
}

func (r ApiPutAssetTypeByNameRequest) Execute() (*AssetType, *http.Response, error) {
	return r.ApiService.PutAssetTypeByNameExecute(r)
}

/*
PutAssetTypeByName Update an asset type

Update an asset type.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param assetTypeName The name of the asset type
 @return ApiPutAssetTypeByNameRequest
*/
func (a *AssetTypesAPIService) PutAssetTypeByName(ctx context.Context, assetTypeName string) ApiPutAssetTypeByNameRequest {
	return ApiPutAssetTypeByNameRequest{
		ApiService:    a,
		ctx:           ctx,
		assetTypeName: assetTypeName,
	}
}

// Execute executes the request
//  @return AssetType
func (a *AssetTypesAPIService) PutAssetTypeByNameExecute(r ApiPutAssetTypeByNameRequest) (*AssetType, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AssetType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssetTypesAPIService.PutAssetTypeByName")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/asset-types/{asset-type-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"asset-type-name"+"}", url.PathEscape(parameterValueToString(r.assetTypeName, "assetTypeName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.assetType == nil {
		return localVarReturnValue, nil, reportError("assetType is required and must be specified")
	}

	if r.expansions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "expansions", r.expansions, "csv")
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
	localVarPostBody = r.assetType
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

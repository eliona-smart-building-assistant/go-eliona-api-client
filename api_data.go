/*
Eliona API

API to access Eliona Smart Building Assistant

API version: 2.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)


// DataApiService DataApi service
type DataApiService service

type ApiGetAggregatedDataRequest struct {
	ctx context.Context
	ApiService *DataApiService
	from *time.Time
	until *time.Time
	assetId *int32
	dataSubtype *string
	assetTypeName *string
}

// Filter by lower date time limit
func (r ApiGetAggregatedDataRequest) From(from time.Time) ApiGetAggregatedDataRequest {
	r.from = &from
	return r
}

// Filter by upper date time limit
func (r ApiGetAggregatedDataRequest) Until(until time.Time) ApiGetAggregatedDataRequest {
	r.until = &until
	return r
}

// Filter for a specific asset id
func (r ApiGetAggregatedDataRequest) AssetId(assetId int32) ApiGetAggregatedDataRequest {
	r.assetId = &assetId
	return r
}

// Filter for a specific type of asset data
func (r ApiGetAggregatedDataRequest) DataSubtype(dataSubtype string) ApiGetAggregatedDataRequest {
	r.dataSubtype = &dataSubtype
	return r
}

// Filter the name of the asset type
func (r ApiGetAggregatedDataRequest) AssetTypeName(assetTypeName string) ApiGetAggregatedDataRequest {
	r.assetTypeName = &assetTypeName
	return r
}

func (r ApiGetAggregatedDataRequest) Execute() ([]AggregatedData, *http.Response, error) {
	return r.ApiService.GetAggregatedDataExecute(r)
}

/*
GetAggregatedData Get aggregated data

Gets aggregated data sets which combines a set of data points for a defined periodical raster

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAggregatedDataRequest
*/
func (a *DataApiService) GetAggregatedData(ctx context.Context) ApiGetAggregatedDataRequest {
	return ApiGetAggregatedDataRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []AggregatedData
func (a *DataApiService) GetAggregatedDataExecute(r ApiGetAggregatedDataRequest) ([]AggregatedData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []AggregatedData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataApiService.GetAggregatedData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/aggregated-data"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.from != nil {
		localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	}
	if r.until != nil {
		localVarQueryParams.Add("until", parameterToString(*r.until, ""))
	}
	if r.assetId != nil {
		localVarQueryParams.Add("assetId", parameterToString(*r.assetId, ""))
	}
	if r.dataSubtype != nil {
		localVarQueryParams.Add("dataSubtype", parameterToString(*r.dataSubtype, ""))
	}
	if r.assetTypeName != nil {
		localVarQueryParams.Add("assetTypeName", parameterToString(*r.assetTypeName, ""))
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

type ApiGetDataRequest struct {
	ctx context.Context
	ApiService *DataApiService
	assetId *int32
	dataSubtype *string
	assetTypeName *string
}

// Filter for a specific asset id
func (r ApiGetDataRequest) AssetId(assetId int32) ApiGetDataRequest {
	r.assetId = &assetId
	return r
}

// Filter for a specific type of asset data
func (r ApiGetDataRequest) DataSubtype(dataSubtype string) ApiGetDataRequest {
	r.dataSubtype = &dataSubtype
	return r
}

// Filter the name of the asset type
func (r ApiGetDataRequest) AssetTypeName(assetTypeName string) ApiGetDataRequest {
	r.assetTypeName = &assetTypeName
	return r
}

func (r ApiGetDataRequest) Execute() ([]Data, *http.Response, error) {
	return r.ApiService.GetDataExecute(r)
}

/*
GetData Gets all data

Gets information about data for assets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetDataRequest
*/
func (a *DataApiService) GetData(ctx context.Context) ApiGetDataRequest {
	return ApiGetDataRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Data
func (a *DataApiService) GetDataExecute(r ApiGetDataRequest) ([]Data, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Data
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataApiService.GetData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/data"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.assetId != nil {
		localVarQueryParams.Add("assetId", parameterToString(*r.assetId, ""))
	}
	if r.dataSubtype != nil {
		localVarQueryParams.Add("dataSubtype", parameterToString(*r.dataSubtype, ""))
	}
	if r.assetTypeName != nil {
		localVarQueryParams.Add("assetTypeName", parameterToString(*r.assetTypeName, ""))
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

type ApiGetDataTrendsRequest struct {
	ctx context.Context
	ApiService *DataApiService
	from *time.Time
	until *time.Time
	assetId *int32
	dataSubtype *string
	assetTypeName *string
}

// Filter by lower date time limit
func (r ApiGetDataTrendsRequest) From(from time.Time) ApiGetDataTrendsRequest {
	r.from = &from
	return r
}

// Filter by upper date time limit
func (r ApiGetDataTrendsRequest) Until(until time.Time) ApiGetDataTrendsRequest {
	r.until = &until
	return r
}

// Filter for a specific asset id
func (r ApiGetDataTrendsRequest) AssetId(assetId int32) ApiGetDataTrendsRequest {
	r.assetId = &assetId
	return r
}

// Filter for a specific type of asset data
func (r ApiGetDataTrendsRequest) DataSubtype(dataSubtype string) ApiGetDataTrendsRequest {
	r.dataSubtype = &dataSubtype
	return r
}

// Filter the name of the asset type
func (r ApiGetDataTrendsRequest) AssetTypeName(assetTypeName string) ApiGetDataTrendsRequest {
	r.assetTypeName = &assetTypeName
	return r
}

func (r ApiGetDataTrendsRequest) Execute() ([]Data, *http.Response, error) {
	return r.ApiService.GetDataTrendsExecute(r)
}

/*
GetDataTrends Get trend of historical data

Gets trend information about historical data for assets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetDataTrendsRequest
*/
func (a *DataApiService) GetDataTrends(ctx context.Context) ApiGetDataTrendsRequest {
	return ApiGetDataTrendsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Data
func (a *DataApiService) GetDataTrendsExecute(r ApiGetDataTrendsRequest) ([]Data, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Data
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataApiService.GetDataTrends")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/data-trends"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.from != nil {
		localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	}
	if r.until != nil {
		localVarQueryParams.Add("until", parameterToString(*r.until, ""))
	}
	if r.assetId != nil {
		localVarQueryParams.Add("assetId", parameterToString(*r.assetId, ""))
	}
	if r.dataSubtype != nil {
		localVarQueryParams.Add("dataSubtype", parameterToString(*r.dataSubtype, ""))
	}
	if r.assetTypeName != nil {
		localVarQueryParams.Add("assetTypeName", parameterToString(*r.assetTypeName, ""))
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

type ApiListenDataRequest struct {
	ctx context.Context
	ApiService *DataApiService
	assetId *int32
	dataSubtype *string
}

// Filter for a specific asset id
func (r ApiListenDataRequest) AssetId(assetId int32) ApiListenDataRequest {
	r.assetId = &assetId
	return r
}

// Filter for a specific type of asset data
func (r ApiListenDataRequest) DataSubtype(dataSubtype string) ApiListenDataRequest {
	r.dataSubtype = &dataSubtype
	return r
}

func (r ApiListenDataRequest) Execute() (*Data, *http.Response, error) {
	return r.ApiService.ListenDataExecute(r)
}

/*
ListenData WebSocket connection for asset data changes

Open a WebSocket connection to get informed when new asset data is written or anything changes.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListenDataRequest
*/
func (a *DataApiService) ListenData(ctx context.Context) ApiListenDataRequest {
	return ApiListenDataRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Data
func (a *DataApiService) ListenDataExecute(r ApiListenDataRequest) (*Data, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Data
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataApiService.ListenData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/data-listener"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.assetId != nil {
		localVarQueryParams.Add("assetId", parameterToString(*r.assetId, ""))
	}
	if r.dataSubtype != nil {
		localVarQueryParams.Add("dataSubtype", parameterToString(*r.dataSubtype, ""))
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

type ApiPutDataRequest struct {
	ctx context.Context
	ApiService *DataApiService
	data *Data
}

func (r ApiPutDataRequest) Data(data Data) ApiPutDataRequest {
	r.data = &data
	return r
}

func (r ApiPutDataRequest) Execute() (*http.Response, error) {
	return r.ApiService.PutDataExecute(r)
}

/*
PutData Create or update asset data

Create new asset data or update data if already exists.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPutDataRequest
*/
func (a *DataApiService) PutData(ctx context.Context) ApiPutDataRequest {
	return ApiPutDataRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *DataApiService) PutDataExecute(r ApiPutDataRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataApiService.PutData")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/data"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.data == nil {
		return nil, reportError("data is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.data
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

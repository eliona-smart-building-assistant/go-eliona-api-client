/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.6.12
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
)

// DataAPIService DataAPI service
type DataAPIService service

type ApiGetDataRequest struct {
	ctx           context.Context
	ApiService    *DataAPIService
	assetId       *int32
	parentAssetId *int32
	dataSubtype   *string
	assetTypeName *string
}

// Filter for a specific asset id
func (r ApiGetDataRequest) AssetId(assetId int32) ApiGetDataRequest {
	r.assetId = &assetId
	return r
}

// Filter for a specific parent asset id
func (r ApiGetDataRequest) ParentAssetId(parentAssetId int32) ApiGetDataRequest {
	r.parentAssetId = &parentAssetId
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
func (a *DataAPIService) GetData(ctx context.Context) ApiGetDataRequest {
	return ApiGetDataRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []Data
func (a *DataAPIService) GetDataExecute(r ApiGetDataRequest) ([]Data, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []Data
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataAPIService.GetData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/data"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.assetId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "assetId", r.assetId, "")
	}
	if r.parentAssetId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "parentAssetId", r.parentAssetId, "")
	}
	if r.dataSubtype != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dataSubtype", r.dataSubtype, "")
	}
	if r.assetTypeName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "assetTypeName", r.assetTypeName, "")
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

type ApiGetDataAggregatedRequest struct {
	ctx               context.Context
	ApiService        *DataAPIService
	fromDate          *string
	toDate            *string
	assetId           *int32
	dataSubtype       *string
	assetTypeName     *string
	aggregationId     *int32
	aggregationRaster *string
}

// Filter by lower date time (RFC3339) limit inclusive
func (r ApiGetDataAggregatedRequest) FromDate(fromDate string) ApiGetDataAggregatedRequest {
	r.fromDate = &fromDate
	return r
}

// Filter by upper date time (RFC3339) limit exclusive
func (r ApiGetDataAggregatedRequest) ToDate(toDate string) ApiGetDataAggregatedRequest {
	r.toDate = &toDate
	return r
}

// Filter for a specific asset id
func (r ApiGetDataAggregatedRequest) AssetId(assetId int32) ApiGetDataAggregatedRequest {
	r.assetId = &assetId
	return r
}

// Filter for a specific type of asset data
func (r ApiGetDataAggregatedRequest) DataSubtype(dataSubtype string) ApiGetDataAggregatedRequest {
	r.dataSubtype = &dataSubtype
	return r
}

// Filter the name of the asset type
func (r ApiGetDataAggregatedRequest) AssetTypeName(assetTypeName string) ApiGetDataAggregatedRequest {
	r.assetTypeName = &assetTypeName
	return r
}

// Filter for a specific aggregation id
func (r ApiGetDataAggregatedRequest) AggregationId(aggregationId int32) ApiGetDataAggregatedRequest {
	r.aggregationId = &aggregationId
	return r
}

// Aggregation calculation interval
func (r ApiGetDataAggregatedRequest) AggregationRaster(aggregationRaster string) ApiGetDataAggregatedRequest {
	r.aggregationRaster = &aggregationRaster
	return r
}

func (r ApiGetDataAggregatedRequest) Execute() ([]DataAggregated, *http.Response, error) {
	return r.ApiService.GetDataAggregatedExecute(r)
}

/*
GetDataAggregated Get aggregated data

Gets aggregated data sets which combines a set of data points for a defined periodical raster

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetDataAggregatedRequest
*/
func (a *DataAPIService) GetDataAggregated(ctx context.Context) ApiGetDataAggregatedRequest {
	return ApiGetDataAggregatedRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []DataAggregated
func (a *DataAPIService) GetDataAggregatedExecute(r ApiGetDataAggregatedRequest) ([]DataAggregated, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []DataAggregated
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataAPIService.GetDataAggregated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/data-aggregated"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fromDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fromDate", r.fromDate, "")
	}
	if r.toDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "toDate", r.toDate, "")
	}
	if r.assetId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "assetId", r.assetId, "")
	}
	if r.dataSubtype != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dataSubtype", r.dataSubtype, "")
	}
	if r.assetTypeName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "assetTypeName", r.assetTypeName, "")
	}
	if r.aggregationId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "aggregationId", r.aggregationId, "")
	}
	if r.aggregationRaster != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "aggregationRaster", r.aggregationRaster, "")
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

type ApiGetDataTrendsRequest struct {
	ctx           context.Context
	ApiService    *DataAPIService
	fromDate      *string
	toDate        *string
	assetId       *int32
	dataSubtype   *string
	assetTypeName *string
}

// Filter by lower date time (RFC3339) limit inclusive
func (r ApiGetDataTrendsRequest) FromDate(fromDate string) ApiGetDataTrendsRequest {
	r.fromDate = &fromDate
	return r
}

// Filter by upper date time (RFC3339) limit exclusive
func (r ApiGetDataTrendsRequest) ToDate(toDate string) ApiGetDataTrendsRequest {
	r.toDate = &toDate
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
func (a *DataAPIService) GetDataTrends(ctx context.Context) ApiGetDataTrendsRequest {
	return ApiGetDataTrendsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []Data
func (a *DataAPIService) GetDataTrendsExecute(r ApiGetDataTrendsRequest) ([]Data, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []Data
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataAPIService.GetDataTrends")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/data-trends"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fromDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fromDate", r.fromDate, "")
	}
	if r.toDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "toDate", r.toDate, "")
	}
	if r.assetId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "assetId", r.assetId, "")
	}
	if r.dataSubtype != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dataSubtype", r.dataSubtype, "")
	}
	if r.assetTypeName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "assetTypeName", r.assetTypeName, "")
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

type ApiListenDataRequest struct {
	ctx         context.Context
	ApiService  *DataAPIService
	assetId     *int32
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

func (r ApiListenDataRequest) Execute() (*DataListen, *http.Response, error) {
	return r.ApiService.ListenDataExecute(r)
}

/*
ListenData WebSocket connection for asset data changes

Open a WebSocket connection to get informed when new asset data is written or anything changes.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListenDataRequest
*/
func (a *DataAPIService) ListenData(ctx context.Context) ApiListenDataRequest {
	return ApiListenDataRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DataListen
func (a *DataAPIService) ListenDataExecute(r ApiListenDataRequest) (*DataListen, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DataListen
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataAPIService.ListenData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/data-listener"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.assetId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "assetId", r.assetId, "")
	}
	if r.dataSubtype != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dataSubtype", r.dataSubtype, "")
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

type ApiPutBulkDataRequest struct {
	ctx        context.Context
	ApiService *DataAPIService
	data       *[]Data
	directMode *string
}

func (r ApiPutBulkDataRequest) Data(data []Data) ApiPutBulkDataRequest {
	r.data = &data
	return r
}

// Executes the operation directly without using other services.
func (r ApiPutBulkDataRequest) DirectMode(directMode string) ApiPutBulkDataRequest {
	r.directMode = &directMode
	return r
}

func (r ApiPutBulkDataRequest) Execute() (*http.Response, error) {
	return r.ApiService.PutBulkDataExecute(r)
}

/*
PutBulkData Create or update multiple asset data

Create multiple asset data or update data if already exists. Uses the unique combination of asset id and subtype for updating.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPutBulkDataRequest
*/
func (a *DataAPIService) PutBulkData(ctx context.Context) ApiPutBulkDataRequest {
	return ApiPutBulkDataRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *DataAPIService) PutBulkDataExecute(r ApiPutBulkDataRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataAPIService.PutBulkData")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/data-bulk"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.data == nil {
		return nil, reportError("data is required and must be specified")
	}

	if r.directMode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "directMode", r.directMode, "")
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

type ApiPutDataRequest struct {
	ctx        context.Context
	ApiService *DataAPIService
	data       *Data
	directMode *string
}

func (r ApiPutDataRequest) Data(data Data) ApiPutDataRequest {
	r.data = &data
	return r
}

// Executes the operation directly without using other services.
func (r ApiPutDataRequest) DirectMode(directMode string) ApiPutDataRequest {
	r.directMode = &directMode
	return r
}

func (r ApiPutDataRequest) Execute() (*http.Response, error) {
	return r.ApiService.PutDataExecute(r)
}

/*
PutData Create or update asset data

Create new asset data or update data if already exists. Uses the unique combination of asset id and subtype for updating.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPutDataRequest
*/
func (a *DataAPIService) PutData(ctx context.Context) ApiPutDataRequest {
	return ApiPutDataRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *DataAPIService) PutDataExecute(r ApiPutDataRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataAPIService.PutData")
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

	if r.directMode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "directMode", r.directMode, "")
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

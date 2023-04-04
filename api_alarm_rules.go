/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.4.12
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

// AlarmRulesApiService AlarmRulesApi service
type AlarmRulesApiService service

type ApiGetAlarmRuleByIdRequest struct {
	ctx         context.Context
	ApiService  *AlarmRulesApiService
	alarmRuleId int32
	expansions  *[]string
}

// List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;.
func (r ApiGetAlarmRuleByIdRequest) Expansions(expansions []string) ApiGetAlarmRuleByIdRequest {
	r.expansions = &expansions
	return r
}

func (r ApiGetAlarmRuleByIdRequest) Execute() (*AlarmRule, *http.Response, error) {
	return r.ApiService.GetAlarmRuleByIdExecute(r)
}

/*
GetAlarmRuleById Information about an alarm rule

Gets information about an alarm rule.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param alarmRuleId The id of the alarm rule
 @return ApiGetAlarmRuleByIdRequest
*/
func (a *AlarmRulesApiService) GetAlarmRuleById(ctx context.Context, alarmRuleId int32) ApiGetAlarmRuleByIdRequest {
	return ApiGetAlarmRuleByIdRequest{
		ApiService:  a,
		ctx:         ctx,
		alarmRuleId: alarmRuleId,
	}
}

// Execute executes the request
//  @return AlarmRule
func (a *AlarmRulesApiService) GetAlarmRuleByIdExecute(r ApiGetAlarmRuleByIdRequest) (*AlarmRule, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlarmRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlarmRulesApiService.GetAlarmRuleById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/alarm-rules/{alarm-rule-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"alarm-rule-id"+"}", url.PathEscape(parameterValueToString(r.alarmRuleId, "alarmRuleId")), -1)

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

type ApiGetAlarmRulesRequest struct {
	ctx        context.Context
	ApiService *AlarmRulesApiService
	expansions *[]string
}

// List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;.
func (r ApiGetAlarmRulesRequest) Expansions(expansions []string) ApiGetAlarmRulesRequest {
	r.expansions = &expansions
	return r
}

func (r ApiGetAlarmRulesRequest) Execute() ([]AlarmRule, *http.Response, error) {
	return r.ApiService.GetAlarmRulesExecute(r)
}

/*
GetAlarmRules Information about alarm rules

Gets information about alarm rules.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAlarmRulesRequest
*/
func (a *AlarmRulesApiService) GetAlarmRules(ctx context.Context) ApiGetAlarmRulesRequest {
	return ApiGetAlarmRulesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return []AlarmRule
func (a *AlarmRulesApiService) GetAlarmRulesExecute(r ApiGetAlarmRulesRequest) ([]AlarmRule, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []AlarmRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlarmRulesApiService.GetAlarmRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/alarm-rules"

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

type ApiPostAlarmRuleRequest struct {
	ctx        context.Context
	ApiService *AlarmRulesApiService
	alarmRule  *AlarmRule
}

func (r ApiPostAlarmRuleRequest) AlarmRule(alarmRule AlarmRule) ApiPostAlarmRuleRequest {
	r.alarmRule = &alarmRule
	return r
}

func (r ApiPostAlarmRuleRequest) Execute() (*AlarmRule, *http.Response, error) {
	return r.ApiService.PostAlarmRuleExecute(r)
}

/*
PostAlarmRule Create an alarm rule

Create a new alarm rule.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostAlarmRuleRequest
*/
func (a *AlarmRulesApiService) PostAlarmRule(ctx context.Context) ApiPostAlarmRuleRequest {
	return ApiPostAlarmRuleRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return AlarmRule
func (a *AlarmRulesApiService) PostAlarmRuleExecute(r ApiPostAlarmRuleRequest) (*AlarmRule, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlarmRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlarmRulesApiService.PostAlarmRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/alarm-rules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.alarmRule == nil {
		return localVarReturnValue, nil, reportError("alarmRule is required and must be specified")
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
	localVarPostBody = r.alarmRule
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

type ApiPutAlarmRuleRequest struct {
	ctx        context.Context
	ApiService *AlarmRulesApiService
	alarmRule  *AlarmRule
}

func (r ApiPutAlarmRuleRequest) AlarmRule(alarmRule AlarmRule) ApiPutAlarmRuleRequest {
	r.alarmRule = &alarmRule
	return r
}

func (r ApiPutAlarmRuleRequest) Execute() (*AlarmRule, *http.Response, error) {
	return r.ApiService.PutAlarmRuleExecute(r)
}

/*
PutAlarmRule Create or update an alarm rule

Deprecated - Use POST /alarm-rules to create and PUT /alarm-rules/{alarm-rule-id} to update

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPutAlarmRuleRequest

Deprecated
*/
func (a *AlarmRulesApiService) PutAlarmRule(ctx context.Context) ApiPutAlarmRuleRequest {
	return ApiPutAlarmRuleRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return AlarmRule
// Deprecated
func (a *AlarmRulesApiService) PutAlarmRuleExecute(r ApiPutAlarmRuleRequest) (*AlarmRule, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlarmRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlarmRulesApiService.PutAlarmRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/alarm-rules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.alarmRule == nil {
		return localVarReturnValue, nil, reportError("alarmRule is required and must be specified")
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
	localVarPostBody = r.alarmRule
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

type ApiPutAlarmRuleByIdRequest struct {
	ctx         context.Context
	ApiService  *AlarmRulesApiService
	alarmRuleId int32
	alarmRule   *AlarmRule
}

func (r ApiPutAlarmRuleByIdRequest) AlarmRule(alarmRule AlarmRule) ApiPutAlarmRuleByIdRequest {
	r.alarmRule = &alarmRule
	return r
}

func (r ApiPutAlarmRuleByIdRequest) Execute() (*AlarmRule, *http.Response, error) {
	return r.ApiService.PutAlarmRuleByIdExecute(r)
}

/*
PutAlarmRuleById Update an alarm rule

Update an alarm rule.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param alarmRuleId The id of the alarm rule
 @return ApiPutAlarmRuleByIdRequest
*/
func (a *AlarmRulesApiService) PutAlarmRuleById(ctx context.Context, alarmRuleId int32) ApiPutAlarmRuleByIdRequest {
	return ApiPutAlarmRuleByIdRequest{
		ApiService:  a,
		ctx:         ctx,
		alarmRuleId: alarmRuleId,
	}
}

// Execute executes the request
//  @return AlarmRule
func (a *AlarmRulesApiService) PutAlarmRuleByIdExecute(r ApiPutAlarmRuleByIdRequest) (*AlarmRule, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlarmRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlarmRulesApiService.PutAlarmRuleById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/alarm-rules/{alarm-rule-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"alarm-rule-id"+"}", url.PathEscape(parameterValueToString(r.alarmRuleId, "alarmRuleId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.alarmRule == nil {
		return localVarReturnValue, nil, reportError("alarmRule is required and must be specified")
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
	localVarPostBody = r.alarmRule
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

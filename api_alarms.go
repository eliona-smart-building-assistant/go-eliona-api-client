/*
Eliona API

API to access Eliona Smart Building Assistant

API version: 2.1.0
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


// AlarmsApiService AlarmsApi service
type AlarmsApiService service

type ApiGetAlarmByIdRequest struct {
	ctx context.Context
	ApiService *AlarmsApiService
	alarmRuleId int32
	expansions *[]string
}

// List of referenced data to load. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;.
func (r ApiGetAlarmByIdRequest) Expansions(expansions []string) ApiGetAlarmByIdRequest {
	r.expansions = &expansions
	return r
}

func (r ApiGetAlarmByIdRequest) Execute() (*Alarm, *http.Response, error) {
	return r.ApiService.GetAlarmByIdExecute(r)
}

/*
GetAlarmById Information about alarm

Gets information about alarm.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param alarmRuleId The id of the alarm rule
 @return ApiGetAlarmByIdRequest
*/
func (a *AlarmsApiService) GetAlarmById(ctx context.Context, alarmRuleId int32) ApiGetAlarmByIdRequest {
	return ApiGetAlarmByIdRequest{
		ApiService: a,
		ctx: ctx,
		alarmRuleId: alarmRuleId,
	}
}

// Execute executes the request
//  @return Alarm
func (a *AlarmsApiService) GetAlarmByIdExecute(r ApiGetAlarmByIdRequest) (*Alarm, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Alarm
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlarmsApiService.GetAlarmById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/alarms/{alarm-rule-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"alarm-rule-id"+"}", url.PathEscape(parameterToString(r.alarmRuleId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.expansions != nil {
		localVarQueryParams.Add("expansions", parameterToString(*r.expansions, "csv"))
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

type ApiGetAlarmHistoryByIdRequest struct {
	ctx context.Context
	ApiService *AlarmsApiService
	alarmRuleId int32
	expansions *[]string
}

// List of referenced data to load. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;.
func (r ApiGetAlarmHistoryByIdRequest) Expansions(expansions []string) ApiGetAlarmHistoryByIdRequest {
	r.expansions = &expansions
	return r
}

func (r ApiGetAlarmHistoryByIdRequest) Execute() ([]Alarm, *http.Response, error) {
	return r.ApiService.GetAlarmHistoryByIdExecute(r)
}

/*
GetAlarmHistoryById Information about alarm history

Gets information about alarm over the entire time. This includes current alarm and history.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param alarmRuleId The id of the alarm rule
 @return ApiGetAlarmHistoryByIdRequest
*/
func (a *AlarmsApiService) GetAlarmHistoryById(ctx context.Context, alarmRuleId int32) ApiGetAlarmHistoryByIdRequest {
	return ApiGetAlarmHistoryByIdRequest{
		ApiService: a,
		ctx: ctx,
		alarmRuleId: alarmRuleId,
	}
}

// Execute executes the request
//  @return []Alarm
func (a *AlarmsApiService) GetAlarmHistoryByIdExecute(r ApiGetAlarmHistoryByIdRequest) ([]Alarm, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Alarm
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlarmsApiService.GetAlarmHistoryById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/alarms-history/{alarm-rule-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"alarm-rule-id"+"}", url.PathEscape(parameterToString(r.alarmRuleId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.expansions != nil {
		localVarQueryParams.Add("expansions", parameterToString(*r.expansions, "csv"))
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

type ApiGetAlarmsRequest struct {
	ctx context.Context
	ApiService *AlarmsApiService
	projectId *string
	expansions *[]string
}

// Filters for a specific project
func (r ApiGetAlarmsRequest) ProjectId(projectId string) ApiGetAlarmsRequest {
	r.projectId = &projectId
	return r
}

// List of referenced data to load. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;.
func (r ApiGetAlarmsRequest) Expansions(expansions []string) ApiGetAlarmsRequest {
	r.expansions = &expansions
	return r
}

func (r ApiGetAlarmsRequest) Execute() ([]Alarm, *http.Response, error) {
	return r.ApiService.GetAlarmsExecute(r)
}

/*
GetAlarms Information about alarms

Gets information about alarms

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAlarmsRequest
*/
func (a *AlarmsApiService) GetAlarms(ctx context.Context) ApiGetAlarmsRequest {
	return ApiGetAlarmsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Alarm
func (a *AlarmsApiService) GetAlarmsExecute(r ApiGetAlarmsRequest) ([]Alarm, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Alarm
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlarmsApiService.GetAlarms")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/alarms"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.projectId != nil {
		localVarQueryParams.Add("projectId", parameterToString(*r.projectId, ""))
	}
	if r.expansions != nil {
		localVarQueryParams.Add("expansions", parameterToString(*r.expansions, "csv"))
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

type ApiGetAlarmsHistoryRequest struct {
	ctx context.Context
	ApiService *AlarmsApiService
	projectId *string
	expansions *[]string
}

// Filters for a specific project
func (r ApiGetAlarmsHistoryRequest) ProjectId(projectId string) ApiGetAlarmsHistoryRequest {
	r.projectId = &projectId
	return r
}

// List of referenced data to load. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;.
func (r ApiGetAlarmsHistoryRequest) Expansions(expansions []string) ApiGetAlarmsHistoryRequest {
	r.expansions = &expansions
	return r
}

func (r ApiGetAlarmsHistoryRequest) Execute() ([]Alarm, *http.Response, error) {
	return r.ApiService.GetAlarmsHistoryExecute(r)
}

/*
GetAlarmsHistory Information about alarms history

Gets information about alarms over the entire time. This includes current alarms and alarms, which are already processed.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAlarmsHistoryRequest
*/
func (a *AlarmsApiService) GetAlarmsHistory(ctx context.Context) ApiGetAlarmsHistoryRequest {
	return ApiGetAlarmsHistoryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Alarm
func (a *AlarmsApiService) GetAlarmsHistoryExecute(r ApiGetAlarmsHistoryRequest) ([]Alarm, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Alarm
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlarmsApiService.GetAlarmsHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/alarms-history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.projectId != nil {
		localVarQueryParams.Add("projectId", parameterToString(*r.projectId, ""))
	}
	if r.expansions != nil {
		localVarQueryParams.Add("expansions", parameterToString(*r.expansions, "csv"))
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

type ApiGetHighestAlarmsRequest struct {
	ctx context.Context
	ApiService *AlarmsApiService
	projectId *string
	expansions *[]string
}

// Filters for a specific project
func (r ApiGetHighestAlarmsRequest) ProjectId(projectId string) ApiGetHighestAlarmsRequest {
	r.projectId = &projectId
	return r
}

// List of referenced data to load. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;.
func (r ApiGetHighestAlarmsRequest) Expansions(expansions []string) ApiGetHighestAlarmsRequest {
	r.expansions = &expansions
	return r
}

func (r ApiGetHighestAlarmsRequest) Execute() ([]Alarm, *http.Response, error) {
	return r.ApiService.GetHighestAlarmsExecute(r)
}

/*
GetHighestAlarms Information about most prioritized alarms

Gets information about an alarms with the highest priority for each asset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetHighestAlarmsRequest
*/
func (a *AlarmsApiService) GetHighestAlarms(ctx context.Context) ApiGetHighestAlarmsRequest {
	return ApiGetHighestAlarmsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Alarm
func (a *AlarmsApiService) GetHighestAlarmsExecute(r ApiGetHighestAlarmsRequest) ([]Alarm, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Alarm
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlarmsApiService.GetHighestAlarms")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/alarms-highest"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.projectId != nil {
		localVarQueryParams.Add("projectId", parameterToString(*r.projectId, ""))
	}
	if r.expansions != nil {
		localVarQueryParams.Add("expansions", parameterToString(*r.expansions, "csv"))
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

type ApiPatchAlarmByIdRequest struct {
	ctx context.Context
	ApiService *AlarmsApiService
	alarmRuleId int32
	acknowledged *bool
	acknowledgeText *string
}

// Marks the alarm as acknowledged by setting the acknowledge timestamp to now.
func (r ApiPatchAlarmByIdRequest) Acknowledged(acknowledged bool) ApiPatchAlarmByIdRequest {
	r.acknowledged = &acknowledged
	return r
}

// Sets the text for acknowledgement
func (r ApiPatchAlarmByIdRequest) AcknowledgeText(acknowledgeText string) ApiPatchAlarmByIdRequest {
	r.acknowledgeText = &acknowledgeText
	return r
}

func (r ApiPatchAlarmByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PatchAlarmByIdExecute(r)
}

/*
PatchAlarmById Update alarm

Update properties of alarm for given id.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param alarmRuleId The id of the alarm rule
 @return ApiPatchAlarmByIdRequest
*/
func (a *AlarmsApiService) PatchAlarmById(ctx context.Context, alarmRuleId int32) ApiPatchAlarmByIdRequest {
	return ApiPatchAlarmByIdRequest{
		ApiService: a,
		ctx: ctx,
		alarmRuleId: alarmRuleId,
	}
}

// Execute executes the request
func (a *AlarmsApiService) PatchAlarmByIdExecute(r ApiPatchAlarmByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlarmsApiService.PatchAlarmById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/alarms/{alarm-rule-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"alarm-rule-id"+"}", url.PathEscape(parameterToString(r.alarmRuleId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.acknowledged == nil {
		return nil, reportError("acknowledged is required and must be specified")
	}

	localVarQueryParams.Add("acknowledged", parameterToString(*r.acknowledged, ""))
	if r.acknowledgeText != nil {
		localVarQueryParams.Add("acknowledgeText", parameterToString(*r.acknowledgeText, ""))
	}
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
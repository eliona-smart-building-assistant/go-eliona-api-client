/*
Eliona API

API to access Eliona Smart Building Assistant

API version: 2.4.4
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


// WidgetsApiService WidgetsApi service
type WidgetsApiService service

type ApiGetDashboardWidgetsRequest struct {
	ctx context.Context
	ApiService *WidgetsApiService
	dashboardId int32
	expansions *[]string
}

// List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;.
func (r ApiGetDashboardWidgetsRequest) Expansions(expansions []string) ApiGetDashboardWidgetsRequest {
	r.expansions = &expansions
	return r
}

func (r ApiGetDashboardWidgetsRequest) Execute() (*Widget, *http.Response, error) {
	return r.ApiService.GetDashboardWidgetsExecute(r)
}

/*
GetDashboardWidgets Information about widgets on dashboard

Gets information about widgets on a dashboard.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dashboardId The id of the dashboard
 @return ApiGetDashboardWidgetsRequest
*/
func (a *WidgetsApiService) GetDashboardWidgets(ctx context.Context, dashboardId int32) ApiGetDashboardWidgetsRequest {
	return ApiGetDashboardWidgetsRequest{
		ApiService: a,
		ctx: ctx,
		dashboardId: dashboardId,
	}
}

// Execute executes the request
//  @return Widget
func (a *WidgetsApiService) GetDashboardWidgetsExecute(r ApiGetDashboardWidgetsRequest) (*Widget, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Widget
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WidgetsApiService.GetDashboardWidgets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dashboards/{dashboard-id}/widgets"
	localVarPath = strings.Replace(localVarPath, "{"+"dashboard-id"+"}", url.PathEscape(parameterToString(r.dashboardId, "")), -1)

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

type ApiPostDashboardWidgetRequest struct {
	ctx context.Context
	ApiService *WidgetsApiService
	dashboardId int32
	widget *Widget
	expansions *[]string
}

func (r ApiPostDashboardWidgetRequest) Widget(widget Widget) ApiPostDashboardWidgetRequest {
	r.widget = &widget
	return r
}

// List of referenced data to load, insert or update. Each entry defines the full qualified name of the field to be expanded as follows &#39;ObjectName.fieldName&#39;.
func (r ApiPostDashboardWidgetRequest) Expansions(expansions []string) ApiPostDashboardWidgetRequest {
	r.expansions = &expansions
	return r
}

func (r ApiPostDashboardWidgetRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostDashboardWidgetExecute(r)
}

/*
PostDashboardWidget Adds widget to dashboard

Create a new widget and add this to a dashboard

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dashboardId The id of the dashboard
 @return ApiPostDashboardWidgetRequest
*/
func (a *WidgetsApiService) PostDashboardWidget(ctx context.Context, dashboardId int32) ApiPostDashboardWidgetRequest {
	return ApiPostDashboardWidgetRequest{
		ApiService: a,
		ctx: ctx,
		dashboardId: dashboardId,
	}
}

// Execute executes the request
func (a *WidgetsApiService) PostDashboardWidgetExecute(r ApiPostDashboardWidgetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WidgetsApiService.PostDashboardWidget")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dashboards/{dashboard-id}/widgets"
	localVarPath = strings.Replace(localVarPath, "{"+"dashboard-id"+"}", url.PathEscape(parameterToString(r.dashboardId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.widget == nil {
		return nil, reportError("widget is required and must be specified")
	}

	if r.expansions != nil {
		localVarQueryParams.Add("expansions", parameterToString(*r.expansions, "csv"))
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
	localVarPostBody = r.widget
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

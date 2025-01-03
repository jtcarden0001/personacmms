/*
PersonaCMMS API

This is the Personal Computer Maintenance Management System REST API.

API version: 1.0
Contact: greenrivercodelabs@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// UsageTriggersAPIService UsageTriggersAPI service
type UsageTriggersAPIService service

type ApiAssetsAssetIdTasksTaskIdUsageTriggersGetRequest struct {
	ctx context.Context
	ApiService *UsageTriggersAPIService
	assetId string
	taskId string
}

func (r ApiAssetsAssetIdTasksTaskIdUsageTriggersGetRequest) Execute() ([]TypesUsageTrigger, *http.Response, error) {
	return r.ApiService.AssetsAssetIdTasksTaskIdUsageTriggersGetExecute(r)
}

/*
AssetsAssetIdTasksTaskIdUsageTriggersGet List usage triggers

List usage triggers for a task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param assetId Asset Id
 @param taskId Asset Task Id
 @return ApiAssetsAssetIdTasksTaskIdUsageTriggersGetRequest
*/
func (a *UsageTriggersAPIService) AssetsAssetIdTasksTaskIdUsageTriggersGet(ctx context.Context, assetId string, taskId string) ApiAssetsAssetIdTasksTaskIdUsageTriggersGetRequest {
	return ApiAssetsAssetIdTasksTaskIdUsageTriggersGetRequest{
		ApiService: a,
		ctx: ctx,
		assetId: assetId,
		taskId: taskId,
	}
}

// Execute executes the request
//  @return []TypesUsageTrigger
func (a *UsageTriggersAPIService) AssetsAssetIdTasksTaskIdUsageTriggersGetExecute(r ApiAssetsAssetIdTasksTaskIdUsageTriggersGetRequest) ([]TypesUsageTrigger, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []TypesUsageTrigger
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageTriggersAPIService.AssetsAssetIdTasksTaskIdUsageTriggersGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/assets/{assetId}/tasks/{taskId}/usage-triggers"
	localVarPath = strings.Replace(localVarPath, "{"+"assetId"+"}", url.PathEscape(parameterValueToString(r.assetId, "assetId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"taskId"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiAssetsAssetIdTasksTaskIdUsageTriggersPostRequest struct {
	ctx context.Context
	ApiService *UsageTriggersAPIService
	assetId string
	taskId string
	typesUsageTrigger *TypesUsageTrigger
}

// Usage Trigger object
func (r ApiAssetsAssetIdTasksTaskIdUsageTriggersPostRequest) TypesUsageTrigger(typesUsageTrigger TypesUsageTrigger) ApiAssetsAssetIdTasksTaskIdUsageTriggersPostRequest {
	r.typesUsageTrigger = &typesUsageTrigger
	return r
}

func (r ApiAssetsAssetIdTasksTaskIdUsageTriggersPostRequest) Execute() (*TypesUsageTrigger, *http.Response, error) {
	return r.ApiService.AssetsAssetIdTasksTaskIdUsageTriggersPostExecute(r)
}

/*
AssetsAssetIdTasksTaskIdUsageTriggersPost Create a usage trigger

Create a usage trigger for a task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param assetId Asset Id
 @param taskId Asset Task Id
 @return ApiAssetsAssetIdTasksTaskIdUsageTriggersPostRequest
*/
func (a *UsageTriggersAPIService) AssetsAssetIdTasksTaskIdUsageTriggersPost(ctx context.Context, assetId string, taskId string) ApiAssetsAssetIdTasksTaskIdUsageTriggersPostRequest {
	return ApiAssetsAssetIdTasksTaskIdUsageTriggersPostRequest{
		ApiService: a,
		ctx: ctx,
		assetId: assetId,
		taskId: taskId,
	}
}

// Execute executes the request
//  @return TypesUsageTrigger
func (a *UsageTriggersAPIService) AssetsAssetIdTasksTaskIdUsageTriggersPostExecute(r ApiAssetsAssetIdTasksTaskIdUsageTriggersPostRequest) (*TypesUsageTrigger, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TypesUsageTrigger
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageTriggersAPIService.AssetsAssetIdTasksTaskIdUsageTriggersPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/assets/{assetId}/tasks/{taskId}/usage-triggers"
	localVarPath = strings.Replace(localVarPath, "{"+"assetId"+"}", url.PathEscape(parameterValueToString(r.assetId, "assetId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"taskId"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.typesUsageTrigger == nil {
		return localVarReturnValue, nil, reportError("typesUsageTrigger is required and must be specified")
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
	localVarPostBody = r.typesUsageTrigger
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiAssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdDeleteRequest struct {
	ctx context.Context
	ApiService *UsageTriggersAPIService
	assetId string
	taskId string
	usageTriggerId string
}

func (r ApiAssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdDeleteExecute(r)
}

/*
AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdDelete Delete a usage trigger

Delete a usage trigger

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param assetId Asset Id
 @param taskId Asset Task Id
 @param usageTriggerId Usage Trigger Id
 @return ApiAssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdDeleteRequest
*/
func (a *UsageTriggersAPIService) AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdDelete(ctx context.Context, assetId string, taskId string, usageTriggerId string) ApiAssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdDeleteRequest {
	return ApiAssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		assetId: assetId,
		taskId: taskId,
		usageTriggerId: usageTriggerId,
	}
}

// Execute executes the request
func (a *UsageTriggersAPIService) AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdDeleteExecute(r ApiAssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageTriggersAPIService.AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/assets/{assetId}/tasks/{taskId}/usage-triggers/{usageTriggerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"assetId"+"}", url.PathEscape(parameterValueToString(r.assetId, "assetId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"taskId"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"usageTriggerId"+"}", url.PathEscape(parameterValueToString(r.usageTriggerId, "usageTriggerId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"*/*"}

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiAssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdGetRequest struct {
	ctx context.Context
	ApiService *UsageTriggersAPIService
	assetId string
	taskId string
	usageTriggerId string
}

func (r ApiAssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdGetRequest) Execute() (*TypesUsageTrigger, *http.Response, error) {
	return r.ApiService.AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdGetExecute(r)
}

/*
AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdGet Get a usage trigger

Get a usage trigger

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param assetId Asset Id
 @param taskId Asset Task Id
 @param usageTriggerId Usage Trigger Id
 @return ApiAssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdGetRequest
*/
func (a *UsageTriggersAPIService) AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdGet(ctx context.Context, assetId string, taskId string, usageTriggerId string) ApiAssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdGetRequest {
	return ApiAssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdGetRequest{
		ApiService: a,
		ctx: ctx,
		assetId: assetId,
		taskId: taskId,
		usageTriggerId: usageTriggerId,
	}
}

// Execute executes the request
//  @return TypesUsageTrigger
func (a *UsageTriggersAPIService) AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdGetExecute(r ApiAssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdGetRequest) (*TypesUsageTrigger, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TypesUsageTrigger
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageTriggersAPIService.AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/assets/{assetId}/tasks/{taskId}/usage-triggers/{usageTriggerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"assetId"+"}", url.PathEscape(parameterValueToString(r.assetId, "assetId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"taskId"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"usageTriggerId"+"}", url.PathEscape(parameterValueToString(r.usageTriggerId, "usageTriggerId")), -1)

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiAssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdPutRequest struct {
	ctx context.Context
	ApiService *UsageTriggersAPIService
	assetId string
	taskId string
	usageTriggerId string
	typesUsageTrigger *TypesUsageTrigger
}

// Usage Trigger object
func (r ApiAssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdPutRequest) TypesUsageTrigger(typesUsageTrigger TypesUsageTrigger) ApiAssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdPutRequest {
	r.typesUsageTrigger = &typesUsageTrigger
	return r
}

func (r ApiAssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdPutRequest) Execute() (*TypesUsageTrigger, *http.Response, error) {
	return r.ApiService.AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdPutExecute(r)
}

/*
AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdPut Update a usage trigger

Update a usage trigger

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param assetId Asset Id
 @param taskId Asset Task Id
 @param usageTriggerId Usage Trigger Id
 @return ApiAssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdPutRequest
*/
func (a *UsageTriggersAPIService) AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdPut(ctx context.Context, assetId string, taskId string, usageTriggerId string) ApiAssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdPutRequest {
	return ApiAssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdPutRequest{
		ApiService: a,
		ctx: ctx,
		assetId: assetId,
		taskId: taskId,
		usageTriggerId: usageTriggerId,
	}
}

// Execute executes the request
//  @return TypesUsageTrigger
func (a *UsageTriggersAPIService) AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdPutExecute(r ApiAssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdPutRequest) (*TypesUsageTrigger, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TypesUsageTrigger
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageTriggersAPIService.AssetsAssetIdTasksTaskIdUsageTriggersUsageTriggerIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/assets/{assetId}/tasks/{taskId}/usage-triggers/{usageTriggerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"assetId"+"}", url.PathEscape(parameterValueToString(r.assetId, "assetId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"taskId"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"usageTriggerId"+"}", url.PathEscape(parameterValueToString(r.usageTriggerId, "usageTriggerId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.typesUsageTrigger == nil {
		return localVarReturnValue, nil, reportError("typesUsageTrigger is required and must be specified")
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
	localVarPostBody = r.typesUsageTrigger
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

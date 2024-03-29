/*
SIRCLO Open API

SIRCLO Open API

API version: 1.0.0
Contact: dev@sirclo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sirclo

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type SalesChannelAPI interface {

	/*
	RegisterSalesChannel Method for RegisterSalesChannel

	Register sales channel

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiRegisterSalesChannelRequest
	*/
	RegisterSalesChannel(ctx context.Context) ApiRegisterSalesChannelRequest

	// RegisterSalesChannelExecute executes the request
	//  @return RegisteredSalesChannel
	RegisterSalesChannelExecute(r ApiRegisterSalesChannelRequest) (*RegisteredSalesChannel, *http.Response, error)

	/*
	UpdateSalesChannel Method for UpdateSalesChannel

	Update Sales Channel

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelCode Channel code
	@return ApiUpdateSalesChannelRequest
	*/
	UpdateSalesChannel(ctx context.Context, channelCode string) ApiUpdateSalesChannelRequest

	// UpdateSalesChannelExecute executes the request
	//  @return SalesChannel
	UpdateSalesChannelExecute(r ApiUpdateSalesChannelRequest) (*SalesChannel, *http.Response, error)
}

// SalesChannelAPIService SalesChannelAPI service
type SalesChannelAPIService service

type ApiRegisterSalesChannelRequest struct {
	ctx context.Context
	ApiService SalesChannelAPI
	registerSalesChannelInput *RegisterSalesChannelInput
}

func (r ApiRegisterSalesChannelRequest) RegisterSalesChannelInput(registerSalesChannelInput RegisterSalesChannelInput) ApiRegisterSalesChannelRequest {
	r.registerSalesChannelInput = &registerSalesChannelInput
	return r
}

func (r ApiRegisterSalesChannelRequest) Execute() (*RegisteredSalesChannel, *http.Response, error) {
	return r.ApiService.RegisterSalesChannelExecute(r)
}

/*
RegisterSalesChannel Method for RegisterSalesChannel

Register sales channel

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRegisterSalesChannelRequest
*/
func (a *SalesChannelAPIService) RegisterSalesChannel(ctx context.Context) ApiRegisterSalesChannelRequest {
	return ApiRegisterSalesChannelRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return RegisteredSalesChannel
func (a *SalesChannelAPIService) RegisterSalesChannelExecute(r ApiRegisterSalesChannelRequest) (*RegisteredSalesChannel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RegisteredSalesChannel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SalesChannelAPIService.RegisterSalesChannel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/saleschannel/register"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.registerSalesChannelInput == nil {
		return localVarReturnValue, nil, reportError("registerSalesChannelInput is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/plain"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.registerSalesChannelInput
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
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiUpdateSalesChannelRequest struct {
	ctx context.Context
	ApiService SalesChannelAPI
	channelCode string
	updateSalesChannelInput *UpdateSalesChannelInput
}

func (r ApiUpdateSalesChannelRequest) UpdateSalesChannelInput(updateSalesChannelInput UpdateSalesChannelInput) ApiUpdateSalesChannelRequest {
	r.updateSalesChannelInput = &updateSalesChannelInput
	return r
}

func (r ApiUpdateSalesChannelRequest) Execute() (*SalesChannel, *http.Response, error) {
	return r.ApiService.UpdateSalesChannelExecute(r)
}

/*
UpdateSalesChannel Method for UpdateSalesChannel

Update Sales Channel

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channelCode Channel code
 @return ApiUpdateSalesChannelRequest
*/
func (a *SalesChannelAPIService) UpdateSalesChannel(ctx context.Context, channelCode string) ApiUpdateSalesChannelRequest {
	return ApiUpdateSalesChannelRequest{
		ApiService: a,
		ctx: ctx,
		channelCode: channelCode,
	}
}

// Execute executes the request
//  @return SalesChannel
func (a *SalesChannelAPIService) UpdateSalesChannelExecute(r ApiUpdateSalesChannelRequest) (*SalesChannel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SalesChannel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SalesChannelAPIService.UpdateSalesChannel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/saleschannel/update/{channelCode}"
	localVarPath = strings.Replace(localVarPath, "{"+"channelCode"+"}", url.PathEscape(parameterValueToString(r.channelCode, "channelCode")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateSalesChannelInput == nil {
		return localVarReturnValue, nil, reportError("updateSalesChannelInput is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/plain"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateSalesChannelInput
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["appKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
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
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

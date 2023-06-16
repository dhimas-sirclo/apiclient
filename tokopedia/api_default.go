/*
Tokopedia

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
Contact: dev@sirclo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tokopedia

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type DefaultApi interface {

	/*
	InventoryV1FsFsIdProductInfoGet Method for InventoryV1FsFsIdProductInfoGet

	This method will retrieve single product information either by product id as parameter (choose one of those two parameters to use) from related fs_id.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param fsId Fulfillment service unique identifier
	@return ApiInventoryV1FsFsIdProductInfoGetRequest
	*/
	InventoryV1FsFsIdProductInfoGet(ctx context.Context, fsId int64) ApiInventoryV1FsFsIdProductInfoGetRequest

	// InventoryV1FsFsIdProductInfoGetExecute executes the request
	//  @return GetProductInfoResponse
	InventoryV1FsFsIdProductInfoGetExecute(r ApiInventoryV1FsFsIdProductInfoGetRequest) (*GetProductInfoResponse, *http.Response, error)

	/*
	InventoryV1FsFsIdProductVariantProductIdGet Method for InventoryV1FsFsIdProductVariantProductIdGet

	This endpoint retrieves a list of variants related to a product_id.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param fsId Fulfillment service unique identifier
	@param productId Product unique identifier
	@return ApiInventoryV1FsFsIdProductVariantProductIdGetRequest
	*/
	InventoryV1FsFsIdProductVariantProductIdGet(ctx context.Context, fsId int64, productId int64) ApiInventoryV1FsFsIdProductVariantProductIdGetRequest

	// InventoryV1FsFsIdProductVariantProductIdGetExecute executes the request
	//  @return GetProductVariantResponse
	InventoryV1FsFsIdProductVariantProductIdGetExecute(r ApiInventoryV1FsFsIdProductVariantProductIdGetRequest) (*GetProductVariantResponse, *http.Response, error)

	/*
	InventoryV2FsFsIdCategoryGetVariantGet Method for InventoryV2FsFsIdCategoryGetVariantGet

	This endpoint retrieves a list of variants related to a category_id. Use this API as main source to retrieve the newest category variant data.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param fsId Fulfillment service unique identifier
	@return ApiInventoryV2FsFsIdCategoryGetVariantGetRequest
	*/
	InventoryV2FsFsIdCategoryGetVariantGet(ctx context.Context, fsId int64) ApiInventoryV2FsFsIdCategoryGetVariantGetRequest

	// InventoryV2FsFsIdCategoryGetVariantGetExecute executes the request
	//  @return GetCategoryVariantResponse
	InventoryV2FsFsIdCategoryGetVariantGetExecute(r ApiInventoryV2FsFsIdCategoryGetVariantGetRequest) (*GetCategoryVariantResponse, *http.Response, error)
}

// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiInventoryV1FsFsIdProductInfoGetRequest struct {
	ctx context.Context
	ApiService DefaultApi
	fsId int64
	product *string
	productUrl *string
	sku *string
	shopId *int64
	page *int64
	perPage *int64
	lastSort *string
}

// Can input more than one product_id
func (r ApiInventoryV1FsFsIdProductInfoGetRequest) Product(product string) ApiInventoryV1FsFsIdProductInfoGetRequest {
	r.product = &product
	return r
}

// Can input more than one product_url
func (r ApiInventoryV1FsFsIdProductInfoGetRequest) ProductUrl(productUrl string) ApiInventoryV1FsFsIdProductInfoGetRequest {
	r.productUrl = &productUrl
	return r
}

// Product’s SKU
func (r ApiInventoryV1FsFsIdProductInfoGetRequest) Sku(sku string) ApiInventoryV1FsFsIdProductInfoGetRequest {
	r.sku = &sku
	return r
}

// Shop Identifier
func (r ApiInventoryV1FsFsIdProductInfoGetRequest) ShopId(shopId int64) ApiInventoryV1FsFsIdProductInfoGetRequest {
	r.shopId = &shopId
	return r
}

// Determine which page the order list should start. The minimal value is 1. Page (required if shop_id is filled)
func (r ApiInventoryV1FsFsIdProductInfoGetRequest) Page(page int64) ApiInventoryV1FsFsIdProductInfoGetRequest {
	r.page = &page
	return r
}

// Page per item (required if shop_id is filled). Maximun items are 50 for 1 page
func (r ApiInventoryV1FsFsIdProductInfoGetRequest) PerPage(perPage int64) ApiInventoryV1FsFsIdProductInfoGetRequest {
	r.perPage = &perPage
	return r
}

// This parameter is required when the product exceeds 10.000 products
func (r ApiInventoryV1FsFsIdProductInfoGetRequest) LastSort(lastSort string) ApiInventoryV1FsFsIdProductInfoGetRequest {
	r.lastSort = &lastSort
	return r
}

func (r ApiInventoryV1FsFsIdProductInfoGetRequest) Execute() (*GetProductInfoResponse, *http.Response, error) {
	return r.ApiService.InventoryV1FsFsIdProductInfoGetExecute(r)
}

/*
InventoryV1FsFsIdProductInfoGet Method for InventoryV1FsFsIdProductInfoGet

This method will retrieve single product information either by product id as parameter (choose one of those two parameters to use) from related fs_id.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fsId Fulfillment service unique identifier
 @return ApiInventoryV1FsFsIdProductInfoGetRequest
*/
func (a *DefaultApiService) InventoryV1FsFsIdProductInfoGet(ctx context.Context, fsId int64) ApiInventoryV1FsFsIdProductInfoGetRequest {
	return ApiInventoryV1FsFsIdProductInfoGetRequest{
		ApiService: a,
		ctx: ctx,
		fsId: fsId,
	}
}

// Execute executes the request
//  @return GetProductInfoResponse
func (a *DefaultApiService) InventoryV1FsFsIdProductInfoGetExecute(r ApiInventoryV1FsFsIdProductInfoGetRequest) (*GetProductInfoResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetProductInfoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.InventoryV1FsFsIdProductInfoGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory/v1/fs/{fs_id}/product/info"
	localVarPath = strings.Replace(localVarPath, "{"+"fs_id"+"}", url.PathEscape(parameterValueToString(r.fsId, "fsId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.product != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "product", r.product, "")
	}
	if r.productUrl != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "product_url", r.productUrl, "")
	}
	if r.sku != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sku", r.sku, "")
	}
	if r.shopId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "shop_id", r.shopId, "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "per_page", r.perPage, "")
	}
	if r.lastSort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "last_sort", r.lastSort, "")
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
			var v BaseErrorResponse
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

type ApiInventoryV1FsFsIdProductVariantProductIdGetRequest struct {
	ctx context.Context
	ApiService DefaultApi
	fsId int64
	productId int64
}

func (r ApiInventoryV1FsFsIdProductVariantProductIdGetRequest) Execute() (*GetProductVariantResponse, *http.Response, error) {
	return r.ApiService.InventoryV1FsFsIdProductVariantProductIdGetExecute(r)
}

/*
InventoryV1FsFsIdProductVariantProductIdGet Method for InventoryV1FsFsIdProductVariantProductIdGet

This endpoint retrieves a list of variants related to a product_id.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fsId Fulfillment service unique identifier
 @param productId Product unique identifier
 @return ApiInventoryV1FsFsIdProductVariantProductIdGetRequest
*/
func (a *DefaultApiService) InventoryV1FsFsIdProductVariantProductIdGet(ctx context.Context, fsId int64, productId int64) ApiInventoryV1FsFsIdProductVariantProductIdGetRequest {
	return ApiInventoryV1FsFsIdProductVariantProductIdGetRequest{
		ApiService: a,
		ctx: ctx,
		fsId: fsId,
		productId: productId,
	}
}

// Execute executes the request
//  @return GetProductVariantResponse
func (a *DefaultApiService) InventoryV1FsFsIdProductVariantProductIdGetExecute(r ApiInventoryV1FsFsIdProductVariantProductIdGetRequest) (*GetProductVariantResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetProductVariantResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.InventoryV1FsFsIdProductVariantProductIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory/v1/fs/{fs_id}/product/variant/{product_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"fs_id"+"}", url.PathEscape(parameterValueToString(r.fsId, "fsId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"product_id"+"}", url.PathEscape(parameterValueToString(r.productId, "productId")), -1)

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
			var v BaseErrorResponse
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

type ApiInventoryV2FsFsIdCategoryGetVariantGetRequest struct {
	ctx context.Context
	ApiService DefaultApi
	fsId int64
	catId *int64
}

// Category unique identifier
func (r ApiInventoryV2FsFsIdCategoryGetVariantGetRequest) CatId(catId int64) ApiInventoryV2FsFsIdCategoryGetVariantGetRequest {
	r.catId = &catId
	return r
}

func (r ApiInventoryV2FsFsIdCategoryGetVariantGetRequest) Execute() (*GetCategoryVariantResponse, *http.Response, error) {
	return r.ApiService.InventoryV2FsFsIdCategoryGetVariantGetExecute(r)
}

/*
InventoryV2FsFsIdCategoryGetVariantGet Method for InventoryV2FsFsIdCategoryGetVariantGet

This endpoint retrieves a list of variants related to a category_id. Use this API as main source to retrieve the newest category variant data.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fsId Fulfillment service unique identifier
 @return ApiInventoryV2FsFsIdCategoryGetVariantGetRequest
*/
func (a *DefaultApiService) InventoryV2FsFsIdCategoryGetVariantGet(ctx context.Context, fsId int64) ApiInventoryV2FsFsIdCategoryGetVariantGetRequest {
	return ApiInventoryV2FsFsIdCategoryGetVariantGetRequest{
		ApiService: a,
		ctx: ctx,
		fsId: fsId,
	}
}

// Execute executes the request
//  @return GetCategoryVariantResponse
func (a *DefaultApiService) InventoryV2FsFsIdCategoryGetVariantGetExecute(r ApiInventoryV2FsFsIdCategoryGetVariantGetRequest) (*GetCategoryVariantResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetCategoryVariantResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.InventoryV2FsFsIdCategoryGetVariantGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory/v2/fs/{fs_id}/category/get_variant"
	localVarPath = strings.Replace(localVarPath, "{"+"fs_id"+"}", url.PathEscape(parameterValueToString(r.fsId, "fsId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.catId == nil {
		return localVarReturnValue, nil, reportError("catId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "cat_id", r.catId, "")
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
			var v BaseErrorResponse
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

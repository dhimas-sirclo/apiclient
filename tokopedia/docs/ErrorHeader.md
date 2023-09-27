# ErrorHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessTime** | Pointer to **float64** |  | [optional] 
**Messages** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **string** | Error codes: * BDL_DLV_001 - Invalid request body, please check again - Invalid request body, please check again * BDL_DLV_002 - Invalid field %s format, value %v should be %s - Invalid params, please check again * BDL_DLV_013 - Filter type must be 1 or 2 and cannot be empty - Invalid filter value, please check again * BDL_USC_001 - FS ID is not associated with Shop ID or Warehouse ID - fs_id is not associated, please check again * BDL_USC_002 - Shop Info not found - Shop information cannot be found * BDL_USC_003 - Too many or too few bundle items given for this type of bundle - The number of bundle item is not right for this type of bundle * BDL_USC_004 - 2 or more same products in bundle - Multiple bundle must consist of minimum two different products * BDL_USC_005 - Different products in single bundle - Single bundle only accept same products * BDL_USC_006 - Please use different min order for each bundle item - Please use different min order * BDL_USC_007 - Min Order of single type bundle must be greater than 1 - Please set the minimum order more than 1 * BDL_USC_009 - Bundle item price must be lower than its original price - Please set the bundle item price lower than the original price * BDL_USC_010 - Bundle item stock cannot be empty - Please add the bundle item stock * BDL_USC_012 - One of the product variant is invalid - There is invalid product variant on the bundle * CMP_DLV_001 - Invalid field %s format, value %v should be %s - There is invalid request format * CMP_DLV_002 - Product IDs must be numeric and comma separated - Invalid Product IDs format, please kindly check * CMP_DLV_003 - Value of field %s exceed limit %s&#x3D;%d - There is invalid request format * CMP_DLV_004 - product_id cannot be empty - Product ID is empty, please kindly check * CMP_GRPC_001 - Failed Send GRPC Request - Failed sending request to upstream * CMP_GRPC_003 - Data Not Found - Data not found from upstream * CMP_GRPC_004 - Failed Marshalling JSON - Failed processing request to upstream * CMP_GRPC_005 - Failed Unmarshal JSON - Failed processing response from upstream * CMP_USC_001 - fs_id Is Not Associated With Shop ID or Warehouse ID - FS ID is not associated * CMP_USC_002 - Some Product ID Is Not Own By Current Shop ID - Some Product ID is not owned by current Shop ID * CMP_USC_003 - Invalid Product ID Format - Invalid Product ID format, please kindly check * ETS_USC_001 - fs_id Is Not Associated With Shop ID or Warehouse ID - FS ID is not associated, please kindly check * ORD_API_001 - Failed To Initialize Request - Failed sending request to upstream * ORD_API_003 - Failed To Fetching Request - Failed fetching response from upstream * ORD_API_004 - Failed To Reading Response Body - Failed processing response from upstream * ORD_API_005 - Failed To Parsing Json Response - Failed processing response from upstream * ORD_API_006 - There Are Error From Ext Service - Failed getting response from upstream * ORD_API_008 - Data Not Found - Order data not found * ORD_API_009 - Error Data Order Too Big, Must Reduce Time Interval - Response is too large, please reduce the time interval * ORD_API_010 - Partial Fulfillment Quantity must be lower than ordered - POF request quantity must be lower than initial order quantity * ORD_API_011 - Partial Fulfillment Request have been made before - POF request has already been made * ORD_API_012 - Order Bundle is not eligible for partial order fulfillment - POF is not eligible with bundling * ORD_API_013 - There&#39;s no change in the product quantity for partial order fulfillment - POF request quantity is the same as initial order quantity * ORD_API_016 - FS ID not authorized to reject cancel request - FS ID is not eligible to reject the buyer request cancellation * ORD_DLV_001 - fs_id cannot be empty - fs_id field is empty, please check again * ORD_DLV_002 - invalid fs_id format - Invalid fs_id format, please check again * ORD_DLV_003 - error order id and invoice no is empty - Both Order ID and Invoice Number is empty * ORD_DLV_004 - error must choose either order id or invoice no as parameter - Please choose either Order ID or Invoice Number to be filled * ORD_DLV_005 - invalid order_id format - Invalid order_id format, please check again * ORD_DLV_006 - invalid shop_id format - Invalid shop_id format, please check again * ORD_DLV_007 - invalid warehouse_id format - Invalid warehouse_id format, please check again * ORD_DLV_008 - please choose one between shop id or warehouse id - Please choose either Shop ID or Warehouse ID to be filled * ORD_DLV_009 - from_date cannot be empty - from_date is empty, please check again * ORD_DLV_010 - invalid from_date format - Invalid from_date format, please check again * ORD_DLV_011 - to_date cannot be empty - to_date is empty, please check again * ORD_DLV_012 - invalid to_date format - Invalid to_date format, please check again * ORD_DLV_013 - Date Range must be less than %d days * ORD_DLV_014 - page cannot be empty - page is empty, please check again * ORD_DLV_015 - invalid page format - Invalid page format, please check again * ORD_DLV_016 - per_page cannot be empty - per_page is empty, please check again * ORD_DLV_017 - invalid per_page format - Invalid per_page format, please check again * ORD_DLV_018 - invalid status format - Invalid status format, please check again * ORD_DLV_019 - failed read body request - Failed read body request, please check again * ORD_DLV_020 - wrong json format - Wrong JSON body format, please check again * ORD_DLV_021 - shipping ref number cannot be empty for confirm shipping - Shipping Ref Number is empty and is needed for confirm shipping, please check again * ORD_DLV_022 - Shop ID cannot be empty for set delivered - Shop ID is empty and is needed for set delivered, please check again * ORD_DLV_023 - order status is not supported - Order Status is not supported * ORD_DLV_025 - invalid next_order_id format - Invalid next_order_id format, please check again * ORD_DLV_028 - invalid order_type param value - Invalid order_type param value, please check again * ORD_DLV_029 - invalid order_type param format - Invalid order_type format, please check again * ORD_USC_001 - Order Status not eligible to accept - Order status is not eligible to be changed to accept * ORD_USC_002 - Order ID is not fulfillment * ORD_USC_003 - Order ID from warehouse id is not authenticate - Order ID is not associated with Warehouse ID * ORD_USC_004 - Order ID need to be fulfilled by toko cabang * ORD_USC_005 - Order ID from shop id is not authenticate - Order ID is not associated with Shop ID * ORD_USC_006 - Admin ID Not Found From Shop ID * ORD_USC_007 - Error When Request Accept Order To Ext Service - Failed getting response from upstream * ORD_USC_008 - Order Status not eligible to reject - Order status it not eligible to changed to reject * ORD_USC_009 - Mandatory To Fill Reason - It is required to fill reject reason * ORD_USC_010 - Failed Getting Products - Failed getting product information * ORD_USC_011 - No Matching product id * ORD_USC_012 - Mandatory To Fill Shop Close End Date and Note - It is required to fill shop close end date and note * ORD_USC_013 - Shop Close End Date Present But Invalid - Invalid shop close end date value * ORD_USC_014 - Error When Request Reject Order To Ext Service - Failed getting response from upstream * ORD_USC_015 - FS ID Is Not Associated With Shop ID or Warehouse ID - FS ID is not associated with Shop ID or Warehouse ID * ORD_USC_016 - Order ID Status Must Be 500 (In Shipping Process) to Update Into Set Delivered * ORD_USC_018 - Order ID from warehouse id is not authenticate - Order ID is not associated with Warehouse ID * ORD_USC_019 - Order need to get request cancellation first from buyer * ORD_USC_020 - You are not allowed to reject this order * ORD_USC_024 - FS ID not associated with Shop ID - FS ID is not associated with Shop ID * PRD_API_001 - Failed To Initialize Request - Failed initialize request to upstream * PRD_API_003 - Failed To Fetching Request - Failed fetching request from upstream * PRD_API_004 - Failed To Reading Response Body - Error on processing data * PRD_API_005 - Failed To Parsing Json Response - Error on processing data * PRD_API_006 - There Are Error From Ext Service - There is error from external service * PRD_API_007 - Failed Marshalling JSON - Error on processing data * PRD_CACHE_001 - Failed retrieve new upload id - Error on processing data * PRD_CACHE_002 - Failed Marshall Cache Key Value - Error on processing data * PRD_CACHE_003 - Failed To Make Cache - Error on processing data * PRD_CACHE_005 - Data Not Found - Error on processing data * PRD_CACHE_006 - Failed to retrieve cache data - Error on processing data * PRD_CACHE_007 - Failed Unmarshal cache response - Error on processing data * PRD_DB_001 - Failed To Query DB Data - Failed getting data from database * PRD_DB_002 - Failed To Scan DB Data - Failed getting data from database * PRD_DB_003 - Data Not Found - Failed getting data from database * PRD_DLV_001 - fs_id cannot be empty - fs_id field is empty, please check again * PRD_DLV_002 - invalid fs_id format - fs_id is in the wrong format, please check again * PRD_DLV_003 - invalid shop_id format - shop_id is in the wrong format, please check again * PRD_DLV_005 - page cannot be empty when shop_id or warehouse_id is filled - fs_id is in the wrong format, please check * PRD_DLV_015 - failed read body request - Failed reading the body request * PRD_DLV_016 - shop_id cannot be empty - shop_id field is empty, please check again * PRD_DLV_018 - Max allowed products per-update are %d products - The request has exceed the max allowed product price edit per request * PRD_DLV_019 - invalid warehouse_id format - warehouse_id is in the wrong format, please check again * PRD_DLV_038 - Value %s of field %s is not allowed, the allowed fields are %s - There is something wrong with the request body * PRD_DLV_042 - Invalid field %s format, value %v should be %s - There is something wrong with the request body * PRD_DLV_043 - Invalid request body, please check again - There is something wrong with the request body * PRD_DLV_052 - warehouse_id cannot be empty - warehouse_id is empty, please check again * PRD_DLV_055 - warehouse_id is not eligible to perform this operation - warehouse_id is not eligible, please check again * PRD_GRPC_001 - Failed Send GRPC Request - Failed sending request to upstream * PRD_GRPC_002 - There Are Error From Ext Service - Failed getting response from upstream * PRD_GRPC_003 - Data Not Found - Product not found * PRD_USC_001 - fs_id Is Not Associated With Shop ID or Warehouse ID - fs_id not associated, please check again * PRD_USC_008 - Mandatory To Fill Warehouse ID - warehouse_id is empty, please check again * PRD_USC_009 - Warehouse is not owned by toko cabang * PRD_USC_010 - Warehouse is owned by toko cabang * PRD_USC_011 - Partner ID not found - Failed to get Partner ID * PRD_USC_012 - Shop ID Not Match - Shop ID did not match with warehouse data * PRD_USC_013 - FS Type cannot use warehouse - Failed to precess data * PRD_USC_015 - Failed Epoch converter today - Error on processing data * PRD_USC_021 - Warehouse Data Not Found - Warehouse ID not found * PRD_USC_024 - Error Convert number format into string - Failed getting response from upstream * PRD_USC_025 - Shop ID is not associated with upload id - Shop ID not associated, please check again * PRD_USC_027 - Invalid action parameter value - invalid action value, please check again * PRD_USC_028 - Product Status Not In Active or Warehouse. Edit Process Canceled - Unable to edit product because the status is not active or warehouse * PRD_USC_029 - Warehouse ID Information Not Found - Failed to get warehouse information * RBAC_MDLW_001 - FS ID is not associated with Shop ID or Warehouse ID - fs_id not assosiated, please check again * RBAC_MDLW_002 - This shop owner has not given permission for your app to use this API on the shop - Application does not have shop permission * RBAC_USC_008 - App is not associated with shop - Application not assosiated, please check again * RBAC_USC_012 - This shop owner has not given permission for your app to use this API on the shop - Application does not have shop permission * SHP_API_003 - Failed To Fetching Request - Failed fetching response from upstream * SHP_DLV_001 - FS ID is not valid - Invalid FS ID, please kindly check * SHP_DLV_002 - Shop ID is not valid - Invalid Shop ID, please kindly check * SHP_DLV_003 - Page format is not valid - Invalid Page format, please kindly check * SHP_DLV_004 - Per Page format is not valid - Invalid Per Page format, please kindly check * SHP_DLV_005 - Date format is not valid - Invalid Date format, please kindly check * SHP_DLV_007 - Date Range must be less than %d days - Invalid Date Range value, please kindly check * SHP_DLV_008 - failed read body request - Failed read body request, please kindly check * SHP_DLV_009 - invalid hide_zero format - Invalid hide_zero format, please kindly check * SHP_DLV_010 - page_count value exceed limit - Invalid page_count value, please kindly check * SHP_DLV_011 - invalid display value - Invalid display value, please kindly check * SHP_DLV_012 - fs_id cannot be empty - fs_id is empty, please kindly check * SHP_DLV_013 - shop_id cannot be empty - shop_id is empty, please kindly check * SHP_DLV_014 - invalid page_count format - invalid page_count format, please kindly check * PRD_DLV_045 - Field %s should not be empty, please fill the empty field - Please check the request params again * SHP_GRPC_001 - Failed Send GRPC Request - Failed sending request to upstream * SHP_GRPC_002 - There Are Error From Ext Service - Failed getting response from upstream * SHP_GRPC_003 - Data Not Found - Failed finding data from upstream * SHP_USC_001 - FS ID is not associated with Shop ID - FS ID is not connected with the Shop ID, please kindly check * SHP_USC_002 - Shop Owner is not found - Shop owner info not found, please kindly check * SHP_USC_004 - Shop information is empty - Failed getting shop info, please kindly check * SPE_DLV_001 - fs_id cannot be empty - FS ID is empty * SPE_DLV_005 - Invalid field %s format, value %v should be %s - Invalid field format, please kindly check * SPE_DLV_006 - Invalid request body, please check again - Invalid request body, please kindly check * SPE_DLV_007 - Value of field %s exceed limit %s&#x3D;%d - Invalid field value, please kindly check * SPE_DLV_008 - Value %s of field %s is not allowed, the allowed fields are %s - Invalid field value, please kindly check * SPE_GRPC_001 - Failed Send GRPC Request - Failed sending request to upstream * SPE_GRPC_003 - Data Not Found - Data not found * SPE_GRPC_004 - Failed Marshalling JSON - Failed processing request to upstream * SPE_GRPC_005 - Failed Unmarshal JSON - Failed processing response from upstream * SPE_USC_001 - fs_id Is Not Associated With Shop ID or Warehouse ID - FS ID is not associated, please kindly check * SPE_USC_002 - Product ID does not belong to Shop ID or Warehouse ID - Product ID is not associated, please kindly check * SPE_USC_003 - Warehouse ID Information Not Found - Warehouse information not found, please kindly check * SPE_USC_004 - Cannot View Slash Price Information Parent Product Variant * VRT_DLV_001 - fs_id cannot be empty - There is invalid request params * VRT_DLV_002 - invalid fs_id format - There is invalid request params * VRT_DLV_004 - Invalid cat_id format - There is invalid request params * VRT_DLV_005 - invalid product_id format - There is invalid request params * VRT_USC_001 - fs_id Is Not Associated With Shop ID or Warehouse ID - fs_id not associated, please check again * VRT_USC_002 - Category ID Not Found - Category ID not found, please check again  | [optional] 

## Methods

### NewErrorHeader

`func NewErrorHeader() *ErrorHeader`

NewErrorHeader instantiates a new ErrorHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorHeaderWithDefaults

`func NewErrorHeaderWithDefaults() *ErrorHeader`

NewErrorHeaderWithDefaults instantiates a new ErrorHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessTime

`func (o *ErrorHeader) GetProcessTime() float64`

GetProcessTime returns the ProcessTime field if non-nil, zero value otherwise.

### GetProcessTimeOk

`func (o *ErrorHeader) GetProcessTimeOk() (*float64, bool)`

GetProcessTimeOk returns a tuple with the ProcessTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessTime

`func (o *ErrorHeader) SetProcessTime(v float64)`

SetProcessTime sets ProcessTime field to given value.

### HasProcessTime

`func (o *ErrorHeader) HasProcessTime() bool`

HasProcessTime returns a boolean if a field has been set.

### GetMessages

`func (o *ErrorHeader) GetMessages() string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ErrorHeader) GetMessagesOk() (*string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ErrorHeader) SetMessages(v string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ErrorHeader) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetReason

`func (o *ErrorHeader) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ErrorHeader) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ErrorHeader) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ErrorHeader) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetErrorCode

`func (o *ErrorHeader) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ErrorHeader) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ErrorHeader) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ErrorHeader) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



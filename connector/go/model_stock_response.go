/*
 * MP Connector API
 *
 * MP Connector API
 *
 * API version: v1.0.0
 * Contact: dev@sirclo.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package connector

type StockResponse struct {

	WarehouseId string `json:"warehouse_id,omitempty"`

	Qty float32 `json:"qty,omitempty"`

	Success bool `json:"success,omitempty"`

	Message string `json:"message,omitempty"`
}
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

type UpdateStockInput struct {

	ProductId string `json:"product_id"`

	VariantId string `json:"variant_id"`

	Stocks []StockInput `json:"stocks,omitempty"`
}

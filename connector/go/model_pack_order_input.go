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

type PackOrderInput struct {

	OrderId string `json:"order_id"`

	AirwaybillNumber string `json:"airwaybill_number,omitempty"`
}

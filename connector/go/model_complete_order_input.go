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

type CompleteOrderInput struct {

	OrderId string `json:"order_id"`

	ReceivedBy string `json:"received_by,omitempty"`
}

/*
 * Quatrix API
 *
 * File Transfer And Sharing API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package quatrix

type ShareRecipientsResp struct {

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Email string `json:"email,omitempty"`

	HasKey bool `json:"has_key,omitempty"`
}
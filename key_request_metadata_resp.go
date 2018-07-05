/*
 * Quatrix API
 *
 * File Transfer And Sharing API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package quatrix

type KeyRequestMetadataResp struct {

	Id string `json:"id,omitempty"`

	SenderName string `json:"sender_name,omitempty"`

	SenderEmail string `json:"sender_email,omitempty"`

	UserName string `json:"user_name,omitempty"`

	UserEmail string `json:"user_email,omitempty"`

	Locale string `json:"locale,omitempty"`

	IsContact bool `json:"is_contact,omitempty"`
}
/*
 * Quatrix API
 *
 * File Transfer And Sharing API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package quatrix

type UserPermissionResp struct {

	UserId string `json:"user_id,omitempty"`

	Operations float32 `json:"operations,omitempty"`

	Owns bool `json:"owns,omitempty"`

	Notify bool `json:"notify,omitempty"`
}

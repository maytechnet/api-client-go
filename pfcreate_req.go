/*
 * Quatrix API
 *
 * File Transfer And Sharing API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package quatrix

type PfcreateReq struct {

	Created float32 `json:"created,omitempty"`

	Name string `json:"name"`

	FileId string `json:"file_id"`

	UsersPermissions []UserPermissionReq `json:"users_permissions"`

	Notify bool `json:"notify,omitempty"`
}

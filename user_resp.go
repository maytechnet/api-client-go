/*
 * Quatrix API
 *
 * File Transfer And Sharing API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package quatrix

type UserResp struct {

	Id string `json:"id,omitempty"`

	ParentId string `json:"parent_id,omitempty"`

	HomeId string `json:"home_id,omitempty"`

	HomeName string `json:"home_name,omitempty"`

	SuperAdmin string `json:"super_admin,omitempty"`

	Name string `json:"name,omitempty"`

	Email string `json:"email,omitempty"`

	Status string `json:"status,omitempty"`

	Quota float32 `json:"quota,omitempty"`

	Created float32 `json:"created,omitempty"`

	Modified float32 `json:"modified,omitempty"`

	Groups []Group `json:"groups,omitempty"`

	UserOperations float32 `json:"user_operations,omitempty"`

	EffectiveOperations float32 `json:"effective_operations,omitempty"`

	HomeOperations float32 `json:"home_operations,omitempty"`

	Language string `json:"language,omitempty"`

	StorageId string `json:"storage_id,omitempty"`

	HasKey bool `json:"has_key,omitempty"`

	Services []ShortUserService `json:"services,omitempty"`

	UniqueLogin string `json:"unique_login,omitempty"`

	Readonly bool `json:"readonly,omitempty"`

	AuthMethods []string `json:"auth_methods,omitempty"`

	LastLogin float32 `json:"last_login,omitempty"`

	StorageUsed float32 `json:"storage_used,omitempty"`

	ForcedAuth []string `json:"forced_auth,omitempty"`
}
/*
 * Quatrix API
 *
 * File Transfer And Sharing API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package quatrix

type UserDeleteReq struct {

	Ids []string `json:"ids"`

	DeleteHome bool `json:"delete_home,omitempty"`
}

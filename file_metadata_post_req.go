/*
 * Quatrix API
 *
 * File Transfer And Sharing API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package quatrix

type FileMetadataPostReq struct {

	Id string `json:"id"`

	Mtime float32 `json:"mtime,omitempty"`
}

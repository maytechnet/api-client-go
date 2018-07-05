/*
 * Quatrix API
 *
 * File Transfer And Sharing API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package quatrix

type FilesReturnUploadLinkReq struct {

	ParentId string `json:"parent_id"`

	FileSize float32 `json:"file_size,omitempty"`
}
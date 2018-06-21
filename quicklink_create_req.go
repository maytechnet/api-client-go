/*
 * Quatrix API
 *
 * File Transfer And Sharing API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package quatrix

type QuicklinkCreateReq struct {

	Subject string `json:"subject,omitempty"`

	Files []string `json:"files"`

	Pin string `json:"pin,omitempty"`
}

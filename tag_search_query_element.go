/*
 * Quatrix API
 *
 * File Transfer And Sharing API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package quatrix

type TagSearchQueryElement struct {

	And []SearchTag `json:"_and,omitempty"`

	Or []SearchTag `json:"_or,omitempty"`
}
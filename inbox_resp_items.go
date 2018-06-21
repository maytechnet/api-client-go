/*
 * Quatrix API
 *
 * File Transfer And Sharing API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package quatrix

type InboxRespItems struct {

	Id string `json:"id,omitempty"`

	// Sender name
	SenderName string `json:"sender_name,omitempty"`

	Activates float32 `json:"activates,omitempty"`

	IsReply bool `json:"is_reply,omitempty"`

	Replied bool `json:"replied,omitempty"`

	Subject string `json:"subject,omitempty"`

	Type_ string `json:"type,omitempty"`

	FileRequest bool `json:"file_request,omitempty"`

	Expires float32 `json:"expires,omitempty"`

	Protected bool `json:"protected,omitempty"`
}

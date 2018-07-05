/*
 * Quatrix API
 *
 * File Transfer And Sharing API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package quatrix

type ShareCreateReq struct {

	FolderId string `json:"folder_id"`

	Files []string `json:"files"`

	Contacts []string `json:"contacts"`

	Subject string `json:"subject,omitempty"`

	Message string `json:"message,omitempty"`

	Expires float32 `json:"expires,omitempty"`

	Activates float32 `json:"activates,omitempty"`

	ReturnFiles bool `json:"return_files,omitempty"`

	PgpEncrypted bool `json:"pgp_encrypted,omitempty"`

	ReturnPgpEncrypted bool `json:"return_pgp_encrypted,omitempty"`

	MessageSignature string `json:"message_signature,omitempty"`

	Notify bool `json:"notify,omitempty"`

	ShareType string `json:"share_type,omitempty"`

	Pin string `json:"pin,omitempty"`

	// create share with sending email or just create share
	SendEmail bool `json:"send_email,omitempty"`
}
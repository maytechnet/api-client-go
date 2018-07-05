/*
 * Quatrix API
 *
 * File Transfer And Sharing API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package quatrix

type FilesReturnMetadataResp struct {

	Id string `json:"id,omitempty"`

	Subject string `json:"subject,omitempty"`

	SenderName string `json:"sender_name,omitempty"`

	SenderKey string `json:"sender_key,omitempty"`

	SenderEmail string `json:"sender_email,omitempty"`

	UserName string `json:"user_name,omitempty"`

	UserEmail string `json:"user_email,omitempty"`

	Locale string `json:"locale,omitempty"`

	ReturnPgpEncrypted bool `json:"return_pgp_encrypted,omitempty"`

	Title string `json:"title,omitempty"`

	PgpEnabled bool `json:"pgp_enabled,omitempty"`

	ShareType string `json:"share_type,omitempty"`

	FolderId string `json:"folder_id,omitempty"`

	Activates float32 `json:"activates,omitempty"`

	Expires float32 `json:"expires,omitempty"`

	FilesPerShare float32 `json:"files_per_share,omitempty"`

	Message string `json:"message,omitempty"`
}
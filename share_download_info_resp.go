/*
 * Quatrix API
 *
 * Download and upload files or folders, share them with predefined security options, manage your account and profile settings and a lot more functionalities can be easily integrated into your application using our Quatrix APIs. Learn more how to authenticate the Quatrix session, how to construct JSON formatted API calls and what responses to expect in our [API Guide](https://docs.maytech.net/display/MD/API+Guide).
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package quatrix

type ShareDownloadInfoResp struct {

	Id string `json:"id,omitempty"`

	Message string `json:"message,omitempty"`

	Subject string `json:"subject,omitempty"`

	Activates float32 `json:"activates,omitempty"`

	Expires float32 `json:"expires,omitempty"`

	SenderName string `json:"sender_name,omitempty"`

	SenderEmail string `json:"sender_email,omitempty"`

	// user name from current session
	UserName string `json:"user_name,omitempty"`

	// user email from current session
	UserEmail string `json:"user_email,omitempty"`

	Locale string `json:"locale,omitempty"`

	UserId string `json:"user_id,omitempty"`

	// PGP protected share
	PgpEncrypted bool `json:"pgp_encrypted,omitempty"`

	// for PGP
	PrivateKey string `json:"private_key,omitempty"`

	// It shows that the user from the current session has PGP key. If True - the user will get his private_key.
	PgpEnabled bool `json:"pgp_enabled,omitempty"`

	Status string `json:"status,omitempty"`

	ShareType string `json:"share_type,omitempty"`

	Title string `json:"title,omitempty"`

	Expired bool `json:"expired,omitempty"`

	IsReply bool `json:"is_reply,omitempty"`

	PinProtected bool `json:"pin_protected,omitempty"`
}

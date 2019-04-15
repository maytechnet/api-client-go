/*
 * Quatrix API
 *
 * Download and upload files or folders, share them with predefined security options, manage your account and profile settings and a lot more functionalities can be easily integrated into your application using our Quatrix APIs. Learn more how to authenticate the Quatrix session, how to construct JSON formatted API calls and what responses to expect in our [API Guide](https://docs.maytech.net/display/MD/API+Guide).
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package quatrix

type SettingsResp struct {

	Title string `json:"title,omitempty"`

	Bcc []string `json:"bcc,omitempty"`

	BillingEmails []string `json:"billing_emails,omitempty"`

	EmailFooter string `json:"email_footer,omitempty"`

	Modified float32 `json:"modified,omitempty"`

	Language string `json:"language,omitempty"`

	PgpEnabled bool `json:"pgp_enabled,omitempty"`

	ShareTypes *interface{} `json:"share_types,omitempty"`
}

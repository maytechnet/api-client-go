/*
 * Quatrix API
 *
 * File Transfer And Sharing API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package quatrix

type SettingsSetReq struct {

	Title string `json:"title,omitempty"`

	Bcc []string `json:"bcc,omitempty"`

	BillingEmails []string `json:"billing_emails,omitempty"`

	EmailFooter string `json:"email_footer,omitempty"`

	Language string `json:"language,omitempty"`

	EnablePgp bool `json:"enable_pgp,omitempty"`

	VirusScan bool `json:"virus_scan,omitempty"`

	ShareTypes *interface{} `json:"share_types,omitempty"`

	AuthMethods []string `json:"auth_methods,omitempty"`
}
/*
 * Quatrix API
 *
 * Download and upload files or folders, share them with predefined security options, manage your account and profile settings and a lot more functionalities can be easily integrated into your application using our Quatrix APIs. Learn more how to authenticate the Quatrix session, how to construct JSON formatted API calls and what responses to expect in our [API Guide](https://docs.maytech.net/display/MD/API+Guide).
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package quatrix

type SettingssetShareTypes struct {

	PublicShare bool `json:"public_share,omitempty"`

	TrackedShare bool `json:"tracked_share,omitempty"`

	RestrictedShare bool `json:"restricted_share,omitempty"`
}

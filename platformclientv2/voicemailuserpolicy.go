package platformclientv2
import (
	"time"
	"encoding/json"
)

// Voicemailuserpolicy
type Voicemailuserpolicy struct { 
	// Enabled - Whether the user has voicemail enabled
	Enabled *bool `json:"enabled,omitempty"`


	// AlertTimeoutSeconds - The number of seconds to ring the user's phone before a call is transfered to voicemail
	AlertTimeoutSeconds *int `json:"alertTimeoutSeconds,omitempty"`


	// Pin - The user's PIN to access their voicemail. This property is only used for updates and never provided otherwise to ensure security
	Pin *string `json:"pin,omitempty"`


	// ModifiedDate - The date the policy was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`


	// SendEmailNotifications - Whether email notifications are sent to the user when a new voicemail is received
	SendEmailNotifications *bool `json:"sendEmailNotifications,omitempty"`

}

// String returns a JSON representation of the model
func (o *Voicemailuserpolicy) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

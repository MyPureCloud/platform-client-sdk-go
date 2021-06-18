package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Voicemailorganizationpolicy
type Voicemailorganizationpolicy struct { 
	// Enabled - Whether voicemail is enable for this organization
	Enabled *bool `json:"enabled,omitempty"`


	// AlertTimeoutSeconds - The organization's default number of seconds to ring a user's phone before a call is transfered to voicemail
	AlertTimeoutSeconds *int `json:"alertTimeoutSeconds,omitempty"`


	// PinConfiguration - The configuration for user PINs to access their voicemail from a phone
	PinConfiguration *Pinconfiguration `json:"pinConfiguration,omitempty"`


	// VoicemailExtension - The extension for voicemail retrieval.  The default value is *86.
	VoicemailExtension *string `json:"voicemailExtension,omitempty"`


	// PinRequired - If this is true, a PIN is required when accessing a user's voicemail from a phone.
	PinRequired *bool `json:"pinRequired,omitempty"`


	// SendEmailNotifications - Whether email notifications are sent for new voicemails in the organization. If false, new voicemail email notifications are not be sent for the organization overriding any user or group setting.
	SendEmailNotifications *bool `json:"sendEmailNotifications,omitempty"`


	// DisableEmailPii - Removes any PII from emails. This overrides any analogous group configuration value. This is always true if HIPAA is enabled or unknown for an organization.
	DisableEmailPii *bool `json:"disableEmailPii,omitempty"`


	// ModifiedDate - The date the policy was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Voicemailorganizationpolicy) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

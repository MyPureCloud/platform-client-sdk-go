package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Voicemailgrouppolicy
type Voicemailgrouppolicy struct { 
	// Name
	Name *string `json:"name,omitempty"`


	// Group - The group associated with the policy
	Group *Group `json:"group,omitempty"`


	// Enabled - Whether voicemail is enabled for the group
	Enabled *bool `json:"enabled,omitempty"`


	// SendEmailNotifications - Whether email notifications are sent to group members when a new voicemail is received
	SendEmailNotifications *bool `json:"sendEmailNotifications,omitempty"`


	// DisableEmailPii - Removes any PII from group emails. This is overridden by the analogous organization configuration value. This is always true if HIPAA is enabled or unknown for an organization.
	DisableEmailPii *bool `json:"disableEmailPii,omitempty"`


	// RotateCallsSecs - How many seconds to ring before rotating to the next member in the group
	RotateCallsSecs *int `json:"rotateCallsSecs,omitempty"`


	// StopRingingAfterRotations - How many rotations to go through
	StopRingingAfterRotations *int `json:"stopRingingAfterRotations,omitempty"`


	// OverflowGroupId -  A fallback group to contact when all of the members in this group did not answer the call.
	OverflowGroupId *string `json:"overflowGroupId,omitempty"`


	// GroupAlertType - Specifies if the members in this group should be contacted randomly, in a specific order, or by round-robin.
	GroupAlertType *string `json:"groupAlertType,omitempty"`

}

// String returns a JSON representation of the model
func (o *Voicemailgrouppolicy) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

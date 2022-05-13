package platformclientv2
import (
	"github.com/leekchan/timeutil"
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


	// OverflowGroupId - A fallback group to contact when all of the members in this group did not answer the call.
	OverflowGroupId *string `json:"overflowGroupId,omitempty"`


	// GroupAlertType - Specifies if the members in this group should be contacted randomly, in a specific order, or by round-robin.
	GroupAlertType *string `json:"groupAlertType,omitempty"`


	// InteractiveResponsePromptId - The prompt to use when connecting a user to a Group Ring call
	InteractiveResponsePromptId *string `json:"interactiveResponsePromptId,omitempty"`

}

func (o *Voicemailgrouppolicy) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Voicemailgrouppolicy
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Group *Group `json:"group,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		SendEmailNotifications *bool `json:"sendEmailNotifications,omitempty"`
		
		DisableEmailPii *bool `json:"disableEmailPii,omitempty"`
		
		RotateCallsSecs *int `json:"rotateCallsSecs,omitempty"`
		
		StopRingingAfterRotations *int `json:"stopRingingAfterRotations,omitempty"`
		
		OverflowGroupId *string `json:"overflowGroupId,omitempty"`
		
		GroupAlertType *string `json:"groupAlertType,omitempty"`
		
		InteractiveResponsePromptId *string `json:"interactiveResponsePromptId,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Group: o.Group,
		
		Enabled: o.Enabled,
		
		SendEmailNotifications: o.SendEmailNotifications,
		
		DisableEmailPii: o.DisableEmailPii,
		
		RotateCallsSecs: o.RotateCallsSecs,
		
		StopRingingAfterRotations: o.StopRingingAfterRotations,
		
		OverflowGroupId: o.OverflowGroupId,
		
		GroupAlertType: o.GroupAlertType,
		
		InteractiveResponsePromptId: o.InteractiveResponsePromptId,
		Alias:    (*Alias)(o),
	})
}

func (o *Voicemailgrouppolicy) UnmarshalJSON(b []byte) error {
	var VoicemailgrouppolicyMap map[string]interface{}
	err := json.Unmarshal(b, &VoicemailgrouppolicyMap)
	if err != nil {
		return err
	}
	
	if Name, ok := VoicemailgrouppolicyMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Group, ok := VoicemailgrouppolicyMap["group"].(map[string]interface{}); ok {
		GroupString, _ := json.Marshal(Group)
		json.Unmarshal(GroupString, &o.Group)
	}
	
	if Enabled, ok := VoicemailgrouppolicyMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if SendEmailNotifications, ok := VoicemailgrouppolicyMap["sendEmailNotifications"].(bool); ok {
		o.SendEmailNotifications = &SendEmailNotifications
	}
    
	if DisableEmailPii, ok := VoicemailgrouppolicyMap["disableEmailPii"].(bool); ok {
		o.DisableEmailPii = &DisableEmailPii
	}
    
	if RotateCallsSecs, ok := VoicemailgrouppolicyMap["rotateCallsSecs"].(float64); ok {
		RotateCallsSecsInt := int(RotateCallsSecs)
		o.RotateCallsSecs = &RotateCallsSecsInt
	}
	
	if StopRingingAfterRotations, ok := VoicemailgrouppolicyMap["stopRingingAfterRotations"].(float64); ok {
		StopRingingAfterRotationsInt := int(StopRingingAfterRotations)
		o.StopRingingAfterRotations = &StopRingingAfterRotationsInt
	}
	
	if OverflowGroupId, ok := VoicemailgrouppolicyMap["overflowGroupId"].(string); ok {
		o.OverflowGroupId = &OverflowGroupId
	}
    
	if GroupAlertType, ok := VoicemailgrouppolicyMap["groupAlertType"].(string); ok {
		o.GroupAlertType = &GroupAlertType
	}
    
	if InteractiveResponsePromptId, ok := VoicemailgrouppolicyMap["interactiveResponsePromptId"].(string); ok {
		o.InteractiveResponsePromptId = &InteractiveResponsePromptId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Voicemailgrouppolicy) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

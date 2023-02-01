package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Voicemailgrouppolicy
type Voicemailgrouppolicy struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

	// IncludeEmailTranscriptions - Whether to include the voicemail transcription in a group notification email
	IncludeEmailTranscriptions *bool `json:"includeEmailTranscriptions,omitempty"`

	// LanguagePreference - The language preference for the group.  Used for group voicemail transcription
	LanguagePreference *string `json:"languagePreference,omitempty"`

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

	// InteractiveResponseRequired - Whether user should be prompted with a confirmation prompt when connecting to a Group Ring call
	InteractiveResponseRequired *bool `json:"interactiveResponseRequired,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Voicemailgrouppolicy) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Voicemailgrouppolicy) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Voicemailgrouppolicy
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Group *Group `json:"group,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		SendEmailNotifications *bool `json:"sendEmailNotifications,omitempty"`
		
		DisableEmailPii *bool `json:"disableEmailPii,omitempty"`
		
		IncludeEmailTranscriptions *bool `json:"includeEmailTranscriptions,omitempty"`
		
		LanguagePreference *string `json:"languagePreference,omitempty"`
		
		RotateCallsSecs *int `json:"rotateCallsSecs,omitempty"`
		
		StopRingingAfterRotations *int `json:"stopRingingAfterRotations,omitempty"`
		
		OverflowGroupId *string `json:"overflowGroupId,omitempty"`
		
		GroupAlertType *string `json:"groupAlertType,omitempty"`
		
		InteractiveResponsePromptId *string `json:"interactiveResponsePromptId,omitempty"`
		
		InteractiveResponseRequired *bool `json:"interactiveResponseRequired,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		Group: o.Group,
		
		Enabled: o.Enabled,
		
		SendEmailNotifications: o.SendEmailNotifications,
		
		DisableEmailPii: o.DisableEmailPii,
		
		IncludeEmailTranscriptions: o.IncludeEmailTranscriptions,
		
		LanguagePreference: o.LanguagePreference,
		
		RotateCallsSecs: o.RotateCallsSecs,
		
		StopRingingAfterRotations: o.StopRingingAfterRotations,
		
		OverflowGroupId: o.OverflowGroupId,
		
		GroupAlertType: o.GroupAlertType,
		
		InteractiveResponsePromptId: o.InteractiveResponsePromptId,
		
		InteractiveResponseRequired: o.InteractiveResponseRequired,
		Alias:    (Alias)(o),
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
    
	if IncludeEmailTranscriptions, ok := VoicemailgrouppolicyMap["includeEmailTranscriptions"].(bool); ok {
		o.IncludeEmailTranscriptions = &IncludeEmailTranscriptions
	}
    
	if LanguagePreference, ok := VoicemailgrouppolicyMap["languagePreference"].(string); ok {
		o.LanguagePreference = &LanguagePreference
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
    
	if InteractiveResponseRequired, ok := VoicemailgrouppolicyMap["interactiveResponseRequired"].(bool); ok {
		o.InteractiveResponseRequired = &InteractiveResponseRequired
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Voicemailgrouppolicy) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

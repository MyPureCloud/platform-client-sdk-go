package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Voicemailorganizationpolicy
type Voicemailorganizationpolicy struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Enabled - Whether voicemail is enabled for this organization
	Enabled *bool `json:"enabled,omitempty"`

	// AlertTimeoutSeconds - The organization's default number of seconds to ring a user's phone before a call is transferred to voicemail
	AlertTimeoutSeconds *int `json:"alertTimeoutSeconds,omitempty"`

	// PinConfiguration - The configuration for user PINs to access their voicemail from a phone
	PinConfiguration *Pinconfiguration `json:"pinConfiguration,omitempty"`

	// VoicemailExtension - The extension for voicemail retrieval.  The default value is *86.
	VoicemailExtension *string `json:"voicemailExtension,omitempty"`

	// PinRequired - If this is true, a PIN is required when accessing a user's voicemail from a phone.
	PinRequired *bool `json:"pinRequired,omitempty"`

	// InteractiveResponseRequired - Whether user should be prompted with a confirmation prompt when connecting to a Group Ring call
	InteractiveResponseRequired *bool `json:"interactiveResponseRequired,omitempty"`

	// SendEmailNotifications - Whether email notifications are sent for new voicemails in the organization. If false, new voicemail email notifications are not be sent for the organization overriding any user or group setting.
	SendEmailNotifications *bool `json:"sendEmailNotifications,omitempty"`

	// IncludeEmailTranscriptions - Whether to include the voicemail transcription in the notification email
	IncludeEmailTranscriptions *bool `json:"includeEmailTranscriptions,omitempty"`

	// DisableEmailPii - Removes any PII from emails. This overrides any analogous group configuration value. This is always true if HIPAA is enabled or unknown for an organization.
	DisableEmailPii *bool `json:"disableEmailPii,omitempty"`

	// MaximumRecordingTimeSeconds - Default value for the maximum length of time in seconds of a recorded voicemail
	MaximumRecordingTimeSeconds *int `json:"maximumRecordingTimeSeconds,omitempty"`

	// ModifiedDate - The date the policy was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Voicemailorganizationpolicy) SetField(field string, fieldValue interface{}) {
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

func (o Voicemailorganizationpolicy) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ModifiedDate", }
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
	type Alias Voicemailorganizationpolicy
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		AlertTimeoutSeconds *int `json:"alertTimeoutSeconds,omitempty"`
		
		PinConfiguration *Pinconfiguration `json:"pinConfiguration,omitempty"`
		
		VoicemailExtension *string `json:"voicemailExtension,omitempty"`
		
		PinRequired *bool `json:"pinRequired,omitempty"`
		
		InteractiveResponseRequired *bool `json:"interactiveResponseRequired,omitempty"`
		
		SendEmailNotifications *bool `json:"sendEmailNotifications,omitempty"`
		
		IncludeEmailTranscriptions *bool `json:"includeEmailTranscriptions,omitempty"`
		
		DisableEmailPii *bool `json:"disableEmailPii,omitempty"`
		
		MaximumRecordingTimeSeconds *int `json:"maximumRecordingTimeSeconds,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		Alias
	}{ 
		Enabled: o.Enabled,
		
		AlertTimeoutSeconds: o.AlertTimeoutSeconds,
		
		PinConfiguration: o.PinConfiguration,
		
		VoicemailExtension: o.VoicemailExtension,
		
		PinRequired: o.PinRequired,
		
		InteractiveResponseRequired: o.InteractiveResponseRequired,
		
		SendEmailNotifications: o.SendEmailNotifications,
		
		IncludeEmailTranscriptions: o.IncludeEmailTranscriptions,
		
		DisableEmailPii: o.DisableEmailPii,
		
		MaximumRecordingTimeSeconds: o.MaximumRecordingTimeSeconds,
		
		ModifiedDate: ModifiedDate,
		Alias:    (Alias)(o),
	})
}

func (o *Voicemailorganizationpolicy) UnmarshalJSON(b []byte) error {
	var VoicemailorganizationpolicyMap map[string]interface{}
	err := json.Unmarshal(b, &VoicemailorganizationpolicyMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := VoicemailorganizationpolicyMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if AlertTimeoutSeconds, ok := VoicemailorganizationpolicyMap["alertTimeoutSeconds"].(float64); ok {
		AlertTimeoutSecondsInt := int(AlertTimeoutSeconds)
		o.AlertTimeoutSeconds = &AlertTimeoutSecondsInt
	}
	
	if PinConfiguration, ok := VoicemailorganizationpolicyMap["pinConfiguration"].(map[string]interface{}); ok {
		PinConfigurationString, _ := json.Marshal(PinConfiguration)
		json.Unmarshal(PinConfigurationString, &o.PinConfiguration)
	}
	
	if VoicemailExtension, ok := VoicemailorganizationpolicyMap["voicemailExtension"].(string); ok {
		o.VoicemailExtension = &VoicemailExtension
	}
    
	if PinRequired, ok := VoicemailorganizationpolicyMap["pinRequired"].(bool); ok {
		o.PinRequired = &PinRequired
	}
    
	if InteractiveResponseRequired, ok := VoicemailorganizationpolicyMap["interactiveResponseRequired"].(bool); ok {
		o.InteractiveResponseRequired = &InteractiveResponseRequired
	}
    
	if SendEmailNotifications, ok := VoicemailorganizationpolicyMap["sendEmailNotifications"].(bool); ok {
		o.SendEmailNotifications = &SendEmailNotifications
	}
    
	if IncludeEmailTranscriptions, ok := VoicemailorganizationpolicyMap["includeEmailTranscriptions"].(bool); ok {
		o.IncludeEmailTranscriptions = &IncludeEmailTranscriptions
	}
    
	if DisableEmailPii, ok := VoicemailorganizationpolicyMap["disableEmailPii"].(bool); ok {
		o.DisableEmailPii = &DisableEmailPii
	}
    
	if MaximumRecordingTimeSeconds, ok := VoicemailorganizationpolicyMap["maximumRecordingTimeSeconds"].(float64); ok {
		MaximumRecordingTimeSecondsInt := int(MaximumRecordingTimeSeconds)
		o.MaximumRecordingTimeSeconds = &MaximumRecordingTimeSecondsInt
	}
	
	if modifiedDateString, ok := VoicemailorganizationpolicyMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Voicemailorganizationpolicy) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

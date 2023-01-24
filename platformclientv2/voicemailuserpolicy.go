package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Voicemailuserpolicy
type Voicemailuserpolicy struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Voicemailuserpolicy) SetField(field string, fieldValue interface{}) {
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

func (o Voicemailuserpolicy) MarshalJSON() ([]byte, error) {
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
			if contains(dateTimeFields, fieldName) {
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
	type Alias Voicemailuserpolicy
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		AlertTimeoutSeconds *int `json:"alertTimeoutSeconds,omitempty"`
		
		Pin *string `json:"pin,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		SendEmailNotifications *bool `json:"sendEmailNotifications,omitempty"`
		Alias
	}{ 
		Enabled: o.Enabled,
		
		AlertTimeoutSeconds: o.AlertTimeoutSeconds,
		
		Pin: o.Pin,
		
		ModifiedDate: ModifiedDate,
		
		SendEmailNotifications: o.SendEmailNotifications,
		Alias:    (Alias)(o),
	})
}

func (o *Voicemailuserpolicy) UnmarshalJSON(b []byte) error {
	var VoicemailuserpolicyMap map[string]interface{}
	err := json.Unmarshal(b, &VoicemailuserpolicyMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := VoicemailuserpolicyMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if AlertTimeoutSeconds, ok := VoicemailuserpolicyMap["alertTimeoutSeconds"].(float64); ok {
		AlertTimeoutSecondsInt := int(AlertTimeoutSeconds)
		o.AlertTimeoutSeconds = &AlertTimeoutSecondsInt
	}
	
	if Pin, ok := VoicemailuserpolicyMap["pin"].(string); ok {
		o.Pin = &Pin
	}
    
	if modifiedDateString, ok := VoicemailuserpolicyMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	
	if SendEmailNotifications, ok := VoicemailuserpolicyMap["sendEmailNotifications"].(bool); ok {
		o.SendEmailNotifications = &SendEmailNotifications
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Voicemailuserpolicy) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

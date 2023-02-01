package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationthreadingwindow
type Conversationthreadingwindow struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Settings - The conversation threading window timeout (Minutes) for each messaging type
	Settings *[]Conversationthreadingwindowsetting `json:"settings,omitempty"`

	// DefaultTimeoutMinutes - The default conversation threading window timeout (Minutes)
	DefaultTimeoutMinutes *int `json:"defaultTimeoutMinutes,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationthreadingwindow) SetField(field string, fieldValue interface{}) {
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

func (o Conversationthreadingwindow) MarshalJSON() ([]byte, error) {
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
	type Alias Conversationthreadingwindow
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Settings *[]Conversationthreadingwindowsetting `json:"settings,omitempty"`
		
		DefaultTimeoutMinutes *int `json:"defaultTimeoutMinutes,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Settings: o.Settings,
		
		DefaultTimeoutMinutes: o.DefaultTimeoutMinutes,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationthreadingwindow) UnmarshalJSON(b []byte) error {
	var ConversationthreadingwindowMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationthreadingwindowMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationthreadingwindowMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Settings, ok := ConversationthreadingwindowMap["settings"].([]interface{}); ok {
		SettingsString, _ := json.Marshal(Settings)
		json.Unmarshal(SettingsString, &o.Settings)
	}
	
	if DefaultTimeoutMinutes, ok := ConversationthreadingwindowMap["defaultTimeoutMinutes"].(float64); ok {
		DefaultTimeoutMinutesInt := int(DefaultTimeoutMinutes)
		o.DefaultTimeoutMinutes = &DefaultTimeoutMinutesInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationthreadingwindow) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

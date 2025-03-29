package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Emailthreadingsettings
type Emailthreadingsettings struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// StartNewConversationOnSubjectChange - This setting controls whether a new conversation is started if the subject of an inbound email is different from the subject of the current conversation. RE: and FWD: prefixes in any language are ignored.
	StartNewConversationOnSubjectChange *bool `json:"startNewConversationOnSubjectChange,omitempty"`

	// TimeoutInMinutes - In minutes, how long an email conversation should keep being threaded after being disconnected.
	TimeoutInMinutes *int `json:"timeoutInMinutes,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Emailthreadingsettings) SetField(field string, fieldValue interface{}) {
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

func (o Emailthreadingsettings) MarshalJSON() ([]byte, error) {
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
	type Alias Emailthreadingsettings
	
	return json.Marshal(&struct { 
		StartNewConversationOnSubjectChange *bool `json:"startNewConversationOnSubjectChange,omitempty"`
		
		TimeoutInMinutes *int `json:"timeoutInMinutes,omitempty"`
		Alias
	}{ 
		StartNewConversationOnSubjectChange: o.StartNewConversationOnSubjectChange,
		
		TimeoutInMinutes: o.TimeoutInMinutes,
		Alias:    (Alias)(o),
	})
}

func (o *Emailthreadingsettings) UnmarshalJSON(b []byte) error {
	var EmailthreadingsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &EmailthreadingsettingsMap)
	if err != nil {
		return err
	}
	
	if StartNewConversationOnSubjectChange, ok := EmailthreadingsettingsMap["startNewConversationOnSubjectChange"].(bool); ok {
		o.StartNewConversationOnSubjectChange = &StartNewConversationOnSubjectChange
	}
    
	if TimeoutInMinutes, ok := EmailthreadingsettingsMap["timeoutInMinutes"].(float64); ok {
		TimeoutInMinutesInt := int(TimeoutInMinutes)
		o.TimeoutInMinutes = &TimeoutInMinutesInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Emailthreadingsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

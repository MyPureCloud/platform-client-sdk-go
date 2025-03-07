package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Queuemediasettings
type Queuemediasettings struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Call - The queue media settings for call interactions.
	Call *Mediasettings `json:"call,omitempty"`

	// Callback - The queue media settings for callback interactions.
	Callback *Callbackmediasettings `json:"callback,omitempty"`

	// Chat - The queue media settings for chat interactions.
	Chat *Mediasettings `json:"chat,omitempty"`

	// Email - The queue media settings for email interactions.
	Email *Emailmediasettings `json:"email,omitempty"`

	// Message - The queue media settings for message interactions.
	Message *Messagemediasettings `json:"message,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Queuemediasettings) SetField(field string, fieldValue interface{}) {
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

func (o Queuemediasettings) MarshalJSON() ([]byte, error) {
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
	type Alias Queuemediasettings
	
	return json.Marshal(&struct { 
		Call *Mediasettings `json:"call,omitempty"`
		
		Callback *Callbackmediasettings `json:"callback,omitempty"`
		
		Chat *Mediasettings `json:"chat,omitempty"`
		
		Email *Emailmediasettings `json:"email,omitempty"`
		
		Message *Messagemediasettings `json:"message,omitempty"`
		Alias
	}{ 
		Call: o.Call,
		
		Callback: o.Callback,
		
		Chat: o.Chat,
		
		Email: o.Email,
		
		Message: o.Message,
		Alias:    (Alias)(o),
	})
}

func (o *Queuemediasettings) UnmarshalJSON(b []byte) error {
	var QueuemediasettingsMap map[string]interface{}
	err := json.Unmarshal(b, &QueuemediasettingsMap)
	if err != nil {
		return err
	}
	
	if Call, ok := QueuemediasettingsMap["call"].(map[string]interface{}); ok {
		CallString, _ := json.Marshal(Call)
		json.Unmarshal(CallString, &o.Call)
	}
	
	if Callback, ok := QueuemediasettingsMap["callback"].(map[string]interface{}); ok {
		CallbackString, _ := json.Marshal(Callback)
		json.Unmarshal(CallbackString, &o.Callback)
	}
	
	if Chat, ok := QueuemediasettingsMap["chat"].(map[string]interface{}); ok {
		ChatString, _ := json.Marshal(Chat)
		json.Unmarshal(ChatString, &o.Chat)
	}
	
	if Email, ok := QueuemediasettingsMap["email"].(map[string]interface{}); ok {
		EmailString, _ := json.Marshal(Email)
		json.Unmarshal(EmailString, &o.Email)
	}
	
	if Message, ok := QueuemediasettingsMap["message"].(map[string]interface{}); ok {
		MessageString, _ := json.Marshal(Message)
		json.Unmarshal(MessageString, &o.Message)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queuemediasettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

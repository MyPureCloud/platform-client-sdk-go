package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Mediapolicies
type Mediapolicies struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// CallPolicy - Conditions and actions for calls
	CallPolicy *Callmediapolicy `json:"callPolicy,omitempty"`

	// ChatPolicy - Conditions and actions for chats
	ChatPolicy *Chatmediapolicy `json:"chatPolicy,omitempty"`

	// EmailPolicy - Conditions and actions for emails
	EmailPolicy *Emailmediapolicy `json:"emailPolicy,omitempty"`

	// MessagePolicy - Conditions and actions for messages
	MessagePolicy *Messagemediapolicy `json:"messagePolicy,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Mediapolicies) SetField(field string, fieldValue interface{}) {
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

func (o Mediapolicies) MarshalJSON() ([]byte, error) {
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
	type Alias Mediapolicies
	
	return json.Marshal(&struct { 
		CallPolicy *Callmediapolicy `json:"callPolicy,omitempty"`
		
		ChatPolicy *Chatmediapolicy `json:"chatPolicy,omitempty"`
		
		EmailPolicy *Emailmediapolicy `json:"emailPolicy,omitempty"`
		
		MessagePolicy *Messagemediapolicy `json:"messagePolicy,omitempty"`
		Alias
	}{ 
		CallPolicy: o.CallPolicy,
		
		ChatPolicy: o.ChatPolicy,
		
		EmailPolicy: o.EmailPolicy,
		
		MessagePolicy: o.MessagePolicy,
		Alias:    (Alias)(o),
	})
}

func (o *Mediapolicies) UnmarshalJSON(b []byte) error {
	var MediapoliciesMap map[string]interface{}
	err := json.Unmarshal(b, &MediapoliciesMap)
	if err != nil {
		return err
	}
	
	if CallPolicy, ok := MediapoliciesMap["callPolicy"].(map[string]interface{}); ok {
		CallPolicyString, _ := json.Marshal(CallPolicy)
		json.Unmarshal(CallPolicyString, &o.CallPolicy)
	}
	
	if ChatPolicy, ok := MediapoliciesMap["chatPolicy"].(map[string]interface{}); ok {
		ChatPolicyString, _ := json.Marshal(ChatPolicy)
		json.Unmarshal(ChatPolicyString, &o.ChatPolicy)
	}
	
	if EmailPolicy, ok := MediapoliciesMap["emailPolicy"].(map[string]interface{}); ok {
		EmailPolicyString, _ := json.Marshal(EmailPolicy)
		json.Unmarshal(EmailPolicyString, &o.EmailPolicy)
	}
	
	if MessagePolicy, ok := MediapoliciesMap["messagePolicy"].(map[string]interface{}); ok {
		MessagePolicyString, _ := json.Marshal(MessagePolicy)
		json.Unmarshal(MessagePolicyString, &o.MessagePolicy)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Mediapolicies) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentownedrouting
type Agentownedrouting struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EnableAgentOwnedCallbacks - Indicates if Agent Owned Callbacks are enabled for the queue
	EnableAgentOwnedCallbacks *bool `json:"enableAgentOwnedCallbacks,omitempty"`

	// MaxOwnedCallbackHours - The max amount of time a callback can be owned (in hours); Allowable range 1 - 168 hour(s) (inclusive)
	MaxOwnedCallbackHours *int `json:"maxOwnedCallbackHours,omitempty"`

	// MaxOwnedCallbackDelayHours - The max amount of time a callback can be scheduled out into the future (in hours); Allowable range 1 - 720 hour(s) (inclusive)
	MaxOwnedCallbackDelayHours *int `json:"maxOwnedCallbackDelayHours,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Agentownedrouting) SetField(field string, fieldValue interface{}) {
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

func (o Agentownedrouting) MarshalJSON() ([]byte, error) {
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
	type Alias Agentownedrouting
	
	return json.Marshal(&struct { 
		EnableAgentOwnedCallbacks *bool `json:"enableAgentOwnedCallbacks,omitempty"`
		
		MaxOwnedCallbackHours *int `json:"maxOwnedCallbackHours,omitempty"`
		
		MaxOwnedCallbackDelayHours *int `json:"maxOwnedCallbackDelayHours,omitempty"`
		Alias
	}{ 
		EnableAgentOwnedCallbacks: o.EnableAgentOwnedCallbacks,
		
		MaxOwnedCallbackHours: o.MaxOwnedCallbackHours,
		
		MaxOwnedCallbackDelayHours: o.MaxOwnedCallbackDelayHours,
		Alias:    (Alias)(o),
	})
}

func (o *Agentownedrouting) UnmarshalJSON(b []byte) error {
	var AgentownedroutingMap map[string]interface{}
	err := json.Unmarshal(b, &AgentownedroutingMap)
	if err != nil {
		return err
	}
	
	if EnableAgentOwnedCallbacks, ok := AgentownedroutingMap["enableAgentOwnedCallbacks"].(bool); ok {
		o.EnableAgentOwnedCallbacks = &EnableAgentOwnedCallbacks
	}
    
	if MaxOwnedCallbackHours, ok := AgentownedroutingMap["maxOwnedCallbackHours"].(float64); ok {
		MaxOwnedCallbackHoursInt := int(MaxOwnedCallbackHours)
		o.MaxOwnedCallbackHours = &MaxOwnedCallbackHoursInt
	}
	
	if MaxOwnedCallbackDelayHours, ok := AgentownedroutingMap["maxOwnedCallbackDelayHours"].(float64); ok {
		MaxOwnedCallbackDelayHoursInt := int(MaxOwnedCallbackDelayHours)
		o.MaxOwnedCallbackDelayHours = &MaxOwnedCallbackDelayHoursInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Agentownedrouting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

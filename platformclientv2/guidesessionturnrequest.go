package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Guidesessionturnrequest - Request for a guide session turn
type Guidesessionturnrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// InputEvent - The input event for this turn.
	InputEvent *Guidesessioninputevent `json:"inputEvent,omitempty"`

	// LanguageCode - The language code for this turn.
	LanguageCode *string `json:"languageCode,omitempty"`

	// Version - The version for this turn.
	Version *string `json:"version,omitempty"`

	// InputVariables - The input variables for this turn.
	InputVariables *[]Guidesessionvariable `json:"inputVariables,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Guidesessionturnrequest) SetField(field string, fieldValue interface{}) {
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

func (o Guidesessionturnrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Guidesessionturnrequest
	
	return json.Marshal(&struct { 
		InputEvent *Guidesessioninputevent `json:"inputEvent,omitempty"`
		
		LanguageCode *string `json:"languageCode,omitempty"`
		
		Version *string `json:"version,omitempty"`
		
		InputVariables *[]Guidesessionvariable `json:"inputVariables,omitempty"`
		Alias
	}{ 
		InputEvent: o.InputEvent,
		
		LanguageCode: o.LanguageCode,
		
		Version: o.Version,
		
		InputVariables: o.InputVariables,
		Alias:    (Alias)(o),
	})
}

func (o *Guidesessionturnrequest) UnmarshalJSON(b []byte) error {
	var GuidesessionturnrequestMap map[string]interface{}
	err := json.Unmarshal(b, &GuidesessionturnrequestMap)
	if err != nil {
		return err
	}
	
	if InputEvent, ok := GuidesessionturnrequestMap["inputEvent"].(map[string]interface{}); ok {
		InputEventString, _ := json.Marshal(InputEvent)
		json.Unmarshal(InputEventString, &o.InputEvent)
	}
	
	if LanguageCode, ok := GuidesessionturnrequestMap["languageCode"].(string); ok {
		o.LanguageCode = &LanguageCode
	}
    
	if Version, ok := GuidesessionturnrequestMap["version"].(string); ok {
		o.Version = &Version
	}
    
	if InputVariables, ok := GuidesessionturnrequestMap["inputVariables"].([]interface{}); ok {
		InputVariablesString, _ := json.Marshal(InputVariables)
		json.Unmarshal(InputVariablesString, &o.InputVariables)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Guidesessionturnrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

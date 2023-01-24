package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotflowturnrequest - Settings for a turn request to a bot flow.
type Textbotflowturnrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// PreviousTurn - The reference to a previous turn if appropriate, used to avoid race conditions.
	PreviousTurn *Textbotturnreference `json:"previousTurn,omitempty"`

	// InputEventType - Indicates the type of input event being requested. If appropriate, fill out the matching user input object details on this request.
	InputEventType *string `json:"inputEventType,omitempty"`

	// InputEventUserInput - The data for the input event of this turn if it is a user input event. Only one inputEvent may be set.
	InputEventUserInput *Textbotuserinputevent `json:"inputEventUserInput,omitempty"`

	// InputEventError - The data for the input event of this turn if it is an error event. Only one inputEvent may be set.
	InputEventError *Textboterrorinputevent `json:"inputEventError,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Textbotflowturnrequest) SetField(field string, fieldValue interface{}) {
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

func (o Textbotflowturnrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Textbotflowturnrequest
	
	return json.Marshal(&struct { 
		PreviousTurn *Textbotturnreference `json:"previousTurn,omitempty"`
		
		InputEventType *string `json:"inputEventType,omitempty"`
		
		InputEventUserInput *Textbotuserinputevent `json:"inputEventUserInput,omitempty"`
		
		InputEventError *Textboterrorinputevent `json:"inputEventError,omitempty"`
		Alias
	}{ 
		PreviousTurn: o.PreviousTurn,
		
		InputEventType: o.InputEventType,
		
		InputEventUserInput: o.InputEventUserInput,
		
		InputEventError: o.InputEventError,
		Alias:    (Alias)(o),
	})
}

func (o *Textbotflowturnrequest) UnmarshalJSON(b []byte) error {
	var TextbotflowturnrequestMap map[string]interface{}
	err := json.Unmarshal(b, &TextbotflowturnrequestMap)
	if err != nil {
		return err
	}
	
	if PreviousTurn, ok := TextbotflowturnrequestMap["previousTurn"].(map[string]interface{}); ok {
		PreviousTurnString, _ := json.Marshal(PreviousTurn)
		json.Unmarshal(PreviousTurnString, &o.PreviousTurn)
	}
	
	if InputEventType, ok := TextbotflowturnrequestMap["inputEventType"].(string); ok {
		o.InputEventType = &InputEventType
	}
    
	if InputEventUserInput, ok := TextbotflowturnrequestMap["inputEventUserInput"].(map[string]interface{}); ok {
		InputEventUserInputString, _ := json.Marshal(InputEventUserInput)
		json.Unmarshal(InputEventUserInputString, &o.InputEventUserInput)
	}
	
	if InputEventError, ok := TextbotflowturnrequestMap["inputEventError"].(map[string]interface{}); ok {
		InputEventErrorString, _ := json.Marshal(InputEventError)
		json.Unmarshal(InputEventErrorString, &o.InputEventError)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Textbotflowturnrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

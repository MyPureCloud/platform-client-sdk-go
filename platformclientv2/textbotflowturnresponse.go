package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotflowturnresponse - Information related to a success bot flow turn request.
type Textbotflowturnresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The ID of the bot flow turn. If additional turns are needed, supply this ID as the previous turn in your next turn request.
	Id *string `json:"id,omitempty"`

	// PreviousTurn - The reference to a previous turn, if applicable.
	PreviousTurn *Textbotturnreference `json:"previousTurn,omitempty"`

	// Prompts - The output prompts for this turn.
	Prompts *Textbotoutputprompts `json:"prompts,omitempty"`

	// NextActionType - Indicates the suggested next action. If appropriate, the matching output event object includes additional information.
	NextActionType *string `json:"nextActionType,omitempty"`

	// NextActionDisconnect - The next action directive for this turn if it is a Disconnect type.
	NextActionDisconnect *Textbotdisconnectaction `json:"nextActionDisconnect,omitempty"`

	// NextActionWaitForInput - The next action directive for this turn if it is a WaitForInput type.
	NextActionWaitForInput *Textbotwaitforinputaction `json:"nextActionWaitForInput,omitempty"`

	// NextActionExit - The next action directive for this turn if it is an Exit type.
	NextActionExit *Textbotexitaction `json:"nextActionExit,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Textbotflowturnresponse) SetField(field string, fieldValue interface{}) {
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

func (o Textbotflowturnresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Textbotflowturnresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		PreviousTurn *Textbotturnreference `json:"previousTurn,omitempty"`
		
		Prompts *Textbotoutputprompts `json:"prompts,omitempty"`
		
		NextActionType *string `json:"nextActionType,omitempty"`
		
		NextActionDisconnect *Textbotdisconnectaction `json:"nextActionDisconnect,omitempty"`
		
		NextActionWaitForInput *Textbotwaitforinputaction `json:"nextActionWaitForInput,omitempty"`
		
		NextActionExit *Textbotexitaction `json:"nextActionExit,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		PreviousTurn: o.PreviousTurn,
		
		Prompts: o.Prompts,
		
		NextActionType: o.NextActionType,
		
		NextActionDisconnect: o.NextActionDisconnect,
		
		NextActionWaitForInput: o.NextActionWaitForInput,
		
		NextActionExit: o.NextActionExit,
		Alias:    (Alias)(o),
	})
}

func (o *Textbotflowturnresponse) UnmarshalJSON(b []byte) error {
	var TextbotflowturnresponseMap map[string]interface{}
	err := json.Unmarshal(b, &TextbotflowturnresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TextbotflowturnresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if PreviousTurn, ok := TextbotflowturnresponseMap["previousTurn"].(map[string]interface{}); ok {
		PreviousTurnString, _ := json.Marshal(PreviousTurn)
		json.Unmarshal(PreviousTurnString, &o.PreviousTurn)
	}
	
	if Prompts, ok := TextbotflowturnresponseMap["prompts"].(map[string]interface{}); ok {
		PromptsString, _ := json.Marshal(Prompts)
		json.Unmarshal(PromptsString, &o.Prompts)
	}
	
	if NextActionType, ok := TextbotflowturnresponseMap["nextActionType"].(string); ok {
		o.NextActionType = &NextActionType
	}
    
	if NextActionDisconnect, ok := TextbotflowturnresponseMap["nextActionDisconnect"].(map[string]interface{}); ok {
		NextActionDisconnectString, _ := json.Marshal(NextActionDisconnect)
		json.Unmarshal(NextActionDisconnectString, &o.NextActionDisconnect)
	}
	
	if NextActionWaitForInput, ok := TextbotflowturnresponseMap["nextActionWaitForInput"].(map[string]interface{}); ok {
		NextActionWaitForInputString, _ := json.Marshal(NextActionWaitForInput)
		json.Unmarshal(NextActionWaitForInputString, &o.NextActionWaitForInput)
	}
	
	if NextActionExit, ok := TextbotflowturnresponseMap["nextActionExit"].(map[string]interface{}); ok {
		NextActionExitString, _ := json.Marshal(NextActionExit)
		json.Unmarshal(NextActionExitString, &o.NextActionExit)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Textbotflowturnresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

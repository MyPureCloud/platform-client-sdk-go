package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Agenticvirtualagentrepetitioncheck - A rule that detects repeated user or agent messages and adds a corrective instruction.
type Agenticvirtualagentrepetitioncheck struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType - Whether this check looks for repetition in user messages or agent responses.
	VarType *string `json:"type,omitempty"`

	// Messages - The number of prior messages of the specified type to compare for repetition.
	Messages *int `json:"messages,omitempty"`

	// Similarity - The similarity category compared to the Levenshtein result that triggers this check's instruction.
	Similarity *string `json:"similarity,omitempty"`

	// Instruction - The instruction added to the virtual agent's turn when message similarity matches the configured category.
	Instruction *string `json:"instruction,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Agenticvirtualagentrepetitioncheck) SetField(field string, fieldValue interface{}) {
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

func (o Agenticvirtualagentrepetitioncheck) MarshalJSON() ([]byte, error) {
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
	type Alias Agenticvirtualagentrepetitioncheck
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Messages *int `json:"messages,omitempty"`
		
		Similarity *string `json:"similarity,omitempty"`
		
		Instruction *string `json:"instruction,omitempty"`
		Alias
	}{ 
		VarType: o.VarType,
		
		Messages: o.Messages,
		
		Similarity: o.Similarity,
		
		Instruction: o.Instruction,
		Alias:    (Alias)(o),
	})
}

func (o *Agenticvirtualagentrepetitioncheck) UnmarshalJSON(b []byte) error {
	var AgenticvirtualagentrepetitioncheckMap map[string]interface{}
	err := json.Unmarshal(b, &AgenticvirtualagentrepetitioncheckMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := AgenticvirtualagentrepetitioncheckMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Messages, ok := AgenticvirtualagentrepetitioncheckMap["messages"].(float64); ok {
		MessagesInt := int(Messages)
		o.Messages = &MessagesInt
	}
	
	if Similarity, ok := AgenticvirtualagentrepetitioncheckMap["similarity"].(string); ok {
		o.Similarity = &Similarity
	}
    
	if Instruction, ok := AgenticvirtualagentrepetitioncheckMap["instruction"].(string); ok {
		o.Instruction = &Instruction
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentrepetitioncheck) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

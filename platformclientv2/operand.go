package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Operand
type Operand struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType - The Operand type of the category
	VarType *string `json:"type,omitempty"`

	// Occurrence - The minimum number of occurrences of the defined operand type
	Occurrence *int `json:"occurrence,omitempty"`

	// Inverted - Applies a NOT modifier to the operand group
	Inverted *bool `json:"inverted,omitempty"`

	// Term - Filter interaction by word(s)
	Term *Term `json:"term,omitempty"`

	// TopicId - Filter interaction by topic ID
	TopicId *string `json:"topicId,omitempty"`

	// VoiceSecondsPosition - Dictates when the operand must occur in a voice interaction
	VoiceSecondsPosition *Operandposition `json:"voiceSecondsPosition,omitempty"`

	// DigitalWordsPosition - Dictates when the operand must occur in a digital interaction
	DigitalWordsPosition *Operandposition `json:"digitalWordsPosition,omitempty"`

	// InfixOperator - Defines a logical operation that is applied on the current operand, against the following operand
	InfixOperator *Infixoperator `json:"infixOperator,omitempty"`

	// Operands - Contains a new level of operands
	Operands *[]Operand `json:"operands,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Operand) SetField(field string, fieldValue interface{}) {
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

func (o Operand) MarshalJSON() ([]byte, error) {
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
	type Alias Operand
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Occurrence *int `json:"occurrence,omitempty"`
		
		Inverted *bool `json:"inverted,omitempty"`
		
		Term *Term `json:"term,omitempty"`
		
		TopicId *string `json:"topicId,omitempty"`
		
		VoiceSecondsPosition *Operandposition `json:"voiceSecondsPosition,omitempty"`
		
		DigitalWordsPosition *Operandposition `json:"digitalWordsPosition,omitempty"`
		
		InfixOperator *Infixoperator `json:"infixOperator,omitempty"`
		
		Operands *[]Operand `json:"operands,omitempty"`
		Alias
	}{ 
		VarType: o.VarType,
		
		Occurrence: o.Occurrence,
		
		Inverted: o.Inverted,
		
		Term: o.Term,
		
		TopicId: o.TopicId,
		
		VoiceSecondsPosition: o.VoiceSecondsPosition,
		
		DigitalWordsPosition: o.DigitalWordsPosition,
		
		InfixOperator: o.InfixOperator,
		
		Operands: o.Operands,
		Alias:    (Alias)(o),
	})
}

func (o *Operand) UnmarshalJSON(b []byte) error {
	var OperandMap map[string]interface{}
	err := json.Unmarshal(b, &OperandMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := OperandMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Occurrence, ok := OperandMap["occurrence"].(float64); ok {
		OccurrenceInt := int(Occurrence)
		o.Occurrence = &OccurrenceInt
	}
	
	if Inverted, ok := OperandMap["inverted"].(bool); ok {
		o.Inverted = &Inverted
	}
    
	if Term, ok := OperandMap["term"].(map[string]interface{}); ok {
		TermString, _ := json.Marshal(Term)
		json.Unmarshal(TermString, &o.Term)
	}
	
	if TopicId, ok := OperandMap["topicId"].(string); ok {
		o.TopicId = &TopicId
	}
    
	if VoiceSecondsPosition, ok := OperandMap["voiceSecondsPosition"].(map[string]interface{}); ok {
		VoiceSecondsPositionString, _ := json.Marshal(VoiceSecondsPosition)
		json.Unmarshal(VoiceSecondsPositionString, &o.VoiceSecondsPosition)
	}
	
	if DigitalWordsPosition, ok := OperandMap["digitalWordsPosition"].(map[string]interface{}); ok {
		DigitalWordsPositionString, _ := json.Marshal(DigitalWordsPosition)
		json.Unmarshal(DigitalWordsPositionString, &o.DigitalWordsPosition)
	}
	
	if InfixOperator, ok := OperandMap["infixOperator"].(map[string]interface{}); ok {
		InfixOperatorString, _ := json.Marshal(InfixOperator)
		json.Unmarshal(InfixOperatorString, &o.InfixOperator)
	}
	
	if Operands, ok := OperandMap["operands"].([]interface{}); ok {
		OperandsString, _ := json.Marshal(Operands)
		json.Unmarshal(OperandsString, &o.Operands)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Operand) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

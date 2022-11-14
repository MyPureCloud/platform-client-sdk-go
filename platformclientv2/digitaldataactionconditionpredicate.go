package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Digitaldataactionconditionpredicate
type Digitaldataactionconditionpredicate struct { 
	// OutputField - The name of an output field from the data action's output to use for this condition
	OutputField *string `json:"outputField,omitempty"`


	// OutputOperator - The operation with which to evaluate this condition
	OutputOperator *string `json:"outputOperator,omitempty"`


	// ComparisonValue - The value to compare against for this condition
	ComparisonValue *string `json:"comparisonValue,omitempty"`


	// Inverted - If true, inverts the result of evaluating this Predicate. Default is false.
	Inverted *bool `json:"inverted,omitempty"`


	// OutputFieldMissingResolution - The result of this predicate if the requested output field is missing from the data action's result
	OutputFieldMissingResolution *bool `json:"outputFieldMissingResolution,omitempty"`


	// ValueType - The data type the value should be treated as.
	ValueType *string `json:"valueType,omitempty"`

}

func (o *Digitaldataactionconditionpredicate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Digitaldataactionconditionpredicate
	
	return json.Marshal(&struct { 
		OutputField *string `json:"outputField,omitempty"`
		
		OutputOperator *string `json:"outputOperator,omitempty"`
		
		ComparisonValue *string `json:"comparisonValue,omitempty"`
		
		Inverted *bool `json:"inverted,omitempty"`
		
		OutputFieldMissingResolution *bool `json:"outputFieldMissingResolution,omitempty"`
		
		ValueType *string `json:"valueType,omitempty"`
		*Alias
	}{ 
		OutputField: o.OutputField,
		
		OutputOperator: o.OutputOperator,
		
		ComparisonValue: o.ComparisonValue,
		
		Inverted: o.Inverted,
		
		OutputFieldMissingResolution: o.OutputFieldMissingResolution,
		
		ValueType: o.ValueType,
		Alias:    (*Alias)(o),
	})
}

func (o *Digitaldataactionconditionpredicate) UnmarshalJSON(b []byte) error {
	var DigitaldataactionconditionpredicateMap map[string]interface{}
	err := json.Unmarshal(b, &DigitaldataactionconditionpredicateMap)
	if err != nil {
		return err
	}
	
	if OutputField, ok := DigitaldataactionconditionpredicateMap["outputField"].(string); ok {
		o.OutputField = &OutputField
	}
    
	if OutputOperator, ok := DigitaldataactionconditionpredicateMap["outputOperator"].(string); ok {
		o.OutputOperator = &OutputOperator
	}
    
	if ComparisonValue, ok := DigitaldataactionconditionpredicateMap["comparisonValue"].(string); ok {
		o.ComparisonValue = &ComparisonValue
	}
    
	if Inverted, ok := DigitaldataactionconditionpredicateMap["inverted"].(bool); ok {
		o.Inverted = &Inverted
	}
    
	if OutputFieldMissingResolution, ok := DigitaldataactionconditionpredicateMap["outputFieldMissingResolution"].(bool); ok {
		o.OutputFieldMissingResolution = &OutputFieldMissingResolution
	}
    
	if ValueType, ok := DigitaldataactionconditionpredicateMap["valueType"].(string); ok {
		o.ValueType = &ValueType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Digitaldataactionconditionpredicate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

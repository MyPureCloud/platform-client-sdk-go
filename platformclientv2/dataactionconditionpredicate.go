package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dataactionconditionpredicate
type Dataactionconditionpredicate struct { 
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

}

func (o *Dataactionconditionpredicate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dataactionconditionpredicate
	
	return json.Marshal(&struct { 
		OutputField *string `json:"outputField,omitempty"`
		
		OutputOperator *string `json:"outputOperator,omitempty"`
		
		ComparisonValue *string `json:"comparisonValue,omitempty"`
		
		Inverted *bool `json:"inverted,omitempty"`
		
		OutputFieldMissingResolution *bool `json:"outputFieldMissingResolution,omitempty"`
		*Alias
	}{ 
		OutputField: o.OutputField,
		
		OutputOperator: o.OutputOperator,
		
		ComparisonValue: o.ComparisonValue,
		
		Inverted: o.Inverted,
		
		OutputFieldMissingResolution: o.OutputFieldMissingResolution,
		Alias:    (*Alias)(o),
	})
}

func (o *Dataactionconditionpredicate) UnmarshalJSON(b []byte) error {
	var DataactionconditionpredicateMap map[string]interface{}
	err := json.Unmarshal(b, &DataactionconditionpredicateMap)
	if err != nil {
		return err
	}
	
	if OutputField, ok := DataactionconditionpredicateMap["outputField"].(string); ok {
		o.OutputField = &OutputField
	}
    
	if OutputOperator, ok := DataactionconditionpredicateMap["outputOperator"].(string); ok {
		o.OutputOperator = &OutputOperator
	}
    
	if ComparisonValue, ok := DataactionconditionpredicateMap["comparisonValue"].(string); ok {
		o.ComparisonValue = &ComparisonValue
	}
    
	if Inverted, ok := DataactionconditionpredicateMap["inverted"].(bool); ok {
		o.Inverted = &Inverted
	}
    
	if OutputFieldMissingResolution, ok := DataactionconditionpredicateMap["outputFieldMissingResolution"].(bool); ok {
		o.OutputFieldMissingResolution = &OutputFieldMissingResolution
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Dataactionconditionpredicate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

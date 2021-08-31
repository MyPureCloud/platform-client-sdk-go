package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerrulesetconfigchangedataactionconditionpredicate
type Dialerrulesetconfigchangedataactionconditionpredicate struct { 
	// OutputField
	OutputField *string `json:"outputField,omitempty"`


	// OutputOperator
	OutputOperator *string `json:"outputOperator,omitempty"`


	// ComparisonValue
	ComparisonValue *string `json:"comparisonValue,omitempty"`


	// OutputFieldMissingResolution
	OutputFieldMissingResolution *bool `json:"outputFieldMissingResolution,omitempty"`


	// Inverted
	Inverted *bool `json:"inverted,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (o *Dialerrulesetconfigchangedataactionconditionpredicate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialerrulesetconfigchangedataactionconditionpredicate
	
	return json.Marshal(&struct { 
		OutputField *string `json:"outputField,omitempty"`
		
		OutputOperator *string `json:"outputOperator,omitempty"`
		
		ComparisonValue *string `json:"comparisonValue,omitempty"`
		
		OutputFieldMissingResolution *bool `json:"outputFieldMissingResolution,omitempty"`
		
		Inverted *bool `json:"inverted,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		OutputField: o.OutputField,
		
		OutputOperator: o.OutputOperator,
		
		ComparisonValue: o.ComparisonValue,
		
		OutputFieldMissingResolution: o.OutputFieldMissingResolution,
		
		Inverted: o.Inverted,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialerrulesetconfigchangedataactionconditionpredicate) UnmarshalJSON(b []byte) error {
	var DialerrulesetconfigchangedataactionconditionpredicateMap map[string]interface{}
	err := json.Unmarshal(b, &DialerrulesetconfigchangedataactionconditionpredicateMap)
	if err != nil {
		return err
	}
	
	if OutputField, ok := DialerrulesetconfigchangedataactionconditionpredicateMap["outputField"].(string); ok {
		o.OutputField = &OutputField
	}
	
	if OutputOperator, ok := DialerrulesetconfigchangedataactionconditionpredicateMap["outputOperator"].(string); ok {
		o.OutputOperator = &OutputOperator
	}
	
	if ComparisonValue, ok := DialerrulesetconfigchangedataactionconditionpredicateMap["comparisonValue"].(string); ok {
		o.ComparisonValue = &ComparisonValue
	}
	
	if OutputFieldMissingResolution, ok := DialerrulesetconfigchangedataactionconditionpredicateMap["outputFieldMissingResolution"].(bool); ok {
		o.OutputFieldMissingResolution = &OutputFieldMissingResolution
	}
	
	if Inverted, ok := DialerrulesetconfigchangedataactionconditionpredicateMap["inverted"].(bool); ok {
		o.Inverted = &Inverted
	}
	
	if AdditionalProperties, ok := DialerrulesetconfigchangedataactionconditionpredicateMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialerrulesetconfigchangedataactionconditionpredicate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

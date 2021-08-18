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

func (u *Dialerrulesetconfigchangedataactionconditionpredicate) MarshalJSON() ([]byte, error) {
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
		OutputField: u.OutputField,
		
		OutputOperator: u.OutputOperator,
		
		ComparisonValue: u.ComparisonValue,
		
		OutputFieldMissingResolution: u.OutputFieldMissingResolution,
		
		Inverted: u.Inverted,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialerrulesetconfigchangedataactionconditionpredicate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

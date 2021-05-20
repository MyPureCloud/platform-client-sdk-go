package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Dialerrulesetconfigchangedataactionconditionpredicate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

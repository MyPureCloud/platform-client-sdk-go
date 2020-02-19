package platformclientv2
import (
	"encoding/json"
)

// Suggestsearchcriteria
type Suggestsearchcriteria struct { 
	// EndValue - The end value of the range. This field is used for range search types.
	EndValue *string `json:"endValue,omitempty"`


	// Values - A list of values for the search to match against
	Values *[]string `json:"values,omitempty"`


	// StartValue - The start value of the range. This field is used for range search types.
	StartValue *string `json:"startValue,omitempty"`


	// Fields - Field names to search against
	Fields *[]string `json:"fields,omitempty"`


	// Value - A value for the search to match against
	Value *string `json:"value,omitempty"`


	// Operator - How to apply this search criteria against other criteria
	Operator *string `json:"operator,omitempty"`


	// Group - Groups multiple conditions
	Group *[]Suggestsearchcriteria `json:"group,omitempty"`

}

// String returns a JSON representation of the model
func (o *Suggestsearchcriteria) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
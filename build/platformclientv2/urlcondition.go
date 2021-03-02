package platformclientv2
import (
	"encoding/json"
)

// Urlcondition
type Urlcondition struct { 
	// Values - The URL condition value.
	Values *[]string `json:"values,omitempty"`


	// Operator - The comparison operator.
	Operator *string `json:"operator,omitempty"`

}

// String returns a JSON representation of the model
func (o *Urlcondition) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

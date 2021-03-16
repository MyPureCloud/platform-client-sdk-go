package platformclientv2
import (
	"encoding/json"
)

// Botaggregatequeryclause
type Botaggregatequeryclause struct { 
	// VarType - Boolean operation to apply to the provided predicates
	VarType *string `json:"type,omitempty"`


	// Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
	Predicates *[]Botaggregatequerypredicate `json:"predicates,omitempty"`

}

// String returns a JSON representation of the model
func (o *Botaggregatequeryclause) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

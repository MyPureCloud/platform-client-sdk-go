package platformclientv2
import (
	"encoding/json"
)

// Useraggregatequeryclause
type Useraggregatequeryclause struct { 
	// VarType - Boolean operation to apply to the provided predicates
	VarType *string `json:"type,omitempty"`


	// Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
	Predicates *[]Useraggregatequerypredicate `json:"predicates,omitempty"`

}

// String returns a JSON representation of the model
func (o *Useraggregatequeryclause) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

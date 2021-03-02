package platformclientv2
import (
	"encoding/json"
)

// Flowobservationqueryclause
type Flowobservationqueryclause struct { 
	// VarType - Boolean operation to apply to the provided predicates
	VarType *string `json:"type,omitempty"`


	// Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
	Predicates *[]Flowobservationquerypredicate `json:"predicates,omitempty"`

}

// String returns a JSON representation of the model
func (o *Flowobservationqueryclause) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

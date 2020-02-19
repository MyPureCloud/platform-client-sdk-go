package platformclientv2
import (
	"encoding/json"
)

// Queueobservationqueryclause
type Queueobservationqueryclause struct { 
	// VarType - Boolean operation to apply to the provided predicates
	VarType *string `json:"type,omitempty"`


	// Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
	Predicates *[]Queueobservationquerypredicate `json:"predicates,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueobservationqueryclause) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

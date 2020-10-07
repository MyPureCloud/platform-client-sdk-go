package platformclientv2
import (
	"encoding/json"
)

// Transcriptaggregatequeryclause
type Transcriptaggregatequeryclause struct { 
	// VarType - Boolean operation to apply to the provided predicates
	VarType *string `json:"type,omitempty"`


	// Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
	Predicates *[]Transcriptaggregatequerypredicate `json:"predicates,omitempty"`

}

// String returns a JSON representation of the model
func (o *Transcriptaggregatequeryclause) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

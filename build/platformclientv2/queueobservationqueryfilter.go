package platformclientv2
import (
	"encoding/json"
)

// Queueobservationqueryfilter
type Queueobservationqueryfilter struct { 
	// VarType - Boolean operation to apply to the provided predicates and clauses
	VarType *string `json:"type,omitempty"`


	// Clauses - Boolean 'and/or' logic with up to two-levels of nesting
	Clauses *[]Queueobservationqueryclause `json:"clauses,omitempty"`


	// Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
	Predicates *[]Queueobservationquerypredicate `json:"predicates,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueobservationqueryfilter) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

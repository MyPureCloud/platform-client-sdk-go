package platformclientv2
import (
	"encoding/json"
)

// Journeyaggregatequeryfilter
type Journeyaggregatequeryfilter struct { 
	// VarType - Boolean operation to apply to the provided predicates and clauses
	VarType *string `json:"type,omitempty"`


	// Clauses - Boolean 'and/or' logic with up to two-levels of nesting
	Clauses *[]Journeyaggregatequeryclause `json:"clauses,omitempty"`


	// Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
	Predicates *[]Journeyaggregatequerypredicate `json:"predicates,omitempty"`

}

// String returns a JSON representation of the model
func (o *Journeyaggregatequeryfilter) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

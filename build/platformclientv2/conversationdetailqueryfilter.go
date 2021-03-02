package platformclientv2
import (
	"encoding/json"
)

// Conversationdetailqueryfilter
type Conversationdetailqueryfilter struct { 
	// VarType - Boolean operation to apply to the provided predicates and clauses
	VarType *string `json:"type,omitempty"`


	// Clauses - Boolean 'and/or' logic with up to two-levels of nesting
	Clauses *[]Conversationdetailqueryclause `json:"clauses,omitempty"`


	// Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
	Predicates *[]Conversationdetailquerypredicate `json:"predicates,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationdetailqueryfilter) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

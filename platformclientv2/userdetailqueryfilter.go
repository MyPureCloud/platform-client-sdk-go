package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Userdetailqueryfilter
type Userdetailqueryfilter struct { 
	// VarType - Boolean operation to apply to the provided predicates and clauses
	VarType *string `json:"type,omitempty"`


	// Clauses - Boolean 'and/or' logic with up to two-levels of nesting
	Clauses *[]Userdetailqueryclause `json:"clauses,omitempty"`


	// Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
	Predicates *[]Userdetailquerypredicate `json:"predicates,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userdetailqueryfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

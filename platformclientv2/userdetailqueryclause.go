package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Userdetailqueryclause
type Userdetailqueryclause struct { 
	// VarType - Boolean operation to apply to the provided predicates
	VarType *string `json:"type,omitempty"`


	// Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
	Predicates *[]Userdetailquerypredicate `json:"predicates,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userdetailqueryclause) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

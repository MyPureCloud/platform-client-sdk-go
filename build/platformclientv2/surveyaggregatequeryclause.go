package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Surveyaggregatequeryclause
type Surveyaggregatequeryclause struct { 
	// VarType - Boolean operation to apply to the provided predicates
	VarType *string `json:"type,omitempty"`


	// Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
	Predicates *[]Surveyaggregatequerypredicate `json:"predicates,omitempty"`

}

// String returns a JSON representation of the model
func (o *Surveyaggregatequeryclause) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

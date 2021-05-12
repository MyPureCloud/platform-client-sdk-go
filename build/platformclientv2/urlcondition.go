package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Urlcondition
type Urlcondition struct { 
	// Values - The URL condition value.
	Values *[]string `json:"values,omitempty"`


	// Operator - The comparison operator.
	Operator *string `json:"operator,omitempty"`

}

// String returns a JSON representation of the model
func (o *Urlcondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

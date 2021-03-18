package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Responsefilter - Used to filter response queries
type Responsefilter struct { 
	// Name - Field to filter on. Allowed values are 'name' and 'libraryId.
	Name *string `json:"name,omitempty"`


	// Operator - Filter operation: IN, EQUALS, NOTEQUALS.
	Operator *string `json:"operator,omitempty"`


	// Values - Values to filter on.
	Values *[]string `json:"values,omitempty"`

}

// String returns a JSON representation of the model
func (o *Responsefilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

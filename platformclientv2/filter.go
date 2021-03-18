package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Filter
type Filter struct { 
	// Name - The name of the field by which to filter.
	Name *string `json:"name,omitempty"`


	// VarType - The type of the filter, DATE or STRING.
	VarType *string `json:"type,omitempty"`


	// Operator - The operation that the filter performs.
	Operator *string `json:"operator,omitempty"`


	// Values - The values to make the filter comparison against.
	Values *[]string `json:"values,omitempty"`

}

// String returns a JSON representation of the model
func (o *Filter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

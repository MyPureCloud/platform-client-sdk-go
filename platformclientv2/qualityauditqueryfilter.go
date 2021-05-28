package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Qualityauditqueryfilter
type Qualityauditqueryfilter struct { 
	// Property - Name of the property to filter.
	Property *string `json:"property,omitempty"`


	// Value - Value of the property to filter.
	Value *string `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Qualityauditqueryfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

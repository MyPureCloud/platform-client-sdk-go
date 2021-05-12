package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Propertychange
type Propertychange struct { 
	// Property - The property that was changed
	Property *string `json:"property,omitempty"`


	// OldValues - Previous values for the property.
	OldValues *[]string `json:"oldValues,omitempty"`


	// NewValues - New values for the property.
	NewValues *[]string `json:"newValues,omitempty"`

}

// String returns a JSON representation of the model
func (o *Propertychange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

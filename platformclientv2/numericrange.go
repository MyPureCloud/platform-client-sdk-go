package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Numericrange
type Numericrange struct { 
	// Gt - Greater than
	Gt *float32 `json:"gt,omitempty"`


	// Gte - Greater than or equal to
	Gte *float32 `json:"gte,omitempty"`


	// Lt - Less than
	Lt *float32 `json:"lt,omitempty"`


	// Lte - Less than or equal to
	Lte *float32 `json:"lte,omitempty"`

}

// String returns a JSON representation of the model
func (o *Numericrange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

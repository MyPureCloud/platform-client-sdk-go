package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Aggregationrange
type Aggregationrange struct { 
	// Gte - Greater than or equal to
	Gte *float32 `json:"gte,omitempty"`


	// Lt - Less than
	Lt *float32 `json:"lt,omitempty"`

}

// String returns a JSON representation of the model
func (o *Aggregationrange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Routingconversationattributes
type Routingconversationattributes struct { 
	// Priority
	Priority *int `json:"priority,omitempty"`

}

// String returns a JSON representation of the model
func (o *Routingconversationattributes) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

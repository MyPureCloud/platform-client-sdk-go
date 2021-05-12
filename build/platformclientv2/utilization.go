package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Utilization
type Utilization struct { 
	// Utilization - Map of media type to utilization settings.  Valid media types include call, callback, chat, email, and message.
	Utilization *map[string]Mediautilization `json:"utilization,omitempty"`

}

// String returns a JSON representation of the model
func (o *Utilization) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

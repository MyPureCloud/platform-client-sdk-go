package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Configurationoverrides
type Configurationoverrides struct { 
	// Priority - Indicates whether or not the contact will be placed in front of the queue or at the end of the queue.
	Priority *bool `json:"priority,omitempty"`

}

// String returns a JSON representation of the model
func (o *Configurationoverrides) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

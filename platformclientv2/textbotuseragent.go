package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotuseragent - Information about the caller executing a bot flow.
type Textbotuseragent struct { 
	// Name - The name of the user agent.
	Name *string `json:"name,omitempty"`

}

// String returns a JSON representation of the model
func (o *Textbotuseragent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

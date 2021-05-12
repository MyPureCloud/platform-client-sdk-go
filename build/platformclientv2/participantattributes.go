package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Participantattributes
type Participantattributes struct { 
	// Attributes - The map of attribute keys to values.
	Attributes *map[string]string `json:"attributes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Participantattributes) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

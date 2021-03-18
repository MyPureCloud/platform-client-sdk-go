package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Consulttransferupdate
type Consulttransferupdate struct { 
	// SpeakTo - Determines to whom the initiating participant is speaking.
	SpeakTo *string `json:"speakTo,omitempty"`

}

// String returns a JSON representation of the model
func (o *Consulttransferupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

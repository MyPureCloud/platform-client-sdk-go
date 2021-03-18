package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Nludetectioninput
type Nludetectioninput struct { 
	// Text - The text to perform NLU detection on.
	Text *string `json:"text,omitempty"`

}

// String returns a JSON representation of the model
func (o *Nludetectioninput) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

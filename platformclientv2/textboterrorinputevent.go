package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Textboterrorinputevent - Settings for an input event to the bot flow indicating an error has occurred.
type Textboterrorinputevent struct { 
	// Code - The error code.
	Code *string `json:"code,omitempty"`


	// Message - The error message.
	Message *string `json:"message,omitempty"`

}

// String returns a JSON representation of the model
func (o *Textboterrorinputevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

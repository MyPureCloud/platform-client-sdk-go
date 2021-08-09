package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotuserinputevent - Settings for an input event to the bot flow indicating user input is available.
type Textbotuserinputevent struct { 
	// Mode - The input mode.
	Mode *string `json:"mode,omitempty"`


	// Alternatives - The input alternatives.
	Alternatives *[]Textbotuserinputalternative `json:"alternatives,omitempty"`

}

// String returns a JSON representation of the model
func (o *Textbotuserinputevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

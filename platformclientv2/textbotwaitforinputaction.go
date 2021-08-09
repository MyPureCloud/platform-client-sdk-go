package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotwaitforinputaction - Settings for a next-action of waiting for additional user input and sending the data as an input action to the bot flow.
type Textbotwaitforinputaction struct { 
	// ModeConstraints - The mode constraints for the user input.
	ModeConstraints *Textbotmodeconstraints `json:"modeConstraints,omitempty"`

}

// String returns a JSON representation of the model
func (o *Textbotwaitforinputaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

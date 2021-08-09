package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotinputoutputdata - Input/Output data related to a bot flow which is exiting gracefully.
type Textbotinputoutputdata struct { 
	// Variables - The input/output variables using the format as appropriate for the variable data type in the flow definition.
	Variables *map[string]interface{} `json:"variables,omitempty"`

}

// String returns a JSON representation of the model
func (o *Textbotinputoutputdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotinputoutputdata - Input/Output data related to a bot flow which is exiting gracefully.
type Textbotinputoutputdata struct { 
	// Variables - The input/output variables using the format as appropriate for the variable data type in the flow definition.
	Variables *map[string]interface{} `json:"variables,omitempty"`

}

func (u *Textbotinputoutputdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbotinputoutputdata

	

	return json.Marshal(&struct { 
		Variables *map[string]interface{} `json:"variables,omitempty"`
		*Alias
	}{ 
		Variables: u.Variables,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Textbotinputoutputdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

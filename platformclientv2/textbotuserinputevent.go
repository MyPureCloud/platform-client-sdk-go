package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Textbotuserinputevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbotuserinputevent

	

	return json.Marshal(&struct { 
		Mode *string `json:"mode,omitempty"`
		
		Alternatives *[]Textbotuserinputalternative `json:"alternatives,omitempty"`
		*Alias
	}{ 
		Mode: u.Mode,
		
		Alternatives: u.Alternatives,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Textbotuserinputevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

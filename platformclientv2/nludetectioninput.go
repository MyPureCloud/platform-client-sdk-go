package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Nludetectioninput
type Nludetectioninput struct { 
	// Text - The text to perform NLU detection on.
	Text *string `json:"text,omitempty"`

}

func (u *Nludetectioninput) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nludetectioninput

	

	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		*Alias
	}{ 
		Text: u.Text,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Nludetectioninput) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

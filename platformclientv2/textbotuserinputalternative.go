package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotuserinputalternative - User input data used in a bot flow turn.
type Textbotuserinputalternative struct { 
	// Transcript - The user input transcript.
	Transcript *Textbottranscript `json:"transcript,omitempty"`

}

func (u *Textbotuserinputalternative) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbotuserinputalternative

	

	return json.Marshal(&struct { 
		Transcript *Textbottranscript `json:"transcript,omitempty"`
		*Alias
	}{ 
		Transcript: u.Transcript,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Textbotuserinputalternative) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Consulttransferupdate
type Consulttransferupdate struct { 
	// SpeakTo - Determines to whom the initiating participant is speaking.
	SpeakTo *string `json:"speakTo,omitempty"`

}

func (u *Consulttransferupdate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Consulttransferupdate

	

	return json.Marshal(&struct { 
		SpeakTo *string `json:"speakTo,omitempty"`
		*Alias
	}{ 
		SpeakTo: u.SpeakTo,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Consulttransferupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

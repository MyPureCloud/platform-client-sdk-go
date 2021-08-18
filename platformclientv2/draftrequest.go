package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Draftrequest
type Draftrequest struct { 
	// Intents - Draft intent object.
	Intents *[]Draftintents `json:"intents,omitempty"`

}

func (u *Draftrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Draftrequest

	

	return json.Marshal(&struct { 
		Intents *[]Draftintents `json:"intents,omitempty"`
		*Alias
	}{ 
		Intents: u.Intents,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Draftrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

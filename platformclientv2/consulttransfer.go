package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Consulttransfer
type Consulttransfer struct { 
	// SpeakTo - Determines to whom the initiating participant is speaking. Defaults to DESTINATION
	SpeakTo *string `json:"speakTo,omitempty"`


	// Destination - Destination phone number and name.
	Destination *Destination `json:"destination,omitempty"`

}

func (u *Consulttransfer) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Consulttransfer

	

	return json.Marshal(&struct { 
		SpeakTo *string `json:"speakTo,omitempty"`
		
		Destination *Destination `json:"destination,omitempty"`
		*Alias
	}{ 
		SpeakTo: u.SpeakTo,
		
		Destination: u.Destination,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Consulttransfer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

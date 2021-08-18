package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Maxparticipants
type Maxparticipants struct { 
	// MaxParticipants - The maximum number of participants that are allowed on a conversation.
	MaxParticipants *int `json:"maxParticipants,omitempty"`

}

func (u *Maxparticipants) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Maxparticipants

	

	return json.Marshal(&struct { 
		MaxParticipants *int `json:"maxParticipants,omitempty"`
		*Alias
	}{ 
		MaxParticipants: u.MaxParticipants,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Maxparticipants) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

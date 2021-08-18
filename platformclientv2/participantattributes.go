package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Participantattributes
type Participantattributes struct { 
	// Attributes - The map of attribute keys to values.
	Attributes *map[string]string `json:"attributes,omitempty"`

}

func (u *Participantattributes) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Participantattributes

	

	return json.Marshal(&struct { 
		Attributes *map[string]string `json:"attributes,omitempty"`
		*Alias
	}{ 
		Attributes: u.Attributes,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Participantattributes) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

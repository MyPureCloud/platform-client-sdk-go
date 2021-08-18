package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Configurationoverrides
type Configurationoverrides struct { 
	// Priority - Indicates whether or not the contact will be placed in front of the queue or at the end of the queue.
	Priority *bool `json:"priority,omitempty"`

}

func (u *Configurationoverrides) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Configurationoverrides

	

	return json.Marshal(&struct { 
		Priority *bool `json:"priority,omitempty"`
		*Alias
	}{ 
		Priority: u.Priority,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Configurationoverrides) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

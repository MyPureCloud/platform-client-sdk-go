package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Utilization
type Utilization struct { 
	// Utilization - Map of media type to utilization settings.  Valid media types include call, callback, chat, email, and message.
	Utilization *map[string]Mediautilization `json:"utilization,omitempty"`

}

func (u *Utilization) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Utilization

	

	return json.Marshal(&struct { 
		Utilization *map[string]Mediautilization `json:"utilization,omitempty"`
		*Alias
	}{ 
		Utilization: u.Utilization,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Utilization) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

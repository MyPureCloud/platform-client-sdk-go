package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentmaxutilization
type Agentmaxutilization struct { 
	// Utilization - Map of media type to utilization settings.  Valid media types include call, callback, chat, email, and message.
	Utilization *map[string]Mediautilization `json:"utilization,omitempty"`


	// Level
	Level *string `json:"level,omitempty"`

}

func (u *Agentmaxutilization) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Agentmaxutilization

	

	return json.Marshal(&struct { 
		Utilization *map[string]Mediautilization `json:"utilization,omitempty"`
		
		Level *string `json:"level,omitempty"`
		*Alias
	}{ 
		Utilization: u.Utilization,
		
		Level: u.Level,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Agentmaxutilization) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

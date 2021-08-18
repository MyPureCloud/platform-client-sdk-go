package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Agent
type Agent struct { 
	// Stage - The current stage for this agent
	Stage *string `json:"stage,omitempty"`

}

func (u *Agent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Agent

	

	return json.Marshal(&struct { 
		Stage *string `json:"stage,omitempty"`
		*Alias
	}{ 
		Stage: u.Stage,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Agent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Murescheduleresultwrapper
type Murescheduleresultwrapper struct { 
	// AgentSchedules - The list of agent schedules
	AgentSchedules *[]Buagentschedulerescheduleresponse `json:"agentSchedules,omitempty"`

}

func (u *Murescheduleresultwrapper) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Murescheduleresultwrapper

	

	return json.Marshal(&struct { 
		AgentSchedules *[]Buagentschedulerescheduleresponse `json:"agentSchedules,omitempty"`
		*Alias
	}{ 
		AgentSchedules: u.AgentSchedules,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Murescheduleresultwrapper) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

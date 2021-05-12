package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Murescheduleresultwrapper
type Murescheduleresultwrapper struct { 
	// AgentSchedules - The list of agent schedules
	AgentSchedules *[]Buagentschedulerescheduleresponse `json:"agentSchedules,omitempty"`

}

// String returns a JSON representation of the model
func (o *Murescheduleresultwrapper) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"encoding/json"
)

// Murescheduleresultwrapper
type Murescheduleresultwrapper struct { 
	// AgentSchedules - The list of agent schedules
	AgentSchedules *[]Buagentschedulerescheduleresponse `json:"agentSchedules,omitempty"`

}

// String returns a JSON representation of the model
func (o *Murescheduleresultwrapper) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

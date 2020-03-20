package platformclientv2
import (
	"encoding/json"
)

// Unscheduledagentwarning
type Unscheduledagentwarning struct { 
	// Agent - The agent for which this warning applies
	Agent *Userreference `json:"agent,omitempty"`


	// UnscheduledReason - The reason this agent was not scheduled
	UnscheduledReason *string `json:"unscheduledReason,omitempty"`

}

// String returns a JSON representation of the model
func (o *Unscheduledagentwarning) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

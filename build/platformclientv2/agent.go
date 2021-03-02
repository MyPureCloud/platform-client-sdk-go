package platformclientv2
import (
	"encoding/json"
)

// Agent
type Agent struct { 
	// Stage - The current stage for this agent
	Stage *string `json:"stage,omitempty"`

}

// String returns a JSON representation of the model
func (o *Agent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

package platformclientv2
import (
	"encoding/json"
)

// Scoredagent
type Scoredagent struct { 
	// Agent - The agent
	Agent *Addressableentityref `json:"agent,omitempty"`


	// Score - Agent's score for the current conversation, from 0 - 100, higher being better
	Score *int `json:"score,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scoredagent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

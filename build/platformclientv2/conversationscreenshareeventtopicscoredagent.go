package platformclientv2
import (
	"encoding/json"
)

// Conversationscreenshareeventtopicscoredagent
type Conversationscreenshareeventtopicscoredagent struct { 
	// Agent
	Agent *Conversationscreenshareeventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int32 `json:"score,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationscreenshareeventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

package platformclientv2
import (
	"encoding/json"
)

// Conversationchateventtopicscoredagent
type Conversationchateventtopicscoredagent struct { 
	// Agent
	Agent *Conversationchateventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int32 `json:"score,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationchateventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

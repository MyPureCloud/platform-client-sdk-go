package platformclientv2
import (
	"encoding/json"
)

// Conversationemaileventtopicscoredagent
type Conversationemaileventtopicscoredagent struct { 
	// Agent
	Agent *Conversationemaileventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int32 `json:"score,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationemaileventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

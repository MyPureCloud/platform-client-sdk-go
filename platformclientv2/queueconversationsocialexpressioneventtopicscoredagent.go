package platformclientv2
import (
	"encoding/json"
)

// Queueconversationsocialexpressioneventtopicscoredagent
type Queueconversationsocialexpressioneventtopicscoredagent struct { 
	// Agent
	Agent *Queueconversationsocialexpressioneventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int `json:"score,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

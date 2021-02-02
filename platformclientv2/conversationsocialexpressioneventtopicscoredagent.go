package platformclientv2
import (
	"encoding/json"
)

// Conversationsocialexpressioneventtopicscoredagent
type Conversationsocialexpressioneventtopicscoredagent struct { 
	// Agent
	Agent *Conversationsocialexpressioneventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int `json:"score,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationsocialexpressioneventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

package platformclientv2
import (
	"encoding/json"
)

// Conversationvideoeventtopicscoredagent
type Conversationvideoeventtopicscoredagent struct { 
	// Agent
	Agent *Conversationvideoeventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int `json:"score,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationvideoeventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

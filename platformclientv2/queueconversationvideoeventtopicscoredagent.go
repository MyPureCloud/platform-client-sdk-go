package platformclientv2
import (
	"encoding/json"
)

// Queueconversationvideoeventtopicscoredagent
type Queueconversationvideoeventtopicscoredagent struct { 
	// Agent
	Agent *Queueconversationvideoeventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int `json:"score,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

package platformclientv2
import (
	"encoding/json"
)

// Queueconversationeventtopicscoredagent
type Queueconversationeventtopicscoredagent struct { 
	// Agent
	Agent *Queueconversationeventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int `json:"score,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationeventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

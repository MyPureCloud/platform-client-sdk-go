package platformclientv2
import (
	"encoding/json"
)

// Conversationeventtopicscoredagent
type Conversationeventtopicscoredagent struct { 
	// Agent
	Agent *Conversationeventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int `json:"score,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationeventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

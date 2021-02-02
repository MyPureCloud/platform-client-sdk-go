package platformclientv2
import (
	"encoding/json"
)

// Conversationcalleventtopicscoredagent
type Conversationcalleventtopicscoredagent struct { 
	// Agent
	Agent *Conversationcalleventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int `json:"score,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationcalleventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

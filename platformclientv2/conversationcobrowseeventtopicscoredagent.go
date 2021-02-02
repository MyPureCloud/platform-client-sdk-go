package platformclientv2
import (
	"encoding/json"
)

// Conversationcobrowseeventtopicscoredagent
type Conversationcobrowseeventtopicscoredagent struct { 
	// Agent
	Agent *Conversationcobrowseeventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int `json:"score,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationcobrowseeventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

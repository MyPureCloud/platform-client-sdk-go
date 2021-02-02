package platformclientv2
import (
	"encoding/json"
)

// Queueconversationcobrowseeventtopicscoredagent
type Queueconversationcobrowseeventtopicscoredagent struct { 
	// Agent
	Agent *Queueconversationcobrowseeventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int `json:"score,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationcobrowseeventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

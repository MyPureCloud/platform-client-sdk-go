package platformclientv2
import (
	"encoding/json"
)

// Queueconversationcalleventtopicscoredagent
type Queueconversationcalleventtopicscoredagent struct { 
	// Agent
	Agent *Queueconversationcalleventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int32 `json:"score,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationcalleventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

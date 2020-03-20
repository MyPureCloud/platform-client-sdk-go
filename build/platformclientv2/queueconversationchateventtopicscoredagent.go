package platformclientv2
import (
	"encoding/json"
)

// Queueconversationchateventtopicscoredagent
type Queueconversationchateventtopicscoredagent struct { 
	// Agent
	Agent *Queueconversationchateventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int32 `json:"score,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationchateventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

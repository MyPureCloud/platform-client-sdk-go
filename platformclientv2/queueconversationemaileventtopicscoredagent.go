package platformclientv2
import (
	"encoding/json"
)

// Queueconversationemaileventtopicscoredagent
type Queueconversationemaileventtopicscoredagent struct { 
	// Agent
	Agent *Queueconversationemaileventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int32 `json:"score,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationemaileventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

package platformclientv2
import (
	"encoding/json"
)

// Queueconversationmessageeventtopicscoredagent
type Queueconversationmessageeventtopicscoredagent struct { 
	// Agent
	Agent *Queueconversationmessageeventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int `json:"score,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationmessageeventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

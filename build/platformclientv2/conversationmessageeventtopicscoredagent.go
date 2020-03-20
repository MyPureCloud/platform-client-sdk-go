package platformclientv2
import (
	"encoding/json"
)

// Conversationmessageeventtopicscoredagent
type Conversationmessageeventtopicscoredagent struct { 
	// Agent
	Agent *Conversationmessageeventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int32 `json:"score,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationmessageeventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

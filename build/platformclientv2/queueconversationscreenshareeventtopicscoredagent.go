package platformclientv2
import (
	"encoding/json"
)

// Queueconversationscreenshareeventtopicscoredagent
type Queueconversationscreenshareeventtopicscoredagent struct { 
	// Agent
	Agent *Queueconversationscreenshareeventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int32 `json:"score,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationscreenshareeventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

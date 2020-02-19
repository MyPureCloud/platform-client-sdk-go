package platformclientv2
import (
	"encoding/json"
)

// Queueconversationcallbackeventtopicscoredagent
type Queueconversationcallbackeventtopicscoredagent struct { 
	// Agent
	Agent *Queueconversationcallbackeventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int32 `json:"score,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationcallbackeventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

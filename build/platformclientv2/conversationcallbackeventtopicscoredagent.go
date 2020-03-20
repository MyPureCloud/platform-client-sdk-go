package platformclientv2
import (
	"encoding/json"
)

// Conversationcallbackeventtopicscoredagent
type Conversationcallbackeventtopicscoredagent struct { 
	// Agent
	Agent *Conversationcallbackeventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int32 `json:"score,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationcallbackeventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

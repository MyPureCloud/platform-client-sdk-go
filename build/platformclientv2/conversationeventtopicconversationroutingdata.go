package platformclientv2
import (
	"encoding/json"
)

// Conversationeventtopicconversationroutingdata
type Conversationeventtopicconversationroutingdata struct { 
	// Queue
	Queue *Conversationeventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Conversationeventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int32 `json:"priority,omitempty"`


	// Skills
	Skills *[]Conversationeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Conversationeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

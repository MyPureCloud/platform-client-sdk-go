package platformclientv2
import (
	"encoding/json"
)

// Conversationcalleventtopicconversationroutingdata
type Conversationcalleventtopicconversationroutingdata struct { 
	// Queue
	Queue *Conversationcalleventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Conversationcalleventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Conversationcalleventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Conversationcalleventtopicscoredagent `json:"scoredAgents,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationcalleventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

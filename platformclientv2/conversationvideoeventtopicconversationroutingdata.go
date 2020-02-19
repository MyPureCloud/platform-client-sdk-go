package platformclientv2
import (
	"encoding/json"
)

// Conversationvideoeventtopicconversationroutingdata
type Conversationvideoeventtopicconversationroutingdata struct { 
	// Queue
	Queue *Conversationvideoeventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Conversationvideoeventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int32 `json:"priority,omitempty"`


	// Skills
	Skills *[]Conversationvideoeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Conversationvideoeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationvideoeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

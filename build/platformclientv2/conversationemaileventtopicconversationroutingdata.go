package platformclientv2
import (
	"encoding/json"
)

// Conversationemaileventtopicconversationroutingdata
type Conversationemaileventtopicconversationroutingdata struct { 
	// Queue
	Queue *Conversationemaileventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Conversationemaileventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Conversationemaileventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Conversationemaileventtopicscoredagent `json:"scoredAgents,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationemaileventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

package platformclientv2
import (
	"encoding/json"
)

// Queueconversationemaileventtopicconversationroutingdata
type Queueconversationemaileventtopicconversationroutingdata struct { 
	// Queue
	Queue *Queueconversationemaileventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Queueconversationemaileventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Queueconversationemaileventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Queueconversationemaileventtopicscoredagent `json:"scoredAgents,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationemaileventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

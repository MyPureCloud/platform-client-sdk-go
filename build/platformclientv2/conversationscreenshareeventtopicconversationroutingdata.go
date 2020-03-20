package platformclientv2
import (
	"encoding/json"
)

// Conversationscreenshareeventtopicconversationroutingdata
type Conversationscreenshareeventtopicconversationroutingdata struct { 
	// Queue
	Queue *Conversationscreenshareeventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Conversationscreenshareeventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int32 `json:"priority,omitempty"`


	// Skills
	Skills *[]Conversationscreenshareeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Conversationscreenshareeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationscreenshareeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

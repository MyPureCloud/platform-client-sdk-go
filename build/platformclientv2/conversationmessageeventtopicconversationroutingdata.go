package platformclientv2
import (
	"encoding/json"
)

// Conversationmessageeventtopicconversationroutingdata
type Conversationmessageeventtopicconversationroutingdata struct { 
	// Queue
	Queue *Conversationmessageeventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Conversationmessageeventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Conversationmessageeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Conversationmessageeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationmessageeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

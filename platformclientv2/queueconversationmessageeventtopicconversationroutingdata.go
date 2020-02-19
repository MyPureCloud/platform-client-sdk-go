package platformclientv2
import (
	"encoding/json"
)

// Queueconversationmessageeventtopicconversationroutingdata
type Queueconversationmessageeventtopicconversationroutingdata struct { 
	// Queue
	Queue *Queueconversationmessageeventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Queueconversationmessageeventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int32 `json:"priority,omitempty"`


	// Skills
	Skills *[]Queueconversationmessageeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Queueconversationmessageeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationmessageeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

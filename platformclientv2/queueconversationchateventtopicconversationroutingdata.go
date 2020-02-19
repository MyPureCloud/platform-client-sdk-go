package platformclientv2
import (
	"encoding/json"
)

// Queueconversationchateventtopicconversationroutingdata
type Queueconversationchateventtopicconversationroutingdata struct { 
	// Queue
	Queue *Queueconversationchateventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Queueconversationchateventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int32 `json:"priority,omitempty"`


	// Skills
	Skills *[]Queueconversationchateventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Queueconversationchateventtopicscoredagent `json:"scoredAgents,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationchateventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

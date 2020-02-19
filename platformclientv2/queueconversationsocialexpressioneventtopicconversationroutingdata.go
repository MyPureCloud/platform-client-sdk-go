package platformclientv2
import (
	"encoding/json"
)

// Queueconversationsocialexpressioneventtopicconversationroutingdata
type Queueconversationsocialexpressioneventtopicconversationroutingdata struct { 
	// Queue
	Queue *Queueconversationsocialexpressioneventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Queueconversationsocialexpressioneventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int32 `json:"priority,omitempty"`


	// Skills
	Skills *[]Queueconversationsocialexpressioneventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Queueconversationsocialexpressioneventtopicscoredagent `json:"scoredAgents,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

package platformclientv2
import (
	"encoding/json"
)

// Conversationcallbackeventtopicconversationroutingdata
type Conversationcallbackeventtopicconversationroutingdata struct { 
	// Queue
	Queue *Conversationcallbackeventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Conversationcallbackeventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int32 `json:"priority,omitempty"`


	// Skills
	Skills *[]Conversationcallbackeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Conversationcallbackeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationcallbackeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

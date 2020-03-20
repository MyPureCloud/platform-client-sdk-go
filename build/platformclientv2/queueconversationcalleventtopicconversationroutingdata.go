package platformclientv2
import (
	"encoding/json"
)

// Queueconversationcalleventtopicconversationroutingdata
type Queueconversationcalleventtopicconversationroutingdata struct { 
	// Queue
	Queue *Queueconversationcalleventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Queueconversationcalleventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int32 `json:"priority,omitempty"`


	// Skills
	Skills *[]Queueconversationcalleventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Queueconversationcalleventtopicscoredagent `json:"scoredAgents,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationcalleventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

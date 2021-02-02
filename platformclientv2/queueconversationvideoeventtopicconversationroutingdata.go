package platformclientv2
import (
	"encoding/json"
)

// Queueconversationvideoeventtopicconversationroutingdata
type Queueconversationvideoeventtopicconversationroutingdata struct { 
	// Queue
	Queue *Queueconversationvideoeventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Queueconversationvideoeventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Queueconversationvideoeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Queueconversationvideoeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

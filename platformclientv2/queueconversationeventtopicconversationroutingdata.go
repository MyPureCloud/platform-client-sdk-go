package platformclientv2
import (
	"encoding/json"
)

// Queueconversationeventtopicconversationroutingdata
type Queueconversationeventtopicconversationroutingdata struct { 
	// Queue
	Queue *Queueconversationeventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Queueconversationeventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Queueconversationeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Queueconversationeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

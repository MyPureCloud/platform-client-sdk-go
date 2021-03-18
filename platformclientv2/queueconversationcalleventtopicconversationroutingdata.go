package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcalleventtopicconversationroutingdata
type Queueconversationcalleventtopicconversationroutingdata struct { 
	// Queue
	Queue *Queueconversationcalleventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Queueconversationcalleventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Queueconversationcalleventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Queueconversationcalleventtopicscoredagent `json:"scoredAgents,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationcalleventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

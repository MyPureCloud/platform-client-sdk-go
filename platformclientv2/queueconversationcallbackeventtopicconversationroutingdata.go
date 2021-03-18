package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcallbackeventtopicconversationroutingdata
type Queueconversationcallbackeventtopicconversationroutingdata struct { 
	// Queue
	Queue *Queueconversationcallbackeventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Queueconversationcallbackeventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Queueconversationcallbackeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Queueconversationcallbackeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationcallbackeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

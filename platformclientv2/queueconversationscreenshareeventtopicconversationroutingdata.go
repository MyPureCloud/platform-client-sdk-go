package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationscreenshareeventtopicconversationroutingdata
type Queueconversationscreenshareeventtopicconversationroutingdata struct { 
	// Queue
	Queue *Queueconversationscreenshareeventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Queueconversationscreenshareeventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Queueconversationscreenshareeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Queueconversationscreenshareeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationscreenshareeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

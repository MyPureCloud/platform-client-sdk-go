package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationeventtopicconversationroutingdata
type Conversationeventtopicconversationroutingdata struct { 
	// Queue
	Queue *Conversationeventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Conversationeventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Conversationeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Conversationeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

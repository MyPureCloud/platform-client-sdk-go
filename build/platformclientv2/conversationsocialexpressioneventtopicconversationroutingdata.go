package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationsocialexpressioneventtopicconversationroutingdata
type Conversationsocialexpressioneventtopicconversationroutingdata struct { 
	// Queue
	Queue *Conversationsocialexpressioneventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Conversationsocialexpressioneventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Conversationsocialexpressioneventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Conversationsocialexpressioneventtopicscoredagent `json:"scoredAgents,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationsocialexpressioneventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

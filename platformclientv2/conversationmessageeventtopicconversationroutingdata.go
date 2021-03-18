package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmessageeventtopicconversationroutingdata
type Conversationmessageeventtopicconversationroutingdata struct { 
	// Queue
	Queue *Conversationmessageeventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Conversationmessageeventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Conversationmessageeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Conversationmessageeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationmessageeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

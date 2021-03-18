package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcobrowseeventtopicconversationroutingdata
type Queueconversationcobrowseeventtopicconversationroutingdata struct { 
	// Queue
	Queue *Queueconversationcobrowseeventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Queueconversationcobrowseeventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Queueconversationcobrowseeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Queueconversationcobrowseeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationcobrowseeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

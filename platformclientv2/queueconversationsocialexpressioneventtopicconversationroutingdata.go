package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationsocialexpressioneventtopicconversationroutingdata
type Queueconversationsocialexpressioneventtopicconversationroutingdata struct { 
	// Queue
	Queue *Queueconversationsocialexpressioneventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Queueconversationsocialexpressioneventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Queueconversationsocialexpressioneventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Queueconversationsocialexpressioneventtopicscoredagent `json:"scoredAgents,omitempty"`

}

func (u *Queueconversationsocialexpressioneventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationsocialexpressioneventtopicconversationroutingdata

	

	return json.Marshal(&struct { 
		Queue *Queueconversationsocialexpressioneventtopicurireference `json:"queue,omitempty"`
		
		Language *Queueconversationsocialexpressioneventtopicurireference `json:"language,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Queueconversationsocialexpressioneventtopicurireference `json:"skills,omitempty"`
		
		ScoredAgents *[]Queueconversationsocialexpressioneventtopicscoredagent `json:"scoredAgents,omitempty"`
		*Alias
	}{ 
		Queue: u.Queue,
		
		Language: u.Language,
		
		Priority: u.Priority,
		
		Skills: u.Skills,
		
		ScoredAgents: u.ScoredAgents,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

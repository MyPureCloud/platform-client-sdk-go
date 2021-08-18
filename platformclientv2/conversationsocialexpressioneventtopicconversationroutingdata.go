package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Conversationsocialexpressioneventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationsocialexpressioneventtopicconversationroutingdata

	

	return json.Marshal(&struct { 
		Queue *Conversationsocialexpressioneventtopicurireference `json:"queue,omitempty"`
		
		Language *Conversationsocialexpressioneventtopicurireference `json:"language,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Conversationsocialexpressioneventtopicurireference `json:"skills,omitempty"`
		
		ScoredAgents *[]Conversationsocialexpressioneventtopicscoredagent `json:"scoredAgents,omitempty"`
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
func (o *Conversationsocialexpressioneventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

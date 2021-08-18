package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationchateventtopicconversationroutingdata
type Conversationchateventtopicconversationroutingdata struct { 
	// Queue
	Queue *Conversationchateventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Conversationchateventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Conversationchateventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Conversationchateventtopicscoredagent `json:"scoredAgents,omitempty"`

}

func (u *Conversationchateventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationchateventtopicconversationroutingdata

	

	return json.Marshal(&struct { 
		Queue *Conversationchateventtopicurireference `json:"queue,omitempty"`
		
		Language *Conversationchateventtopicurireference `json:"language,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Conversationchateventtopicurireference `json:"skills,omitempty"`
		
		ScoredAgents *[]Conversationchateventtopicscoredagent `json:"scoredAgents,omitempty"`
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
func (o *Conversationchateventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

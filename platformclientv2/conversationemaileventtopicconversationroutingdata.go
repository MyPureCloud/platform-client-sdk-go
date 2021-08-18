package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationemaileventtopicconversationroutingdata
type Conversationemaileventtopicconversationroutingdata struct { 
	// Queue
	Queue *Conversationemaileventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Conversationemaileventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Conversationemaileventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Conversationemaileventtopicscoredagent `json:"scoredAgents,omitempty"`

}

func (u *Conversationemaileventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationemaileventtopicconversationroutingdata

	

	return json.Marshal(&struct { 
		Queue *Conversationemaileventtopicurireference `json:"queue,omitempty"`
		
		Language *Conversationemaileventtopicurireference `json:"language,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Conversationemaileventtopicurireference `json:"skills,omitempty"`
		
		ScoredAgents *[]Conversationemaileventtopicscoredagent `json:"scoredAgents,omitempty"`
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
func (o *Conversationemaileventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

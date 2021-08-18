package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationscreenshareeventtopicconversationroutingdata
type Conversationscreenshareeventtopicconversationroutingdata struct { 
	// Queue
	Queue *Conversationscreenshareeventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Conversationscreenshareeventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Conversationscreenshareeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Conversationscreenshareeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

func (u *Conversationscreenshareeventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationscreenshareeventtopicconversationroutingdata

	

	return json.Marshal(&struct { 
		Queue *Conversationscreenshareeventtopicurireference `json:"queue,omitempty"`
		
		Language *Conversationscreenshareeventtopicurireference `json:"language,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Conversationscreenshareeventtopicurireference `json:"skills,omitempty"`
		
		ScoredAgents *[]Conversationscreenshareeventtopicscoredagent `json:"scoredAgents,omitempty"`
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
func (o *Conversationscreenshareeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

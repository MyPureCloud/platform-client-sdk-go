package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationvideoeventtopicconversationroutingdata
type Conversationvideoeventtopicconversationroutingdata struct { 
	// Queue
	Queue *Conversationvideoeventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Conversationvideoeventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Conversationvideoeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Conversationvideoeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

func (u *Conversationvideoeventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationvideoeventtopicconversationroutingdata

	

	return json.Marshal(&struct { 
		Queue *Conversationvideoeventtopicurireference `json:"queue,omitempty"`
		
		Language *Conversationvideoeventtopicurireference `json:"language,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Conversationvideoeventtopicurireference `json:"skills,omitempty"`
		
		ScoredAgents *[]Conversationvideoeventtopicscoredagent `json:"scoredAgents,omitempty"`
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
func (o *Conversationvideoeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

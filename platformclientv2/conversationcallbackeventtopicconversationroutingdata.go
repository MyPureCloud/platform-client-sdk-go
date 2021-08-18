package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcallbackeventtopicconversationroutingdata
type Conversationcallbackeventtopicconversationroutingdata struct { 
	// Queue
	Queue *Conversationcallbackeventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Conversationcallbackeventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Conversationcallbackeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Conversationcallbackeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

func (u *Conversationcallbackeventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcallbackeventtopicconversationroutingdata

	

	return json.Marshal(&struct { 
		Queue *Conversationcallbackeventtopicurireference `json:"queue,omitempty"`
		
		Language *Conversationcallbackeventtopicurireference `json:"language,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Conversationcallbackeventtopicurireference `json:"skills,omitempty"`
		
		ScoredAgents *[]Conversationcallbackeventtopicscoredagent `json:"scoredAgents,omitempty"`
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
func (o *Conversationcallbackeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

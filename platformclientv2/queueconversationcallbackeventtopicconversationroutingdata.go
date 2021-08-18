package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcallbackeventtopicconversationroutingdata
type Queueconversationcallbackeventtopicconversationroutingdata struct { 
	// Queue
	Queue *Queueconversationcallbackeventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Queueconversationcallbackeventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Queueconversationcallbackeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Queueconversationcallbackeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

func (u *Queueconversationcallbackeventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationcallbackeventtopicconversationroutingdata

	

	return json.Marshal(&struct { 
		Queue *Queueconversationcallbackeventtopicurireference `json:"queue,omitempty"`
		
		Language *Queueconversationcallbackeventtopicurireference `json:"language,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Queueconversationcallbackeventtopicurireference `json:"skills,omitempty"`
		
		ScoredAgents *[]Queueconversationcallbackeventtopicscoredagent `json:"scoredAgents,omitempty"`
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
func (o *Queueconversationcallbackeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

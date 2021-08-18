package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationemaileventtopicconversationroutingdata
type Queueconversationemaileventtopicconversationroutingdata struct { 
	// Queue
	Queue *Queueconversationemaileventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Queueconversationemaileventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Queueconversationemaileventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Queueconversationemaileventtopicscoredagent `json:"scoredAgents,omitempty"`

}

func (u *Queueconversationemaileventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationemaileventtopicconversationroutingdata

	

	return json.Marshal(&struct { 
		Queue *Queueconversationemaileventtopicurireference `json:"queue,omitempty"`
		
		Language *Queueconversationemaileventtopicurireference `json:"language,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Queueconversationemaileventtopicurireference `json:"skills,omitempty"`
		
		ScoredAgents *[]Queueconversationemaileventtopicscoredagent `json:"scoredAgents,omitempty"`
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
func (o *Queueconversationemaileventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

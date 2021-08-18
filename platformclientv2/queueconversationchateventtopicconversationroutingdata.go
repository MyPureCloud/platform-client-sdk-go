package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationchateventtopicconversationroutingdata
type Queueconversationchateventtopicconversationroutingdata struct { 
	// Queue
	Queue *Queueconversationchateventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Queueconversationchateventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Queueconversationchateventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Queueconversationchateventtopicscoredagent `json:"scoredAgents,omitempty"`

}

func (u *Queueconversationchateventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationchateventtopicconversationroutingdata

	

	return json.Marshal(&struct { 
		Queue *Queueconversationchateventtopicurireference `json:"queue,omitempty"`
		
		Language *Queueconversationchateventtopicurireference `json:"language,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Queueconversationchateventtopicurireference `json:"skills,omitempty"`
		
		ScoredAgents *[]Queueconversationchateventtopicscoredagent `json:"scoredAgents,omitempty"`
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
func (o *Queueconversationchateventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

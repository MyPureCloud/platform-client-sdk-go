package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationvideoeventtopicconversationroutingdata
type Queueconversationvideoeventtopicconversationroutingdata struct { 
	// Queue
	Queue *Queueconversationvideoeventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Queueconversationvideoeventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Queueconversationvideoeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Queueconversationvideoeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

func (u *Queueconversationvideoeventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationvideoeventtopicconversationroutingdata

	

	return json.Marshal(&struct { 
		Queue *Queueconversationvideoeventtopicurireference `json:"queue,omitempty"`
		
		Language *Queueconversationvideoeventtopicurireference `json:"language,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Queueconversationvideoeventtopicurireference `json:"skills,omitempty"`
		
		ScoredAgents *[]Queueconversationvideoeventtopicscoredagent `json:"scoredAgents,omitempty"`
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
func (o *Queueconversationvideoeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

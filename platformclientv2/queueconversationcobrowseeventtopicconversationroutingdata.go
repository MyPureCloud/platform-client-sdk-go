package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcobrowseeventtopicconversationroutingdata
type Queueconversationcobrowseeventtopicconversationroutingdata struct { 
	// Queue
	Queue *Queueconversationcobrowseeventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Queueconversationcobrowseeventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Queueconversationcobrowseeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Queueconversationcobrowseeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

func (u *Queueconversationcobrowseeventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationcobrowseeventtopicconversationroutingdata

	

	return json.Marshal(&struct { 
		Queue *Queueconversationcobrowseeventtopicurireference `json:"queue,omitempty"`
		
		Language *Queueconversationcobrowseeventtopicurireference `json:"language,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Queueconversationcobrowseeventtopicurireference `json:"skills,omitempty"`
		
		ScoredAgents *[]Queueconversationcobrowseeventtopicscoredagent `json:"scoredAgents,omitempty"`
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
func (o *Queueconversationcobrowseeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

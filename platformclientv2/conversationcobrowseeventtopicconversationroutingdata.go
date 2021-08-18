package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcobrowseeventtopicconversationroutingdata
type Conversationcobrowseeventtopicconversationroutingdata struct { 
	// Queue
	Queue *Conversationcobrowseeventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Conversationcobrowseeventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Conversationcobrowseeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Conversationcobrowseeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

func (u *Conversationcobrowseeventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcobrowseeventtopicconversationroutingdata

	

	return json.Marshal(&struct { 
		Queue *Conversationcobrowseeventtopicurireference `json:"queue,omitempty"`
		
		Language *Conversationcobrowseeventtopicurireference `json:"language,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Conversationcobrowseeventtopicurireference `json:"skills,omitempty"`
		
		ScoredAgents *[]Conversationcobrowseeventtopicscoredagent `json:"scoredAgents,omitempty"`
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
func (o *Conversationcobrowseeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

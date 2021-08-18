package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcalleventtopicconversationroutingdata
type Conversationcalleventtopicconversationroutingdata struct { 
	// Queue
	Queue *Conversationcalleventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Conversationcalleventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Conversationcalleventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Conversationcalleventtopicscoredagent `json:"scoredAgents,omitempty"`

}

func (u *Conversationcalleventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcalleventtopicconversationroutingdata

	

	return json.Marshal(&struct { 
		Queue *Conversationcalleventtopicurireference `json:"queue,omitempty"`
		
		Language *Conversationcalleventtopicurireference `json:"language,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Conversationcalleventtopicurireference `json:"skills,omitempty"`
		
		ScoredAgents *[]Conversationcalleventtopicscoredagent `json:"scoredAgents,omitempty"`
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
func (o *Conversationcalleventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

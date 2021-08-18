package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcalleventtopicconversationroutingdata
type Queueconversationcalleventtopicconversationroutingdata struct { 
	// Queue
	Queue *Queueconversationcalleventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Queueconversationcalleventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Queueconversationcalleventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Queueconversationcalleventtopicscoredagent `json:"scoredAgents,omitempty"`

}

func (u *Queueconversationcalleventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationcalleventtopicconversationroutingdata

	

	return json.Marshal(&struct { 
		Queue *Queueconversationcalleventtopicurireference `json:"queue,omitempty"`
		
		Language *Queueconversationcalleventtopicurireference `json:"language,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Queueconversationcalleventtopicurireference `json:"skills,omitempty"`
		
		ScoredAgents *[]Queueconversationcalleventtopicscoredagent `json:"scoredAgents,omitempty"`
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
func (o *Queueconversationcalleventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

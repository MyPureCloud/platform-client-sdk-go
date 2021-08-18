package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationmessageeventtopicconversationroutingdata
type Queueconversationmessageeventtopicconversationroutingdata struct { 
	// Queue
	Queue *Queueconversationmessageeventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Queueconversationmessageeventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Queueconversationmessageeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Queueconversationmessageeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

func (u *Queueconversationmessageeventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationmessageeventtopicconversationroutingdata

	

	return json.Marshal(&struct { 
		Queue *Queueconversationmessageeventtopicurireference `json:"queue,omitempty"`
		
		Language *Queueconversationmessageeventtopicurireference `json:"language,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Queueconversationmessageeventtopicurireference `json:"skills,omitempty"`
		
		ScoredAgents *[]Queueconversationmessageeventtopicscoredagent `json:"scoredAgents,omitempty"`
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
func (o *Queueconversationmessageeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

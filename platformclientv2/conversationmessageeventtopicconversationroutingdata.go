package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmessageeventtopicconversationroutingdata
type Conversationmessageeventtopicconversationroutingdata struct { 
	// Queue
	Queue *Conversationmessageeventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Conversationmessageeventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Conversationmessageeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Conversationmessageeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

func (u *Conversationmessageeventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationmessageeventtopicconversationroutingdata

	

	return json.Marshal(&struct { 
		Queue *Conversationmessageeventtopicurireference `json:"queue,omitempty"`
		
		Language *Conversationmessageeventtopicurireference `json:"language,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Conversationmessageeventtopicurireference `json:"skills,omitempty"`
		
		ScoredAgents *[]Conversationmessageeventtopicscoredagent `json:"scoredAgents,omitempty"`
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
func (o *Conversationmessageeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

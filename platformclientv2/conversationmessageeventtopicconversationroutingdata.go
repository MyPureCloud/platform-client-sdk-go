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

func (o *Conversationmessageeventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
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
		Queue: o.Queue,
		
		Language: o.Language,
		
		Priority: o.Priority,
		
		Skills: o.Skills,
		
		ScoredAgents: o.ScoredAgents,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationmessageeventtopicconversationroutingdata) UnmarshalJSON(b []byte) error {
	var ConversationmessageeventtopicconversationroutingdataMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationmessageeventtopicconversationroutingdataMap)
	if err != nil {
		return err
	}
	
	if Queue, ok := ConversationmessageeventtopicconversationroutingdataMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Language, ok := ConversationmessageeventtopicconversationroutingdataMap["language"].(map[string]interface{}); ok {
		LanguageString, _ := json.Marshal(Language)
		json.Unmarshal(LanguageString, &o.Language)
	}
	
	if Priority, ok := ConversationmessageeventtopicconversationroutingdataMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if Skills, ok := ConversationmessageeventtopicconversationroutingdataMap["skills"].([]interface{}); ok {
		SkillsString, _ := json.Marshal(Skills)
		json.Unmarshal(SkillsString, &o.Skills)
	}
	
	if ScoredAgents, ok := ConversationmessageeventtopicconversationroutingdataMap["scoredAgents"].([]interface{}); ok {
		ScoredAgentsString, _ := json.Marshal(ScoredAgents)
		json.Unmarshal(ScoredAgentsString, &o.ScoredAgents)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationmessageeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

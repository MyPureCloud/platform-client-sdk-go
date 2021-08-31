package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationsocialexpressioneventtopicconversationroutingdata
type Conversationsocialexpressioneventtopicconversationroutingdata struct { 
	// Queue
	Queue *Conversationsocialexpressioneventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Conversationsocialexpressioneventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Conversationsocialexpressioneventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Conversationsocialexpressioneventtopicscoredagent `json:"scoredAgents,omitempty"`

}

func (o *Conversationsocialexpressioneventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationsocialexpressioneventtopicconversationroutingdata
	
	return json.Marshal(&struct { 
		Queue *Conversationsocialexpressioneventtopicurireference `json:"queue,omitempty"`
		
		Language *Conversationsocialexpressioneventtopicurireference `json:"language,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Conversationsocialexpressioneventtopicurireference `json:"skills,omitempty"`
		
		ScoredAgents *[]Conversationsocialexpressioneventtopicscoredagent `json:"scoredAgents,omitempty"`
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

func (o *Conversationsocialexpressioneventtopicconversationroutingdata) UnmarshalJSON(b []byte) error {
	var ConversationsocialexpressioneventtopicconversationroutingdataMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationsocialexpressioneventtopicconversationroutingdataMap)
	if err != nil {
		return err
	}
	
	if Queue, ok := ConversationsocialexpressioneventtopicconversationroutingdataMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Language, ok := ConversationsocialexpressioneventtopicconversationroutingdataMap["language"].(map[string]interface{}); ok {
		LanguageString, _ := json.Marshal(Language)
		json.Unmarshal(LanguageString, &o.Language)
	}
	
	if Priority, ok := ConversationsocialexpressioneventtopicconversationroutingdataMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if Skills, ok := ConversationsocialexpressioneventtopicconversationroutingdataMap["skills"].([]interface{}); ok {
		SkillsString, _ := json.Marshal(Skills)
		json.Unmarshal(SkillsString, &o.Skills)
	}
	
	if ScoredAgents, ok := ConversationsocialexpressioneventtopicconversationroutingdataMap["scoredAgents"].([]interface{}); ok {
		ScoredAgentsString, _ := json.Marshal(ScoredAgents)
		json.Unmarshal(ScoredAgentsString, &o.ScoredAgents)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationsocialexpressioneventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

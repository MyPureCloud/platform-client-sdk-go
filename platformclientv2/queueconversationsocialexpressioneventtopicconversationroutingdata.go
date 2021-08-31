package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationsocialexpressioneventtopicconversationroutingdata
type Queueconversationsocialexpressioneventtopicconversationroutingdata struct { 
	// Queue
	Queue *Queueconversationsocialexpressioneventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Queueconversationsocialexpressioneventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Queueconversationsocialexpressioneventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Queueconversationsocialexpressioneventtopicscoredagent `json:"scoredAgents,omitempty"`

}

func (o *Queueconversationsocialexpressioneventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationsocialexpressioneventtopicconversationroutingdata
	
	return json.Marshal(&struct { 
		Queue *Queueconversationsocialexpressioneventtopicurireference `json:"queue,omitempty"`
		
		Language *Queueconversationsocialexpressioneventtopicurireference `json:"language,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Queueconversationsocialexpressioneventtopicurireference `json:"skills,omitempty"`
		
		ScoredAgents *[]Queueconversationsocialexpressioneventtopicscoredagent `json:"scoredAgents,omitempty"`
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

func (o *Queueconversationsocialexpressioneventtopicconversationroutingdata) UnmarshalJSON(b []byte) error {
	var QueueconversationsocialexpressioneventtopicconversationroutingdataMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationsocialexpressioneventtopicconversationroutingdataMap)
	if err != nil {
		return err
	}
	
	if Queue, ok := QueueconversationsocialexpressioneventtopicconversationroutingdataMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Language, ok := QueueconversationsocialexpressioneventtopicconversationroutingdataMap["language"].(map[string]interface{}); ok {
		LanguageString, _ := json.Marshal(Language)
		json.Unmarshal(LanguageString, &o.Language)
	}
	
	if Priority, ok := QueueconversationsocialexpressioneventtopicconversationroutingdataMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if Skills, ok := QueueconversationsocialexpressioneventtopicconversationroutingdataMap["skills"].([]interface{}); ok {
		SkillsString, _ := json.Marshal(Skills)
		json.Unmarshal(SkillsString, &o.Skills)
	}
	
	if ScoredAgents, ok := QueueconversationsocialexpressioneventtopicconversationroutingdataMap["scoredAgents"].([]interface{}); ok {
		ScoredAgentsString, _ := json.Marshal(ScoredAgents)
		json.Unmarshal(ScoredAgentsString, &o.ScoredAgents)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

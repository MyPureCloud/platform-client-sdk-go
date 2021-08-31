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

func (o *Conversationcobrowseeventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
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
		Queue: o.Queue,
		
		Language: o.Language,
		
		Priority: o.Priority,
		
		Skills: o.Skills,
		
		ScoredAgents: o.ScoredAgents,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationcobrowseeventtopicconversationroutingdata) UnmarshalJSON(b []byte) error {
	var ConversationcobrowseeventtopicconversationroutingdataMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcobrowseeventtopicconversationroutingdataMap)
	if err != nil {
		return err
	}
	
	if Queue, ok := ConversationcobrowseeventtopicconversationroutingdataMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Language, ok := ConversationcobrowseeventtopicconversationroutingdataMap["language"].(map[string]interface{}); ok {
		LanguageString, _ := json.Marshal(Language)
		json.Unmarshal(LanguageString, &o.Language)
	}
	
	if Priority, ok := ConversationcobrowseeventtopicconversationroutingdataMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if Skills, ok := ConversationcobrowseeventtopicconversationroutingdataMap["skills"].([]interface{}); ok {
		SkillsString, _ := json.Marshal(Skills)
		json.Unmarshal(SkillsString, &o.Skills)
	}
	
	if ScoredAgents, ok := ConversationcobrowseeventtopicconversationroutingdataMap["scoredAgents"].([]interface{}); ok {
		ScoredAgentsString, _ := json.Marshal(ScoredAgents)
		json.Unmarshal(ScoredAgentsString, &o.ScoredAgents)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcobrowseeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

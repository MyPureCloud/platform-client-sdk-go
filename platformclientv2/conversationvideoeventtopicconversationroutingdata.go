package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationvideoeventtopicconversationroutingdata
type Conversationvideoeventtopicconversationroutingdata struct { 
	// Queue
	Queue *Conversationvideoeventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Conversationvideoeventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Conversationvideoeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Conversationvideoeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

func (o *Conversationvideoeventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationvideoeventtopicconversationroutingdata
	
	return json.Marshal(&struct { 
		Queue *Conversationvideoeventtopicurireference `json:"queue,omitempty"`
		
		Language *Conversationvideoeventtopicurireference `json:"language,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Conversationvideoeventtopicurireference `json:"skills,omitempty"`
		
		ScoredAgents *[]Conversationvideoeventtopicscoredagent `json:"scoredAgents,omitempty"`
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

func (o *Conversationvideoeventtopicconversationroutingdata) UnmarshalJSON(b []byte) error {
	var ConversationvideoeventtopicconversationroutingdataMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationvideoeventtopicconversationroutingdataMap)
	if err != nil {
		return err
	}
	
	if Queue, ok := ConversationvideoeventtopicconversationroutingdataMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Language, ok := ConversationvideoeventtopicconversationroutingdataMap["language"].(map[string]interface{}); ok {
		LanguageString, _ := json.Marshal(Language)
		json.Unmarshal(LanguageString, &o.Language)
	}
	
	if Priority, ok := ConversationvideoeventtopicconversationroutingdataMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if Skills, ok := ConversationvideoeventtopicconversationroutingdataMap["skills"].([]interface{}); ok {
		SkillsString, _ := json.Marshal(Skills)
		json.Unmarshal(SkillsString, &o.Skills)
	}
	
	if ScoredAgents, ok := ConversationvideoeventtopicconversationroutingdataMap["scoredAgents"].([]interface{}); ok {
		ScoredAgentsString, _ := json.Marshal(ScoredAgents)
		json.Unmarshal(ScoredAgentsString, &o.ScoredAgents)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationvideoeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

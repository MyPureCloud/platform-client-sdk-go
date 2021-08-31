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

func (o *Conversationcalleventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
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
		Queue: o.Queue,
		
		Language: o.Language,
		
		Priority: o.Priority,
		
		Skills: o.Skills,
		
		ScoredAgents: o.ScoredAgents,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationcalleventtopicconversationroutingdata) UnmarshalJSON(b []byte) error {
	var ConversationcalleventtopicconversationroutingdataMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcalleventtopicconversationroutingdataMap)
	if err != nil {
		return err
	}
	
	if Queue, ok := ConversationcalleventtopicconversationroutingdataMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Language, ok := ConversationcalleventtopicconversationroutingdataMap["language"].(map[string]interface{}); ok {
		LanguageString, _ := json.Marshal(Language)
		json.Unmarshal(LanguageString, &o.Language)
	}
	
	if Priority, ok := ConversationcalleventtopicconversationroutingdataMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if Skills, ok := ConversationcalleventtopicconversationroutingdataMap["skills"].([]interface{}); ok {
		SkillsString, _ := json.Marshal(Skills)
		json.Unmarshal(SkillsString, &o.Skills)
	}
	
	if ScoredAgents, ok := ConversationcalleventtopicconversationroutingdataMap["scoredAgents"].([]interface{}); ok {
		ScoredAgentsString, _ := json.Marshal(ScoredAgents)
		json.Unmarshal(ScoredAgentsString, &o.ScoredAgents)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcalleventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

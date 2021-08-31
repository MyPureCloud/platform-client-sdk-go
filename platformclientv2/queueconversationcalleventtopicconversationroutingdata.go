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

func (o *Queueconversationcalleventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
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
		Queue: o.Queue,
		
		Language: o.Language,
		
		Priority: o.Priority,
		
		Skills: o.Skills,
		
		ScoredAgents: o.ScoredAgents,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationcalleventtopicconversationroutingdata) UnmarshalJSON(b []byte) error {
	var QueueconversationcalleventtopicconversationroutingdataMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationcalleventtopicconversationroutingdataMap)
	if err != nil {
		return err
	}
	
	if Queue, ok := QueueconversationcalleventtopicconversationroutingdataMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Language, ok := QueueconversationcalleventtopicconversationroutingdataMap["language"].(map[string]interface{}); ok {
		LanguageString, _ := json.Marshal(Language)
		json.Unmarshal(LanguageString, &o.Language)
	}
	
	if Priority, ok := QueueconversationcalleventtopicconversationroutingdataMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if Skills, ok := QueueconversationcalleventtopicconversationroutingdataMap["skills"].([]interface{}); ok {
		SkillsString, _ := json.Marshal(Skills)
		json.Unmarshal(SkillsString, &o.Skills)
	}
	
	if ScoredAgents, ok := QueueconversationcalleventtopicconversationroutingdataMap["scoredAgents"].([]interface{}); ok {
		ScoredAgentsString, _ := json.Marshal(ScoredAgents)
		json.Unmarshal(ScoredAgentsString, &o.ScoredAgents)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationcalleventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

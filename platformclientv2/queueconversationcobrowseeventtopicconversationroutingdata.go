package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcobrowseeventtopicconversationroutingdata
type Queueconversationcobrowseeventtopicconversationroutingdata struct { 
	// Queue
	Queue *Queueconversationcobrowseeventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Queueconversationcobrowseeventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Queueconversationcobrowseeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Queueconversationcobrowseeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

func (o *Queueconversationcobrowseeventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationcobrowseeventtopicconversationroutingdata
	
	return json.Marshal(&struct { 
		Queue *Queueconversationcobrowseeventtopicurireference `json:"queue,omitempty"`
		
		Language *Queueconversationcobrowseeventtopicurireference `json:"language,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Queueconversationcobrowseeventtopicurireference `json:"skills,omitempty"`
		
		ScoredAgents *[]Queueconversationcobrowseeventtopicscoredagent `json:"scoredAgents,omitempty"`
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

func (o *Queueconversationcobrowseeventtopicconversationroutingdata) UnmarshalJSON(b []byte) error {
	var QueueconversationcobrowseeventtopicconversationroutingdataMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationcobrowseeventtopicconversationroutingdataMap)
	if err != nil {
		return err
	}
	
	if Queue, ok := QueueconversationcobrowseeventtopicconversationroutingdataMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Language, ok := QueueconversationcobrowseeventtopicconversationroutingdataMap["language"].(map[string]interface{}); ok {
		LanguageString, _ := json.Marshal(Language)
		json.Unmarshal(LanguageString, &o.Language)
	}
	
	if Priority, ok := QueueconversationcobrowseeventtopicconversationroutingdataMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if Skills, ok := QueueconversationcobrowseeventtopicconversationroutingdataMap["skills"].([]interface{}); ok {
		SkillsString, _ := json.Marshal(Skills)
		json.Unmarshal(SkillsString, &o.Skills)
	}
	
	if ScoredAgents, ok := QueueconversationcobrowseeventtopicconversationroutingdataMap["scoredAgents"].([]interface{}); ok {
		ScoredAgentsString, _ := json.Marshal(ScoredAgents)
		json.Unmarshal(ScoredAgentsString, &o.ScoredAgents)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationcobrowseeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationvideoeventtopicconversationroutingdata
type Queueconversationvideoeventtopicconversationroutingdata struct { 
	// Queue
	Queue *Queueconversationvideoeventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Queueconversationvideoeventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Queueconversationvideoeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Queueconversationvideoeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

func (o *Queueconversationvideoeventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationvideoeventtopicconversationroutingdata
	
	return json.Marshal(&struct { 
		Queue *Queueconversationvideoeventtopicurireference `json:"queue,omitempty"`
		
		Language *Queueconversationvideoeventtopicurireference `json:"language,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Queueconversationvideoeventtopicurireference `json:"skills,omitempty"`
		
		ScoredAgents *[]Queueconversationvideoeventtopicscoredagent `json:"scoredAgents,omitempty"`
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

func (o *Queueconversationvideoeventtopicconversationroutingdata) UnmarshalJSON(b []byte) error {
	var QueueconversationvideoeventtopicconversationroutingdataMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationvideoeventtopicconversationroutingdataMap)
	if err != nil {
		return err
	}
	
	if Queue, ok := QueueconversationvideoeventtopicconversationroutingdataMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Language, ok := QueueconversationvideoeventtopicconversationroutingdataMap["language"].(map[string]interface{}); ok {
		LanguageString, _ := json.Marshal(Language)
		json.Unmarshal(LanguageString, &o.Language)
	}
	
	if Priority, ok := QueueconversationvideoeventtopicconversationroutingdataMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if Skills, ok := QueueconversationvideoeventtopicconversationroutingdataMap["skills"].([]interface{}); ok {
		SkillsString, _ := json.Marshal(Skills)
		json.Unmarshal(SkillsString, &o.Skills)
	}
	
	if ScoredAgents, ok := QueueconversationvideoeventtopicconversationroutingdataMap["scoredAgents"].([]interface{}); ok {
		ScoredAgentsString, _ := json.Marshal(ScoredAgents)
		json.Unmarshal(ScoredAgentsString, &o.ScoredAgents)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

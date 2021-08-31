package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationscreenshareeventtopicconversationroutingdata
type Queueconversationscreenshareeventtopicconversationroutingdata struct { 
	// Queue
	Queue *Queueconversationscreenshareeventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Queueconversationscreenshareeventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Queueconversationscreenshareeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Queueconversationscreenshareeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

func (o *Queueconversationscreenshareeventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationscreenshareeventtopicconversationroutingdata
	
	return json.Marshal(&struct { 
		Queue *Queueconversationscreenshareeventtopicurireference `json:"queue,omitempty"`
		
		Language *Queueconversationscreenshareeventtopicurireference `json:"language,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Queueconversationscreenshareeventtopicurireference `json:"skills,omitempty"`
		
		ScoredAgents *[]Queueconversationscreenshareeventtopicscoredagent `json:"scoredAgents,omitempty"`
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

func (o *Queueconversationscreenshareeventtopicconversationroutingdata) UnmarshalJSON(b []byte) error {
	var QueueconversationscreenshareeventtopicconversationroutingdataMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationscreenshareeventtopicconversationroutingdataMap)
	if err != nil {
		return err
	}
	
	if Queue, ok := QueueconversationscreenshareeventtopicconversationroutingdataMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Language, ok := QueueconversationscreenshareeventtopicconversationroutingdataMap["language"].(map[string]interface{}); ok {
		LanguageString, _ := json.Marshal(Language)
		json.Unmarshal(LanguageString, &o.Language)
	}
	
	if Priority, ok := QueueconversationscreenshareeventtopicconversationroutingdataMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if Skills, ok := QueueconversationscreenshareeventtopicconversationroutingdataMap["skills"].([]interface{}); ok {
		SkillsString, _ := json.Marshal(Skills)
		json.Unmarshal(SkillsString, &o.Skills)
	}
	
	if ScoredAgents, ok := QueueconversationscreenshareeventtopicconversationroutingdataMap["scoredAgents"].([]interface{}); ok {
		ScoredAgentsString, _ := json.Marshal(ScoredAgents)
		json.Unmarshal(ScoredAgentsString, &o.ScoredAgents)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationscreenshareeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

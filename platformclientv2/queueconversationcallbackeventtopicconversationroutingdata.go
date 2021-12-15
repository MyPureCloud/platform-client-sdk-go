package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcallbackeventtopicconversationroutingdata
type Queueconversationcallbackeventtopicconversationroutingdata struct { 
	// Queue - A UriReference for a resource
	Queue *Queueconversationcallbackeventtopicurireference `json:"queue,omitempty"`


	// Language - A UriReference for a resource
	Language *Queueconversationcallbackeventtopicurireference `json:"language,omitempty"`


	// Priority - The priority of the conversation to use for routing decisions
	Priority *int `json:"priority,omitempty"`


	// Skills - The skills to use for routing decisions
	Skills *[]Queueconversationcallbackeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents - A collection of agents and their assigned scores for this conversation (0 - 100, higher being better), for use in routing to preferred agents
	ScoredAgents *[]Queueconversationcallbackeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

func (o *Queueconversationcallbackeventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationcallbackeventtopicconversationroutingdata
	
	return json.Marshal(&struct { 
		Queue *Queueconversationcallbackeventtopicurireference `json:"queue,omitempty"`
		
		Language *Queueconversationcallbackeventtopicurireference `json:"language,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Queueconversationcallbackeventtopicurireference `json:"skills,omitempty"`
		
		ScoredAgents *[]Queueconversationcallbackeventtopicscoredagent `json:"scoredAgents,omitempty"`
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

func (o *Queueconversationcallbackeventtopicconversationroutingdata) UnmarshalJSON(b []byte) error {
	var QueueconversationcallbackeventtopicconversationroutingdataMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationcallbackeventtopicconversationroutingdataMap)
	if err != nil {
		return err
	}
	
	if Queue, ok := QueueconversationcallbackeventtopicconversationroutingdataMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Language, ok := QueueconversationcallbackeventtopicconversationroutingdataMap["language"].(map[string]interface{}); ok {
		LanguageString, _ := json.Marshal(Language)
		json.Unmarshal(LanguageString, &o.Language)
	}
	
	if Priority, ok := QueueconversationcallbackeventtopicconversationroutingdataMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if Skills, ok := QueueconversationcallbackeventtopicconversationroutingdataMap["skills"].([]interface{}); ok {
		SkillsString, _ := json.Marshal(Skills)
		json.Unmarshal(SkillsString, &o.Skills)
	}
	
	if ScoredAgents, ok := QueueconversationcallbackeventtopicconversationroutingdataMap["scoredAgents"].([]interface{}); ok {
		ScoredAgentsString, _ := json.Marshal(ScoredAgents)
		json.Unmarshal(ScoredAgentsString, &o.ScoredAgents)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationcallbackeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

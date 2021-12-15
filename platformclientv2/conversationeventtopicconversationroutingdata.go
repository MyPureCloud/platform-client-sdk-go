package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationeventtopicconversationroutingdata - Information on how a communication should be routed to an agent.
type Conversationeventtopicconversationroutingdata struct { 
	// Queue
	Queue *Conversationeventtopicurireference `json:"queue,omitempty"`


	// Language - A UriReference for a resource
	Language *Conversationeventtopicurireference `json:"language,omitempty"`


	// Priority - The priority of the conversation to use for routing decisions
	Priority *int `json:"priority,omitempty"`


	// Skills - The skills to use for routing decisions
	Skills *[]Conversationeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents - A collection of agents and their assigned scores for this conversation (0 - 100, higher being better), for use in routing to preferred agents
	ScoredAgents *[]Conversationeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

func (o *Conversationeventtopicconversationroutingdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationeventtopicconversationroutingdata
	
	return json.Marshal(&struct { 
		Queue *Conversationeventtopicurireference `json:"queue,omitempty"`
		
		Language *Conversationeventtopicurireference `json:"language,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Conversationeventtopicurireference `json:"skills,omitempty"`
		
		ScoredAgents *[]Conversationeventtopicscoredagent `json:"scoredAgents,omitempty"`
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

func (o *Conversationeventtopicconversationroutingdata) UnmarshalJSON(b []byte) error {
	var ConversationeventtopicconversationroutingdataMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationeventtopicconversationroutingdataMap)
	if err != nil {
		return err
	}
	
	if Queue, ok := ConversationeventtopicconversationroutingdataMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Language, ok := ConversationeventtopicconversationroutingdataMap["language"].(map[string]interface{}); ok {
		LanguageString, _ := json.Marshal(Language)
		json.Unmarshal(LanguageString, &o.Language)
	}
	
	if Priority, ok := ConversationeventtopicconversationroutingdataMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if Skills, ok := ConversationeventtopicconversationroutingdataMap["skills"].([]interface{}); ok {
		SkillsString, _ := json.Marshal(Skills)
		json.Unmarshal(SkillsString, &o.Skills)
	}
	
	if ScoredAgents, ok := ConversationeventtopicconversationroutingdataMap["scoredAgents"].([]interface{}); ok {
		ScoredAgentsString, _ := json.Marshal(ScoredAgents)
		json.Unmarshal(ScoredAgentsString, &o.ScoredAgents)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

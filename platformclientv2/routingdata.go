package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Routingdata
type Routingdata struct { 
	// QueueId - The identifier of the routing queue
	QueueId *string `json:"queueId,omitempty"`


	// LanguageId - The identifier of a language to be considered in routing
	LanguageId *string `json:"languageId,omitempty"`


	// Priority - The priority for routing
	Priority *int `json:"priority,omitempty"`


	// SkillIds - A list of skill identifiers to be considered in routing
	SkillIds *[]string `json:"skillIds,omitempty"`


	// PreferredAgentIds - A list of agents to be preferred in routing
	PreferredAgentIds *[]string `json:"preferredAgentIds,omitempty"`


	// ScoredAgents - A list of scored agents for routing decisions
	ScoredAgents *[]Scoredagent `json:"scoredAgents,omitempty"`


	// RoutingFlags - An array of flags indicating how the conversation should be routed
	RoutingFlags *[]string `json:"routingFlags,omitempty"`

}

func (o *Routingdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Routingdata
	
	return json.Marshal(&struct { 
		QueueId *string `json:"queueId,omitempty"`
		
		LanguageId *string `json:"languageId,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		SkillIds *[]string `json:"skillIds,omitempty"`
		
		PreferredAgentIds *[]string `json:"preferredAgentIds,omitempty"`
		
		ScoredAgents *[]Scoredagent `json:"scoredAgents,omitempty"`
		
		RoutingFlags *[]string `json:"routingFlags,omitempty"`
		*Alias
	}{ 
		QueueId: o.QueueId,
		
		LanguageId: o.LanguageId,
		
		Priority: o.Priority,
		
		SkillIds: o.SkillIds,
		
		PreferredAgentIds: o.PreferredAgentIds,
		
		ScoredAgents: o.ScoredAgents,
		
		RoutingFlags: o.RoutingFlags,
		Alias:    (*Alias)(o),
	})
}

func (o *Routingdata) UnmarshalJSON(b []byte) error {
	var RoutingdataMap map[string]interface{}
	err := json.Unmarshal(b, &RoutingdataMap)
	if err != nil {
		return err
	}
	
	if QueueId, ok := RoutingdataMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if LanguageId, ok := RoutingdataMap["languageId"].(string); ok {
		o.LanguageId = &LanguageId
	}
    
	if Priority, ok := RoutingdataMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if SkillIds, ok := RoutingdataMap["skillIds"].([]interface{}); ok {
		SkillIdsString, _ := json.Marshal(SkillIds)
		json.Unmarshal(SkillIdsString, &o.SkillIds)
	}
	
	if PreferredAgentIds, ok := RoutingdataMap["preferredAgentIds"].([]interface{}); ok {
		PreferredAgentIdsString, _ := json.Marshal(PreferredAgentIds)
		json.Unmarshal(PreferredAgentIdsString, &o.PreferredAgentIds)
	}
	
	if ScoredAgents, ok := RoutingdataMap["scoredAgents"].([]interface{}); ok {
		ScoredAgentsString, _ := json.Marshal(ScoredAgents)
		json.Unmarshal(ScoredAgentsString, &o.ScoredAgents)
	}
	
	if RoutingFlags, ok := RoutingdataMap["routingFlags"].([]interface{}); ok {
		RoutingFlagsString, _ := json.Marshal(RoutingFlags)
		json.Unmarshal(RoutingFlagsString, &o.RoutingFlags)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Routingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

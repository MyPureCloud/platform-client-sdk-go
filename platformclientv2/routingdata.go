package platformclientv2
import (
	"encoding/json"
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

}

// String returns a JSON representation of the model
func (o *Routingdata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

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

}

func (u *Routingdata) MarshalJSON() ([]byte, error) {
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
		*Alias
	}{ 
		QueueId: u.QueueId,
		
		LanguageId: u.LanguageId,
		
		Priority: u.Priority,
		
		SkillIds: u.SkillIds,
		
		PreferredAgentIds: u.PreferredAgentIds,
		
		ScoredAgents: u.ScoredAgents,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Routingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

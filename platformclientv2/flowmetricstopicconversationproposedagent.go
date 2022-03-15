package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowmetricstopicconversationproposedagent
type Flowmetricstopicconversationproposedagent struct { 
	// AgentRank - Proposed agent rank for this conversation from predictive routing (lower is better)
	AgentRank *int `json:"agentRank,omitempty"`


	// ProposedAgentId - Unique identifier for the agent that was proposed by predictive routing
	ProposedAgentId *string `json:"proposedAgentId,omitempty"`

}

func (o *Flowmetricstopicconversationproposedagent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Flowmetricstopicconversationproposedagent
	
	return json.Marshal(&struct { 
		AgentRank *int `json:"agentRank,omitempty"`
		
		ProposedAgentId *string `json:"proposedAgentId,omitempty"`
		*Alias
	}{ 
		AgentRank: o.AgentRank,
		
		ProposedAgentId: o.ProposedAgentId,
		Alias:    (*Alias)(o),
	})
}

func (o *Flowmetricstopicconversationproposedagent) UnmarshalJSON(b []byte) error {
	var FlowmetricstopicconversationproposedagentMap map[string]interface{}
	err := json.Unmarshal(b, &FlowmetricstopicconversationproposedagentMap)
	if err != nil {
		return err
	}
	
	if AgentRank, ok := FlowmetricstopicconversationproposedagentMap["agentRank"].(float64); ok {
		AgentRankInt := int(AgentRank)
		o.AgentRank = &AgentRankInt
	}
	
	if ProposedAgentId, ok := FlowmetricstopicconversationproposedagentMap["proposedAgentId"].(string); ok {
		o.ProposedAgentId = &ProposedAgentId
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Flowmetricstopicconversationproposedagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

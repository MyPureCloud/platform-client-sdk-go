package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsproposedagent
type Analyticsproposedagent struct { 
	// ProposedAgentId - Unique identifier of an agent that was proposed by predictive routing
	ProposedAgentId *string `json:"proposedAgentId,omitempty"`


	// AgentRank - Proposed agent rank for this conversation from predictive routing (lower is better)
	AgentRank *int `json:"agentRank,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsproposedagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsproposedagent
type Analyticsproposedagent struct { 
	// AgentRank - Proposed agent rank for this conversation from predictive routing (lower is better)
	AgentRank *int `json:"agentRank,omitempty"`


	// ProposedAgentId - Unique identifier for the agent that was proposed by predictive routing
	ProposedAgentId *string `json:"proposedAgentId,omitempty"`

}

func (u *Analyticsproposedagent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsproposedagent

	

	return json.Marshal(&struct { 
		AgentRank *int `json:"agentRank,omitempty"`
		
		ProposedAgentId *string `json:"proposedAgentId,omitempty"`
		*Alias
	}{ 
		AgentRank: u.AgentRank,
		
		ProposedAgentId: u.ProposedAgentId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Analyticsproposedagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsscoredagent
type Analyticsscoredagent struct { 
	// ScoredAgentId - Unique identifier of an agent that was scored for this conversation
	ScoredAgentId *string `json:"scoredAgentId,omitempty"`


	// AgentScore - Assigned agent score for this conversation (0 - 100, higher being better)
	AgentScore *int `json:"agentScore,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsscoredagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"encoding/json"
)

// Analyticsscoredagent
type Analyticsscoredagent struct { 
	// ScoredAgentId - Unique identifier of an agent that was scored for this conversation
	ScoredAgentId *string `json:"scoredAgentId,omitempty"`


	// AgentScore - Assigned agent score for this conversation (0 - 100, higher being better)
	AgentScore *int32 `json:"agentScore,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsscoredagent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

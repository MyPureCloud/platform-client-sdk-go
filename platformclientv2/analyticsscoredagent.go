package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsscoredagent
type Analyticsscoredagent struct { 
	// AgentScore - Assigned agent score for this conversation (0 - 100, higher being better)
	AgentScore *int `json:"agentScore,omitempty"`


	// ScoredAgentId - Unique identifier for the agent that was scored for this conversation
	ScoredAgentId *string `json:"scoredAgentId,omitempty"`

}

func (o *Analyticsscoredagent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsscoredagent
	
	return json.Marshal(&struct { 
		AgentScore *int `json:"agentScore,omitempty"`
		
		ScoredAgentId *string `json:"scoredAgentId,omitempty"`
		*Alias
	}{ 
		AgentScore: o.AgentScore,
		
		ScoredAgentId: o.ScoredAgentId,
		Alias:    (*Alias)(o),
	})
}

func (o *Analyticsscoredagent) UnmarshalJSON(b []byte) error {
	var AnalyticsscoredagentMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsscoredagentMap)
	if err != nil {
		return err
	}
	
	if AgentScore, ok := AnalyticsscoredagentMap["agentScore"].(float64); ok {
		AgentScoreInt := int(AgentScore)
		o.AgentScore = &AgentScoreInt
	}
	
	if ScoredAgentId, ok := AnalyticsscoredagentMap["scoredAgentId"].(string); ok {
		o.ScoredAgentId = &ScoredAgentId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsscoredagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowmetricstopicconversationscoredagent
type Flowmetricstopicconversationscoredagent struct { 
	// AgentScore - Assigned agent score for this conversation (0 - 100, higher being better)
	AgentScore *int `json:"agentScore,omitempty"`


	// ScoredAgentId - Unique identifier for the agent that was scored for this conversation
	ScoredAgentId *string `json:"scoredAgentId,omitempty"`

}

func (o *Flowmetricstopicconversationscoredagent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Flowmetricstopicconversationscoredagent
	
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

func (o *Flowmetricstopicconversationscoredagent) UnmarshalJSON(b []byte) error {
	var FlowmetricstopicconversationscoredagentMap map[string]interface{}
	err := json.Unmarshal(b, &FlowmetricstopicconversationscoredagentMap)
	if err != nil {
		return err
	}
	
	if AgentScore, ok := FlowmetricstopicconversationscoredagentMap["agentScore"].(float64); ok {
		AgentScoreInt := int(AgentScore)
		o.AgentScore = &AgentScoreInt
	}
	
	if ScoredAgentId, ok := FlowmetricstopicconversationscoredagentMap["scoredAgentId"].(string); ok {
		o.ScoredAgentId = &ScoredAgentId
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Flowmetricstopicconversationscoredagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationsocialexpressioneventtopicscoredagent
type Conversationsocialexpressioneventtopicscoredagent struct { 
	// Agent
	Agent *Conversationsocialexpressioneventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int `json:"score,omitempty"`

}

func (o *Conversationsocialexpressioneventtopicscoredagent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationsocialexpressioneventtopicscoredagent
	
	return json.Marshal(&struct { 
		Agent *Conversationsocialexpressioneventtopicurireference `json:"agent,omitempty"`
		
		Score *int `json:"score,omitempty"`
		*Alias
	}{ 
		Agent: o.Agent,
		
		Score: o.Score,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationsocialexpressioneventtopicscoredagent) UnmarshalJSON(b []byte) error {
	var ConversationsocialexpressioneventtopicscoredagentMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationsocialexpressioneventtopicscoredagentMap)
	if err != nil {
		return err
	}
	
	if Agent, ok := ConversationsocialexpressioneventtopicscoredagentMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if Score, ok := ConversationsocialexpressioneventtopicscoredagentMap["score"].(float64); ok {
		ScoreInt := int(Score)
		o.Score = &ScoreInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationsocialexpressioneventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

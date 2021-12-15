package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationsocialexpressioneventtopicscoredagent
type Queueconversationsocialexpressioneventtopicscoredagent struct { 
	// Agent - A UriReference for a resource
	Agent *Queueconversationsocialexpressioneventtopicurireference `json:"agent,omitempty"`


	// Score - Agent's score for the current conversation, from 0 - 100, higher being better
	Score *int `json:"score,omitempty"`

}

func (o *Queueconversationsocialexpressioneventtopicscoredagent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationsocialexpressioneventtopicscoredagent
	
	return json.Marshal(&struct { 
		Agent *Queueconversationsocialexpressioneventtopicurireference `json:"agent,omitempty"`
		
		Score *int `json:"score,omitempty"`
		*Alias
	}{ 
		Agent: o.Agent,
		
		Score: o.Score,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationsocialexpressioneventtopicscoredagent) UnmarshalJSON(b []byte) error {
	var QueueconversationsocialexpressioneventtopicscoredagentMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationsocialexpressioneventtopicscoredagentMap)
	if err != nil {
		return err
	}
	
	if Agent, ok := QueueconversationsocialexpressioneventtopicscoredagentMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if Score, ok := QueueconversationsocialexpressioneventtopicscoredagentMap["score"].(float64); ok {
		ScoreInt := int(Score)
		o.Score = &ScoreInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

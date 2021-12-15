package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationemaileventtopicscoredagent
type Queueconversationemaileventtopicscoredagent struct { 
	// Agent - A UriReference for a resource
	Agent *Queueconversationemaileventtopicurireference `json:"agent,omitempty"`


	// Score - Agent's score for the current conversation, from 0 - 100, higher being better
	Score *int `json:"score,omitempty"`

}

func (o *Queueconversationemaileventtopicscoredagent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationemaileventtopicscoredagent
	
	return json.Marshal(&struct { 
		Agent *Queueconversationemaileventtopicurireference `json:"agent,omitempty"`
		
		Score *int `json:"score,omitempty"`
		*Alias
	}{ 
		Agent: o.Agent,
		
		Score: o.Score,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationemaileventtopicscoredagent) UnmarshalJSON(b []byte) error {
	var QueueconversationemaileventtopicscoredagentMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationemaileventtopicscoredagentMap)
	if err != nil {
		return err
	}
	
	if Agent, ok := QueueconversationemaileventtopicscoredagentMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if Score, ok := QueueconversationemaileventtopicscoredagentMap["score"].(float64); ok {
		ScoreInt := int(Score)
		o.Score = &ScoreInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationemaileventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

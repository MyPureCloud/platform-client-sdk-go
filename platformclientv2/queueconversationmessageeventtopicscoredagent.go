package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationmessageeventtopicscoredagent
type Queueconversationmessageeventtopicscoredagent struct { 
	// Agent
	Agent *Queueconversationmessageeventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int `json:"score,omitempty"`

}

func (o *Queueconversationmessageeventtopicscoredagent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationmessageeventtopicscoredagent
	
	return json.Marshal(&struct { 
		Agent *Queueconversationmessageeventtopicurireference `json:"agent,omitempty"`
		
		Score *int `json:"score,omitempty"`
		*Alias
	}{ 
		Agent: o.Agent,
		
		Score: o.Score,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationmessageeventtopicscoredagent) UnmarshalJSON(b []byte) error {
	var QueueconversationmessageeventtopicscoredagentMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationmessageeventtopicscoredagentMap)
	if err != nil {
		return err
	}
	
	if Agent, ok := QueueconversationmessageeventtopicscoredagentMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if Score, ok := QueueconversationmessageeventtopicscoredagentMap["score"].(float64); ok {
		ScoreInt := int(Score)
		o.Score = &ScoreInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationmessageeventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

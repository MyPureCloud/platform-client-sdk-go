package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmessageeventtopicscoredagent
type Conversationmessageeventtopicscoredagent struct { 
	// Agent
	Agent *Conversationmessageeventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int `json:"score,omitempty"`

}

func (o *Conversationmessageeventtopicscoredagent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationmessageeventtopicscoredagent
	
	return json.Marshal(&struct { 
		Agent *Conversationmessageeventtopicurireference `json:"agent,omitempty"`
		
		Score *int `json:"score,omitempty"`
		*Alias
	}{ 
		Agent: o.Agent,
		
		Score: o.Score,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationmessageeventtopicscoredagent) UnmarshalJSON(b []byte) error {
	var ConversationmessageeventtopicscoredagentMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationmessageeventtopicscoredagentMap)
	if err != nil {
		return err
	}
	
	if Agent, ok := ConversationmessageeventtopicscoredagentMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if Score, ok := ConversationmessageeventtopicscoredagentMap["score"].(float64); ok {
		ScoreInt := int(Score)
		o.Score = &ScoreInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationmessageeventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

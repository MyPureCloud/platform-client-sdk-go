package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationscreenshareeventtopicscoredagent
type Conversationscreenshareeventtopicscoredagent struct { 
	// Agent
	Agent *Conversationscreenshareeventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int `json:"score,omitempty"`

}

func (o *Conversationscreenshareeventtopicscoredagent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationscreenshareeventtopicscoredagent
	
	return json.Marshal(&struct { 
		Agent *Conversationscreenshareeventtopicurireference `json:"agent,omitempty"`
		
		Score *int `json:"score,omitempty"`
		*Alias
	}{ 
		Agent: o.Agent,
		
		Score: o.Score,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationscreenshareeventtopicscoredagent) UnmarshalJSON(b []byte) error {
	var ConversationscreenshareeventtopicscoredagentMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationscreenshareeventtopicscoredagentMap)
	if err != nil {
		return err
	}
	
	if Agent, ok := ConversationscreenshareeventtopicscoredagentMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if Score, ok := ConversationscreenshareeventtopicscoredagentMap["score"].(float64); ok {
		ScoreInt := int(Score)
		o.Score = &ScoreInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationscreenshareeventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

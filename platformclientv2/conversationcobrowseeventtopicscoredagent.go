package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcobrowseeventtopicscoredagent
type Conversationcobrowseeventtopicscoredagent struct { 
	// Agent
	Agent *Conversationcobrowseeventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int `json:"score,omitempty"`

}

func (o *Conversationcobrowseeventtopicscoredagent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcobrowseeventtopicscoredagent
	
	return json.Marshal(&struct { 
		Agent *Conversationcobrowseeventtopicurireference `json:"agent,omitempty"`
		
		Score *int `json:"score,omitempty"`
		*Alias
	}{ 
		Agent: o.Agent,
		
		Score: o.Score,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationcobrowseeventtopicscoredagent) UnmarshalJSON(b []byte) error {
	var ConversationcobrowseeventtopicscoredagentMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcobrowseeventtopicscoredagentMap)
	if err != nil {
		return err
	}
	
	if Agent, ok := ConversationcobrowseeventtopicscoredagentMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if Score, ok := ConversationcobrowseeventtopicscoredagentMap["score"].(float64); ok {
		ScoreInt := int(Score)
		o.Score = &ScoreInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcobrowseeventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

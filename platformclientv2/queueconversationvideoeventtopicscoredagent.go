package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationvideoeventtopicscoredagent
type Queueconversationvideoeventtopicscoredagent struct { 
	// Agent
	Agent *Queueconversationvideoeventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int `json:"score,omitempty"`

}

func (o *Queueconversationvideoeventtopicscoredagent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationvideoeventtopicscoredagent
	
	return json.Marshal(&struct { 
		Agent *Queueconversationvideoeventtopicurireference `json:"agent,omitempty"`
		
		Score *int `json:"score,omitempty"`
		*Alias
	}{ 
		Agent: o.Agent,
		
		Score: o.Score,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationvideoeventtopicscoredagent) UnmarshalJSON(b []byte) error {
	var QueueconversationvideoeventtopicscoredagentMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationvideoeventtopicscoredagentMap)
	if err != nil {
		return err
	}
	
	if Agent, ok := QueueconversationvideoeventtopicscoredagentMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if Score, ok := QueueconversationvideoeventtopicscoredagentMap["score"].(float64); ok {
		ScoreInt := int(Score)
		o.Score = &ScoreInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcallbackeventtopicscoredagent
type Queueconversationcallbackeventtopicscoredagent struct { 
	// Agent
	Agent *Queueconversationcallbackeventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int `json:"score,omitempty"`

}

func (o *Queueconversationcallbackeventtopicscoredagent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationcallbackeventtopicscoredagent
	
	return json.Marshal(&struct { 
		Agent *Queueconversationcallbackeventtopicurireference `json:"agent,omitempty"`
		
		Score *int `json:"score,omitempty"`
		*Alias
	}{ 
		Agent: o.Agent,
		
		Score: o.Score,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationcallbackeventtopicscoredagent) UnmarshalJSON(b []byte) error {
	var QueueconversationcallbackeventtopicscoredagentMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationcallbackeventtopicscoredagentMap)
	if err != nil {
		return err
	}
	
	if Agent, ok := QueueconversationcallbackeventtopicscoredagentMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if Score, ok := QueueconversationcallbackeventtopicscoredagentMap["score"].(float64); ok {
		ScoreInt := int(Score)
		o.Score = &ScoreInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationcallbackeventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

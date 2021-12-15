package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationscreenshareeventtopicscoredagent
type Queueconversationscreenshareeventtopicscoredagent struct { 
	// Agent - A UriReference for a resource
	Agent *Queueconversationscreenshareeventtopicurireference `json:"agent,omitempty"`


	// Score - Agent's score for the current conversation, from 0 - 100, higher being better
	Score *int `json:"score,omitempty"`

}

func (o *Queueconversationscreenshareeventtopicscoredagent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationscreenshareeventtopicscoredagent
	
	return json.Marshal(&struct { 
		Agent *Queueconversationscreenshareeventtopicurireference `json:"agent,omitempty"`
		
		Score *int `json:"score,omitempty"`
		*Alias
	}{ 
		Agent: o.Agent,
		
		Score: o.Score,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationscreenshareeventtopicscoredagent) UnmarshalJSON(b []byte) error {
	var QueueconversationscreenshareeventtopicscoredagentMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationscreenshareeventtopicscoredagentMap)
	if err != nil {
		return err
	}
	
	if Agent, ok := QueueconversationscreenshareeventtopicscoredagentMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if Score, ok := QueueconversationscreenshareeventtopicscoredagentMap["score"].(float64); ok {
		ScoreInt := int(Score)
		o.Score = &ScoreInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationscreenshareeventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

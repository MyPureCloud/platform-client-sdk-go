package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Agent
type Agent struct { 
	// Stage - The current stage for this agent
	Stage *string `json:"stage,omitempty"`

}

func (o *Agent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Agent
	
	return json.Marshal(&struct { 
		Stage *string `json:"stage,omitempty"`
		*Alias
	}{ 
		Stage: o.Stage,
		Alias:    (*Alias)(o),
	})
}

func (o *Agent) UnmarshalJSON(b []byte) error {
	var AgentMap map[string]interface{}
	err := json.Unmarshal(b, &AgentMap)
	if err != nil {
		return err
	}
	
	if Stage, ok := AgentMap["stage"].(string); ok {
		o.Stage = &Stage
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Agent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

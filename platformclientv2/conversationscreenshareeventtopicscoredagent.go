package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Conversationscreenshareeventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

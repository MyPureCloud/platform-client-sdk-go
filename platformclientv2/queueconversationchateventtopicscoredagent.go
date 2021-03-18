package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationchateventtopicscoredagent
type Queueconversationchateventtopicscoredagent struct { 
	// Agent
	Agent *Queueconversationchateventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int `json:"score,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationchateventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

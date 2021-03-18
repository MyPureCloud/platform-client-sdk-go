package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationsocialexpressioneventtopicscoredagent
type Queueconversationsocialexpressioneventtopicscoredagent struct { 
	// Agent
	Agent *Queueconversationsocialexpressioneventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int `json:"score,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

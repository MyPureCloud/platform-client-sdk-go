package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationeventtopicscoredagent
type Conversationeventtopicscoredagent struct { 
	// Agent
	Agent *Conversationeventtopicurireference `json:"agent,omitempty"`


	// Score
	Score *int `json:"score,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationeventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Conversationcobrowseeventtopicscoredagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

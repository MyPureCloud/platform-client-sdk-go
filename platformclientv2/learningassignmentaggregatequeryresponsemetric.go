package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentaggregatequeryresponsemetric
type Learningassignmentaggregatequeryresponsemetric struct { 
	// Metric - The metric this applies to
	Metric *string `json:"metric,omitempty"`


	// Stats - The aggregated values for this metric
	Stats *Learningassignmentaggregatequeryresponsestats `json:"stats,omitempty"`

}

// String returns a JSON representation of the model
func (o *Learningassignmentaggregatequeryresponsemetric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

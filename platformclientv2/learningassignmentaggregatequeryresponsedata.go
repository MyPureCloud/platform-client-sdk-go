package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentaggregatequeryresponsedata
type Learningassignmentaggregatequeryresponsedata struct { 
	// Interval - Specifies the range of due dates to be used for filtering. A maximum of 1 year can be specified in the range. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`


	// Metrics - The list of aggregated metrics
	Metrics *[]Learningassignmentaggregatequeryresponsemetric `json:"metrics,omitempty"`

}

// String returns a JSON representation of the model
func (o *Learningassignmentaggregatequeryresponsedata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

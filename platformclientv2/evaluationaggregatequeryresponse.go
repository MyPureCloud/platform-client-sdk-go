package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationaggregatequeryresponse
type Evaluationaggregatequeryresponse struct { 
	// Results
	Results *[]Evaluationaggregatedatacontainer `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Evaluationaggregatequeryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

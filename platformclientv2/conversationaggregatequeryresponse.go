package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationaggregatequeryresponse
type Conversationaggregatequeryresponse struct { 
	// Results
	Results *[]Conversationaggregatedatacontainer `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationaggregatequeryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

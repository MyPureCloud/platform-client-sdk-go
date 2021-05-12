package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Flowaggregatequeryresponse
type Flowaggregatequeryresponse struct { 
	// Results
	Results *[]Flowaggregatedatacontainer `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Flowaggregatequeryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

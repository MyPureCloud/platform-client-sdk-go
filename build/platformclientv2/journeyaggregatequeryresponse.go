package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyaggregatequeryresponse
type Journeyaggregatequeryresponse struct { 
	// Results
	Results *[]Journeyaggregatedatacontainer `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Journeyaggregatequeryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

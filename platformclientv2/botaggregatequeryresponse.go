package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Botaggregatequeryresponse
type Botaggregatequeryresponse struct { 
	// Results
	Results *[]Botaggregatedatacontainer `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Botaggregatequeryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

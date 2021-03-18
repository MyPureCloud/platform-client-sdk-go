package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Transcriptaggregatequeryresponse
type Transcriptaggregatequeryresponse struct { 
	// Results
	Results *[]Transcriptaggregatedatacontainer `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Transcriptaggregatequeryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

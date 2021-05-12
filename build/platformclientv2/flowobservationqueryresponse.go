package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Flowobservationqueryresponse
type Flowobservationqueryresponse struct { 
	// Results
	Results *[]Flowobservationdatacontainer `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Flowobservationqueryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

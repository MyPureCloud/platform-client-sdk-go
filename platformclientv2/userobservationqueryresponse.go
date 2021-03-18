package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Userobservationqueryresponse
type Userobservationqueryresponse struct { 
	// Results
	Results *[]Userobservationdatacontainer `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userobservationqueryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

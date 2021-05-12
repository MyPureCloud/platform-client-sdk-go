package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Coachingappointmentaggregateresponse
type Coachingappointmentaggregateresponse struct { 
	// Results - The results of the query
	Results *[]Queryresponsegroupeddata `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Coachingappointmentaggregateresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

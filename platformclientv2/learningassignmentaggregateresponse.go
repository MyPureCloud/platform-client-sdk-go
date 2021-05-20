package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentaggregateresponse
type Learningassignmentaggregateresponse struct { 
	// Results - The results of the query
	Results *[]Learningassignmentaggregatequeryresponsegroupeddata `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Learningassignmentaggregateresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Developmentactivityaggregateresponse
type Developmentactivityaggregateresponse struct { 
	// Results - The results of the query
	Results *[]Developmentactivityaggregatequeryresponsegroupeddata `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Developmentactivityaggregateresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

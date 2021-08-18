package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Developmentactivityaggregateresponse
type Developmentactivityaggregateresponse struct { 
	// Results - The results of the query
	Results *[]Developmentactivityaggregatequeryresponsegroupeddata `json:"results,omitempty"`

}

func (u *Developmentactivityaggregateresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Developmentactivityaggregateresponse

	

	return json.Marshal(&struct { 
		Results *[]Developmentactivityaggregatequeryresponsegroupeddata `json:"results,omitempty"`
		*Alias
	}{ 
		Results: u.Results,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Developmentactivityaggregateresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

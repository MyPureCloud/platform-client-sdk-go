package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyaggregatequeryresponse
type Journeyaggregatequeryresponse struct { 
	// Results
	Results *[]Journeyaggregatedatacontainer `json:"results,omitempty"`

}

func (u *Journeyaggregatequeryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeyaggregatequeryresponse

	

	return json.Marshal(&struct { 
		Results *[]Journeyaggregatedatacontainer `json:"results,omitempty"`
		*Alias
	}{ 
		Results: u.Results,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Journeyaggregatequeryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

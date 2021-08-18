package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Transcriptaggregatequeryresponse
type Transcriptaggregatequeryresponse struct { 
	// Results
	Results *[]Transcriptaggregatedatacontainer `json:"results,omitempty"`

}

func (u *Transcriptaggregatequeryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Transcriptaggregatequeryresponse

	

	return json.Marshal(&struct { 
		Results *[]Transcriptaggregatedatacontainer `json:"results,omitempty"`
		*Alias
	}{ 
		Results: u.Results,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Transcriptaggregatequeryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

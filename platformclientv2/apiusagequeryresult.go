package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Apiusagequeryresult
type Apiusagequeryresult struct { 
	// Results - Query results
	Results *[]Apiusagerow `json:"results,omitempty"`


	// QueryStatus - Query status
	QueryStatus *string `json:"queryStatus,omitempty"`

}

func (u *Apiusagequeryresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Apiusagequeryresult

	

	return json.Marshal(&struct { 
		Results *[]Apiusagerow `json:"results,omitempty"`
		
		QueryStatus *string `json:"queryStatus,omitempty"`
		*Alias
	}{ 
		Results: u.Results,
		
		QueryStatus: u.QueryStatus,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Apiusagequeryresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

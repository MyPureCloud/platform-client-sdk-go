package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Forecastservicelevelresponse
type Forecastservicelevelresponse struct { 
	// Percent - The percent of calls to answer in the number of seconds defined
	Percent *int `json:"percent,omitempty"`


	// Seconds - The number of seconds to define for the percent of calls to be answered
	Seconds *int `json:"seconds,omitempty"`

}

func (u *Forecastservicelevelresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Forecastservicelevelresponse

	

	return json.Marshal(&struct { 
		Percent *int `json:"percent,omitempty"`
		
		Seconds *int `json:"seconds,omitempty"`
		*Alias
	}{ 
		Percent: u.Percent,
		
		Seconds: u.Seconds,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Forecastservicelevelresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Forecastservicelevelresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

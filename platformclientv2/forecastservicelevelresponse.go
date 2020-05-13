package platformclientv2
import (
	"encoding/json"
)

// Forecastservicelevelresponse
type Forecastservicelevelresponse struct { 
	// Percent - The percent of calls to answer in the number of seconds defined
	Percent *int32 `json:"percent,omitempty"`


	// Seconds - The number of seconds to define for the percent of calls to be answered
	Seconds *int32 `json:"seconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Forecastservicelevelresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

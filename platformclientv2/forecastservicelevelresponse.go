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

func (o *Forecastservicelevelresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Forecastservicelevelresponse
	
	return json.Marshal(&struct { 
		Percent *int `json:"percent,omitempty"`
		
		Seconds *int `json:"seconds,omitempty"`
		*Alias
	}{ 
		Percent: o.Percent,
		
		Seconds: o.Seconds,
		Alias:    (*Alias)(o),
	})
}

func (o *Forecastservicelevelresponse) UnmarshalJSON(b []byte) error {
	var ForecastservicelevelresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ForecastservicelevelresponseMap)
	if err != nil {
		return err
	}
	
	if Percent, ok := ForecastservicelevelresponseMap["percent"].(float64); ok {
		PercentInt := int(Percent)
		o.Percent = &PercentInt
	}
	
	if Seconds, ok := ForecastservicelevelresponseMap["seconds"].(float64); ok {
		SecondsInt := int(Seconds)
		o.Seconds = &SecondsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Forecastservicelevelresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Forecastaveragespeedofanswerresponse
type Forecastaveragespeedofanswerresponse struct { 
	// Seconds - the average speed of answer goal in seconds
	Seconds *int `json:"seconds,omitempty"`

}

func (o *Forecastaveragespeedofanswerresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Forecastaveragespeedofanswerresponse
	
	return json.Marshal(&struct { 
		Seconds *int `json:"seconds,omitempty"`
		*Alias
	}{ 
		Seconds: o.Seconds,
		Alias:    (*Alias)(o),
	})
}

func (o *Forecastaveragespeedofanswerresponse) UnmarshalJSON(b []byte) error {
	var ForecastaveragespeedofanswerresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ForecastaveragespeedofanswerresponseMap)
	if err != nil {
		return err
	}
	
	if Seconds, ok := ForecastaveragespeedofanswerresponseMap["seconds"].(float64); ok {
		SecondsInt := int(Seconds)
		o.Seconds = &SecondsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Forecastaveragespeedofanswerresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

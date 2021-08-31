package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Forecastabandonrateresponse
type Forecastabandonrateresponse struct { 
	// Percent - The target percent abandon rate goal
	Percent *int `json:"percent,omitempty"`

}

func (o *Forecastabandonrateresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Forecastabandonrateresponse
	
	return json.Marshal(&struct { 
		Percent *int `json:"percent,omitempty"`
		*Alias
	}{ 
		Percent: o.Percent,
		Alias:    (*Alias)(o),
	})
}

func (o *Forecastabandonrateresponse) UnmarshalJSON(b []byte) error {
	var ForecastabandonrateresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ForecastabandonrateresponseMap)
	if err != nil {
		return err
	}
	
	if Percent, ok := ForecastabandonrateresponseMap["percent"].(float64); ok {
		PercentInt := int(Percent)
		o.Percent = &PercentInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Forecastabandonrateresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

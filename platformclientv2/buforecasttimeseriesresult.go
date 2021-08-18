package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buforecasttimeseriesresult
type Buforecasttimeseriesresult struct { 
	// Metric - The metric this result applies to
	Metric *string `json:"metric,omitempty"`


	// ForecastingMethod - The forecasting method that was used for this metric
	ForecastingMethod *string `json:"forecastingMethod,omitempty"`


	// ForecastType - The forecasting type in this forecast result
	ForecastType *string `json:"forecastType,omitempty"`

}

func (u *Buforecasttimeseriesresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buforecasttimeseriesresult

	

	return json.Marshal(&struct { 
		Metric *string `json:"metric,omitempty"`
		
		ForecastingMethod *string `json:"forecastingMethod,omitempty"`
		
		ForecastType *string `json:"forecastType,omitempty"`
		*Alias
	}{ 
		Metric: u.Metric,
		
		ForecastingMethod: u.ForecastingMethod,
		
		ForecastType: u.ForecastType,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Buforecasttimeseriesresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

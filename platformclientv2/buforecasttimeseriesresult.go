package platformclientv2
import (
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

}

// String returns a JSON representation of the model
func (o *Buforecasttimeseriesresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

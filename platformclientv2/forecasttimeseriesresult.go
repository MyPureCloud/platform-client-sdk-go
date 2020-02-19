package platformclientv2
import (
	"encoding/json"
)

// Forecasttimeseriesresult
type Forecasttimeseriesresult struct { 
	// Metric - The metric this result applies to
	Metric *string `json:"metric,omitempty"`


	// ForecastingMethod - The forecasting method that was used for this metric
	ForecastingMethod *string `json:"forecastingMethod,omitempty"`

}

// String returns a JSON representation of the model
func (o *Forecasttimeseriesresult) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

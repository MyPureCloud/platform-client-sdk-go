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

func (o *Buforecasttimeseriesresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buforecasttimeseriesresult
	
	return json.Marshal(&struct { 
		Metric *string `json:"metric,omitempty"`
		
		ForecastingMethod *string `json:"forecastingMethod,omitempty"`
		
		ForecastType *string `json:"forecastType,omitempty"`
		*Alias
	}{ 
		Metric: o.Metric,
		
		ForecastingMethod: o.ForecastingMethod,
		
		ForecastType: o.ForecastType,
		Alias:    (*Alias)(o),
	})
}

func (o *Buforecasttimeseriesresult) UnmarshalJSON(b []byte) error {
	var BuforecasttimeseriesresultMap map[string]interface{}
	err := json.Unmarshal(b, &BuforecasttimeseriesresultMap)
	if err != nil {
		return err
	}
	
	if Metric, ok := BuforecasttimeseriesresultMap["metric"].(string); ok {
		o.Metric = &Metric
	}
    
	if ForecastingMethod, ok := BuforecasttimeseriesresultMap["forecastingMethod"].(string); ok {
		o.ForecastingMethod = &ForecastingMethod
	}
    
	if ForecastType, ok := BuforecasttimeseriesresultMap["forecastType"].(string); ok {
		o.ForecastType = &ForecastType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Buforecasttimeseriesresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

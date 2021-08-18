package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buintradaydatagroup
type Buintradaydatagroup struct { 
	// MediaType - The media type associated with this intraday group
	MediaType *string `json:"mediaType,omitempty"`


	// ForecastDataSummary - Forecast data summary for this date range
	ForecastDataSummary *Buintradayforecastdata `json:"forecastDataSummary,omitempty"`


	// ForecastDataPerInterval - Forecast data per interval for this date range
	ForecastDataPerInterval *[]Buintradayforecastdata `json:"forecastDataPerInterval,omitempty"`


	// ScheduleDataSummary - Schedule data summary for this date range
	ScheduleDataSummary *Buintradayscheduledata `json:"scheduleDataSummary,omitempty"`


	// ScheduleDataPerInterval - Schedule data per interval for this date range
	ScheduleDataPerInterval *[]Buintradayscheduledata `json:"scheduleDataPerInterval,omitempty"`


	// PerformancePredictionDataSummary - Performance prediction data summary for this date range
	PerformancePredictionDataSummary *Intradayperformancepredictiondata `json:"performancePredictionDataSummary,omitempty"`


	// PerformancePredictionDataPerInterval - Performance prediction data per interval for this date range
	PerformancePredictionDataPerInterval *[]Intradayperformancepredictiondata `json:"performancePredictionDataPerInterval,omitempty"`

}

func (u *Buintradaydatagroup) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buintradaydatagroup

	

	return json.Marshal(&struct { 
		MediaType *string `json:"mediaType,omitempty"`
		
		ForecastDataSummary *Buintradayforecastdata `json:"forecastDataSummary,omitempty"`
		
		ForecastDataPerInterval *[]Buintradayforecastdata `json:"forecastDataPerInterval,omitempty"`
		
		ScheduleDataSummary *Buintradayscheduledata `json:"scheduleDataSummary,omitempty"`
		
		ScheduleDataPerInterval *[]Buintradayscheduledata `json:"scheduleDataPerInterval,omitempty"`
		
		PerformancePredictionDataSummary *Intradayperformancepredictiondata `json:"performancePredictionDataSummary,omitempty"`
		
		PerformancePredictionDataPerInterval *[]Intradayperformancepredictiondata `json:"performancePredictionDataPerInterval,omitempty"`
		*Alias
	}{ 
		MediaType: u.MediaType,
		
		ForecastDataSummary: u.ForecastDataSummary,
		
		ForecastDataPerInterval: u.ForecastDataPerInterval,
		
		ScheduleDataSummary: u.ScheduleDataSummary,
		
		ScheduleDataPerInterval: u.ScheduleDataPerInterval,
		
		PerformancePredictionDataSummary: u.PerformancePredictionDataSummary,
		
		PerformancePredictionDataPerInterval: u.PerformancePredictionDataPerInterval,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Buintradaydatagroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

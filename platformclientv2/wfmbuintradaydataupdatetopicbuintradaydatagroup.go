package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuintradaydataupdatetopicbuintradaydatagroup
type Wfmbuintradaydataupdatetopicbuintradaydatagroup struct { 
	// MediaType
	MediaType *string `json:"mediaType,omitempty"`


	// ForecastDataSummary
	ForecastDataSummary *Wfmbuintradaydataupdatetopicbuintradayforecastdata `json:"forecastDataSummary,omitempty"`


	// ForecastDataPerInterval
	ForecastDataPerInterval *[]Wfmbuintradaydataupdatetopicbuintradayforecastdata `json:"forecastDataPerInterval,omitempty"`


	// ScheduleDataSummary
	ScheduleDataSummary *Wfmbuintradaydataupdatetopicbuintradayscheduledata `json:"scheduleDataSummary,omitempty"`


	// ScheduleDataPerInterval
	ScheduleDataPerInterval *[]Wfmbuintradaydataupdatetopicbuintradayscheduledata `json:"scheduleDataPerInterval,omitempty"`


	// PerformancePredictionDataSummary
	PerformancePredictionDataSummary *Wfmbuintradaydataupdatetopicintradayperformancepredictiondata `json:"performancePredictionDataSummary,omitempty"`


	// PerformancePredictionDataPerInterval
	PerformancePredictionDataPerInterval *[]Wfmbuintradaydataupdatetopicintradayperformancepredictiondata `json:"performancePredictionDataPerInterval,omitempty"`

}

func (u *Wfmbuintradaydataupdatetopicbuintradaydatagroup) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuintradaydataupdatetopicbuintradaydatagroup

	

	return json.Marshal(&struct { 
		MediaType *string `json:"mediaType,omitempty"`
		
		ForecastDataSummary *Wfmbuintradaydataupdatetopicbuintradayforecastdata `json:"forecastDataSummary,omitempty"`
		
		ForecastDataPerInterval *[]Wfmbuintradaydataupdatetopicbuintradayforecastdata `json:"forecastDataPerInterval,omitempty"`
		
		ScheduleDataSummary *Wfmbuintradaydataupdatetopicbuintradayscheduledata `json:"scheduleDataSummary,omitempty"`
		
		ScheduleDataPerInterval *[]Wfmbuintradaydataupdatetopicbuintradayscheduledata `json:"scheduleDataPerInterval,omitempty"`
		
		PerformancePredictionDataSummary *Wfmbuintradaydataupdatetopicintradayperformancepredictiondata `json:"performancePredictionDataSummary,omitempty"`
		
		PerformancePredictionDataPerInterval *[]Wfmbuintradaydataupdatetopicintradayperformancepredictiondata `json:"performancePredictionDataPerInterval,omitempty"`
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
func (o *Wfmbuintradaydataupdatetopicbuintradaydatagroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

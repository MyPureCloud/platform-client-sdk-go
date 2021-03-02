package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Wfmbuintradaydataupdatetopicbuintradaydatagroup) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

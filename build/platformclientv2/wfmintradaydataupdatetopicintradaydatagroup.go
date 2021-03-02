package platformclientv2
import (
	"encoding/json"
)

// Wfmintradaydataupdatetopicintradaydatagroup
type Wfmintradaydataupdatetopicintradaydatagroup struct { 
	// MediaType
	MediaType *string `json:"mediaType,omitempty"`


	// ForecastDataPerInterval
	ForecastDataPerInterval *[]Wfmintradaydataupdatetopicintradayforecastdata `json:"forecastDataPerInterval,omitempty"`


	// ScheduleDataPerInterval
	ScheduleDataPerInterval *[]Wfmintradaydataupdatetopicintradayscheduledata `json:"scheduleDataPerInterval,omitempty"`


	// HistoricalAgentDataPerInterval
	HistoricalAgentDataPerInterval *[]Wfmintradaydataupdatetopicintradayhistoricalagentdata `json:"historicalAgentDataPerInterval,omitempty"`


	// HistoricalQueueDataPerInterval
	HistoricalQueueDataPerInterval *[]Wfmintradaydataupdatetopicintradayhistoricalqueuedata `json:"historicalQueueDataPerInterval,omitempty"`


	// PerformancePredictionAgentDataPerInterval
	PerformancePredictionAgentDataPerInterval *[]Wfmintradaydataupdatetopicintradayperformancepredictionagentdata `json:"performancePredictionAgentDataPerInterval,omitempty"`


	// PerformancePredictionQueueDataPerInterval
	PerformancePredictionQueueDataPerInterval *[]Wfmintradaydataupdatetopicintradayperformancepredictionqueuedata `json:"performancePredictionQueueDataPerInterval,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradaydatagroup) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

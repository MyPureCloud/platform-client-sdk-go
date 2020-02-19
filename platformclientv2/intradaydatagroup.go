package platformclientv2
import (
	"encoding/json"
)

// Intradaydatagroup
type Intradaydatagroup struct { 
	// MediaType - The media type associated with this intraday group
	MediaType *string `json:"mediaType,omitempty"`


	// ForecastDataPerInterval - Forecast data for this date range
	ForecastDataPerInterval *[]Intradayforecastdata `json:"forecastDataPerInterval,omitempty"`


	// ScheduleDataPerInterval - Schedule data for this date range
	ScheduleDataPerInterval *[]Intradayscheduledata `json:"scheduleDataPerInterval,omitempty"`


	// HistoricalAgentDataPerInterval - Historical agent data for this date range
	HistoricalAgentDataPerInterval *[]Intradayhistoricalagentdata `json:"historicalAgentDataPerInterval,omitempty"`


	// HistoricalQueueDataPerInterval - Historical queue data for this date range
	HistoricalQueueDataPerInterval *[]Intradayhistoricalqueuedata `json:"historicalQueueDataPerInterval,omitempty"`


	// PerformancePredictionAgentDataPerInterval - Performance prediction data for this date range
	PerformancePredictionAgentDataPerInterval *[]Intradayperformancepredictionagentdata `json:"performancePredictionAgentDataPerInterval,omitempty"`


	// PerformancePredictionQueueDataPerInterval - Performance prediction data for this date range
	PerformancePredictionQueueDataPerInterval *[]Intradayperformancepredictionqueuedata `json:"performancePredictionQueueDataPerInterval,omitempty"`

}

// String returns a JSON representation of the model
func (o *Intradaydatagroup) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

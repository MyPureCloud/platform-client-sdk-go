package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Wfmintradaydataupdatetopicintradaydatagroup) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmintradaydataupdatetopicintradaydatagroup

	

	return json.Marshal(&struct { 
		MediaType *string `json:"mediaType,omitempty"`
		
		ForecastDataPerInterval *[]Wfmintradaydataupdatetopicintradayforecastdata `json:"forecastDataPerInterval,omitempty"`
		
		ScheduleDataPerInterval *[]Wfmintradaydataupdatetopicintradayscheduledata `json:"scheduleDataPerInterval,omitempty"`
		
		HistoricalAgentDataPerInterval *[]Wfmintradaydataupdatetopicintradayhistoricalagentdata `json:"historicalAgentDataPerInterval,omitempty"`
		
		HistoricalQueueDataPerInterval *[]Wfmintradaydataupdatetopicintradayhistoricalqueuedata `json:"historicalQueueDataPerInterval,omitempty"`
		
		PerformancePredictionAgentDataPerInterval *[]Wfmintradaydataupdatetopicintradayperformancepredictionagentdata `json:"performancePredictionAgentDataPerInterval,omitempty"`
		
		PerformancePredictionQueueDataPerInterval *[]Wfmintradaydataupdatetopicintradayperformancepredictionqueuedata `json:"performancePredictionQueueDataPerInterval,omitempty"`
		*Alias
	}{ 
		MediaType: u.MediaType,
		
		ForecastDataPerInterval: u.ForecastDataPerInterval,
		
		ScheduleDataPerInterval: u.ScheduleDataPerInterval,
		
		HistoricalAgentDataPerInterval: u.HistoricalAgentDataPerInterval,
		
		HistoricalQueueDataPerInterval: u.HistoricalQueueDataPerInterval,
		
		PerformancePredictionAgentDataPerInterval: u.PerformancePredictionAgentDataPerInterval,
		
		PerformancePredictionQueueDataPerInterval: u.PerformancePredictionQueueDataPerInterval,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradaydatagroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

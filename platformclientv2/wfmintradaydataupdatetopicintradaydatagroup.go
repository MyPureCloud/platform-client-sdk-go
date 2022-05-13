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

func (o *Wfmintradaydataupdatetopicintradaydatagroup) MarshalJSON() ([]byte, error) {
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
		MediaType: o.MediaType,
		
		ForecastDataPerInterval: o.ForecastDataPerInterval,
		
		ScheduleDataPerInterval: o.ScheduleDataPerInterval,
		
		HistoricalAgentDataPerInterval: o.HistoricalAgentDataPerInterval,
		
		HistoricalQueueDataPerInterval: o.HistoricalQueueDataPerInterval,
		
		PerformancePredictionAgentDataPerInterval: o.PerformancePredictionAgentDataPerInterval,
		
		PerformancePredictionQueueDataPerInterval: o.PerformancePredictionQueueDataPerInterval,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmintradaydataupdatetopicintradaydatagroup) UnmarshalJSON(b []byte) error {
	var WfmintradaydataupdatetopicintradaydatagroupMap map[string]interface{}
	err := json.Unmarshal(b, &WfmintradaydataupdatetopicintradaydatagroupMap)
	if err != nil {
		return err
	}
	
	if MediaType, ok := WfmintradaydataupdatetopicintradaydatagroupMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if ForecastDataPerInterval, ok := WfmintradaydataupdatetopicintradaydatagroupMap["forecastDataPerInterval"].([]interface{}); ok {
		ForecastDataPerIntervalString, _ := json.Marshal(ForecastDataPerInterval)
		json.Unmarshal(ForecastDataPerIntervalString, &o.ForecastDataPerInterval)
	}
	
	if ScheduleDataPerInterval, ok := WfmintradaydataupdatetopicintradaydatagroupMap["scheduleDataPerInterval"].([]interface{}); ok {
		ScheduleDataPerIntervalString, _ := json.Marshal(ScheduleDataPerInterval)
		json.Unmarshal(ScheduleDataPerIntervalString, &o.ScheduleDataPerInterval)
	}
	
	if HistoricalAgentDataPerInterval, ok := WfmintradaydataupdatetopicintradaydatagroupMap["historicalAgentDataPerInterval"].([]interface{}); ok {
		HistoricalAgentDataPerIntervalString, _ := json.Marshal(HistoricalAgentDataPerInterval)
		json.Unmarshal(HistoricalAgentDataPerIntervalString, &o.HistoricalAgentDataPerInterval)
	}
	
	if HistoricalQueueDataPerInterval, ok := WfmintradaydataupdatetopicintradaydatagroupMap["historicalQueueDataPerInterval"].([]interface{}); ok {
		HistoricalQueueDataPerIntervalString, _ := json.Marshal(HistoricalQueueDataPerInterval)
		json.Unmarshal(HistoricalQueueDataPerIntervalString, &o.HistoricalQueueDataPerInterval)
	}
	
	if PerformancePredictionAgentDataPerInterval, ok := WfmintradaydataupdatetopicintradaydatagroupMap["performancePredictionAgentDataPerInterval"].([]interface{}); ok {
		PerformancePredictionAgentDataPerIntervalString, _ := json.Marshal(PerformancePredictionAgentDataPerInterval)
		json.Unmarshal(PerformancePredictionAgentDataPerIntervalString, &o.PerformancePredictionAgentDataPerInterval)
	}
	
	if PerformancePredictionQueueDataPerInterval, ok := WfmintradaydataupdatetopicintradaydatagroupMap["performancePredictionQueueDataPerInterval"].([]interface{}); ok {
		PerformancePredictionQueueDataPerIntervalString, _ := json.Marshal(PerformancePredictionQueueDataPerInterval)
		json.Unmarshal(PerformancePredictionQueueDataPerIntervalString, &o.PerformancePredictionQueueDataPerInterval)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradaydatagroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

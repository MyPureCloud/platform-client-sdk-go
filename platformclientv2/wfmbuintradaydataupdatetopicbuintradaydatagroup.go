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

func (o *Wfmbuintradaydataupdatetopicbuintradaydatagroup) MarshalJSON() ([]byte, error) {
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
		MediaType: o.MediaType,
		
		ForecastDataSummary: o.ForecastDataSummary,
		
		ForecastDataPerInterval: o.ForecastDataPerInterval,
		
		ScheduleDataSummary: o.ScheduleDataSummary,
		
		ScheduleDataPerInterval: o.ScheduleDataPerInterval,
		
		PerformancePredictionDataSummary: o.PerformancePredictionDataSummary,
		
		PerformancePredictionDataPerInterval: o.PerformancePredictionDataPerInterval,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmbuintradaydataupdatetopicbuintradaydatagroup) UnmarshalJSON(b []byte) error {
	var WfmbuintradaydataupdatetopicbuintradaydatagroupMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbuintradaydataupdatetopicbuintradaydatagroupMap)
	if err != nil {
		return err
	}
	
	if MediaType, ok := WfmbuintradaydataupdatetopicbuintradaydatagroupMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if ForecastDataSummary, ok := WfmbuintradaydataupdatetopicbuintradaydatagroupMap["forecastDataSummary"].(map[string]interface{}); ok {
		ForecastDataSummaryString, _ := json.Marshal(ForecastDataSummary)
		json.Unmarshal(ForecastDataSummaryString, &o.ForecastDataSummary)
	}
	
	if ForecastDataPerInterval, ok := WfmbuintradaydataupdatetopicbuintradaydatagroupMap["forecastDataPerInterval"].([]interface{}); ok {
		ForecastDataPerIntervalString, _ := json.Marshal(ForecastDataPerInterval)
		json.Unmarshal(ForecastDataPerIntervalString, &o.ForecastDataPerInterval)
	}
	
	if ScheduleDataSummary, ok := WfmbuintradaydataupdatetopicbuintradaydatagroupMap["scheduleDataSummary"].(map[string]interface{}); ok {
		ScheduleDataSummaryString, _ := json.Marshal(ScheduleDataSummary)
		json.Unmarshal(ScheduleDataSummaryString, &o.ScheduleDataSummary)
	}
	
	if ScheduleDataPerInterval, ok := WfmbuintradaydataupdatetopicbuintradaydatagroupMap["scheduleDataPerInterval"].([]interface{}); ok {
		ScheduleDataPerIntervalString, _ := json.Marshal(ScheduleDataPerInterval)
		json.Unmarshal(ScheduleDataPerIntervalString, &o.ScheduleDataPerInterval)
	}
	
	if PerformancePredictionDataSummary, ok := WfmbuintradaydataupdatetopicbuintradaydatagroupMap["performancePredictionDataSummary"].(map[string]interface{}); ok {
		PerformancePredictionDataSummaryString, _ := json.Marshal(PerformancePredictionDataSummary)
		json.Unmarshal(PerformancePredictionDataSummaryString, &o.PerformancePredictionDataSummary)
	}
	
	if PerformancePredictionDataPerInterval, ok := WfmbuintradaydataupdatetopicbuintradaydatagroupMap["performancePredictionDataPerInterval"].([]interface{}); ok {
		PerformancePredictionDataPerIntervalString, _ := json.Marshal(PerformancePredictionDataPerInterval)
		json.Unmarshal(PerformancePredictionDataPerIntervalString, &o.PerformancePredictionDataPerInterval)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbuintradaydataupdatetopicbuintradaydatagroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

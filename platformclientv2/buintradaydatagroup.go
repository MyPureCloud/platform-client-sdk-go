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

func (o *Buintradaydatagroup) MarshalJSON() ([]byte, error) {
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

func (o *Buintradaydatagroup) UnmarshalJSON(b []byte) error {
	var BuintradaydatagroupMap map[string]interface{}
	err := json.Unmarshal(b, &BuintradaydatagroupMap)
	if err != nil {
		return err
	}
	
	if MediaType, ok := BuintradaydatagroupMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
	
	if ForecastDataSummary, ok := BuintradaydatagroupMap["forecastDataSummary"].(map[string]interface{}); ok {
		ForecastDataSummaryString, _ := json.Marshal(ForecastDataSummary)
		json.Unmarshal(ForecastDataSummaryString, &o.ForecastDataSummary)
	}
	
	if ForecastDataPerInterval, ok := BuintradaydatagroupMap["forecastDataPerInterval"].([]interface{}); ok {
		ForecastDataPerIntervalString, _ := json.Marshal(ForecastDataPerInterval)
		json.Unmarshal(ForecastDataPerIntervalString, &o.ForecastDataPerInterval)
	}
	
	if ScheduleDataSummary, ok := BuintradaydatagroupMap["scheduleDataSummary"].(map[string]interface{}); ok {
		ScheduleDataSummaryString, _ := json.Marshal(ScheduleDataSummary)
		json.Unmarshal(ScheduleDataSummaryString, &o.ScheduleDataSummary)
	}
	
	if ScheduleDataPerInterval, ok := BuintradaydatagroupMap["scheduleDataPerInterval"].([]interface{}); ok {
		ScheduleDataPerIntervalString, _ := json.Marshal(ScheduleDataPerInterval)
		json.Unmarshal(ScheduleDataPerIntervalString, &o.ScheduleDataPerInterval)
	}
	
	if PerformancePredictionDataSummary, ok := BuintradaydatagroupMap["performancePredictionDataSummary"].(map[string]interface{}); ok {
		PerformancePredictionDataSummaryString, _ := json.Marshal(PerformancePredictionDataSummary)
		json.Unmarshal(PerformancePredictionDataSummaryString, &o.PerformancePredictionDataSummary)
	}
	
	if PerformancePredictionDataPerInterval, ok := BuintradaydatagroupMap["performancePredictionDataPerInterval"].([]interface{}); ok {
		PerformancePredictionDataPerIntervalString, _ := json.Marshal(PerformancePredictionDataPerInterval)
		json.Unmarshal(PerformancePredictionDataPerIntervalString, &o.PerformancePredictionDataPerInterval)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buintradaydatagroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

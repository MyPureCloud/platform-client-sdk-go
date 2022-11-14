package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Historicaladherencequeryresult
type Historicaladherencequeryresult struct { 
	// UserId - The ID of the user for whom the adherence is queried
	UserId *string `json:"userId,omitempty"`


	// ManagementUnitId - The ID of the management unit of the user for whom the adherence is queried
	ManagementUnitId *string `json:"managementUnitId,omitempty"`


	// StartDate - Beginning of the date range that was queried, in ISO-8601 format
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - End of the date range that was queried, in ISO-8601 format. If it was not set, end date will be set to the queried time
	EndDate *time.Time `json:"endDate,omitempty"`


	// AdherencePercentage - Adherence percentage for this user, in the scale of 0 - 100
	AdherencePercentage *float64 `json:"adherencePercentage,omitempty"`


	// ConformancePercentage - Conformance percentage for this user, in the scale of 0 - 100. Conformance percentage can be greater than 100 when the actual on queue time is greater than the scheduled on queue time for the same period.
	ConformancePercentage *float64 `json:"conformancePercentage,omitempty"`


	// Impact - The impact of the current adherence state for this user
	Impact *string `json:"impact,omitempty"`


	// ExceptionInfo - List of adherence exceptions for this user
	ExceptionInfo *[]Historicaladherenceexceptioninfo `json:"exceptionInfo,omitempty"`


	// DayMetrics - Adherence and conformance metrics for days in query range
	DayMetrics *[]Historicaladherencedaymetrics `json:"dayMetrics,omitempty"`


	// ActualsEndDate - The end date of the actual activities in ISO-8601 format.
	ActualsEndDate *time.Time `json:"actualsEndDate,omitempty"`


	// Actuals - List of actual activity with offset for this user
	Actuals *[]Historicaladherenceactuals `json:"actuals,omitempty"`

}

func (o *Historicaladherencequeryresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Historicaladherencequeryresult
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if o.EndDate != nil {
		
		*EndDate = timeutil.Strftime(o.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndDate = nil
	}
	
	ActualsEndDate := new(string)
	if o.ActualsEndDate != nil {
		
		*ActualsEndDate = timeutil.Strftime(o.ActualsEndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ActualsEndDate = nil
	}
	
	return json.Marshal(&struct { 
		UserId *string `json:"userId,omitempty"`
		
		ManagementUnitId *string `json:"managementUnitId,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		AdherencePercentage *float64 `json:"adherencePercentage,omitempty"`
		
		ConformancePercentage *float64 `json:"conformancePercentage,omitempty"`
		
		Impact *string `json:"impact,omitempty"`
		
		ExceptionInfo *[]Historicaladherenceexceptioninfo `json:"exceptionInfo,omitempty"`
		
		DayMetrics *[]Historicaladherencedaymetrics `json:"dayMetrics,omitempty"`
		
		ActualsEndDate *string `json:"actualsEndDate,omitempty"`
		
		Actuals *[]Historicaladherenceactuals `json:"actuals,omitempty"`
		*Alias
	}{ 
		UserId: o.UserId,
		
		ManagementUnitId: o.ManagementUnitId,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		AdherencePercentage: o.AdherencePercentage,
		
		ConformancePercentage: o.ConformancePercentage,
		
		Impact: o.Impact,
		
		ExceptionInfo: o.ExceptionInfo,
		
		DayMetrics: o.DayMetrics,
		
		ActualsEndDate: ActualsEndDate,
		
		Actuals: o.Actuals,
		Alias:    (*Alias)(o),
	})
}

func (o *Historicaladherencequeryresult) UnmarshalJSON(b []byte) error {
	var HistoricaladherencequeryresultMap map[string]interface{}
	err := json.Unmarshal(b, &HistoricaladherencequeryresultMap)
	if err != nil {
		return err
	}
	
	if UserId, ok := HistoricaladherencequeryresultMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if ManagementUnitId, ok := HistoricaladherencequeryresultMap["managementUnitId"].(string); ok {
		o.ManagementUnitId = &ManagementUnitId
	}
    
	if startDateString, ok := HistoricaladherencequeryresultMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := HistoricaladherencequeryresultMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if AdherencePercentage, ok := HistoricaladherencequeryresultMap["adherencePercentage"].(float64); ok {
		o.AdherencePercentage = &AdherencePercentage
	}
    
	if ConformancePercentage, ok := HistoricaladherencequeryresultMap["conformancePercentage"].(float64); ok {
		o.ConformancePercentage = &ConformancePercentage
	}
    
	if Impact, ok := HistoricaladherencequeryresultMap["impact"].(string); ok {
		o.Impact = &Impact
	}
    
	if ExceptionInfo, ok := HistoricaladherencequeryresultMap["exceptionInfo"].([]interface{}); ok {
		ExceptionInfoString, _ := json.Marshal(ExceptionInfo)
		json.Unmarshal(ExceptionInfoString, &o.ExceptionInfo)
	}
	
	if DayMetrics, ok := HistoricaladherencequeryresultMap["dayMetrics"].([]interface{}); ok {
		DayMetricsString, _ := json.Marshal(DayMetrics)
		json.Unmarshal(DayMetricsString, &o.DayMetrics)
	}
	
	if actualsEndDateString, ok := HistoricaladherencequeryresultMap["actualsEndDate"].(string); ok {
		ActualsEndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", actualsEndDateString)
		o.ActualsEndDate = &ActualsEndDate
	}
	
	if Actuals, ok := HistoricaladherencequeryresultMap["actuals"].([]interface{}); ok {
		ActualsString, _ := json.Marshal(Actuals)
		json.Unmarshal(ActualsString, &o.Actuals)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Historicaladherencequeryresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

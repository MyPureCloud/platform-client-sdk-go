package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmhistoricaladherencebulkuserdaymetrics
type Wfmhistoricaladherencebulkuserdaymetrics struct { 
	// DayStartOffsetSeconds - Start of day offset in seconds relative to query start time
	DayStartOffsetSeconds *int `json:"dayStartOffsetSeconds,omitempty"`


	// AdherenceScheduleSeconds - Duration of schedule in seconds included for adherence percentage calculation
	AdherenceScheduleSeconds *int `json:"adherenceScheduleSeconds,omitempty"`


	// ConformanceScheduleSeconds - Total scheduled duration in seconds for OnQueue activities
	ConformanceScheduleSeconds *int `json:"conformanceScheduleSeconds,omitempty"`


	// ConformanceActualSeconds - Total actually worked duration in seconds for OnQueue activities
	ConformanceActualSeconds *int `json:"conformanceActualSeconds,omitempty"`


	// ExceptionCount - Total number of adherence exceptions for this user
	ExceptionCount *int `json:"exceptionCount,omitempty"`


	// ExceptionDurationSeconds - Total duration in seconds of adherence exceptions for this user
	ExceptionDurationSeconds *int `json:"exceptionDurationSeconds,omitempty"`


	// ImpactSeconds - The impact duration in seconds of current adherence state for this user
	ImpactSeconds *int `json:"impactSeconds,omitempty"`


	// ScheduleLengthSeconds - Total duration in seconds for all scheduled activities
	ScheduleLengthSeconds *int `json:"scheduleLengthSeconds,omitempty"`


	// ActualLengthSeconds - Total duration in seconds for all actually worked activities
	ActualLengthSeconds *int `json:"actualLengthSeconds,omitempty"`


	// AdherencePercentage - Total adherence percentage for this user, in the scale of 0 - 100
	AdherencePercentage *float64 `json:"adherencePercentage,omitempty"`


	// ConformancePercentage - Total conformance percentage for this user, in the scale of 0 - 100. Conformance percentage can be greater than 100 when the actual on queue time is greater than the scheduled on queue time for the same period.
	ConformancePercentage *float64 `json:"conformancePercentage,omitempty"`

}

func (o *Wfmhistoricaladherencebulkuserdaymetrics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmhistoricaladherencebulkuserdaymetrics
	
	return json.Marshal(&struct { 
		DayStartOffsetSeconds *int `json:"dayStartOffsetSeconds,omitempty"`
		
		AdherenceScheduleSeconds *int `json:"adherenceScheduleSeconds,omitempty"`
		
		ConformanceScheduleSeconds *int `json:"conformanceScheduleSeconds,omitempty"`
		
		ConformanceActualSeconds *int `json:"conformanceActualSeconds,omitempty"`
		
		ExceptionCount *int `json:"exceptionCount,omitempty"`
		
		ExceptionDurationSeconds *int `json:"exceptionDurationSeconds,omitempty"`
		
		ImpactSeconds *int `json:"impactSeconds,omitempty"`
		
		ScheduleLengthSeconds *int `json:"scheduleLengthSeconds,omitempty"`
		
		ActualLengthSeconds *int `json:"actualLengthSeconds,omitempty"`
		
		AdherencePercentage *float64 `json:"adherencePercentage,omitempty"`
		
		ConformancePercentage *float64 `json:"conformancePercentage,omitempty"`
		*Alias
	}{ 
		DayStartOffsetSeconds: o.DayStartOffsetSeconds,
		
		AdherenceScheduleSeconds: o.AdherenceScheduleSeconds,
		
		ConformanceScheduleSeconds: o.ConformanceScheduleSeconds,
		
		ConformanceActualSeconds: o.ConformanceActualSeconds,
		
		ExceptionCount: o.ExceptionCount,
		
		ExceptionDurationSeconds: o.ExceptionDurationSeconds,
		
		ImpactSeconds: o.ImpactSeconds,
		
		ScheduleLengthSeconds: o.ScheduleLengthSeconds,
		
		ActualLengthSeconds: o.ActualLengthSeconds,
		
		AdherencePercentage: o.AdherencePercentage,
		
		ConformancePercentage: o.ConformancePercentage,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmhistoricaladherencebulkuserdaymetrics) UnmarshalJSON(b []byte) error {
	var WfmhistoricaladherencebulkuserdaymetricsMap map[string]interface{}
	err := json.Unmarshal(b, &WfmhistoricaladherencebulkuserdaymetricsMap)
	if err != nil {
		return err
	}
	
	if DayStartOffsetSeconds, ok := WfmhistoricaladherencebulkuserdaymetricsMap["dayStartOffsetSeconds"].(float64); ok {
		DayStartOffsetSecondsInt := int(DayStartOffsetSeconds)
		o.DayStartOffsetSeconds = &DayStartOffsetSecondsInt
	}
	
	if AdherenceScheduleSeconds, ok := WfmhistoricaladherencebulkuserdaymetricsMap["adherenceScheduleSeconds"].(float64); ok {
		AdherenceScheduleSecondsInt := int(AdherenceScheduleSeconds)
		o.AdherenceScheduleSeconds = &AdherenceScheduleSecondsInt
	}
	
	if ConformanceScheduleSeconds, ok := WfmhistoricaladherencebulkuserdaymetricsMap["conformanceScheduleSeconds"].(float64); ok {
		ConformanceScheduleSecondsInt := int(ConformanceScheduleSeconds)
		o.ConformanceScheduleSeconds = &ConformanceScheduleSecondsInt
	}
	
	if ConformanceActualSeconds, ok := WfmhistoricaladherencebulkuserdaymetricsMap["conformanceActualSeconds"].(float64); ok {
		ConformanceActualSecondsInt := int(ConformanceActualSeconds)
		o.ConformanceActualSeconds = &ConformanceActualSecondsInt
	}
	
	if ExceptionCount, ok := WfmhistoricaladherencebulkuserdaymetricsMap["exceptionCount"].(float64); ok {
		ExceptionCountInt := int(ExceptionCount)
		o.ExceptionCount = &ExceptionCountInt
	}
	
	if ExceptionDurationSeconds, ok := WfmhistoricaladherencebulkuserdaymetricsMap["exceptionDurationSeconds"].(float64); ok {
		ExceptionDurationSecondsInt := int(ExceptionDurationSeconds)
		o.ExceptionDurationSeconds = &ExceptionDurationSecondsInt
	}
	
	if ImpactSeconds, ok := WfmhistoricaladherencebulkuserdaymetricsMap["impactSeconds"].(float64); ok {
		ImpactSecondsInt := int(ImpactSeconds)
		o.ImpactSeconds = &ImpactSecondsInt
	}
	
	if ScheduleLengthSeconds, ok := WfmhistoricaladherencebulkuserdaymetricsMap["scheduleLengthSeconds"].(float64); ok {
		ScheduleLengthSecondsInt := int(ScheduleLengthSeconds)
		o.ScheduleLengthSeconds = &ScheduleLengthSecondsInt
	}
	
	if ActualLengthSeconds, ok := WfmhistoricaladherencebulkuserdaymetricsMap["actualLengthSeconds"].(float64); ok {
		ActualLengthSecondsInt := int(ActualLengthSeconds)
		o.ActualLengthSeconds = &ActualLengthSecondsInt
	}
	
	if AdherencePercentage, ok := WfmhistoricaladherencebulkuserdaymetricsMap["adherencePercentage"].(float64); ok {
		o.AdherencePercentage = &AdherencePercentage
	}
    
	if ConformancePercentage, ok := WfmhistoricaladherencebulkuserdaymetricsMap["conformancePercentage"].(float64); ok {
		o.ConformancePercentage = &ConformancePercentage
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherencebulkuserdaymetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"encoding/json"
)

// Historicaladherencedaymetrics
type Historicaladherencedaymetrics struct { 
	// DayStartOffsetSecs - Start of day offset in seconds relative to query start time
	DayStartOffsetSecs *int32 `json:"dayStartOffsetSecs,omitempty"`


	// AdherenceScheduleSecs - Duration of schedule in seconds included for adherence percentage calculation
	AdherenceScheduleSecs *int32 `json:"adherenceScheduleSecs,omitempty"`


	// ConformanceScheduleSecs - Total scheduled duration in seconds for OnQueue activities
	ConformanceScheduleSecs *int32 `json:"conformanceScheduleSecs,omitempty"`


	// ConformanceActualSecs - Total actually worked duration in seconds for OnQueue activities
	ConformanceActualSecs *int32 `json:"conformanceActualSecs,omitempty"`


	// ExceptionCount - Total number of adherence exceptions for this user
	ExceptionCount *int32 `json:"exceptionCount,omitempty"`


	// ExceptionDurationSecs - Total duration in seconds of adherence exceptions for this user
	ExceptionDurationSecs *int32 `json:"exceptionDurationSecs,omitempty"`


	// ImpactSeconds - The impact duration in seconds of current adherence state for this user
	ImpactSeconds *int32 `json:"impactSeconds,omitempty"`


	// ScheduleLengthSecs - Total duration in seconds for all scheduled activities
	ScheduleLengthSecs *int32 `json:"scheduleLengthSecs,omitempty"`


	// ActualLengthSecs - Total duration in seconds for all actually worked activities
	ActualLengthSecs *int32 `json:"actualLengthSecs,omitempty"`


	// AdherencePercentage - Total adherence percentage for this user, in the scale of 0 - 100
	AdherencePercentage *float64 `json:"adherencePercentage,omitempty"`


	// ConformancePercentage - Total conformance percentage for this user, in the scale of 0 - 100. Conformance percentage can be greater than 100 when the actual on queue time is greater than the scheduled on queue time for the same period.
	ConformancePercentage *float64 `json:"conformancePercentage,omitempty"`

}

// String returns a JSON representation of the model
func (o *Historicaladherencedaymetrics) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

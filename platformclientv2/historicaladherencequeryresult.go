package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Historicaladherencequeryresult
type Historicaladherencequeryresult struct { 
	// UserId - The ID of the user for whom the adherence is queried
	UserId *string `json:"userId,omitempty"`


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


	// Actuals - List of actual activity with offset for this user
	Actuals *[]Historicaladherenceactuals `json:"actuals,omitempty"`

}

// String returns a JSON representation of the model
func (o *Historicaladherencequeryresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

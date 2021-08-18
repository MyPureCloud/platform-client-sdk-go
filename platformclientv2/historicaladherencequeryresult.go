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

func (u *Historicaladherencequeryresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Historicaladherencequeryresult

	
	StartDate := new(string)
	if u.StartDate != nil {
		
		*StartDate = timeutil.Strftime(u.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if u.EndDate != nil {
		
		*EndDate = timeutil.Strftime(u.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndDate = nil
	}
	

	return json.Marshal(&struct { 
		UserId *string `json:"userId,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		AdherencePercentage *float64 `json:"adherencePercentage,omitempty"`
		
		ConformancePercentage *float64 `json:"conformancePercentage,omitempty"`
		
		Impact *string `json:"impact,omitempty"`
		
		ExceptionInfo *[]Historicaladherenceexceptioninfo `json:"exceptionInfo,omitempty"`
		
		DayMetrics *[]Historicaladherencedaymetrics `json:"dayMetrics,omitempty"`
		
		Actuals *[]Historicaladherenceactuals `json:"actuals,omitempty"`
		*Alias
	}{ 
		UserId: u.UserId,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		AdherencePercentage: u.AdherencePercentage,
		
		ConformancePercentage: u.ConformancePercentage,
		
		Impact: u.Impact,
		
		ExceptionInfo: u.ExceptionInfo,
		
		DayMetrics: u.DayMetrics,
		
		Actuals: u.Actuals,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Historicaladherencequeryresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

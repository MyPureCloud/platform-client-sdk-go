package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmhistoricaladherencebulkuserresult
type Wfmhistoricaladherencebulkuserresult struct { 
	// UserId - The ID of the user for whom the adherence is queried
	UserId *string `json:"userId,omitempty"`


	// AdherencePercentage - Adherence percentage for this user, in the scale of 0 - 100
	AdherencePercentage *float64 `json:"adherencePercentage,omitempty"`


	// ConformancePercentage - Conformance percentage for this user, in the scale of 0 - 100. Conformance percentage can be greater than 100 when the actual on queue time is greater than the scheduled on queue time for the same period.
	ConformancePercentage *float64 `json:"conformancePercentage,omitempty"`


	// Impact - The impact of the current adherence state for this user
	Impact *string `json:"impact,omitempty"`


	// ExceptionInfo - List of adherence exceptions for this user
	ExceptionInfo *[]Historicaladherenceexceptioninfo `json:"exceptionInfo,omitempty"`


	// DayMetrics - Adherence and conformance metrics for days in query range
	DayMetrics *[]Wfmhistoricaladherencebulkuserdaymetrics `json:"dayMetrics,omitempty"`

}

func (o *Wfmhistoricaladherencebulkuserresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmhistoricaladherencebulkuserresult
	
	return json.Marshal(&struct { 
		UserId *string `json:"userId,omitempty"`
		
		AdherencePercentage *float64 `json:"adherencePercentage,omitempty"`
		
		ConformancePercentage *float64 `json:"conformancePercentage,omitempty"`
		
		Impact *string `json:"impact,omitempty"`
		
		ExceptionInfo *[]Historicaladherenceexceptioninfo `json:"exceptionInfo,omitempty"`
		
		DayMetrics *[]Wfmhistoricaladherencebulkuserdaymetrics `json:"dayMetrics,omitempty"`
		*Alias
	}{ 
		UserId: o.UserId,
		
		AdherencePercentage: o.AdherencePercentage,
		
		ConformancePercentage: o.ConformancePercentage,
		
		Impact: o.Impact,
		
		ExceptionInfo: o.ExceptionInfo,
		
		DayMetrics: o.DayMetrics,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmhistoricaladherencebulkuserresult) UnmarshalJSON(b []byte) error {
	var WfmhistoricaladherencebulkuserresultMap map[string]interface{}
	err := json.Unmarshal(b, &WfmhistoricaladherencebulkuserresultMap)
	if err != nil {
		return err
	}
	
	if UserId, ok := WfmhistoricaladherencebulkuserresultMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if AdherencePercentage, ok := WfmhistoricaladherencebulkuserresultMap["adherencePercentage"].(float64); ok {
		o.AdherencePercentage = &AdherencePercentage
	}
    
	if ConformancePercentage, ok := WfmhistoricaladherencebulkuserresultMap["conformancePercentage"].(float64); ok {
		o.ConformancePercentage = &ConformancePercentage
	}
    
	if Impact, ok := WfmhistoricaladherencebulkuserresultMap["impact"].(string); ok {
		o.Impact = &Impact
	}
    
	if ExceptionInfo, ok := WfmhistoricaladherencebulkuserresultMap["exceptionInfo"].([]interface{}); ok {
		ExceptionInfoString, _ := json.Marshal(ExceptionInfo)
		json.Unmarshal(ExceptionInfoString, &o.ExceptionInfo)
	}
	
	if DayMetrics, ok := WfmhistoricaladherencebulkuserresultMap["dayMetrics"].([]interface{}); ok {
		DayMetricsString, _ := json.Marshal(DayMetrics)
		json.Unmarshal(DayMetricsString, &o.DayMetrics)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherencebulkuserresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

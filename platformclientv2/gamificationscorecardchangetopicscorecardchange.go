package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Gamificationscorecardchangetopicscorecardchange
type Gamificationscorecardchangetopicscorecardchange struct { 
	// Workday
	Workday *string `json:"workday,omitempty"`


	// DivisionId
	DivisionId *string `json:"divisionId,omitempty"`


	// TeamId
	TeamId *string `json:"teamId,omitempty"`


	// PerformanceProfileId
	PerformanceProfileId *string `json:"performanceProfileId,omitempty"`


	// UserId
	UserId *string `json:"userId,omitempty"`


	// PerformanceMetrics
	PerformanceMetrics *[]Gamificationscorecardchangetopicperformancemetric `json:"performanceMetrics,omitempty"`

}

func (o *Gamificationscorecardchangetopicscorecardchange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Gamificationscorecardchangetopicscorecardchange
	
	return json.Marshal(&struct { 
		Workday *string `json:"workday,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		TeamId *string `json:"teamId,omitempty"`
		
		PerformanceProfileId *string `json:"performanceProfileId,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		PerformanceMetrics *[]Gamificationscorecardchangetopicperformancemetric `json:"performanceMetrics,omitempty"`
		*Alias
	}{ 
		Workday: o.Workday,
		
		DivisionId: o.DivisionId,
		
		TeamId: o.TeamId,
		
		PerformanceProfileId: o.PerformanceProfileId,
		
		UserId: o.UserId,
		
		PerformanceMetrics: o.PerformanceMetrics,
		Alias:    (*Alias)(o),
	})
}

func (o *Gamificationscorecardchangetopicscorecardchange) UnmarshalJSON(b []byte) error {
	var GamificationscorecardchangetopicscorecardchangeMap map[string]interface{}
	err := json.Unmarshal(b, &GamificationscorecardchangetopicscorecardchangeMap)
	if err != nil {
		return err
	}
	
	if Workday, ok := GamificationscorecardchangetopicscorecardchangeMap["workday"].(string); ok {
		o.Workday = &Workday
	}
	
	if DivisionId, ok := GamificationscorecardchangetopicscorecardchangeMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
	
	if TeamId, ok := GamificationscorecardchangetopicscorecardchangeMap["teamId"].(string); ok {
		o.TeamId = &TeamId
	}
	
	if PerformanceProfileId, ok := GamificationscorecardchangetopicscorecardchangeMap["performanceProfileId"].(string); ok {
		o.PerformanceProfileId = &PerformanceProfileId
	}
	
	if UserId, ok := GamificationscorecardchangetopicscorecardchangeMap["userId"].(string); ok {
		o.UserId = &UserId
	}
	
	if PerformanceMetrics, ok := GamificationscorecardchangetopicscorecardchangeMap["performanceMetrics"].([]interface{}); ok {
		PerformanceMetricsString, _ := json.Marshal(PerformanceMetrics)
		json.Unmarshal(PerformanceMetricsString, &o.PerformanceMetrics)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Gamificationscorecardchangetopicscorecardchange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

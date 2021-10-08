package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Singleworkdayaveragepoints
type Singleworkdayaveragepoints struct { 
	// DateWorkday - Queried target workday. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateWorkday *time.Time `json:"dateWorkday,omitempty"`


	// Division - The targeted division for the average points
	Division *Division `json:"division,omitempty"`


	// AveragePoints - The average points per agent earned within the division
	AveragePoints *float64 `json:"averagePoints,omitempty"`


	// PerformanceProfile - The targeted performance profile for the average points
	PerformanceProfile *Addressableentityref `json:"performanceProfile,omitempty"`

}

func (o *Singleworkdayaveragepoints) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Singleworkdayaveragepoints
	
	DateWorkday := new(string)
	if o.DateWorkday != nil {
		*DateWorkday = timeutil.Strftime(o.DateWorkday, "%Y-%m-%d")
	} else {
		DateWorkday = nil
	}
	
	return json.Marshal(&struct { 
		DateWorkday *string `json:"dateWorkday,omitempty"`
		
		Division *Division `json:"division,omitempty"`
		
		AveragePoints *float64 `json:"averagePoints,omitempty"`
		
		PerformanceProfile *Addressableentityref `json:"performanceProfile,omitempty"`
		*Alias
	}{ 
		DateWorkday: DateWorkday,
		
		Division: o.Division,
		
		AveragePoints: o.AveragePoints,
		
		PerformanceProfile: o.PerformanceProfile,
		Alias:    (*Alias)(o),
	})
}

func (o *Singleworkdayaveragepoints) UnmarshalJSON(b []byte) error {
	var SingleworkdayaveragepointsMap map[string]interface{}
	err := json.Unmarshal(b, &SingleworkdayaveragepointsMap)
	if err != nil {
		return err
	}
	
	if dateWorkdayString, ok := SingleworkdayaveragepointsMap["dateWorkday"].(string); ok {
		DateWorkday, _ := time.Parse("2006-01-02", dateWorkdayString)
		o.DateWorkday = &DateWorkday
	}
	
	if Division, ok := SingleworkdayaveragepointsMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if AveragePoints, ok := SingleworkdayaveragepointsMap["averagePoints"].(float64); ok {
		o.AveragePoints = &AveragePoints
	}
	
	if PerformanceProfile, ok := SingleworkdayaveragepointsMap["performanceProfile"].(map[string]interface{}); ok {
		PerformanceProfileString, _ := json.Marshal(PerformanceProfile)
		json.Unmarshal(PerformanceProfileString, &o.PerformanceProfile)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Singleworkdayaveragepoints) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

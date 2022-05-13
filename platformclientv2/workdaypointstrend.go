package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workdaypointstrend
type Workdaypointstrend struct { 
	// DateStartWorkday - The start workday for the query range for the gamification points trend. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateStartWorkday *time.Time `json:"dateStartWorkday,omitempty"`


	// DateEndWorkday - The end workday for the query range for the gamification points trend. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateEndWorkday *time.Time `json:"dateEndWorkday,omitempty"`


	// User - The targeted user for the query
	User *Userreference `json:"user,omitempty"`


	// DayOfWeek - Aggregated for same day comparison
	DayOfWeek *string `json:"dayOfWeek,omitempty"`


	// AveragePoints - The total average points
	AveragePoints *float64 `json:"averagePoints,omitempty"`


	// Trend - Daily points trends
	Trend *[]Workdaypointstrenditem `json:"trend,omitempty"`

}

func (o *Workdaypointstrend) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workdaypointstrend
	
	DateStartWorkday := new(string)
	if o.DateStartWorkday != nil {
		*DateStartWorkday = timeutil.Strftime(o.DateStartWorkday, "%Y-%m-%d")
	} else {
		DateStartWorkday = nil
	}
	
	DateEndWorkday := new(string)
	if o.DateEndWorkday != nil {
		*DateEndWorkday = timeutil.Strftime(o.DateEndWorkday, "%Y-%m-%d")
	} else {
		DateEndWorkday = nil
	}
	
	return json.Marshal(&struct { 
		DateStartWorkday *string `json:"dateStartWorkday,omitempty"`
		
		DateEndWorkday *string `json:"dateEndWorkday,omitempty"`
		
		User *Userreference `json:"user,omitempty"`
		
		DayOfWeek *string `json:"dayOfWeek,omitempty"`
		
		AveragePoints *float64 `json:"averagePoints,omitempty"`
		
		Trend *[]Workdaypointstrenditem `json:"trend,omitempty"`
		*Alias
	}{ 
		DateStartWorkday: DateStartWorkday,
		
		DateEndWorkday: DateEndWorkday,
		
		User: o.User,
		
		DayOfWeek: o.DayOfWeek,
		
		AveragePoints: o.AveragePoints,
		
		Trend: o.Trend,
		Alias:    (*Alias)(o),
	})
}

func (o *Workdaypointstrend) UnmarshalJSON(b []byte) error {
	var WorkdaypointstrendMap map[string]interface{}
	err := json.Unmarshal(b, &WorkdaypointstrendMap)
	if err != nil {
		return err
	}
	
	if dateStartWorkdayString, ok := WorkdaypointstrendMap["dateStartWorkday"].(string); ok {
		DateStartWorkday, _ := time.Parse("2006-01-02", dateStartWorkdayString)
		o.DateStartWorkday = &DateStartWorkday
	}
	
	if dateEndWorkdayString, ok := WorkdaypointstrendMap["dateEndWorkday"].(string); ok {
		DateEndWorkday, _ := time.Parse("2006-01-02", dateEndWorkdayString)
		o.DateEndWorkday = &DateEndWorkday
	}
	
	if User, ok := WorkdaypointstrendMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if DayOfWeek, ok := WorkdaypointstrendMap["dayOfWeek"].(string); ok {
		o.DayOfWeek = &DayOfWeek
	}
    
	if AveragePoints, ok := WorkdaypointstrendMap["averagePoints"].(float64); ok {
		o.AveragePoints = &AveragePoints
	}
    
	if Trend, ok := WorkdaypointstrendMap["trend"].([]interface{}); ok {
		TrendString, _ := json.Marshal(Trend)
		json.Unmarshal(TrendString, &o.Trend)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Workdaypointstrend) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

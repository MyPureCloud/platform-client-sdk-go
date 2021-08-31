package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Overallbestpointsitem
type Overallbestpointsitem struct { 
	// GranularityType - Best points aggregation interval granularity
	GranularityType *string `json:"granularityType,omitempty"`


	// Users - List of associated users with the equal points.
	Users *[]Userreference `json:"users,omitempty"`


	// Count - The count of the user IDs in the list
	Count *int `json:"count,omitempty"`


	// Points - Gamification points
	Points *int `json:"points,omitempty"`


	// DateStartWorkday - Start workday of the best points aggregation interval. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateStartWorkday *time.Time `json:"dateStartWorkday,omitempty"`


	// DateEndWorkday - End workday of the best points aggregation interval. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateEndWorkday *time.Time `json:"dateEndWorkday,omitempty"`

}

func (o *Overallbestpointsitem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Overallbestpointsitem
	
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
		GranularityType *string `json:"granularityType,omitempty"`
		
		Users *[]Userreference `json:"users,omitempty"`
		
		Count *int `json:"count,omitempty"`
		
		Points *int `json:"points,omitempty"`
		
		DateStartWorkday *string `json:"dateStartWorkday,omitempty"`
		
		DateEndWorkday *string `json:"dateEndWorkday,omitempty"`
		*Alias
	}{ 
		GranularityType: o.GranularityType,
		
		Users: o.Users,
		
		Count: o.Count,
		
		Points: o.Points,
		
		DateStartWorkday: DateStartWorkday,
		
		DateEndWorkday: DateEndWorkday,
		Alias:    (*Alias)(o),
	})
}

func (o *Overallbestpointsitem) UnmarshalJSON(b []byte) error {
	var OverallbestpointsitemMap map[string]interface{}
	err := json.Unmarshal(b, &OverallbestpointsitemMap)
	if err != nil {
		return err
	}
	
	if GranularityType, ok := OverallbestpointsitemMap["granularityType"].(string); ok {
		o.GranularityType = &GranularityType
	}
	
	if Users, ok := OverallbestpointsitemMap["users"].([]interface{}); ok {
		UsersString, _ := json.Marshal(Users)
		json.Unmarshal(UsersString, &o.Users)
	}
	
	if Count, ok := OverallbestpointsitemMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	
	if Points, ok := OverallbestpointsitemMap["points"].(float64); ok {
		PointsInt := int(Points)
		o.Points = &PointsInt
	}
	
	if dateStartWorkdayString, ok := OverallbestpointsitemMap["dateStartWorkday"].(string); ok {
		DateStartWorkday, _ := time.Parse("2006-01-02", dateStartWorkdayString)
		o.DateStartWorkday = &DateStartWorkday
	}
	
	if dateEndWorkdayString, ok := OverallbestpointsitemMap["dateEndWorkday"].(string); ok {
		DateEndWorkday, _ := time.Parse("2006-01-02", dateEndWorkdayString)
		o.DateEndWorkday = &DateEndWorkday
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Overallbestpointsitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

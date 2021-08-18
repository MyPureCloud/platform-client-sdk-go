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

func (u *Overallbestpointsitem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Overallbestpointsitem

	
	DateStartWorkday := new(string)
	if u.DateStartWorkday != nil {
		*DateStartWorkday = timeutil.Strftime(u.DateStartWorkday, "%Y-%m-%d")
	} else {
		DateStartWorkday = nil
	}
	
	DateEndWorkday := new(string)
	if u.DateEndWorkday != nil {
		*DateEndWorkday = timeutil.Strftime(u.DateEndWorkday, "%Y-%m-%d")
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
		GranularityType: u.GranularityType,
		
		Users: u.Users,
		
		Count: u.Count,
		
		Points: u.Points,
		
		DateStartWorkday: DateStartWorkday,
		
		DateEndWorkday: DateEndWorkday,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Overallbestpointsitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

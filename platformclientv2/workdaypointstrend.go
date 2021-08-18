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

func (u *Workdaypointstrend) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workdaypointstrend

	
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
		
		User: u.User,
		
		DayOfWeek: u.DayOfWeek,
		
		AveragePoints: u.AveragePoints,
		
		Trend: u.Trend,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Workdaypointstrend) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

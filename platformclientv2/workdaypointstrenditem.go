package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workdaypointstrenditem
type Workdaypointstrenditem struct { 
	// DateWorkday - workday date for the points trend. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateWorkday *time.Time `json:"dateWorkday,omitempty"`


	// Points - workday points for the date
	Points *float64 `json:"points,omitempty"`

}

func (o *Workdaypointstrenditem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workdaypointstrenditem
	
	DateWorkday := new(string)
	if o.DateWorkday != nil {
		*DateWorkday = timeutil.Strftime(o.DateWorkday, "%Y-%m-%d")
	} else {
		DateWorkday = nil
	}
	
	return json.Marshal(&struct { 
		DateWorkday *string `json:"dateWorkday,omitempty"`
		
		Points *float64 `json:"points,omitempty"`
		*Alias
	}{ 
		DateWorkday: DateWorkday,
		
		Points: o.Points,
		Alias:    (*Alias)(o),
	})
}

func (o *Workdaypointstrenditem) UnmarshalJSON(b []byte) error {
	var WorkdaypointstrenditemMap map[string]interface{}
	err := json.Unmarshal(b, &WorkdaypointstrenditemMap)
	if err != nil {
		return err
	}
	
	if dateWorkdayString, ok := WorkdaypointstrenditemMap["dateWorkday"].(string); ok {
		DateWorkday, _ := time.Parse("2006-01-02", dateWorkdayString)
		o.DateWorkday = &DateWorkday
	}
	
	if Points, ok := WorkdaypointstrenditemMap["points"].(float64); ok {
		o.Points = &Points
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Workdaypointstrenditem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

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

func (u *Workdaypointstrenditem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workdaypointstrenditem

	
	DateWorkday := new(string)
	if u.DateWorkday != nil {
		*DateWorkday = timeutil.Strftime(u.DateWorkday, "%Y-%m-%d")
	} else {
		DateWorkday = nil
	}
	

	return json.Marshal(&struct { 
		DateWorkday *string `json:"dateWorkday,omitempty"`
		
		Points *float64 `json:"points,omitempty"`
		*Alias
	}{ 
		DateWorkday: DateWorkday,
		
		Points: u.Points,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Workdaypointstrenditem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

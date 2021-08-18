package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Predictorschedule
type Predictorschedule struct { 
	// ScheduleType - The predictor schedule type.
	ScheduleType *string `json:"scheduleType,omitempty"`


	// DateStarted - DateTime indicating when the predictor schedule was started. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStarted *time.Time `json:"dateStarted,omitempty"`

}

func (u *Predictorschedule) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Predictorschedule

	
	DateStarted := new(string)
	if u.DateStarted != nil {
		
		*DateStarted = timeutil.Strftime(u.DateStarted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStarted = nil
	}
	

	return json.Marshal(&struct { 
		ScheduleType *string `json:"scheduleType,omitempty"`
		
		DateStarted *string `json:"dateStarted,omitempty"`
		*Alias
	}{ 
		ScheduleType: u.ScheduleType,
		
		DateStarted: DateStarted,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Predictorschedule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

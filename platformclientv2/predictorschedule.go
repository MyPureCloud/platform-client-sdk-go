package platformclientv2
import (
	"time"
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

// String returns a JSON representation of the model
func (o *Predictorschedule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

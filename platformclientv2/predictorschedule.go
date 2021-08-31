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

func (o *Predictorschedule) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Predictorschedule
	
	DateStarted := new(string)
	if o.DateStarted != nil {
		
		*DateStarted = timeutil.Strftime(o.DateStarted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStarted = nil
	}
	
	return json.Marshal(&struct { 
		ScheduleType *string `json:"scheduleType,omitempty"`
		
		DateStarted *string `json:"dateStarted,omitempty"`
		*Alias
	}{ 
		ScheduleType: o.ScheduleType,
		
		DateStarted: DateStarted,
		Alias:    (*Alias)(o),
	})
}

func (o *Predictorschedule) UnmarshalJSON(b []byte) error {
	var PredictorscheduleMap map[string]interface{}
	err := json.Unmarshal(b, &PredictorscheduleMap)
	if err != nil {
		return err
	}
	
	if ScheduleType, ok := PredictorscheduleMap["scheduleType"].(string); ok {
		o.ScheduleType = &ScheduleType
	}
	
	if dateStartedString, ok := PredictorscheduleMap["dateStarted"].(string); ok {
		DateStarted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartedString)
		o.DateStarted = &DateStarted
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Predictorschedule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

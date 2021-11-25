package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Waitlistposition
type Waitlistposition struct { 
	// TimeOffRequest - The time off request for this wait list position
	TimeOffRequest *Timeoffrequestreference `json:"timeOffRequest,omitempty"`


	// TimeOffLimit - The time off limit for which time off request is waitlisted
	TimeOffLimit *Timeofflimitreference `json:"timeOffLimit,omitempty"`


	// Date - The date to which this wait list position applies, as defined by the time zone of the business unit. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	Date *time.Time `json:"date,omitempty"`


	// WaitlistPosition - The time off request's position in the waitlist on the date. 1 means time off is the first in the waitlist
	WaitlistPosition *int `json:"waitlistPosition,omitempty"`

}

func (o *Waitlistposition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Waitlistposition
	
	Date := new(string)
	if o.Date != nil {
		*Date = timeutil.Strftime(o.Date, "%Y-%m-%d")
	} else {
		Date = nil
	}
	
	return json.Marshal(&struct { 
		TimeOffRequest *Timeoffrequestreference `json:"timeOffRequest,omitempty"`
		
		TimeOffLimit *Timeofflimitreference `json:"timeOffLimit,omitempty"`
		
		Date *string `json:"date,omitempty"`
		
		WaitlistPosition *int `json:"waitlistPosition,omitempty"`
		*Alias
	}{ 
		TimeOffRequest: o.TimeOffRequest,
		
		TimeOffLimit: o.TimeOffLimit,
		
		Date: Date,
		
		WaitlistPosition: o.WaitlistPosition,
		Alias:    (*Alias)(o),
	})
}

func (o *Waitlistposition) UnmarshalJSON(b []byte) error {
	var WaitlistpositionMap map[string]interface{}
	err := json.Unmarshal(b, &WaitlistpositionMap)
	if err != nil {
		return err
	}
	
	if TimeOffRequest, ok := WaitlistpositionMap["timeOffRequest"].(map[string]interface{}); ok {
		TimeOffRequestString, _ := json.Marshal(TimeOffRequest)
		json.Unmarshal(TimeOffRequestString, &o.TimeOffRequest)
	}
	
	if TimeOffLimit, ok := WaitlistpositionMap["timeOffLimit"].(map[string]interface{}); ok {
		TimeOffLimitString, _ := json.Marshal(TimeOffLimit)
		json.Unmarshal(TimeOffLimitString, &o.TimeOffLimit)
	}
	
	if dateString, ok := WaitlistpositionMap["date"].(string); ok {
		Date, _ := time.Parse("2006-01-02", dateString)
		o.Date = &Date
	}
	
	if WaitlistPosition, ok := WaitlistpositionMap["waitlistPosition"].(float64); ok {
		WaitlistPositionInt := int(WaitlistPosition)
		o.WaitlistPosition = &WaitlistPositionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Waitlistposition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

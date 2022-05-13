package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Availabletimeoffrange - A list of available time off values in minutes and a number of time off requests currently in waitlist for each interval in requested date range, according to granularity.
type Availabletimeoffrange struct { 
	// TimeOffLimit - The time off limit
	TimeOffLimit *Timeofflimitreference `json:"timeOffLimit,omitempty"`


	// StartDate - Start date of the requested date range. The end date is determined by the size of interval list. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	StartDate *time.Time `json:"startDate,omitempty"`


	// Granularity - Granularity choice for time off limit
	Granularity *string `json:"granularity,omitempty"`


	// AvailableMinutesPerInterval - The list of available time off values in minutes per granularity interval
	AvailableMinutesPerInterval *[]int `json:"availableMinutesPerInterval,omitempty"`


	// WaitlistedRequestsPerInterval - The current number of waitlisted time off requests for every interval per granularity
	WaitlistedRequestsPerInterval *[]int `json:"waitlistedRequestsPerInterval,omitempty"`


	// WaitlistEnabled - Whether the time off request can be waitlisted
	WaitlistEnabled *bool `json:"waitlistEnabled,omitempty"`

}

func (o *Availabletimeoffrange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Availabletimeoffrange
	
	StartDate := new(string)
	if o.StartDate != nil {
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%d")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		TimeOffLimit *Timeofflimitreference `json:"timeOffLimit,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		Granularity *string `json:"granularity,omitempty"`
		
		AvailableMinutesPerInterval *[]int `json:"availableMinutesPerInterval,omitempty"`
		
		WaitlistedRequestsPerInterval *[]int `json:"waitlistedRequestsPerInterval,omitempty"`
		
		WaitlistEnabled *bool `json:"waitlistEnabled,omitempty"`
		*Alias
	}{ 
		TimeOffLimit: o.TimeOffLimit,
		
		StartDate: StartDate,
		
		Granularity: o.Granularity,
		
		AvailableMinutesPerInterval: o.AvailableMinutesPerInterval,
		
		WaitlistedRequestsPerInterval: o.WaitlistedRequestsPerInterval,
		
		WaitlistEnabled: o.WaitlistEnabled,
		Alias:    (*Alias)(o),
	})
}

func (o *Availabletimeoffrange) UnmarshalJSON(b []byte) error {
	var AvailabletimeoffrangeMap map[string]interface{}
	err := json.Unmarshal(b, &AvailabletimeoffrangeMap)
	if err != nil {
		return err
	}
	
	if TimeOffLimit, ok := AvailabletimeoffrangeMap["timeOffLimit"].(map[string]interface{}); ok {
		TimeOffLimitString, _ := json.Marshal(TimeOffLimit)
		json.Unmarshal(TimeOffLimitString, &o.TimeOffLimit)
	}
	
	if startDateString, ok := AvailabletimeoffrangeMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02", startDateString)
		o.StartDate = &StartDate
	}
	
	if Granularity, ok := AvailabletimeoffrangeMap["granularity"].(string); ok {
		o.Granularity = &Granularity
	}
    
	if AvailableMinutesPerInterval, ok := AvailabletimeoffrangeMap["availableMinutesPerInterval"].([]interface{}); ok {
		AvailableMinutesPerIntervalString, _ := json.Marshal(AvailableMinutesPerInterval)
		json.Unmarshal(AvailableMinutesPerIntervalString, &o.AvailableMinutesPerInterval)
	}
	
	if WaitlistedRequestsPerInterval, ok := AvailabletimeoffrangeMap["waitlistedRequestsPerInterval"].([]interface{}); ok {
		WaitlistedRequestsPerIntervalString, _ := json.Marshal(WaitlistedRequestsPerInterval)
		json.Unmarshal(WaitlistedRequestsPerIntervalString, &o.WaitlistedRequestsPerInterval)
	}
	
	if WaitlistEnabled, ok := AvailabletimeoffrangeMap["waitlistEnabled"].(bool); ok {
		o.WaitlistEnabled = &WaitlistEnabled
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Availabletimeoffrange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

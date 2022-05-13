package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeofflimitvaluerange - The date range with lists of time off limit, allocated and waitlisted minutes.
type Timeofflimitvaluerange struct { 
	// TimeOffLimit - The ID of the time off limit
	TimeOffLimit *Timeofflimitreference `json:"timeOffLimit,omitempty"`


	// StartDate - Start date of the requested date range, in ISO-8601 format. The end date is determined by the size of interval lists
	StartDate *time.Time `json:"startDate,omitempty"`


	// Granularity - Granularity choice for time off limit
	Granularity *string `json:"granularity,omitempty"`


	// LimitMinutesPerInterval - A list of time off limit values in minutes per granularity interval
	LimitMinutesPerInterval *[]int `json:"limitMinutesPerInterval,omitempty"`


	// AllocatedMinutesPerInterval - A list of allocated time off minutes per granularity interval
	AllocatedMinutesPerInterval *[]int `json:"allocatedMinutesPerInterval,omitempty"`


	// WaitlistedMinutesPerInterval - A list of waitlisted time off minutes per granularity interval
	WaitlistedMinutesPerInterval *[]int `json:"waitlistedMinutesPerInterval,omitempty"`


	// WaitlistedRequestsPerInterval - The current number of waitlisted time off requests for every interval per granularity
	WaitlistedRequestsPerInterval *[]int `json:"waitlistedRequestsPerInterval,omitempty"`


	// Metadata - Version metadata for the time off limit
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

func (o *Timeofflimitvaluerange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Timeofflimitvaluerange
	
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
		
		LimitMinutesPerInterval *[]int `json:"limitMinutesPerInterval,omitempty"`
		
		AllocatedMinutesPerInterval *[]int `json:"allocatedMinutesPerInterval,omitempty"`
		
		WaitlistedMinutesPerInterval *[]int `json:"waitlistedMinutesPerInterval,omitempty"`
		
		WaitlistedRequestsPerInterval *[]int `json:"waitlistedRequestsPerInterval,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		*Alias
	}{ 
		TimeOffLimit: o.TimeOffLimit,
		
		StartDate: StartDate,
		
		Granularity: o.Granularity,
		
		LimitMinutesPerInterval: o.LimitMinutesPerInterval,
		
		AllocatedMinutesPerInterval: o.AllocatedMinutesPerInterval,
		
		WaitlistedMinutesPerInterval: o.WaitlistedMinutesPerInterval,
		
		WaitlistedRequestsPerInterval: o.WaitlistedRequestsPerInterval,
		
		Metadata: o.Metadata,
		Alias:    (*Alias)(o),
	})
}

func (o *Timeofflimitvaluerange) UnmarshalJSON(b []byte) error {
	var TimeofflimitvaluerangeMap map[string]interface{}
	err := json.Unmarshal(b, &TimeofflimitvaluerangeMap)
	if err != nil {
		return err
	}
	
	if TimeOffLimit, ok := TimeofflimitvaluerangeMap["timeOffLimit"].(map[string]interface{}); ok {
		TimeOffLimitString, _ := json.Marshal(TimeOffLimit)
		json.Unmarshal(TimeOffLimitString, &o.TimeOffLimit)
	}
	
	if startDateString, ok := TimeofflimitvaluerangeMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02", startDateString)
		o.StartDate = &StartDate
	}
	
	if Granularity, ok := TimeofflimitvaluerangeMap["granularity"].(string); ok {
		o.Granularity = &Granularity
	}
    
	if LimitMinutesPerInterval, ok := TimeofflimitvaluerangeMap["limitMinutesPerInterval"].([]interface{}); ok {
		LimitMinutesPerIntervalString, _ := json.Marshal(LimitMinutesPerInterval)
		json.Unmarshal(LimitMinutesPerIntervalString, &o.LimitMinutesPerInterval)
	}
	
	if AllocatedMinutesPerInterval, ok := TimeofflimitvaluerangeMap["allocatedMinutesPerInterval"].([]interface{}); ok {
		AllocatedMinutesPerIntervalString, _ := json.Marshal(AllocatedMinutesPerInterval)
		json.Unmarshal(AllocatedMinutesPerIntervalString, &o.AllocatedMinutesPerInterval)
	}
	
	if WaitlistedMinutesPerInterval, ok := TimeofflimitvaluerangeMap["waitlistedMinutesPerInterval"].([]interface{}); ok {
		WaitlistedMinutesPerIntervalString, _ := json.Marshal(WaitlistedMinutesPerInterval)
		json.Unmarshal(WaitlistedMinutesPerIntervalString, &o.WaitlistedMinutesPerInterval)
	}
	
	if WaitlistedRequestsPerInterval, ok := TimeofflimitvaluerangeMap["waitlistedRequestsPerInterval"].([]interface{}); ok {
		WaitlistedRequestsPerIntervalString, _ := json.Marshal(WaitlistedRequestsPerInterval)
		json.Unmarshal(WaitlistedRequestsPerIntervalString, &o.WaitlistedRequestsPerInterval)
	}
	
	if Metadata, ok := TimeofflimitvaluerangeMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Timeofflimitvaluerange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

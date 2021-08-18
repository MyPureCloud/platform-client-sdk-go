package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createworkplanshift - Shift in a work plan
type Createworkplanshift struct { 
	// Name - Name of the shift
	Name *string `json:"name,omitempty"`


	// Days - Days of the week applicable for this shift
	Days *Setwrapperdayofweek `json:"days,omitempty"`


	// FlexibleStartTime - Whether the start time of the shift is flexible
	FlexibleStartTime *bool `json:"flexibleStartTime,omitempty"`


	// ExactStartTimeMinutesFromMidnight - Exact start time of the shift defined as offset minutes from midnight. Used if flexibleStartTime == false
	ExactStartTimeMinutesFromMidnight *int `json:"exactStartTimeMinutesFromMidnight,omitempty"`


	// EarliestStartTimeMinutesFromMidnight - Earliest start time of the shift defined as offset minutes from midnight. Used if flexibleStartTime == true
	EarliestStartTimeMinutesFromMidnight *int `json:"earliestStartTimeMinutesFromMidnight,omitempty"`


	// LatestStartTimeMinutesFromMidnight - Latest start time of the shift defined as offset minutes from midnight. Used if flexibleStartTime == true
	LatestStartTimeMinutesFromMidnight *int `json:"latestStartTimeMinutesFromMidnight,omitempty"`


	// ConstrainStopTime - Whether the latest stop time constraint for the shift is enabled.  Deprecated, use constrainLatestStopTime instead
	ConstrainStopTime *bool `json:"constrainStopTime,omitempty"`


	// ConstrainLatestStopTime - Whether the latest stop time constraint for the shift is enabled
	ConstrainLatestStopTime *bool `json:"constrainLatestStopTime,omitempty"`


	// LatestStopTimeMinutesFromMidnight - Latest stop time of the shift defined as offset minutes from midnight. Used if constrainStopTime == true
	LatestStopTimeMinutesFromMidnight *int `json:"latestStopTimeMinutesFromMidnight,omitempty"`


	// ConstrainEarliestStopTime - Whether the earliest stop time constraint for the shift is enabled
	ConstrainEarliestStopTime *bool `json:"constrainEarliestStopTime,omitempty"`


	// EarliestStopTimeMinutesFromMidnight - This is the earliest time a shift can end
	EarliestStopTimeMinutesFromMidnight *int `json:"earliestStopTimeMinutesFromMidnight,omitempty"`


	// StartIncrementMinutes - Increment in offset minutes that would contribute to different possible start times for the shift. Used if flexibleStartTime == true
	StartIncrementMinutes *int `json:"startIncrementMinutes,omitempty"`


	// FlexiblePaidTime - Whether the paid time setting for the shift is flexible
	FlexiblePaidTime *bool `json:"flexiblePaidTime,omitempty"`


	// ExactPaidTimeMinutes - Exact paid time in minutes configured for the shift. Used if flexiblePaidTime == false
	ExactPaidTimeMinutes *int `json:"exactPaidTimeMinutes,omitempty"`


	// MinimumPaidTimeMinutes - Minimum paid time in minutes configured for the shift. Used if flexiblePaidTime == true
	MinimumPaidTimeMinutes *int `json:"minimumPaidTimeMinutes,omitempty"`


	// MaximumPaidTimeMinutes - Maximum paid time in minutes configured for the shift. Used if flexiblePaidTime == true
	MaximumPaidTimeMinutes *int `json:"maximumPaidTimeMinutes,omitempty"`


	// ConstrainContiguousWorkTime - Whether the contiguous time constraint for the shift is enabled
	ConstrainContiguousWorkTime *bool `json:"constrainContiguousWorkTime,omitempty"`


	// MinimumContiguousWorkTimeMinutes - Minimum contiguous time in minutes configured for the shift. Used if constrainContiguousWorkTime == true
	MinimumContiguousWorkTimeMinutes *int `json:"minimumContiguousWorkTimeMinutes,omitempty"`


	// MaximumContiguousWorkTimeMinutes - Maximum contiguous time in minutes configured for the shift. Used if constrainContiguousWorkTime == true
	MaximumContiguousWorkTimeMinutes *int `json:"maximumContiguousWorkTimeMinutes,omitempty"`


	// Activities - Activities configured for this shift
	Activities *[]Createworkplanactivity `json:"activities,omitempty"`

}

func (u *Createworkplanshift) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createworkplanshift

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Days *Setwrapperdayofweek `json:"days,omitempty"`
		
		FlexibleStartTime *bool `json:"flexibleStartTime,omitempty"`
		
		ExactStartTimeMinutesFromMidnight *int `json:"exactStartTimeMinutesFromMidnight,omitempty"`
		
		EarliestStartTimeMinutesFromMidnight *int `json:"earliestStartTimeMinutesFromMidnight,omitempty"`
		
		LatestStartTimeMinutesFromMidnight *int `json:"latestStartTimeMinutesFromMidnight,omitempty"`
		
		ConstrainStopTime *bool `json:"constrainStopTime,omitempty"`
		
		ConstrainLatestStopTime *bool `json:"constrainLatestStopTime,omitempty"`
		
		LatestStopTimeMinutesFromMidnight *int `json:"latestStopTimeMinutesFromMidnight,omitempty"`
		
		ConstrainEarliestStopTime *bool `json:"constrainEarliestStopTime,omitempty"`
		
		EarliestStopTimeMinutesFromMidnight *int `json:"earliestStopTimeMinutesFromMidnight,omitempty"`
		
		StartIncrementMinutes *int `json:"startIncrementMinutes,omitempty"`
		
		FlexiblePaidTime *bool `json:"flexiblePaidTime,omitempty"`
		
		ExactPaidTimeMinutes *int `json:"exactPaidTimeMinutes,omitempty"`
		
		MinimumPaidTimeMinutes *int `json:"minimumPaidTimeMinutes,omitempty"`
		
		MaximumPaidTimeMinutes *int `json:"maximumPaidTimeMinutes,omitempty"`
		
		ConstrainContiguousWorkTime *bool `json:"constrainContiguousWorkTime,omitempty"`
		
		MinimumContiguousWorkTimeMinutes *int `json:"minimumContiguousWorkTimeMinutes,omitempty"`
		
		MaximumContiguousWorkTimeMinutes *int `json:"maximumContiguousWorkTimeMinutes,omitempty"`
		
		Activities *[]Createworkplanactivity `json:"activities,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Days: u.Days,
		
		FlexibleStartTime: u.FlexibleStartTime,
		
		ExactStartTimeMinutesFromMidnight: u.ExactStartTimeMinutesFromMidnight,
		
		EarliestStartTimeMinutesFromMidnight: u.EarliestStartTimeMinutesFromMidnight,
		
		LatestStartTimeMinutesFromMidnight: u.LatestStartTimeMinutesFromMidnight,
		
		ConstrainStopTime: u.ConstrainStopTime,
		
		ConstrainLatestStopTime: u.ConstrainLatestStopTime,
		
		LatestStopTimeMinutesFromMidnight: u.LatestStopTimeMinutesFromMidnight,
		
		ConstrainEarliestStopTime: u.ConstrainEarliestStopTime,
		
		EarliestStopTimeMinutesFromMidnight: u.EarliestStopTimeMinutesFromMidnight,
		
		StartIncrementMinutes: u.StartIncrementMinutes,
		
		FlexiblePaidTime: u.FlexiblePaidTime,
		
		ExactPaidTimeMinutes: u.ExactPaidTimeMinutes,
		
		MinimumPaidTimeMinutes: u.MinimumPaidTimeMinutes,
		
		MaximumPaidTimeMinutes: u.MaximumPaidTimeMinutes,
		
		ConstrainContiguousWorkTime: u.ConstrainContiguousWorkTime,
		
		MinimumContiguousWorkTimeMinutes: u.MinimumContiguousWorkTimeMinutes,
		
		MaximumContiguousWorkTimeMinutes: u.MaximumContiguousWorkTimeMinutes,
		
		Activities: u.Activities,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Createworkplanshift) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

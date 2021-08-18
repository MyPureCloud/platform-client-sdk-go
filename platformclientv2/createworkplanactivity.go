package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createworkplanactivity - Activity configured for shift in work plan
type Createworkplanactivity struct { 
	// ActivityCodeId - ID of the activity code associated with this activity
	ActivityCodeId *string `json:"activityCodeId,omitempty"`


	// Description - Description of the activity
	Description *string `json:"description,omitempty"`


	// LengthMinutes - Length of the activity in minutes
	LengthMinutes *int `json:"lengthMinutes,omitempty"`


	// StartTimeIsRelativeToShiftStart - Whether the start time of the activity is relative to the start time of the shift it belongs to
	StartTimeIsRelativeToShiftStart *bool `json:"startTimeIsRelativeToShiftStart,omitempty"`


	// FlexibleStartTime - Whether the start time of the activity is flexible
	FlexibleStartTime *bool `json:"flexibleStartTime,omitempty"`


	// EarliestStartTimeMinutes - Earliest activity start in offset minutes relative to shift start time if startTimeIsRelativeToShiftStart == true else its based on midnight. Used if flexibleStartTime == true
	EarliestStartTimeMinutes *int `json:"earliestStartTimeMinutes,omitempty"`


	// LatestStartTimeMinutes - Latest activity start in offset minutes relative to shift start time if startTimeIsRelativeToShiftStart == true else its based on midnight. Used if flexibleStartTime == true
	LatestStartTimeMinutes *int `json:"latestStartTimeMinutes,omitempty"`


	// ExactStartTimeMinutes - Exact activity start in offset minutes relative to shift start time if startTimeIsRelativeToShiftStart == true else its based on midnight. Used if flexibleStartTime == false
	ExactStartTimeMinutes *int `json:"exactStartTimeMinutes,omitempty"`


	// StartTimeIncrementMinutes - Increment in offset minutes that would contribute to different possible start times for the activity
	StartTimeIncrementMinutes *int `json:"startTimeIncrementMinutes,omitempty"`


	// CountsAsPaidTime - Whether the activity is paid
	CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`


	// CountsAsContiguousWorkTime - Whether the activity duration is counted towards contiguous work time
	CountsAsContiguousWorkTime *bool `json:"countsAsContiguousWorkTime,omitempty"`


	// MinimumLengthFromShiftStartMinutes - The minimum duration between shift start and shift item (e.g., break or meal) start in minutes
	MinimumLengthFromShiftStartMinutes *int `json:"minimumLengthFromShiftStartMinutes,omitempty"`


	// MinimumLengthFromShiftEndMinutes - The minimum duration between shift item (e.g., break or meal) end and shift end in minutes
	MinimumLengthFromShiftEndMinutes *int `json:"minimumLengthFromShiftEndMinutes,omitempty"`

}

func (u *Createworkplanactivity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createworkplanactivity

	

	return json.Marshal(&struct { 
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		StartTimeIsRelativeToShiftStart *bool `json:"startTimeIsRelativeToShiftStart,omitempty"`
		
		FlexibleStartTime *bool `json:"flexibleStartTime,omitempty"`
		
		EarliestStartTimeMinutes *int `json:"earliestStartTimeMinutes,omitempty"`
		
		LatestStartTimeMinutes *int `json:"latestStartTimeMinutes,omitempty"`
		
		ExactStartTimeMinutes *int `json:"exactStartTimeMinutes,omitempty"`
		
		StartTimeIncrementMinutes *int `json:"startTimeIncrementMinutes,omitempty"`
		
		CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`
		
		CountsAsContiguousWorkTime *bool `json:"countsAsContiguousWorkTime,omitempty"`
		
		MinimumLengthFromShiftStartMinutes *int `json:"minimumLengthFromShiftStartMinutes,omitempty"`
		
		MinimumLengthFromShiftEndMinutes *int `json:"minimumLengthFromShiftEndMinutes,omitempty"`
		*Alias
	}{ 
		ActivityCodeId: u.ActivityCodeId,
		
		Description: u.Description,
		
		LengthMinutes: u.LengthMinutes,
		
		StartTimeIsRelativeToShiftStart: u.StartTimeIsRelativeToShiftStart,
		
		FlexibleStartTime: u.FlexibleStartTime,
		
		EarliestStartTimeMinutes: u.EarliestStartTimeMinutes,
		
		LatestStartTimeMinutes: u.LatestStartTimeMinutes,
		
		ExactStartTimeMinutes: u.ExactStartTimeMinutes,
		
		StartTimeIncrementMinutes: u.StartTimeIncrementMinutes,
		
		CountsAsPaidTime: u.CountsAsPaidTime,
		
		CountsAsContiguousWorkTime: u.CountsAsContiguousWorkTime,
		
		MinimumLengthFromShiftStartMinutes: u.MinimumLengthFromShiftStartMinutes,
		
		MinimumLengthFromShiftEndMinutes: u.MinimumLengthFromShiftEndMinutes,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Createworkplanactivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

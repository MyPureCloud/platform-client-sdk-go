package platformclientv2
import (
	"encoding/json"
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

}

// String returns a JSON representation of the model
func (o *Createworkplanactivity) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

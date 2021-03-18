package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Workplanactivity - Activity configured for shift in work plan
type Workplanactivity struct { 
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


	// Id - ID of the activity. This is required only for the case of updating an existing activity
	Id *string `json:"id,omitempty"`


	// Delete - If marked true for updating an existing activity, the activity will be permanently deleted
	Delete *bool `json:"delete,omitempty"`

}

// String returns a JSON representation of the model
func (o *Workplanactivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

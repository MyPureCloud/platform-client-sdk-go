package platformclientv2
import (
	"time"
	"encoding/json"
)

// Reschedulerequest
type Reschedulerequest struct { 
	// StartDate - The start date of the range to reschedule in ISO-8601 format
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - The end date of the range to reschedule in ISO-8601 format
	EndDate *time.Time `json:"endDate,omitempty"`


	// AgentIds - The IDs of the agents to reschedule.  Null or empty means all agents on the schedule
	AgentIds *[]string `json:"agentIds,omitempty"`


	// ActivityCodeIds - The IDs of the activity codes to reschedule. Null or empty means all activity codes will be considered
	ActivityCodeIds *[]string `json:"activityCodeIds,omitempty"`


	// DoNotChangeWeeklyPaidTime - Whether to prevent changes to weekly paid time
	DoNotChangeWeeklyPaidTime *bool `json:"doNotChangeWeeklyPaidTime,omitempty"`


	// DoNotChangeDailyPaidTime - Whether to prevent changes to daily paid time
	DoNotChangeDailyPaidTime *bool `json:"doNotChangeDailyPaidTime,omitempty"`


	// DoNotChangeShiftStartTimes - Whether to prevent changes to shift start times
	DoNotChangeShiftStartTimes *bool `json:"doNotChangeShiftStartTimes,omitempty"`


	// DoNotChangeManuallyEditedShifts - Whether to prevent changes to manually edited shifts
	DoNotChangeManuallyEditedShifts *bool `json:"doNotChangeManuallyEditedShifts,omitempty"`

}

// String returns a JSON representation of the model
func (o *Reschedulerequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

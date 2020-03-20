package platformclientv2
import (
	"time"
	"encoding/json"
)

// Reschedulingoptionsresponse
type Reschedulingoptionsresponse struct { 
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


	// ExistingScheduleId - The schedule ID of the schedule to which the results will be applied
	ExistingScheduleId *string `json:"existingScheduleId,omitempty"`


	// ExistingScheduleVersion - The version of the schedule at the time the rescheduling was initiated
	ExistingScheduleVersion *int32 `json:"existingScheduleVersion,omitempty"`

}

// String returns a JSON representation of the model
func (o *Reschedulingoptionsresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

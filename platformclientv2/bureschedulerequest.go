package platformclientv2
import (
	"time"
	"encoding/json"
)

// Bureschedulerequest
type Bureschedulerequest struct { 
	// StartDate - The start of the range to reschedule.  Defaults to the beginning of the schedule. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - The end of the range to reschedule.  Defaults the the end of the schedule. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	EndDate *time.Time `json:"endDate,omitempty"`


	// AgentIds - The IDs of the agents to consider for rescheduling.  Omit to consider all agents in the specified management units.Agents not in the specified management units will be ignored
	AgentIds *[]string `json:"agentIds,omitempty"`


	// ActivityCodeIds - The IDs of the activity codes to consider for rescheduling.  Omit to consider all activity codes
	ActivityCodeIds *[]string `json:"activityCodeIds,omitempty"`


	// ManagementUnitIds - The IDs of the management units to reschedule
	ManagementUnitIds *[]string `json:"managementUnitIds,omitempty"`


	// DoNotChangeWeeklyPaidTime - Instructs the scheduler whether it is allowed to change weekly paid time
	DoNotChangeWeeklyPaidTime *bool `json:"doNotChangeWeeklyPaidTime,omitempty"`


	// DoNotChangeDailyPaidTime - Instructs the scheduler whether it is allowed to change daily paid time
	DoNotChangeDailyPaidTime *bool `json:"doNotChangeDailyPaidTime,omitempty"`


	// DoNotChangeShiftStartTimes - Instructs the scheduler whether it is allowed to change shift start times
	DoNotChangeShiftStartTimes *bool `json:"doNotChangeShiftStartTimes,omitempty"`


	// DoNotChangeManuallyEditedShifts - Instructs the scheduler whether it is allowed to change manually edited shifts
	DoNotChangeManuallyEditedShifts *bool `json:"doNotChangeManuallyEditedShifts,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bureschedulerequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

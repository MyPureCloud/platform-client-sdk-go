package platformclientv2
import (
	"time"
	"encoding/json"
)

// Reschedulingoptionsrunresponse
type Reschedulingoptionsrunresponse struct { 
	// ExistingSchedule - The existing schedule to which this reschedule run applies
	ExistingSchedule *Buschedulereference `json:"existingSchedule,omitempty"`


	// StartDate - The start date of the period to reschedule. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - The end date of the period to reschedule. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	EndDate *time.Time `json:"endDate,omitempty"`


	// ManagementUnits - Per-management unit rescheduling options
	ManagementUnits *[]Reschedulingmanagementunitresponse `json:"managementUnits,omitempty"`


	// AgentCount - The number of agents to be considered in the reschedule
	AgentCount *int32 `json:"agentCount,omitempty"`


	// ActivityCodeIds - The IDs of the activity codes being considered for reschedule
	ActivityCodeIds *[]string `json:"activityCodeIds,omitempty"`


	// DoNotChangeWeeklyPaidTime - Whether weekly paid time is allowed to be changed
	DoNotChangeWeeklyPaidTime *bool `json:"doNotChangeWeeklyPaidTime,omitempty"`


	// DoNotChangeDailyPaidTime - Whether daily paid time is allowed to be changed
	DoNotChangeDailyPaidTime *bool `json:"doNotChangeDailyPaidTime,omitempty"`


	// DoNotChangeShiftStartTimes - Whether shift start times are allowed to be changed
	DoNotChangeShiftStartTimes *bool `json:"doNotChangeShiftStartTimes,omitempty"`


	// DoNotChangeManuallyEditedShifts - Whether manually edited shifts are allowed to be changed
	DoNotChangeManuallyEditedShifts *bool `json:"doNotChangeManuallyEditedShifts,omitempty"`

}

// String returns a JSON representation of the model
func (o *Reschedulingoptionsrunresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

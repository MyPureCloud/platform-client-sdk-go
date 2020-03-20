package platformclientv2
import (
	"time"
	"encoding/json"
)

// Schedulingrunresponse - Information containing details of a schedule run
type Schedulingrunresponse struct { 
	// RunId - ID of the schedule run
	RunId *string `json:"runId,omitempty"`


	// SchedulerRunId - The runId from scheduler service.  Useful for debugging schedule errors
	SchedulerRunId *string `json:"schedulerRunId,omitempty"`


	// IntradayRescheduling - Whether this is the result of a rescheduling request
	IntradayRescheduling *bool `json:"intradayRescheduling,omitempty"`


	// State - Status of the schedule run
	State *string `json:"state,omitempty"`


	// PercentComplete - Completion percentage of the schedule run
	PercentComplete *float64 `json:"percentComplete,omitempty"`


	// TargetWeek - The start date of the week for which the scheduling is done in yyyy-MM-dd format
	TargetWeek *string `json:"targetWeek,omitempty"`


	// ScheduleId - ID of the schedule. Does not apply to reschedule, see reschedulingOptions.existingScheduleId
	ScheduleId *string `json:"scheduleId,omitempty"`


	// ScheduleDescription - Description of the schedule
	ScheduleDescription *string `json:"scheduleDescription,omitempty"`


	// SchedulingStartTime - Start time of the schedule run. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	SchedulingStartTime *time.Time `json:"schedulingStartTime,omitempty"`


	// SchedulingStartedBy - User that started the schedule run
	SchedulingStartedBy *Userreference `json:"schedulingStartedBy,omitempty"`


	// SchedulingCanceledBy - User that canceled the schedule run
	SchedulingCanceledBy *Userreference `json:"schedulingCanceledBy,omitempty"`


	// SchedulingCompletedTime - Time at which the scheduling run was completed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	SchedulingCompletedTime *time.Time `json:"schedulingCompletedTime,omitempty"`


	// ReschedulingOptions - The selected options for the reschedule request. Will always be null if intradayRescheduling is false
	ReschedulingOptions *Reschedulingoptionsresponse `json:"reschedulingOptions,omitempty"`


	// ReschedulingResultExpiration - When the rescheduling result data will expire. Results are kept temporarily as they should be applied as soon as possible after the run finishes.  Will always be null if intradayRescheduling is false. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	ReschedulingResultExpiration *time.Time `json:"reschedulingResultExpiration,omitempty"`


	// Applied - Whether the rescheduling run has been marked applied
	Applied *bool `json:"applied,omitempty"`


	// UnscheduledAgents - Agents that were not scheduled in the rescheduling operation. Will always be null if intradayRescheduling is false
	UnscheduledAgents *[]Unscheduledagentwarning `json:"unscheduledAgents,omitempty"`

}

// String returns a JSON representation of the model
func (o *Schedulingrunresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

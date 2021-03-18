package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Buschedulerun
type Buschedulerun struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// SchedulerRunId - The scheduler run ID.  Reference this value for support
	SchedulerRunId *string `json:"schedulerRunId,omitempty"`


	// IntradayRescheduling - Whether this is an intraday rescheduling run
	IntradayRescheduling *bool `json:"intradayRescheduling,omitempty"`


	// State - The state of the generation run
	State *string `json:"state,omitempty"`


	// WeekCount - The number of weeks spanned by the schedule
	WeekCount *int `json:"weekCount,omitempty"`


	// PercentComplete - Percent completion of the schedule run
	PercentComplete *float64 `json:"percentComplete,omitempty"`


	// TargetWeek - The start date of the target week. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	TargetWeek *time.Time `json:"targetWeek,omitempty"`


	// Schedule - The generated schedule.  Null unless the schedule run is complete
	Schedule *Buschedulereference `json:"schedule,omitempty"`


	// ScheduleDescription - The description of the generated schedule
	ScheduleDescription *string `json:"scheduleDescription,omitempty"`


	// SchedulingStartTime - When the schedule generation run started. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	SchedulingStartTime *time.Time `json:"schedulingStartTime,omitempty"`


	// SchedulingStartedBy - The user who started the scheduling run
	SchedulingStartedBy *Userreference `json:"schedulingStartedBy,omitempty"`


	// SchedulingCanceledBy - The user who canceled the scheduling run, if applicable
	SchedulingCanceledBy *Userreference `json:"schedulingCanceledBy,omitempty"`


	// SchedulingCompletedTime - When the scheduling run was completed, if applicable. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	SchedulingCompletedTime *time.Time `json:"schedulingCompletedTime,omitempty"`


	// MessageCount - The number of schedule generation messages for this schedule generation run
	MessageCount *int `json:"messageCount,omitempty"`


	// ReschedulingOptions - Rescheduling options for this run.  Null unless intradayRescheduling is true
	ReschedulingOptions *Reschedulingoptionsrunresponse `json:"reschedulingOptions,omitempty"`


	// ReschedulingResultExpiration - When the reschedule result will expire.  Null unless intradayRescheduling is true. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ReschedulingResultExpiration *time.Time `json:"reschedulingResultExpiration,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buschedulerun) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

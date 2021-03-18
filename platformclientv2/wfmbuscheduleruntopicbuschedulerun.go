package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuscheduleruntopicbuschedulerun
type Wfmbuscheduleruntopicbuschedulerun struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// PercentComplete
	PercentComplete *float32 `json:"percentComplete,omitempty"`


	// IntradayRescheduling
	IntradayRescheduling *bool `json:"intradayRescheduling,omitempty"`


	// State
	State *string `json:"state,omitempty"`


	// WeekCount
	WeekCount *int `json:"weekCount,omitempty"`


	// Schedule
	Schedule *Wfmbuscheduleruntopicbuschedulereference `json:"schedule,omitempty"`


	// SchedulingCanceledBy
	SchedulingCanceledBy *Wfmbuscheduleruntopicuserreference `json:"schedulingCanceledBy,omitempty"`


	// SchedulingCompletedTime
	SchedulingCompletedTime *string `json:"schedulingCompletedTime,omitempty"`


	// MessageCount
	MessageCount *int `json:"messageCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbuscheduleruntopicbuschedulerun) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

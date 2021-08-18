package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Wfmbuscheduleruntopicbuschedulerun) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuscheduleruntopicbuschedulerun

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		PercentComplete *float32 `json:"percentComplete,omitempty"`
		
		IntradayRescheduling *bool `json:"intradayRescheduling,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		WeekCount *int `json:"weekCount,omitempty"`
		
		Schedule *Wfmbuscheduleruntopicbuschedulereference `json:"schedule,omitempty"`
		
		SchedulingCanceledBy *Wfmbuscheduleruntopicuserreference `json:"schedulingCanceledBy,omitempty"`
		
		SchedulingCompletedTime *string `json:"schedulingCompletedTime,omitempty"`
		
		MessageCount *int `json:"messageCount,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		PercentComplete: u.PercentComplete,
		
		IntradayRescheduling: u.IntradayRescheduling,
		
		State: u.State,
		
		WeekCount: u.WeekCount,
		
		Schedule: u.Schedule,
		
		SchedulingCanceledBy: u.SchedulingCanceledBy,
		
		SchedulingCompletedTime: u.SchedulingCompletedTime,
		
		MessageCount: u.MessageCount,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmbuscheduleruntopicbuschedulerun) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

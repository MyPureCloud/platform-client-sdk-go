package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bureschedulerequest
type Bureschedulerequest struct { 
	// StartDate - The start of the range to reschedule.  Defaults to the beginning of the schedule. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - The end of the range to reschedule.  Defaults the the end of the schedule. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
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

func (u *Bureschedulerequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bureschedulerequest

	
	StartDate := new(string)
	if u.StartDate != nil {
		
		*StartDate = timeutil.Strftime(u.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if u.EndDate != nil {
		
		*EndDate = timeutil.Strftime(u.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndDate = nil
	}
	

	return json.Marshal(&struct { 
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		AgentIds *[]string `json:"agentIds,omitempty"`
		
		ActivityCodeIds *[]string `json:"activityCodeIds,omitempty"`
		
		ManagementUnitIds *[]string `json:"managementUnitIds,omitempty"`
		
		DoNotChangeWeeklyPaidTime *bool `json:"doNotChangeWeeklyPaidTime,omitempty"`
		
		DoNotChangeDailyPaidTime *bool `json:"doNotChangeDailyPaidTime,omitempty"`
		
		DoNotChangeShiftStartTimes *bool `json:"doNotChangeShiftStartTimes,omitempty"`
		
		DoNotChangeManuallyEditedShifts *bool `json:"doNotChangeManuallyEditedShifts,omitempty"`
		*Alias
	}{ 
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		AgentIds: u.AgentIds,
		
		ActivityCodeIds: u.ActivityCodeIds,
		
		ManagementUnitIds: u.ManagementUnitIds,
		
		DoNotChangeWeeklyPaidTime: u.DoNotChangeWeeklyPaidTime,
		
		DoNotChangeDailyPaidTime: u.DoNotChangeDailyPaidTime,
		
		DoNotChangeShiftStartTimes: u.DoNotChangeShiftStartTimes,
		
		DoNotChangeManuallyEditedShifts: u.DoNotChangeManuallyEditedShifts,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Bureschedulerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

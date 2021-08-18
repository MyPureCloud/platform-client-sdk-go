package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Reschedulingoptionsrunresponse
type Reschedulingoptionsrunresponse struct { 
	// ExistingSchedule - The existing schedule to which this reschedule run applies
	ExistingSchedule *Buschedulereference `json:"existingSchedule,omitempty"`


	// StartDate - The start date of the period to reschedule. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - The end date of the period to reschedule. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndDate *time.Time `json:"endDate,omitempty"`


	// ManagementUnits - Per-management unit rescheduling options
	ManagementUnits *[]Reschedulingmanagementunitresponse `json:"managementUnits,omitempty"`


	// AgentCount - The number of agents to be considered in the reschedule
	AgentCount *int `json:"agentCount,omitempty"`


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

func (u *Reschedulingoptionsrunresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reschedulingoptionsrunresponse

	
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
		ExistingSchedule *Buschedulereference `json:"existingSchedule,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		ManagementUnits *[]Reschedulingmanagementunitresponse `json:"managementUnits,omitempty"`
		
		AgentCount *int `json:"agentCount,omitempty"`
		
		ActivityCodeIds *[]string `json:"activityCodeIds,omitempty"`
		
		DoNotChangeWeeklyPaidTime *bool `json:"doNotChangeWeeklyPaidTime,omitempty"`
		
		DoNotChangeDailyPaidTime *bool `json:"doNotChangeDailyPaidTime,omitempty"`
		
		DoNotChangeShiftStartTimes *bool `json:"doNotChangeShiftStartTimes,omitempty"`
		
		DoNotChangeManuallyEditedShifts *bool `json:"doNotChangeManuallyEditedShifts,omitempty"`
		*Alias
	}{ 
		ExistingSchedule: u.ExistingSchedule,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		ManagementUnits: u.ManagementUnits,
		
		AgentCount: u.AgentCount,
		
		ActivityCodeIds: u.ActivityCodeIds,
		
		DoNotChangeWeeklyPaidTime: u.DoNotChangeWeeklyPaidTime,
		
		DoNotChangeDailyPaidTime: u.DoNotChangeDailyPaidTime,
		
		DoNotChangeShiftStartTimes: u.DoNotChangeShiftStartTimes,
		
		DoNotChangeManuallyEditedShifts: u.DoNotChangeManuallyEditedShifts,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Reschedulingoptionsrunresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

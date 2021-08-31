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

func (o *Bureschedulerequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bureschedulerequest
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if o.EndDate != nil {
		
		*EndDate = timeutil.Strftime(o.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		
		AgentIds: o.AgentIds,
		
		ActivityCodeIds: o.ActivityCodeIds,
		
		ManagementUnitIds: o.ManagementUnitIds,
		
		DoNotChangeWeeklyPaidTime: o.DoNotChangeWeeklyPaidTime,
		
		DoNotChangeDailyPaidTime: o.DoNotChangeDailyPaidTime,
		
		DoNotChangeShiftStartTimes: o.DoNotChangeShiftStartTimes,
		
		DoNotChangeManuallyEditedShifts: o.DoNotChangeManuallyEditedShifts,
		Alias:    (*Alias)(o),
	})
}

func (o *Bureschedulerequest) UnmarshalJSON(b []byte) error {
	var BureschedulerequestMap map[string]interface{}
	err := json.Unmarshal(b, &BureschedulerequestMap)
	if err != nil {
		return err
	}
	
	if startDateString, ok := BureschedulerequestMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := BureschedulerequestMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if AgentIds, ok := BureschedulerequestMap["agentIds"].([]interface{}); ok {
		AgentIdsString, _ := json.Marshal(AgentIds)
		json.Unmarshal(AgentIdsString, &o.AgentIds)
	}
	
	if ActivityCodeIds, ok := BureschedulerequestMap["activityCodeIds"].([]interface{}); ok {
		ActivityCodeIdsString, _ := json.Marshal(ActivityCodeIds)
		json.Unmarshal(ActivityCodeIdsString, &o.ActivityCodeIds)
	}
	
	if ManagementUnitIds, ok := BureschedulerequestMap["managementUnitIds"].([]interface{}); ok {
		ManagementUnitIdsString, _ := json.Marshal(ManagementUnitIds)
		json.Unmarshal(ManagementUnitIdsString, &o.ManagementUnitIds)
	}
	
	if DoNotChangeWeeklyPaidTime, ok := BureschedulerequestMap["doNotChangeWeeklyPaidTime"].(bool); ok {
		o.DoNotChangeWeeklyPaidTime = &DoNotChangeWeeklyPaidTime
	}
	
	if DoNotChangeDailyPaidTime, ok := BureschedulerequestMap["doNotChangeDailyPaidTime"].(bool); ok {
		o.DoNotChangeDailyPaidTime = &DoNotChangeDailyPaidTime
	}
	
	if DoNotChangeShiftStartTimes, ok := BureschedulerequestMap["doNotChangeShiftStartTimes"].(bool); ok {
		o.DoNotChangeShiftStartTimes = &DoNotChangeShiftStartTimes
	}
	
	if DoNotChangeManuallyEditedShifts, ok := BureschedulerequestMap["doNotChangeManuallyEditedShifts"].(bool); ok {
		o.DoNotChangeManuallyEditedShifts = &DoNotChangeManuallyEditedShifts
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bureschedulerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

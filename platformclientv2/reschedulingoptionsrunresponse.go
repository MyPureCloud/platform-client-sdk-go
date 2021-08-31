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

func (o *Reschedulingoptionsrunresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reschedulingoptionsrunresponse
	
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
		ExistingSchedule: o.ExistingSchedule,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		ManagementUnits: o.ManagementUnits,
		
		AgentCount: o.AgentCount,
		
		ActivityCodeIds: o.ActivityCodeIds,
		
		DoNotChangeWeeklyPaidTime: o.DoNotChangeWeeklyPaidTime,
		
		DoNotChangeDailyPaidTime: o.DoNotChangeDailyPaidTime,
		
		DoNotChangeShiftStartTimes: o.DoNotChangeShiftStartTimes,
		
		DoNotChangeManuallyEditedShifts: o.DoNotChangeManuallyEditedShifts,
		Alias:    (*Alias)(o),
	})
}

func (o *Reschedulingoptionsrunresponse) UnmarshalJSON(b []byte) error {
	var ReschedulingoptionsrunresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ReschedulingoptionsrunresponseMap)
	if err != nil {
		return err
	}
	
	if ExistingSchedule, ok := ReschedulingoptionsrunresponseMap["existingSchedule"].(map[string]interface{}); ok {
		ExistingScheduleString, _ := json.Marshal(ExistingSchedule)
		json.Unmarshal(ExistingScheduleString, &o.ExistingSchedule)
	}
	
	if startDateString, ok := ReschedulingoptionsrunresponseMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := ReschedulingoptionsrunresponseMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if ManagementUnits, ok := ReschedulingoptionsrunresponseMap["managementUnits"].([]interface{}); ok {
		ManagementUnitsString, _ := json.Marshal(ManagementUnits)
		json.Unmarshal(ManagementUnitsString, &o.ManagementUnits)
	}
	
	if AgentCount, ok := ReschedulingoptionsrunresponseMap["agentCount"].(float64); ok {
		AgentCountInt := int(AgentCount)
		o.AgentCount = &AgentCountInt
	}
	
	if ActivityCodeIds, ok := ReschedulingoptionsrunresponseMap["activityCodeIds"].([]interface{}); ok {
		ActivityCodeIdsString, _ := json.Marshal(ActivityCodeIds)
		json.Unmarshal(ActivityCodeIdsString, &o.ActivityCodeIds)
	}
	
	if DoNotChangeWeeklyPaidTime, ok := ReschedulingoptionsrunresponseMap["doNotChangeWeeklyPaidTime"].(bool); ok {
		o.DoNotChangeWeeklyPaidTime = &DoNotChangeWeeklyPaidTime
	}
	
	if DoNotChangeDailyPaidTime, ok := ReschedulingoptionsrunresponseMap["doNotChangeDailyPaidTime"].(bool); ok {
		o.DoNotChangeDailyPaidTime = &DoNotChangeDailyPaidTime
	}
	
	if DoNotChangeShiftStartTimes, ok := ReschedulingoptionsrunresponseMap["doNotChangeShiftStartTimes"].(bool); ok {
		o.DoNotChangeShiftStartTimes = &DoNotChangeShiftStartTimes
	}
	
	if DoNotChangeManuallyEditedShifts, ok := ReschedulingoptionsrunresponseMap["doNotChangeManuallyEditedShifts"].(bool); ok {
		o.DoNotChangeManuallyEditedShifts = &DoNotChangeManuallyEditedShifts
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Reschedulingoptionsrunresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

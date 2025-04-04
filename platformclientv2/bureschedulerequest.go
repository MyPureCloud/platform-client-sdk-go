package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Bureschedulerequest
type Bureschedulerequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

	// ActivitySmoothingType - Overrides the default BU level activity smoothing type for this reschedule run
	ActivitySmoothingType *string `json:"activitySmoothingType,omitempty"`

	// InduceScheduleVariability - Overrides the default BU level induce schedule variability setting for this reschedule run
	InduceScheduleVariability *bool `json:"induceScheduleVariability,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Bureschedulerequest) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Bureschedulerequest) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartDate","EndDate", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

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
		
		ActivitySmoothingType *string `json:"activitySmoothingType,omitempty"`
		
		InduceScheduleVariability *bool `json:"induceScheduleVariability,omitempty"`
		Alias
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
		
		ActivitySmoothingType: o.ActivitySmoothingType,
		
		InduceScheduleVariability: o.InduceScheduleVariability,
		Alias:    (Alias)(o),
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
    
	if ActivitySmoothingType, ok := BureschedulerequestMap["activitySmoothingType"].(string); ok {
		o.ActivitySmoothingType = &ActivitySmoothingType
	}
    
	if InduceScheduleVariability, ok := BureschedulerequestMap["induceScheduleVariability"].(bool); ok {
		o.InduceScheduleVariability = &InduceScheduleVariability
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Bureschedulerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

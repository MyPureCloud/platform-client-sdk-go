package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Reschedulingoptionsrunresponse
type Reschedulingoptionsrunresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Reschedulingoptionsrunresponse) SetField(field string, fieldValue interface{}) {
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

func (o Reschedulingoptionsrunresponse) MarshalJSON() ([]byte, error) {
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
		Alias
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
		Alias:    (Alias)(o),
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

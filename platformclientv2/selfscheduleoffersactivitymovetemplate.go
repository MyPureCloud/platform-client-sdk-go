package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Selfscheduleoffersactivitymovetemplate
type Selfscheduleoffersactivitymovetemplate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Schedule - The schedule associated with the shift the agent wants to edit.
	Schedule *Buschedulereference `json:"schedule,omitempty"`

	// ShiftId - The ID of the shift to which the self-schedule activity move is associated
	ShiftId *string `json:"shiftId,omitempty"`

	// ReferenceKey - The unique identifier for the offered shift edit that was used for this shift edit
	ReferenceKey *string `json:"referenceKey,omitempty"`

	// StartDate - The start date/time for the shift in ISO-8601 format
	StartDate *time.Time `json:"startDate,omitempty"`

	// LengthMinutes - The duration of this shift in minutes
	LengthMinutes *int `json:"lengthMinutes,omitempty"`

	// ActivityCode - The activity related to this shift
	ActivityCode *Activitycodereference `json:"activityCode,omitempty"`

	// AlternateStartDates - The alternative start dates for this shift
	AlternateStartDates *[]time.Time `json:"alternateStartDates,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Selfscheduleoffersactivitymovetemplate) SetField(field string, fieldValue interface{}) {
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

func (o Selfscheduleoffersactivitymovetemplate) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartDate", }
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
	type Alias Selfscheduleoffersactivitymovetemplate
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		Schedule *Buschedulereference `json:"schedule,omitempty"`
		
		ShiftId *string `json:"shiftId,omitempty"`
		
		ReferenceKey *string `json:"referenceKey,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		ActivityCode *Activitycodereference `json:"activityCode,omitempty"`
		
		AlternateStartDates *[]time.Time `json:"alternateStartDates,omitempty"`
		Alias
	}{ 
		Schedule: o.Schedule,
		
		ShiftId: o.ShiftId,
		
		ReferenceKey: o.ReferenceKey,
		
		StartDate: StartDate,
		
		LengthMinutes: o.LengthMinutes,
		
		ActivityCode: o.ActivityCode,
		
		AlternateStartDates: o.AlternateStartDates,
		Alias:    (Alias)(o),
	})
}

func (o *Selfscheduleoffersactivitymovetemplate) UnmarshalJSON(b []byte) error {
	var SelfscheduleoffersactivitymovetemplateMap map[string]interface{}
	err := json.Unmarshal(b, &SelfscheduleoffersactivitymovetemplateMap)
	if err != nil {
		return err
	}
	
	if Schedule, ok := SelfscheduleoffersactivitymovetemplateMap["schedule"].(map[string]interface{}); ok {
		ScheduleString, _ := json.Marshal(Schedule)
		json.Unmarshal(ScheduleString, &o.Schedule)
	}
	
	if ShiftId, ok := SelfscheduleoffersactivitymovetemplateMap["shiftId"].(string); ok {
		o.ShiftId = &ShiftId
	}
    
	if ReferenceKey, ok := SelfscheduleoffersactivitymovetemplateMap["referenceKey"].(string); ok {
		o.ReferenceKey = &ReferenceKey
	}
    
	if startDateString, ok := SelfscheduleoffersactivitymovetemplateMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if LengthMinutes, ok := SelfscheduleoffersactivitymovetemplateMap["lengthMinutes"].(float64); ok {
		LengthMinutesInt := int(LengthMinutes)
		o.LengthMinutes = &LengthMinutesInt
	}
	
	if ActivityCode, ok := SelfscheduleoffersactivitymovetemplateMap["activityCode"].(map[string]interface{}); ok {
		ActivityCodeString, _ := json.Marshal(ActivityCode)
		json.Unmarshal(ActivityCodeString, &o.ActivityCode)
	}
	
	if AlternateStartDates, ok := SelfscheduleoffersactivitymovetemplateMap["alternateStartDates"].([]interface{}); ok {
		AlternateStartDatesString, _ := json.Marshal(AlternateStartDates)
		json.Unmarshal(AlternateStartDatesString, &o.AlternateStartDates)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Selfscheduleoffersactivitymovetemplate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

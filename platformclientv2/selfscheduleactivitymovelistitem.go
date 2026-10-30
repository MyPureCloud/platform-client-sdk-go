package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Selfscheduleactivitymovelistitem
type Selfscheduleactivitymovelistitem struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Schedule - The schedule to which the shift and activity are associated
	Schedule *Buschedulereference `json:"schedule,omitempty"`

	// ShiftId - The ID of the shift to which this activity is associated
	ShiftId *string `json:"shiftId,omitempty"`

	// State - The state of the self-schedule operation
	State *string `json:"state,omitempty"`

	// WorkflowStatus - The processing status of this self-schedule operation
	WorkflowStatus *string `json:"workflowStatus,omitempty"`

	// StartDate - The original start date of this activity in ISO-8601 format
	StartDate *time.Time `json:"startDate,omitempty"`

	// NewDate - The new start date of this activity in ISO-8601 format
	NewDate *time.Time `json:"newDate,omitempty"`

	// LengthMinutes - The length of this activity in minutes
	LengthMinutes *int `json:"lengthMinutes,omitempty"`

	// ActivityCode - The activity code related to this shift
	ActivityCode *Activitycodereference `json:"activityCode,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Selfscheduleactivitymovelistitem) SetField(field string, fieldValue interface{}) {
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

func (o Selfscheduleactivitymovelistitem) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartDate","NewDate", }
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
	type Alias Selfscheduleactivitymovelistitem
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	NewDate := new(string)
	if o.NewDate != nil {
		
		*NewDate = timeutil.Strftime(o.NewDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		NewDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Schedule *Buschedulereference `json:"schedule,omitempty"`
		
		ShiftId *string `json:"shiftId,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		WorkflowStatus *string `json:"workflowStatus,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		NewDate *string `json:"newDate,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		ActivityCode *Activitycodereference `json:"activityCode,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Schedule: o.Schedule,
		
		ShiftId: o.ShiftId,
		
		State: o.State,
		
		WorkflowStatus: o.WorkflowStatus,
		
		StartDate: StartDate,
		
		NewDate: NewDate,
		
		LengthMinutes: o.LengthMinutes,
		
		ActivityCode: o.ActivityCode,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Selfscheduleactivitymovelistitem) UnmarshalJSON(b []byte) error {
	var SelfscheduleactivitymovelistitemMap map[string]interface{}
	err := json.Unmarshal(b, &SelfscheduleactivitymovelistitemMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SelfscheduleactivitymovelistitemMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Schedule, ok := SelfscheduleactivitymovelistitemMap["schedule"].(map[string]interface{}); ok {
		ScheduleString, _ := json.Marshal(Schedule)
		json.Unmarshal(ScheduleString, &o.Schedule)
	}
	
	if ShiftId, ok := SelfscheduleactivitymovelistitemMap["shiftId"].(string); ok {
		o.ShiftId = &ShiftId
	}
    
	if State, ok := SelfscheduleactivitymovelistitemMap["state"].(string); ok {
		o.State = &State
	}
    
	if WorkflowStatus, ok := SelfscheduleactivitymovelistitemMap["workflowStatus"].(string); ok {
		o.WorkflowStatus = &WorkflowStatus
	}
    
	if startDateString, ok := SelfscheduleactivitymovelistitemMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if newDateString, ok := SelfscheduleactivitymovelistitemMap["newDate"].(string); ok {
		NewDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", newDateString)
		o.NewDate = &NewDate
	}
	
	if LengthMinutes, ok := SelfscheduleactivitymovelistitemMap["lengthMinutes"].(float64); ok {
		LengthMinutesInt := int(LengthMinutes)
		o.LengthMinutes = &LengthMinutesInt
	}
	
	if ActivityCode, ok := SelfscheduleactivitymovelistitemMap["activityCode"].(map[string]interface{}); ok {
		ActivityCodeString, _ := json.Marshal(ActivityCode)
		json.Unmarshal(ActivityCodeString, &o.ActivityCode)
	}
	
	if SelfUri, ok := SelfscheduleactivitymovelistitemMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Selfscheduleactivitymovelistitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

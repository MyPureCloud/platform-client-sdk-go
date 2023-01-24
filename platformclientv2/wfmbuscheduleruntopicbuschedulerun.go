package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuscheduleruntopicbuschedulerun
type Wfmbuscheduleruntopicbuschedulerun struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

	// MessageSeverityCounts
	MessageSeverityCounts *[]Wfmbuscheduleruntopicschedulermessageseveritycount `json:"messageSeverityCounts,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmbuscheduleruntopicbuschedulerun) SetField(field string, fieldValue interface{}) {
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

func (o Wfmbuscheduleruntopicbuschedulerun) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if contains(dateTimeFields, fieldName) {
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
		
		MessageSeverityCounts *[]Wfmbuscheduleruntopicschedulermessageseveritycount `json:"messageSeverityCounts,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		PercentComplete: o.PercentComplete,
		
		IntradayRescheduling: o.IntradayRescheduling,
		
		State: o.State,
		
		WeekCount: o.WeekCount,
		
		Schedule: o.Schedule,
		
		SchedulingCanceledBy: o.SchedulingCanceledBy,
		
		SchedulingCompletedTime: o.SchedulingCompletedTime,
		
		MessageCount: o.MessageCount,
		
		MessageSeverityCounts: o.MessageSeverityCounts,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmbuscheduleruntopicbuschedulerun) UnmarshalJSON(b []byte) error {
	var WfmbuscheduleruntopicbuschedulerunMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbuscheduleruntopicbuschedulerunMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmbuscheduleruntopicbuschedulerunMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if PercentComplete, ok := WfmbuscheduleruntopicbuschedulerunMap["percentComplete"].(float64); ok {
		PercentCompleteFloat32 := float32(PercentComplete)
		o.PercentComplete = &PercentCompleteFloat32
	}
    
	if IntradayRescheduling, ok := WfmbuscheduleruntopicbuschedulerunMap["intradayRescheduling"].(bool); ok {
		o.IntradayRescheduling = &IntradayRescheduling
	}
    
	if State, ok := WfmbuscheduleruntopicbuschedulerunMap["state"].(string); ok {
		o.State = &State
	}
    
	if WeekCount, ok := WfmbuscheduleruntopicbuschedulerunMap["weekCount"].(float64); ok {
		WeekCountInt := int(WeekCount)
		o.WeekCount = &WeekCountInt
	}
	
	if Schedule, ok := WfmbuscheduleruntopicbuschedulerunMap["schedule"].(map[string]interface{}); ok {
		ScheduleString, _ := json.Marshal(Schedule)
		json.Unmarshal(ScheduleString, &o.Schedule)
	}
	
	if SchedulingCanceledBy, ok := WfmbuscheduleruntopicbuschedulerunMap["schedulingCanceledBy"].(map[string]interface{}); ok {
		SchedulingCanceledByString, _ := json.Marshal(SchedulingCanceledBy)
		json.Unmarshal(SchedulingCanceledByString, &o.SchedulingCanceledBy)
	}
	
	if SchedulingCompletedTime, ok := WfmbuscheduleruntopicbuschedulerunMap["schedulingCompletedTime"].(string); ok {
		o.SchedulingCompletedTime = &SchedulingCompletedTime
	}
    
	if MessageCount, ok := WfmbuscheduleruntopicbuschedulerunMap["messageCount"].(float64); ok {
		MessageCountInt := int(MessageCount)
		o.MessageCount = &MessageCountInt
	}
	
	if MessageSeverityCounts, ok := WfmbuscheduleruntopicbuschedulerunMap["messageSeverityCounts"].([]interface{}); ok {
		MessageSeverityCountsString, _ := json.Marshal(MessageSeverityCounts)
		json.Unmarshal(MessageSeverityCountsString, &o.MessageSeverityCounts)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbuscheduleruntopicbuschedulerun) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

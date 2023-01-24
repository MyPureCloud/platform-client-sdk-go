package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Buschedulerun
type Buschedulerun struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// SchedulerRunId - The scheduler run ID.  Reference this value for support
	SchedulerRunId *string `json:"schedulerRunId,omitempty"`

	// IntradayRescheduling - Whether this is an intraday rescheduling run
	IntradayRescheduling *bool `json:"intradayRescheduling,omitempty"`

	// State - The state of the generation run
	State *string `json:"state,omitempty"`

	// WeekCount - The number of weeks spanned by the schedule
	WeekCount *int `json:"weekCount,omitempty"`

	// PercentComplete - Percent completion of the schedule run
	PercentComplete *float64 `json:"percentComplete,omitempty"`

	// TargetWeek - The start date of the target week. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	TargetWeek *time.Time `json:"targetWeek,omitempty"`

	// Schedule - The generated schedule.  Null unless the schedule run is complete
	Schedule *Buschedulereference `json:"schedule,omitempty"`

	// ScheduleDescription - The description of the generated schedule
	ScheduleDescription *string `json:"scheduleDescription,omitempty"`

	// SchedulingStartTime - When the schedule generation run started. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	SchedulingStartTime *time.Time `json:"schedulingStartTime,omitempty"`

	// SchedulingStartedBy - The user who started the scheduling run
	SchedulingStartedBy *Userreference `json:"schedulingStartedBy,omitempty"`

	// SchedulingCanceledBy - The user who canceled the scheduling run, if applicable
	SchedulingCanceledBy *Userreference `json:"schedulingCanceledBy,omitempty"`

	// SchedulingCompletedTime - When the scheduling run was completed, if applicable. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	SchedulingCompletedTime *time.Time `json:"schedulingCompletedTime,omitempty"`

	// MessageCount - The number of schedule generation messages for this schedule generation run
	MessageCount *int `json:"messageCount,omitempty"`

	// MessageSeverityCounts - The list of schedule generation message counts by severity for this schedule generation run
	MessageSeverityCounts *[]Schedulermessageseveritycount `json:"messageSeverityCounts,omitempty"`

	// ReschedulingOptions - Rescheduling options for this run.  Null unless intradayRescheduling is true
	ReschedulingOptions *Reschedulingoptionsrunresponse `json:"reschedulingOptions,omitempty"`

	// ReschedulingResultExpiration - When the reschedule result will expire.  Null unless intradayRescheduling is true. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ReschedulingResultExpiration *time.Time `json:"reschedulingResultExpiration,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Buschedulerun) SetField(field string, fieldValue interface{}) {
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

func (o Buschedulerun) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "SchedulingStartTime","SchedulingCompletedTime","ReschedulingResultExpiration", }
		localDateTimeFields := []string{  }
		dateFields := []string{ "TargetWeek", }

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
	type Alias Buschedulerun
	
	TargetWeek := new(string)
	if o.TargetWeek != nil {
		*TargetWeek = timeutil.Strftime(o.TargetWeek, "%Y-%m-%d")
	} else {
		TargetWeek = nil
	}
	
	SchedulingStartTime := new(string)
	if o.SchedulingStartTime != nil {
		
		*SchedulingStartTime = timeutil.Strftime(o.SchedulingStartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		SchedulingStartTime = nil
	}
	
	SchedulingCompletedTime := new(string)
	if o.SchedulingCompletedTime != nil {
		
		*SchedulingCompletedTime = timeutil.Strftime(o.SchedulingCompletedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		SchedulingCompletedTime = nil
	}
	
	ReschedulingResultExpiration := new(string)
	if o.ReschedulingResultExpiration != nil {
		
		*ReschedulingResultExpiration = timeutil.Strftime(o.ReschedulingResultExpiration, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReschedulingResultExpiration = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SchedulerRunId *string `json:"schedulerRunId,omitempty"`
		
		IntradayRescheduling *bool `json:"intradayRescheduling,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		WeekCount *int `json:"weekCount,omitempty"`
		
		PercentComplete *float64 `json:"percentComplete,omitempty"`
		
		TargetWeek *string `json:"targetWeek,omitempty"`
		
		Schedule *Buschedulereference `json:"schedule,omitempty"`
		
		ScheduleDescription *string `json:"scheduleDescription,omitempty"`
		
		SchedulingStartTime *string `json:"schedulingStartTime,omitempty"`
		
		SchedulingStartedBy *Userreference `json:"schedulingStartedBy,omitempty"`
		
		SchedulingCanceledBy *Userreference `json:"schedulingCanceledBy,omitempty"`
		
		SchedulingCompletedTime *string `json:"schedulingCompletedTime,omitempty"`
		
		MessageCount *int `json:"messageCount,omitempty"`
		
		MessageSeverityCounts *[]Schedulermessageseveritycount `json:"messageSeverityCounts,omitempty"`
		
		ReschedulingOptions *Reschedulingoptionsrunresponse `json:"reschedulingOptions,omitempty"`
		
		ReschedulingResultExpiration *string `json:"reschedulingResultExpiration,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		SchedulerRunId: o.SchedulerRunId,
		
		IntradayRescheduling: o.IntradayRescheduling,
		
		State: o.State,
		
		WeekCount: o.WeekCount,
		
		PercentComplete: o.PercentComplete,
		
		TargetWeek: TargetWeek,
		
		Schedule: o.Schedule,
		
		ScheduleDescription: o.ScheduleDescription,
		
		SchedulingStartTime: SchedulingStartTime,
		
		SchedulingStartedBy: o.SchedulingStartedBy,
		
		SchedulingCanceledBy: o.SchedulingCanceledBy,
		
		SchedulingCompletedTime: SchedulingCompletedTime,
		
		MessageCount: o.MessageCount,
		
		MessageSeverityCounts: o.MessageSeverityCounts,
		
		ReschedulingOptions: o.ReschedulingOptions,
		
		ReschedulingResultExpiration: ReschedulingResultExpiration,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Buschedulerun) UnmarshalJSON(b []byte) error {
	var BuschedulerunMap map[string]interface{}
	err := json.Unmarshal(b, &BuschedulerunMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BuschedulerunMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if SchedulerRunId, ok := BuschedulerunMap["schedulerRunId"].(string); ok {
		o.SchedulerRunId = &SchedulerRunId
	}
    
	if IntradayRescheduling, ok := BuschedulerunMap["intradayRescheduling"].(bool); ok {
		o.IntradayRescheduling = &IntradayRescheduling
	}
    
	if State, ok := BuschedulerunMap["state"].(string); ok {
		o.State = &State
	}
    
	if WeekCount, ok := BuschedulerunMap["weekCount"].(float64); ok {
		WeekCountInt := int(WeekCount)
		o.WeekCount = &WeekCountInt
	}
	
	if PercentComplete, ok := BuschedulerunMap["percentComplete"].(float64); ok {
		o.PercentComplete = &PercentComplete
	}
    
	if targetWeekString, ok := BuschedulerunMap["targetWeek"].(string); ok {
		TargetWeek, _ := time.Parse("2006-01-02", targetWeekString)
		o.TargetWeek = &TargetWeek
	}
	
	if Schedule, ok := BuschedulerunMap["schedule"].(map[string]interface{}); ok {
		ScheduleString, _ := json.Marshal(Schedule)
		json.Unmarshal(ScheduleString, &o.Schedule)
	}
	
	if ScheduleDescription, ok := BuschedulerunMap["scheduleDescription"].(string); ok {
		o.ScheduleDescription = &ScheduleDescription
	}
    
	if schedulingStartTimeString, ok := BuschedulerunMap["schedulingStartTime"].(string); ok {
		SchedulingStartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", schedulingStartTimeString)
		o.SchedulingStartTime = &SchedulingStartTime
	}
	
	if SchedulingStartedBy, ok := BuschedulerunMap["schedulingStartedBy"].(map[string]interface{}); ok {
		SchedulingStartedByString, _ := json.Marshal(SchedulingStartedBy)
		json.Unmarshal(SchedulingStartedByString, &o.SchedulingStartedBy)
	}
	
	if SchedulingCanceledBy, ok := BuschedulerunMap["schedulingCanceledBy"].(map[string]interface{}); ok {
		SchedulingCanceledByString, _ := json.Marshal(SchedulingCanceledBy)
		json.Unmarshal(SchedulingCanceledByString, &o.SchedulingCanceledBy)
	}
	
	if schedulingCompletedTimeString, ok := BuschedulerunMap["schedulingCompletedTime"].(string); ok {
		SchedulingCompletedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", schedulingCompletedTimeString)
		o.SchedulingCompletedTime = &SchedulingCompletedTime
	}
	
	if MessageCount, ok := BuschedulerunMap["messageCount"].(float64); ok {
		MessageCountInt := int(MessageCount)
		o.MessageCount = &MessageCountInt
	}
	
	if MessageSeverityCounts, ok := BuschedulerunMap["messageSeverityCounts"].([]interface{}); ok {
		MessageSeverityCountsString, _ := json.Marshal(MessageSeverityCounts)
		json.Unmarshal(MessageSeverityCountsString, &o.MessageSeverityCounts)
	}
	
	if ReschedulingOptions, ok := BuschedulerunMap["reschedulingOptions"].(map[string]interface{}); ok {
		ReschedulingOptionsString, _ := json.Marshal(ReschedulingOptions)
		json.Unmarshal(ReschedulingOptionsString, &o.ReschedulingOptions)
	}
	
	if reschedulingResultExpirationString, ok := BuschedulerunMap["reschedulingResultExpiration"].(string); ok {
		ReschedulingResultExpiration, _ := time.Parse("2006-01-02T15:04:05.999999Z", reschedulingResultExpirationString)
		o.ReschedulingResultExpiration = &ReschedulingResultExpiration
	}
	
	if SelfUri, ok := BuschedulerunMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Buschedulerun) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

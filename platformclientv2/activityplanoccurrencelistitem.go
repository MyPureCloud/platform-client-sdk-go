package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Activityplanoccurrencelistitem
type Activityplanoccurrencelistitem struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// RunDate - The date the associated activity plan was run, in ISO-8601 format
	RunDate *time.Time `json:"runDate,omitempty"`

	// SessionDateRange - The dates potentially spanned by the sessions in this occurrence, even if no sessions currently exist
	SessionDateRange *Requireddaterange `json:"sessionDateRange,omitempty"`

	// Facilitated - Whether the run sessions are facilitated
	Facilitated *bool `json:"facilitated,omitempty"`

	// SessionCount - The number of sessions that were scheduled
	SessionCount *int `json:"sessionCount,omitempty"`

	// ScheduledAttendeesCount - The number of attendees that were scheduled for this occurrence
	ScheduledAttendeesCount *int `json:"scheduledAttendeesCount,omitempty"`

	// UnscheduledAttendeesCount - The number of attendees that were unable to be scheduled for this occurrence
	UnscheduledAttendeesCount *int `json:"unscheduledAttendeesCount,omitempty"`

	// Edited - Whether this occurrence is edited
	Edited *bool `json:"edited,omitempty"`

	// ViolationCount - The number of violations for this occurrence
	ViolationCount *int `json:"violationCount,omitempty"`

	// AffectedScheduleIndexes - The index of the affected schedule summary in the schedules list
	AffectedScheduleIndexes *[]int `json:"affectedScheduleIndexes,omitempty"`

	// CreatedBy - The user who started the job which created this occurrence. The id may be 'System' if it was an automated process
	CreatedBy *Userreference `json:"createdBy,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Activityplanoccurrencelistitem) SetField(field string, fieldValue interface{}) {
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

func (o Activityplanoccurrencelistitem) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "RunDate", }
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
	type Alias Activityplanoccurrencelistitem
	
	RunDate := new(string)
	if o.RunDate != nil {
		
		*RunDate = timeutil.Strftime(o.RunDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		RunDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		RunDate *string `json:"runDate,omitempty"`
		
		SessionDateRange *Requireddaterange `json:"sessionDateRange,omitempty"`
		
		Facilitated *bool `json:"facilitated,omitempty"`
		
		SessionCount *int `json:"sessionCount,omitempty"`
		
		ScheduledAttendeesCount *int `json:"scheduledAttendeesCount,omitempty"`
		
		UnscheduledAttendeesCount *int `json:"unscheduledAttendeesCount,omitempty"`
		
		Edited *bool `json:"edited,omitempty"`
		
		ViolationCount *int `json:"violationCount,omitempty"`
		
		AffectedScheduleIndexes *[]int `json:"affectedScheduleIndexes,omitempty"`
		
		CreatedBy *Userreference `json:"createdBy,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		RunDate: RunDate,
		
		SessionDateRange: o.SessionDateRange,
		
		Facilitated: o.Facilitated,
		
		SessionCount: o.SessionCount,
		
		ScheduledAttendeesCount: o.ScheduledAttendeesCount,
		
		UnscheduledAttendeesCount: o.UnscheduledAttendeesCount,
		
		Edited: o.Edited,
		
		ViolationCount: o.ViolationCount,
		
		AffectedScheduleIndexes: o.AffectedScheduleIndexes,
		
		CreatedBy: o.CreatedBy,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Activityplanoccurrencelistitem) UnmarshalJSON(b []byte) error {
	var ActivityplanoccurrencelistitemMap map[string]interface{}
	err := json.Unmarshal(b, &ActivityplanoccurrencelistitemMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ActivityplanoccurrencelistitemMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if runDateString, ok := ActivityplanoccurrencelistitemMap["runDate"].(string); ok {
		RunDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", runDateString)
		o.RunDate = &RunDate
	}
	
	if SessionDateRange, ok := ActivityplanoccurrencelistitemMap["sessionDateRange"].(map[string]interface{}); ok {
		SessionDateRangeString, _ := json.Marshal(SessionDateRange)
		json.Unmarshal(SessionDateRangeString, &o.SessionDateRange)
	}
	
	if Facilitated, ok := ActivityplanoccurrencelistitemMap["facilitated"].(bool); ok {
		o.Facilitated = &Facilitated
	}
    
	if SessionCount, ok := ActivityplanoccurrencelistitemMap["sessionCount"].(float64); ok {
		SessionCountInt := int(SessionCount)
		o.SessionCount = &SessionCountInt
	}
	
	if ScheduledAttendeesCount, ok := ActivityplanoccurrencelistitemMap["scheduledAttendeesCount"].(float64); ok {
		ScheduledAttendeesCountInt := int(ScheduledAttendeesCount)
		o.ScheduledAttendeesCount = &ScheduledAttendeesCountInt
	}
	
	if UnscheduledAttendeesCount, ok := ActivityplanoccurrencelistitemMap["unscheduledAttendeesCount"].(float64); ok {
		UnscheduledAttendeesCountInt := int(UnscheduledAttendeesCount)
		o.UnscheduledAttendeesCount = &UnscheduledAttendeesCountInt
	}
	
	if Edited, ok := ActivityplanoccurrencelistitemMap["edited"].(bool); ok {
		o.Edited = &Edited
	}
    
	if ViolationCount, ok := ActivityplanoccurrencelistitemMap["violationCount"].(float64); ok {
		ViolationCountInt := int(ViolationCount)
		o.ViolationCount = &ViolationCountInt
	}
	
	if AffectedScheduleIndexes, ok := ActivityplanoccurrencelistitemMap["affectedScheduleIndexes"].([]interface{}); ok {
		AffectedScheduleIndexesString, _ := json.Marshal(AffectedScheduleIndexes)
		json.Unmarshal(AffectedScheduleIndexesString, &o.AffectedScheduleIndexes)
	}
	
	if CreatedBy, ok := ActivityplanoccurrencelistitemMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if SelfUri, ok := ActivityplanoccurrencelistitemMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Activityplanoccurrencelistitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

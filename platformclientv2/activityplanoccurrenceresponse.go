package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Activityplanoccurrenceresponse
type Activityplanoccurrenceresponse struct { 
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

	// Edited - Whether this occurrence is edited
	Edited *bool `json:"edited,omitempty"`

	// Violations - The violations for this occurrence
	Violations *[]Activityplanoccurrenceviolation `json:"violations,omitempty"`

	// Sessions - The sessions scheduled by the run
	Sessions *[]Activityplanoccurrencesessionwithattendees `json:"sessions,omitempty"`

	// UnscheduledAttendees - The attendees who could not be scheduled for this occurrence
	UnscheduledAttendees *[]Activityplanunscheduledattendee `json:"unscheduledAttendees,omitempty"`

	// AffectedSchedules - Summaries of the schedules affected by this occurrence
	AffectedSchedules *[]Buschedulesummary `json:"affectedSchedules,omitempty"`

	// CreatedBy - The user who started the job which created this occurrence
	CreatedBy *Userreference `json:"createdBy,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Activityplanoccurrenceresponse) SetField(field string, fieldValue interface{}) {
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

func (o Activityplanoccurrenceresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Activityplanoccurrenceresponse
	
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
		
		Edited *bool `json:"edited,omitempty"`
		
		Violations *[]Activityplanoccurrenceviolation `json:"violations,omitempty"`
		
		Sessions *[]Activityplanoccurrencesessionwithattendees `json:"sessions,omitempty"`
		
		UnscheduledAttendees *[]Activityplanunscheduledattendee `json:"unscheduledAttendees,omitempty"`
		
		AffectedSchedules *[]Buschedulesummary `json:"affectedSchedules,omitempty"`
		
		CreatedBy *Userreference `json:"createdBy,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		RunDate: RunDate,
		
		SessionDateRange: o.SessionDateRange,
		
		Facilitated: o.Facilitated,
		
		Edited: o.Edited,
		
		Violations: o.Violations,
		
		Sessions: o.Sessions,
		
		UnscheduledAttendees: o.UnscheduledAttendees,
		
		AffectedSchedules: o.AffectedSchedules,
		
		CreatedBy: o.CreatedBy,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Activityplanoccurrenceresponse) UnmarshalJSON(b []byte) error {
	var ActivityplanoccurrenceresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ActivityplanoccurrenceresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ActivityplanoccurrenceresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if runDateString, ok := ActivityplanoccurrenceresponseMap["runDate"].(string); ok {
		RunDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", runDateString)
		o.RunDate = &RunDate
	}
	
	if SessionDateRange, ok := ActivityplanoccurrenceresponseMap["sessionDateRange"].(map[string]interface{}); ok {
		SessionDateRangeString, _ := json.Marshal(SessionDateRange)
		json.Unmarshal(SessionDateRangeString, &o.SessionDateRange)
	}
	
	if Facilitated, ok := ActivityplanoccurrenceresponseMap["facilitated"].(bool); ok {
		o.Facilitated = &Facilitated
	}
    
	if Edited, ok := ActivityplanoccurrenceresponseMap["edited"].(bool); ok {
		o.Edited = &Edited
	}
    
	if Violations, ok := ActivityplanoccurrenceresponseMap["violations"].([]interface{}); ok {
		ViolationsString, _ := json.Marshal(Violations)
		json.Unmarshal(ViolationsString, &o.Violations)
	}
	
	if Sessions, ok := ActivityplanoccurrenceresponseMap["sessions"].([]interface{}); ok {
		SessionsString, _ := json.Marshal(Sessions)
		json.Unmarshal(SessionsString, &o.Sessions)
	}
	
	if UnscheduledAttendees, ok := ActivityplanoccurrenceresponseMap["unscheduledAttendees"].([]interface{}); ok {
		UnscheduledAttendeesString, _ := json.Marshal(UnscheduledAttendees)
		json.Unmarshal(UnscheduledAttendeesString, &o.UnscheduledAttendees)
	}
	
	if AffectedSchedules, ok := ActivityplanoccurrenceresponseMap["affectedSchedules"].([]interface{}); ok {
		AffectedSchedulesString, _ := json.Marshal(AffectedSchedules)
		json.Unmarshal(AffectedSchedulesString, &o.AffectedSchedules)
	}
	
	if CreatedBy, ok := ActivityplanoccurrenceresponseMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if SelfUri, ok := ActivityplanoccurrenceresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Activityplanoccurrenceresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

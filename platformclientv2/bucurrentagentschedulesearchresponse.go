package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Bucurrentagentschedulesearchresponse
type Bucurrentagentschedulesearchresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// AgentSchedules - The requested agent schedules
	AgentSchedules *[]Buagentschedulesearchresponse `json:"agentSchedules,omitempty"`

	// BusinessUnitTimeZone - The time zone configured for the business unit to which this schedule applies
	BusinessUnitTimeZone *string `json:"businessUnitTimeZone,omitempty"`

	// PublishedSchedules - References to all published week schedules overlapping the start/end date query parameters
	PublishedSchedules *[]Buagentschedulepublishedschedulereference `json:"publishedSchedules,omitempty"`

	// StartDate - The start date of the schedules. Only populated on notifications. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`

	// EndDate - The end date of the schedules. Only populated on notifications. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndDate *time.Time `json:"endDate,omitempty"`

	// Updates - The list of updates for the schedule. Only used in notifications
	Updates *[]Buagentscheduleupdate `json:"updates,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Bucurrentagentschedulesearchresponse) SetField(field string, fieldValue interface{}) {
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

func (o Bucurrentagentschedulesearchresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Bucurrentagentschedulesearchresponse
	
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
		AgentSchedules *[]Buagentschedulesearchresponse `json:"agentSchedules,omitempty"`
		
		BusinessUnitTimeZone *string `json:"businessUnitTimeZone,omitempty"`
		
		PublishedSchedules *[]Buagentschedulepublishedschedulereference `json:"publishedSchedules,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		Updates *[]Buagentscheduleupdate `json:"updates,omitempty"`
		Alias
	}{ 
		AgentSchedules: o.AgentSchedules,
		
		BusinessUnitTimeZone: o.BusinessUnitTimeZone,
		
		PublishedSchedules: o.PublishedSchedules,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		Updates: o.Updates,
		Alias:    (Alias)(o),
	})
}

func (o *Bucurrentagentschedulesearchresponse) UnmarshalJSON(b []byte) error {
	var BucurrentagentschedulesearchresponseMap map[string]interface{}
	err := json.Unmarshal(b, &BucurrentagentschedulesearchresponseMap)
	if err != nil {
		return err
	}
	
	if AgentSchedules, ok := BucurrentagentschedulesearchresponseMap["agentSchedules"].([]interface{}); ok {
		AgentSchedulesString, _ := json.Marshal(AgentSchedules)
		json.Unmarshal(AgentSchedulesString, &o.AgentSchedules)
	}
	
	if BusinessUnitTimeZone, ok := BucurrentagentschedulesearchresponseMap["businessUnitTimeZone"].(string); ok {
		o.BusinessUnitTimeZone = &BusinessUnitTimeZone
	}
    
	if PublishedSchedules, ok := BucurrentagentschedulesearchresponseMap["publishedSchedules"].([]interface{}); ok {
		PublishedSchedulesString, _ := json.Marshal(PublishedSchedules)
		json.Unmarshal(PublishedSchedulesString, &o.PublishedSchedules)
	}
	
	if startDateString, ok := BucurrentagentschedulesearchresponseMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := BucurrentagentschedulesearchresponseMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if Updates, ok := BucurrentagentschedulesearchresponseMap["updates"].([]interface{}); ok {
		UpdatesString, _ := json.Marshal(Updates)
		json.Unmarshal(UpdatesString, &o.Updates)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bucurrentagentschedulesearchresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

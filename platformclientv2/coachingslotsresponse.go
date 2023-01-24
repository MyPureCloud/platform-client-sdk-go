package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Coachingslotsresponse
type Coachingslotsresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// SuggestedSlots - List of slots where coaching appointment can be scheduled
	SuggestedSlots *[]Coachingslot `json:"suggestedSlots,omitempty"`

	// AttendeeSchedules - Periods of availability for attendees to schedule coaching appointment
	AttendeeSchedules *[]Useravailabletimes `json:"attendeeSchedules,omitempty"`

	// FacilitatorSchedules - Periods of availability for facilitators to schedule coaching appointment
	FacilitatorSchedules *[]Useravailabletimes `json:"facilitatorSchedules,omitempty"`

	// WfmScheduleActivities - Detailed data for WFM scheduled activities
	WfmScheduleActivities *[]Wfmscheduleactivity `json:"wfmScheduleActivities,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Coachingslotsresponse) SetField(field string, fieldValue interface{}) {
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

func (o Coachingslotsresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Coachingslotsresponse
	
	return json.Marshal(&struct { 
		SuggestedSlots *[]Coachingslot `json:"suggestedSlots,omitempty"`
		
		AttendeeSchedules *[]Useravailabletimes `json:"attendeeSchedules,omitempty"`
		
		FacilitatorSchedules *[]Useravailabletimes `json:"facilitatorSchedules,omitempty"`
		
		WfmScheduleActivities *[]Wfmscheduleactivity `json:"wfmScheduleActivities,omitempty"`
		Alias
	}{ 
		SuggestedSlots: o.SuggestedSlots,
		
		AttendeeSchedules: o.AttendeeSchedules,
		
		FacilitatorSchedules: o.FacilitatorSchedules,
		
		WfmScheduleActivities: o.WfmScheduleActivities,
		Alias:    (Alias)(o),
	})
}

func (o *Coachingslotsresponse) UnmarshalJSON(b []byte) error {
	var CoachingslotsresponseMap map[string]interface{}
	err := json.Unmarshal(b, &CoachingslotsresponseMap)
	if err != nil {
		return err
	}
	
	if SuggestedSlots, ok := CoachingslotsresponseMap["suggestedSlots"].([]interface{}); ok {
		SuggestedSlotsString, _ := json.Marshal(SuggestedSlots)
		json.Unmarshal(SuggestedSlotsString, &o.SuggestedSlots)
	}
	
	if AttendeeSchedules, ok := CoachingslotsresponseMap["attendeeSchedules"].([]interface{}); ok {
		AttendeeSchedulesString, _ := json.Marshal(AttendeeSchedules)
		json.Unmarshal(AttendeeSchedulesString, &o.AttendeeSchedules)
	}
	
	if FacilitatorSchedules, ok := CoachingslotsresponseMap["facilitatorSchedules"].([]interface{}); ok {
		FacilitatorSchedulesString, _ := json.Marshal(FacilitatorSchedules)
		json.Unmarshal(FacilitatorSchedulesString, &o.FacilitatorSchedules)
	}
	
	if WfmScheduleActivities, ok := CoachingslotsresponseMap["wfmScheduleActivities"].([]interface{}); ok {
		WfmScheduleActivitiesString, _ := json.Marshal(WfmScheduleActivities)
		json.Unmarshal(WfmScheduleActivitiesString, &o.WfmScheduleActivities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Coachingslotsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

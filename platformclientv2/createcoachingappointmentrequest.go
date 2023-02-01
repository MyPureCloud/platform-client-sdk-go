package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Createcoachingappointmentrequest - Create coaching appointment request
type Createcoachingappointmentrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of coaching appointment.
	Name *string `json:"name,omitempty"`

	// Description - The description of coaching appointment.
	Description *string `json:"description,omitempty"`

	// DateStart - The date/time the coaching appointment starts. Times will be rounded down to the minute. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStart *time.Time `json:"dateStart,omitempty"`

	// LengthInMinutes - The duration of coaching appointment in minutes.
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`

	// FacilitatorId - The facilitator ID of coaching appointment.
	FacilitatorId *string `json:"facilitatorId,omitempty"`

	// AttendeeIds - IDs of attendees in the coaching appointment.
	AttendeeIds *[]string `json:"attendeeIds,omitempty"`

	// ConversationIds - IDs of conversations associated with this coaching appointment.
	ConversationIds *[]string `json:"conversationIds,omitempty"`

	// DocumentIds - IDs of documents associated with this coaching appointment.
	DocumentIds *[]string `json:"documentIds,omitempty"`

	// WfmSchedule - The Workforce Management schedule the appointment is associated with.
	WfmSchedule *Wfmschedulereference `json:"wfmSchedule,omitempty"`

	// ExternalLinks - The list of external links related to the appointment
	ExternalLinks *[]string `json:"externalLinks,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Createcoachingappointmentrequest) SetField(field string, fieldValue interface{}) {
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

func (o Createcoachingappointmentrequest) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateStart", }
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
	type Alias Createcoachingappointmentrequest
	
	DateStart := new(string)
	if o.DateStart != nil {
		
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStart = nil
	}
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		FacilitatorId *string `json:"facilitatorId,omitempty"`
		
		AttendeeIds *[]string `json:"attendeeIds,omitempty"`
		
		ConversationIds *[]string `json:"conversationIds,omitempty"`
		
		DocumentIds *[]string `json:"documentIds,omitempty"`
		
		WfmSchedule *Wfmschedulereference `json:"wfmSchedule,omitempty"`
		
		ExternalLinks *[]string `json:"externalLinks,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		Description: o.Description,
		
		DateStart: DateStart,
		
		LengthInMinutes: o.LengthInMinutes,
		
		FacilitatorId: o.FacilitatorId,
		
		AttendeeIds: o.AttendeeIds,
		
		ConversationIds: o.ConversationIds,
		
		DocumentIds: o.DocumentIds,
		
		WfmSchedule: o.WfmSchedule,
		
		ExternalLinks: o.ExternalLinks,
		Alias:    (Alias)(o),
	})
}

func (o *Createcoachingappointmentrequest) UnmarshalJSON(b []byte) error {
	var CreatecoachingappointmentrequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreatecoachingappointmentrequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := CreatecoachingappointmentrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := CreatecoachingappointmentrequestMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if dateStartString, ok := CreatecoachingappointmentrequestMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartString)
		o.DateStart = &DateStart
	}
	
	if LengthInMinutes, ok := CreatecoachingappointmentrequestMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	
	if FacilitatorId, ok := CreatecoachingappointmentrequestMap["facilitatorId"].(string); ok {
		o.FacilitatorId = &FacilitatorId
	}
    
	if AttendeeIds, ok := CreatecoachingappointmentrequestMap["attendeeIds"].([]interface{}); ok {
		AttendeeIdsString, _ := json.Marshal(AttendeeIds)
		json.Unmarshal(AttendeeIdsString, &o.AttendeeIds)
	}
	
	if ConversationIds, ok := CreatecoachingappointmentrequestMap["conversationIds"].([]interface{}); ok {
		ConversationIdsString, _ := json.Marshal(ConversationIds)
		json.Unmarshal(ConversationIdsString, &o.ConversationIds)
	}
	
	if DocumentIds, ok := CreatecoachingappointmentrequestMap["documentIds"].([]interface{}); ok {
		DocumentIdsString, _ := json.Marshal(DocumentIds)
		json.Unmarshal(DocumentIdsString, &o.DocumentIds)
	}
	
	if WfmSchedule, ok := CreatecoachingappointmentrequestMap["wfmSchedule"].(map[string]interface{}); ok {
		WfmScheduleString, _ := json.Marshal(WfmSchedule)
		json.Unmarshal(WfmScheduleString, &o.WfmSchedule)
	}
	
	if ExternalLinks, ok := CreatecoachingappointmentrequestMap["externalLinks"].([]interface{}); ok {
		ExternalLinksString, _ := json.Marshal(ExternalLinks)
		json.Unmarshal(ExternalLinksString, &o.ExternalLinks)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createcoachingappointmentrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

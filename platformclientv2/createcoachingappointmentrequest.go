package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createcoachingappointmentrequest - Create coaching appointment request
type Createcoachingappointmentrequest struct { 
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

func (o *Createcoachingappointmentrequest) MarshalJSON() ([]byte, error) {
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
		*Alias
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
		Alias:    (*Alias)(o),
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

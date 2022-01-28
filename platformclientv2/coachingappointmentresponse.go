package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Coachingappointmentresponse - Coaching appointment response
type Coachingappointmentresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of coaching appointment
	Name *string `json:"name,omitempty"`


	// Description - The description of coaching appointment
	Description *string `json:"description,omitempty"`


	// DateStart - The date/time the coaching appointment starts. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStart *time.Time `json:"dateStart,omitempty"`


	// LengthInMinutes - The duration of coaching appointment in minutes
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`


	// Status - The status of coaching appointment
	Status *string `json:"status,omitempty"`


	// Facilitator - The facilitator of coaching appointment
	Facilitator *Userreference `json:"facilitator,omitempty"`


	// Attendees - The list of attendees attending the coaching
	Attendees *[]Userreference `json:"attendees,omitempty"`


	// CreatedBy - The user who created the coaching appointment
	CreatedBy *Userreference `json:"createdBy,omitempty"`


	// DateCreated - The date/time the coaching appointment was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// ModifiedBy - The last user to modify the coaching appointment
	ModifiedBy *Userreference `json:"modifiedBy,omitempty"`


	// DateModified - The date/time the coaching appointment was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Conversations - The list of conversations associated with coaching appointment.
	Conversations *[]Conversationreference `json:"conversations,omitempty"`


	// Documents - The list of documents associated with coaching appointment.
	Documents *[]Documentreference `json:"documents,omitempty"`


	// IsOverdue - Whether the appointment is overdue.
	IsOverdue *bool `json:"isOverdue,omitempty"`


	// WfmSchedule - The Workforce Management schedule the appointment is associated with.
	WfmSchedule *Wfmschedulereference `json:"wfmSchedule,omitempty"`


	// DateCompleted - The date/time the coaching appointment was set to completed status. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCompleted *time.Time `json:"dateCompleted,omitempty"`


	// ExternalLinks - The list of external links related to the appointment
	ExternalLinks *[]string `json:"externalLinks,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Coachingappointmentresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Coachingappointmentresponse
	
	DateStart := new(string)
	if o.DateStart != nil {
		
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStart = nil
	}
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	DateCompleted := new(string)
	if o.DateCompleted != nil {
		
		*DateCompleted = timeutil.Strftime(o.DateCompleted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCompleted = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Facilitator *Userreference `json:"facilitator,omitempty"`
		
		Attendees *[]Userreference `json:"attendees,omitempty"`
		
		CreatedBy *Userreference `json:"createdBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		ModifiedBy *Userreference `json:"modifiedBy,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Conversations *[]Conversationreference `json:"conversations,omitempty"`
		
		Documents *[]Documentreference `json:"documents,omitempty"`
		
		IsOverdue *bool `json:"isOverdue,omitempty"`
		
		WfmSchedule *Wfmschedulereference `json:"wfmSchedule,omitempty"`
		
		DateCompleted *string `json:"dateCompleted,omitempty"`
		
		ExternalLinks *[]string `json:"externalLinks,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		DateStart: DateStart,
		
		LengthInMinutes: o.LengthInMinutes,
		
		Status: o.Status,
		
		Facilitator: o.Facilitator,
		
		Attendees: o.Attendees,
		
		CreatedBy: o.CreatedBy,
		
		DateCreated: DateCreated,
		
		ModifiedBy: o.ModifiedBy,
		
		DateModified: DateModified,
		
		Conversations: o.Conversations,
		
		Documents: o.Documents,
		
		IsOverdue: o.IsOverdue,
		
		WfmSchedule: o.WfmSchedule,
		
		DateCompleted: DateCompleted,
		
		ExternalLinks: o.ExternalLinks,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Coachingappointmentresponse) UnmarshalJSON(b []byte) error {
	var CoachingappointmentresponseMap map[string]interface{}
	err := json.Unmarshal(b, &CoachingappointmentresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CoachingappointmentresponseMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := CoachingappointmentresponseMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Description, ok := CoachingappointmentresponseMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if dateStartString, ok := CoachingappointmentresponseMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartString)
		o.DateStart = &DateStart
	}
	
	if LengthInMinutes, ok := CoachingappointmentresponseMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	
	if Status, ok := CoachingappointmentresponseMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if Facilitator, ok := CoachingappointmentresponseMap["facilitator"].(map[string]interface{}); ok {
		FacilitatorString, _ := json.Marshal(Facilitator)
		json.Unmarshal(FacilitatorString, &o.Facilitator)
	}
	
	if Attendees, ok := CoachingappointmentresponseMap["attendees"].([]interface{}); ok {
		AttendeesString, _ := json.Marshal(Attendees)
		json.Unmarshal(AttendeesString, &o.Attendees)
	}
	
	if CreatedBy, ok := CoachingappointmentresponseMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if dateCreatedString, ok := CoachingappointmentresponseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if ModifiedBy, ok := CoachingappointmentresponseMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if dateModifiedString, ok := CoachingappointmentresponseMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Conversations, ok := CoachingappointmentresponseMap["conversations"].([]interface{}); ok {
		ConversationsString, _ := json.Marshal(Conversations)
		json.Unmarshal(ConversationsString, &o.Conversations)
	}
	
	if Documents, ok := CoachingappointmentresponseMap["documents"].([]interface{}); ok {
		DocumentsString, _ := json.Marshal(Documents)
		json.Unmarshal(DocumentsString, &o.Documents)
	}
	
	if IsOverdue, ok := CoachingappointmentresponseMap["isOverdue"].(bool); ok {
		o.IsOverdue = &IsOverdue
	}
	
	if WfmSchedule, ok := CoachingappointmentresponseMap["wfmSchedule"].(map[string]interface{}); ok {
		WfmScheduleString, _ := json.Marshal(WfmSchedule)
		json.Unmarshal(WfmScheduleString, &o.WfmSchedule)
	}
	
	if dateCompletedString, ok := CoachingappointmentresponseMap["dateCompleted"].(string); ok {
		DateCompleted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCompletedString)
		o.DateCompleted = &DateCompleted
	}
	
	if ExternalLinks, ok := CoachingappointmentresponseMap["externalLinks"].([]interface{}); ok {
		ExternalLinksString, _ := json.Marshal(ExternalLinks)
		json.Unmarshal(ExternalLinksString, &o.ExternalLinks)
	}
	
	if SelfUri, ok := CoachingappointmentresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Coachingappointmentresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

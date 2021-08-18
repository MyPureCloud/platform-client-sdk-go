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


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Coachingappointmentresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Coachingappointmentresponse

	
	DateStart := new(string)
	if u.DateStart != nil {
		
		*DateStart = timeutil.Strftime(u.DateStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStart = nil
	}
	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
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
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Description: u.Description,
		
		DateStart: DateStart,
		
		LengthInMinutes: u.LengthInMinutes,
		
		Status: u.Status,
		
		Facilitator: u.Facilitator,
		
		Attendees: u.Attendees,
		
		CreatedBy: u.CreatedBy,
		
		DateCreated: DateCreated,
		
		ModifiedBy: u.ModifiedBy,
		
		DateModified: DateModified,
		
		Conversations: u.Conversations,
		
		Documents: u.Documents,
		
		IsOverdue: u.IsOverdue,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Coachingappointmentresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"time"
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

// String returns a JSON representation of the model
func (o *Coachingappointmentresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

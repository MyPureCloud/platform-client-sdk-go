package platformclientv2
import (
	"time"
	"encoding/json"
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

}

// String returns a JSON representation of the model
func (o *Createcoachingappointmentrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

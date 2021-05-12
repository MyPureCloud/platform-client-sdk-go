package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Coachingnotification
type Coachingnotification struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the appointment for this notification.
	Name *string `json:"name,omitempty"`


	// MarkedAsRead - Indicates if notification is read or unread
	MarkedAsRead *bool `json:"markedAsRead,omitempty"`


	// ActionType - Action causing the notification.
	ActionType *string `json:"actionType,omitempty"`


	// Relationship - The relationship of this user to this notification's appointment
	Relationship *string `json:"relationship,omitempty"`


	// DateStart - The start time of the appointment relating to this notification. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStart *time.Time `json:"dateStart,omitempty"`


	// LengthInMinutes - The duration of the appointment on this notification
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`


	// Status - The status of the appointment for this notification
	Status *string `json:"status,omitempty"`


	// User - The user of this notification
	User *Userreference `json:"user,omitempty"`


	// Appointment - The appointment
	Appointment *Coachingappointmentresponse `json:"appointment,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Coachingnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

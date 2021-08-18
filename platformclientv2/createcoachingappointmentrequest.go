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

}

func (u *Createcoachingappointmentrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createcoachingappointmentrequest

	
	DateStart := new(string)
	if u.DateStart != nil {
		
		*DateStart = timeutil.Strftime(u.DateStart, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		*Alias
	}{ 
		Name: u.Name,
		
		Description: u.Description,
		
		DateStart: DateStart,
		
		LengthInMinutes: u.LengthInMinutes,
		
		FacilitatorId: u.FacilitatorId,
		
		AttendeeIds: u.AttendeeIds,
		
		ConversationIds: u.ConversationIds,
		
		DocumentIds: u.DocumentIds,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Createcoachingappointmentrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

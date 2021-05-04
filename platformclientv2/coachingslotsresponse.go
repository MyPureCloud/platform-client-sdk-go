package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Coachingslotsresponse
type Coachingslotsresponse struct { 
	// SuggestedSlots - List of slots where coaching appointment can be scheduled
	SuggestedSlots *[]Coachingslot `json:"suggestedSlots,omitempty"`


	// AttendeeSchedules - Periods of availability for attendees to schedule coaching appointment
	AttendeeSchedules *[]Useravailabletimes `json:"attendeeSchedules,omitempty"`


	// FacilitatorSchedules - Periods of availability for facilitators to schedule coaching appointment
	FacilitatorSchedules *[]Useravailabletimes `json:"facilitatorSchedules,omitempty"`

}

// String returns a JSON representation of the model
func (o *Coachingslotsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

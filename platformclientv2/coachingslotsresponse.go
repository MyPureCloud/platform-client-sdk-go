package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Coachingslotsresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Coachingslotsresponse

	

	return json.Marshal(&struct { 
		SuggestedSlots *[]Coachingslot `json:"suggestedSlots,omitempty"`
		
		AttendeeSchedules *[]Useravailabletimes `json:"attendeeSchedules,omitempty"`
		
		FacilitatorSchedules *[]Useravailabletimes `json:"facilitatorSchedules,omitempty"`
		*Alias
	}{ 
		SuggestedSlots: u.SuggestedSlots,
		
		AttendeeSchedules: u.AttendeeSchedules,
		
		FacilitatorSchedules: u.FacilitatorSchedules,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Coachingslotsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"encoding/json"
)

// Callabletime
type Callabletime struct { 
	// TimeSlots - The time intervals for which it is acceptable to place outbound calls.
	TimeSlots *[]Campaigntimeslot `json:"timeSlots,omitempty"`


	// TimeZoneId - The time zone for the time slots; for example, Africa/Abidjan
	TimeZoneId *string `json:"timeZoneId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Callabletime) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

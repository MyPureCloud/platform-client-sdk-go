package platformclientv2
import (
	"encoding/json"
)

// Timeallowed
type Timeallowed struct { 
	// TimeSlots
	TimeSlots *[]Timeslot `json:"timeSlots,omitempty"`


	// TimeZoneId
	TimeZoneId *string `json:"timeZoneId,omitempty"`


	// Empty
	Empty *bool `json:"empty,omitempty"`

}

// String returns a JSON representation of the model
func (o *Timeallowed) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

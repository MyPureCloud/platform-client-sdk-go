package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Callabletime
type Callabletime struct { 
	// TimeSlots - The time intervals for which it is acceptable to place outbound calls.
	TimeSlots *[]Campaigntimeslot `json:"timeSlots,omitempty"`


	// TimeZoneId - The time zone for the time slots; for example, Africa/Abidjan
	TimeZoneId *string `json:"timeZoneId,omitempty"`

}

func (u *Callabletime) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Callabletime

	

	return json.Marshal(&struct { 
		TimeSlots *[]Campaigntimeslot `json:"timeSlots,omitempty"`
		
		TimeZoneId *string `json:"timeZoneId,omitempty"`
		*Alias
	}{ 
		TimeSlots: u.TimeSlots,
		
		TimeZoneId: u.TimeZoneId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Callabletime) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

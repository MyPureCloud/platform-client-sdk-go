package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Timeallowed) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Timeallowed

	

	return json.Marshal(&struct { 
		TimeSlots *[]Timeslot `json:"timeSlots,omitempty"`
		
		TimeZoneId *string `json:"timeZoneId,omitempty"`
		
		Empty *bool `json:"empty,omitempty"`
		*Alias
	}{ 
		TimeSlots: u.TimeSlots,
		
		TimeZoneId: u.TimeZoneId,
		
		Empty: u.Empty,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Timeallowed) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Coachingslotsrequest
type Coachingslotsrequest struct { 
	// Interval - Range of time to get slots for scheduling coaching appointments. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`


	// LengthInMinutes - The duration of coaching appointment to schedule in 15 minutes granularity up to maximum of 60 minutes
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`


	// AttendeeIds - List of attendees to determine coaching appointment slots
	AttendeeIds *[]string `json:"attendeeIds,omitempty"`


	// FacilitatorIds - List of facilitators to determine coaching appointment slots
	FacilitatorIds *[]string `json:"facilitatorIds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Coachingslotsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

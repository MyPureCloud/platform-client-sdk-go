package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Buagentscheduleshift
type Buagentscheduleshift struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// StartDate - The start date of this shift. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`


	// LengthMinutes - The length of this shift in minutes
	LengthMinutes *int `json:"lengthMinutes,omitempty"`


	// Activities - The activities associated with this shift
	Activities *[]Buagentscheduleactivity `json:"activities,omitempty"`


	// ManuallyEdited - Whether this shift was manually edited. This is only set by clients and is used for rescheduling
	ManuallyEdited *bool `json:"manuallyEdited,omitempty"`


	// Schedule - The schedule to which this shift belongs
	Schedule *Buschedulereference `json:"schedule,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buagentscheduleshift) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

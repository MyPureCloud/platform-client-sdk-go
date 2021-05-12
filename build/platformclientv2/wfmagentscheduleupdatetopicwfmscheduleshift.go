package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmagentscheduleupdatetopicwfmscheduleshift
type Wfmagentscheduleupdatetopicwfmscheduleshift struct { 
	// WeekDate
	WeekDate *string `json:"weekDate,omitempty"`


	// WeekScheduleId
	WeekScheduleId *string `json:"weekScheduleId,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`


	// StartDate
	StartDate *time.Time `json:"startDate,omitempty"`


	// LengthInMinutes
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`


	// Activities
	Activities *[]Wfmagentscheduleupdatetopicwfmscheduleactivity `json:"activities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmagentscheduleupdatetopicwfmscheduleshift) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

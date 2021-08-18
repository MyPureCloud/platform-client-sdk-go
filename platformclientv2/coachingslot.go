package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Coachingslot
type Coachingslot struct { 
	// DateStart - Start date and time of scheduled coaching appointment slot. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStart *time.Time `json:"dateStart,omitempty"`


	// LengthInMinutes - Length of coaching appointment slot in minutes
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`


	// StaffingDifference - Difference between scheduled and forecast headcount for this slot after scheduling the coaching appointment
	StaffingDifference *float64 `json:"staffingDifference,omitempty"`


	// DifferenceRating - Rating based on the staffing difference for scheduled slot
	DifferenceRating *string `json:"differenceRating,omitempty"`


	// WfmSchedule - Workforce Management schedule information associated with the slot
	WfmSchedule *Wfmschedulereference `json:"wfmSchedule,omitempty"`

}

func (u *Coachingslot) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Coachingslot

	
	DateStart := new(string)
	if u.DateStart != nil {
		
		*DateStart = timeutil.Strftime(u.DateStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStart = nil
	}
	

	return json.Marshal(&struct { 
		DateStart *string `json:"dateStart,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		StaffingDifference *float64 `json:"staffingDifference,omitempty"`
		
		DifferenceRating *string `json:"differenceRating,omitempty"`
		
		WfmSchedule *Wfmschedulereference `json:"wfmSchedule,omitempty"`
		*Alias
	}{ 
		DateStart: DateStart,
		
		LengthInMinutes: u.LengthInMinutes,
		
		StaffingDifference: u.StaffingDifference,
		
		DifferenceRating: u.DifferenceRating,
		
		WfmSchedule: u.WfmSchedule,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Coachingslot) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

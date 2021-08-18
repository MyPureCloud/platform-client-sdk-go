package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Wfmagentscheduleupdatetopicwfmscheduleshift) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmagentscheduleupdatetopicwfmscheduleshift

	
	StartDate := new(string)
	if u.StartDate != nil {
		
		*StartDate = timeutil.Strftime(u.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	

	return json.Marshal(&struct { 
		WeekDate *string `json:"weekDate,omitempty"`
		
		WeekScheduleId *string `json:"weekScheduleId,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		Activities *[]Wfmagentscheduleupdatetopicwfmscheduleactivity `json:"activities,omitempty"`
		*Alias
	}{ 
		WeekDate: u.WeekDate,
		
		WeekScheduleId: u.WeekScheduleId,
		
		Id: u.Id,
		
		StartDate: StartDate,
		
		LengthInMinutes: u.LengthInMinutes,
		
		Activities: u.Activities,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmagentscheduleupdatetopicwfmscheduleshift) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

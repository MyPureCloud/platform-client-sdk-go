package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userscheduleshift - Single shift in a user's schedule
type Userscheduleshift struct { 
	// WeekSchedule - The schedule to which this shift belongs
	WeekSchedule *Weekschedulereference `json:"weekSchedule,omitempty"`


	// Id - ID of the schedule shift. This is only for the case of updating and deleting an existing shift
	Id *string `json:"id,omitempty"`


	// StartDate - Start time in UTC for this shift. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`


	// LengthInMinutes - Length of this shift in minutes
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`


	// Activities - List of activities in this shift
	Activities *[]Userscheduleactivity `json:"activities,omitempty"`


	// Delete - If marked true for updating this schedule shift, it will be deleted
	Delete *bool `json:"delete,omitempty"`


	// ManuallyEdited - Whether the shift was set as manually edited
	ManuallyEdited *bool `json:"manuallyEdited,omitempty"`

}

func (u *Userscheduleshift) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userscheduleshift

	
	StartDate := new(string)
	if u.StartDate != nil {
		
		*StartDate = timeutil.Strftime(u.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	

	return json.Marshal(&struct { 
		WeekSchedule *Weekschedulereference `json:"weekSchedule,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		Activities *[]Userscheduleactivity `json:"activities,omitempty"`
		
		Delete *bool `json:"delete,omitempty"`
		
		ManuallyEdited *bool `json:"manuallyEdited,omitempty"`
		*Alias
	}{ 
		WeekSchedule: u.WeekSchedule,
		
		Id: u.Id,
		
		StartDate: StartDate,
		
		LengthInMinutes: u.LengthInMinutes,
		
		Activities: u.Activities,
		
		Delete: u.Delete,
		
		ManuallyEdited: u.ManuallyEdited,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userscheduleshift) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

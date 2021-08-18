package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Buagentscheduleshift) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buagentscheduleshift

	
	StartDate := new(string)
	if u.StartDate != nil {
		
		*StartDate = timeutil.Strftime(u.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		Activities *[]Buagentscheduleactivity `json:"activities,omitempty"`
		
		ManuallyEdited *bool `json:"manuallyEdited,omitempty"`
		
		Schedule *Buschedulereference `json:"schedule,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		StartDate: StartDate,
		
		LengthMinutes: u.LengthMinutes,
		
		Activities: u.Activities,
		
		ManuallyEdited: u.ManuallyEdited,
		
		Schedule: u.Schedule,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Buagentscheduleshift) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

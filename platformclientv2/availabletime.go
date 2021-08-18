package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Availabletime
type Availabletime struct { 
	// DateStart - Start of the availability period. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStart *time.Time `json:"dateStart,omitempty"`


	// LengthInMinutes - Length of availability period in minutes
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`


	// IsPaid - Indicates if this availability period is paid in Workforce Management schedule
	IsPaid *bool `json:"isPaid,omitempty"`


	// ActivityCategory - Workforce Management activity category for this availability period
	ActivityCategory *string `json:"activityCategory,omitempty"`


	// WfmSchedule - Workforce Management schedule information associated with the available time
	WfmSchedule *Wfmschedulereference `json:"wfmSchedule,omitempty"`

}

func (u *Availabletime) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Availabletime

	
	DateStart := new(string)
	if u.DateStart != nil {
		
		*DateStart = timeutil.Strftime(u.DateStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStart = nil
	}
	

	return json.Marshal(&struct { 
		DateStart *string `json:"dateStart,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		IsPaid *bool `json:"isPaid,omitempty"`
		
		ActivityCategory *string `json:"activityCategory,omitempty"`
		
		WfmSchedule *Wfmschedulereference `json:"wfmSchedule,omitempty"`
		*Alias
	}{ 
		DateStart: DateStart,
		
		LengthInMinutes: u.LengthInMinutes,
		
		IsPaid: u.IsPaid,
		
		ActivityCategory: u.ActivityCategory,
		
		WfmSchedule: u.WfmSchedule,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Availabletime) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Shifttradeactivitypreviewresponse
type Shifttradeactivitypreviewresponse struct { 
	// StartDate - The start date and time of this activity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`


	// LengthMinutes - The length in minutes of this activity
	LengthMinutes *int `json:"lengthMinutes,omitempty"`


	// ActivityCodeId - The ID of the activity code for this activity
	ActivityCodeId *string `json:"activityCodeId,omitempty"`


	// CountsAsPaidTime - Whether this activity counts as paid time
	CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`

}

func (u *Shifttradeactivitypreviewresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Shifttradeactivitypreviewresponse

	
	StartDate := new(string)
	if u.StartDate != nil {
		
		*StartDate = timeutil.Strftime(u.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	

	return json.Marshal(&struct { 
		StartDate *string `json:"startDate,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`
		*Alias
	}{ 
		StartDate: StartDate,
		
		LengthMinutes: u.LengthMinutes,
		
		ActivityCodeId: u.ActivityCodeId,
		
		CountsAsPaidTime: u.CountsAsPaidTime,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Shifttradeactivitypreviewresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

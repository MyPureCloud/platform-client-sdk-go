package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmagentscheduleupdatetopicwfmscheduleactivity
type Wfmagentscheduleupdatetopicwfmscheduleactivity struct { 
	// ActivityCodeId
	ActivityCodeId *string `json:"activityCodeId,omitempty"`


	// StartDate
	StartDate *time.Time `json:"startDate,omitempty"`


	// CountsAsPaidTime
	CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`


	// LengthInMinutes
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`


	// TimeOffRequestId
	TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`

}

func (u *Wfmagentscheduleupdatetopicwfmscheduleactivity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmagentscheduleupdatetopicwfmscheduleactivity

	
	StartDate := new(string)
	if u.StartDate != nil {
		
		*StartDate = timeutil.Strftime(u.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	

	return json.Marshal(&struct { 
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`
		
		Description *string `json:"description,omitempty"`
		*Alias
	}{ 
		ActivityCodeId: u.ActivityCodeId,
		
		StartDate: StartDate,
		
		CountsAsPaidTime: u.CountsAsPaidTime,
		
		LengthInMinutes: u.LengthInMinutes,
		
		TimeOffRequestId: u.TimeOffRequestId,
		
		Description: u.Description,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmagentscheduleupdatetopicwfmscheduleactivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userscheduleactivity - Represents a single activity in a user's shift
type Userscheduleactivity struct { 
	// ActivityCodeId - The id for the activity code.  Look up a map of activity codes with the activities route
	ActivityCodeId *string `json:"activityCodeId,omitempty"`


	// StartDate - Start time in UTC for this activity, in ISO-8601 format
	StartDate *time.Time `json:"startDate,omitempty"`


	// LengthInMinutes - Length in minutes for this activity
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`


	// Description - Description for this activity
	Description *string `json:"description,omitempty"`


	// CountsAsPaidTime - Whether this activity is paid
	CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`


	// IsDstFallback - Whether this activity spans a DST fallback
	IsDstFallback *bool `json:"isDstFallback,omitempty"`


	// TimeOffRequestId - Time off request id of this activity
	TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`

}

func (u *Userscheduleactivity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userscheduleactivity

	
	StartDate := new(string)
	if u.StartDate != nil {
		
		*StartDate = timeutil.Strftime(u.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	

	return json.Marshal(&struct { 
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`
		
		IsDstFallback *bool `json:"isDstFallback,omitempty"`
		
		TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`
		*Alias
	}{ 
		ActivityCodeId: u.ActivityCodeId,
		
		StartDate: StartDate,
		
		LengthInMinutes: u.LengthInMinutes,
		
		Description: u.Description,
		
		CountsAsPaidTime: u.CountsAsPaidTime,
		
		IsDstFallback: u.IsDstFallback,
		
		TimeOffRequestId: u.TimeOffRequestId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userscheduleactivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

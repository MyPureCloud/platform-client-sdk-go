package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buagentscheduleactivity
type Buagentscheduleactivity struct { 
	// StartDate - The start date/time of this activity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`


	// LengthMinutes - The length of this activity in minutes
	LengthMinutes *int `json:"lengthMinutes,omitempty"`


	// Description - The description of this activity
	Description *string `json:"description,omitempty"`


	// ActivityCodeId - The ID of the activity code associated with this activity
	ActivityCodeId *string `json:"activityCodeId,omitempty"`


	// Paid - Whether this activity is paid
	Paid *bool `json:"paid,omitempty"`


	// TimeOffRequestId - The ID of the time off request associated with this activity, if applicable
	TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`


	// ExternalActivityId - The ID of the external activity associated with this activity, if applicable
	ExternalActivityId *string `json:"externalActivityId,omitempty"`


	// ExternalActivityType - The type of the external activity associated with this activity, if applicable
	ExternalActivityType *string `json:"externalActivityType,omitempty"`

}

func (u *Buagentscheduleactivity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buagentscheduleactivity

	
	StartDate := new(string)
	if u.StartDate != nil {
		
		*StartDate = timeutil.Strftime(u.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	

	return json.Marshal(&struct { 
		StartDate *string `json:"startDate,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		Paid *bool `json:"paid,omitempty"`
		
		TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`
		
		ExternalActivityId *string `json:"externalActivityId,omitempty"`
		
		ExternalActivityType *string `json:"externalActivityType,omitempty"`
		*Alias
	}{ 
		StartDate: StartDate,
		
		LengthMinutes: u.LengthMinutes,
		
		Description: u.Description,
		
		ActivityCodeId: u.ActivityCodeId,
		
		Paid: u.Paid,
		
		TimeOffRequestId: u.TimeOffRequestId,
		
		ExternalActivityId: u.ExternalActivityId,
		
		ExternalActivityType: u.ExternalActivityType,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Buagentscheduleactivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

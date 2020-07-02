package platformclientv2
import (
	"time"
	"encoding/json"
)

// Buagentscheduleactivity
type Buagentscheduleactivity struct { 
	// StartDate - The start date/time of this activity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	StartDate *time.Time `json:"startDate,omitempty"`


	// LengthMinutes - The length of this activity in minutes
	LengthMinutes *int32 `json:"lengthMinutes,omitempty"`


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

// String returns a JSON representation of the model
func (o *Buagentscheduleactivity) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

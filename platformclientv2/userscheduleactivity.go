package platformclientv2
import (
	"time"
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

// String returns a JSON representation of the model
func (o *Userscheduleactivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"time"
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

// String returns a JSON representation of the model
func (o *Wfmagentscheduleupdatetopicwfmscheduleactivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

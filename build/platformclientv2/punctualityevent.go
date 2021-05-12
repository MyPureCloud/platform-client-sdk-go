package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Punctualityevent
type Punctualityevent struct { 
	// DateScheduleStart - The scheduled activity start time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateScheduleStart *time.Time `json:"dateScheduleStart,omitempty"`


	// DateStart - The time the user started the activity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStart *time.Time `json:"dateStart,omitempty"`


	// LengthMinutes - The length of the activity in minutes
	LengthMinutes *int `json:"lengthMinutes,omitempty"`


	// Description - The description of the activity
	Description *string `json:"description,omitempty"`


	// ActivityCodeId - The ID of the activity code associated with this activity
	ActivityCodeId *string `json:"activityCodeId,omitempty"`


	// ActivityCode - The activity code
	ActivityCode *string `json:"activityCode,omitempty"`


	// Category - The category for the activity
	Category *string `json:"category,omitempty"`


	// Points - The points earned for this activity
	Points *int `json:"points,omitempty"`


	// Delta - Difference between this activity and the last activity in seconds
	Delta *float64 `json:"delta,omitempty"`


	// Bullseye
	Bullseye *bool `json:"bullseye,omitempty"`

}

// String returns a JSON representation of the model
func (o *Punctualityevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

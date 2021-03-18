package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Schedulegenerationwarning - Schedule generation warning
type Schedulegenerationwarning struct { 
	// UserId - ID of the user in the warning
	UserId *string `json:"userId,omitempty"`


	// UserNotLicensed - Whether the user does not have the appropriate license to be scheduled
	UserNotLicensed *bool `json:"userNotLicensed,omitempty"`


	// UnableToMeetMaxDays - Whether the number of scheduled days exceeded the maximum days to schedule defined in the agent work plan
	UnableToMeetMaxDays *bool `json:"unableToMeetMaxDays,omitempty"`


	// UnableToScheduleRequiredDays - Days indicated as required to work in agent work plan where no viable shift was found to schedule
	UnableToScheduleRequiredDays *[]string `json:"unableToScheduleRequiredDays,omitempty"`


	// UnableToMeetMinPaidForTheWeek - Whether the schedule did not meet the minimum paid time for the week defined in the agent work plan
	UnableToMeetMinPaidForTheWeek *bool `json:"unableToMeetMinPaidForTheWeek,omitempty"`


	// UnableToMeetMaxPaidForTheWeek - Whether the schedule exceeded the maximum paid time for the week defined in the agent work plan
	UnableToMeetMaxPaidForTheWeek *bool `json:"unableToMeetMaxPaidForTheWeek,omitempty"`


	// NoNeedDays - Days agent was scheduled but there was no need to meet. The scheduled days have no effect on service levels
	NoNeedDays *[]string `json:"noNeedDays,omitempty"`


	// ShiftsTooCloseTogether - Whether the schedule did not meet the minimum time between shifts defined in the agent work plan
	ShiftsTooCloseTogether *bool `json:"shiftsTooCloseTogether,omitempty"`

}

// String returns a JSON representation of the model
func (o *Schedulegenerationwarning) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

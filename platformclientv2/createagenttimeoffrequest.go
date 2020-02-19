package platformclientv2
import (
	"time"
	"encoding/json"
)

// Createagenttimeoffrequest
type Createagenttimeoffrequest struct { 
	// ActivityCodeId - The ID of the activity code associated with this time off request. Activity code must be of the TimeOff category
	ActivityCodeId *string `json:"activityCodeId,omitempty"`


	// Notes - Notes about the time off request
	Notes *string `json:"notes,omitempty"`


	// FullDayManagementUnitDates - A set of dates in yyyy-MM-dd format.  Should be interpreted in the management unit's configured time zone.
	FullDayManagementUnitDates *[]string `json:"fullDayManagementUnitDates,omitempty"`


	// PartialDayStartDateTimes - A set of start date-times in ISO-8601 format for partial day requests.
	PartialDayStartDateTimes *[]time.Time `json:"partialDayStartDateTimes,omitempty"`


	// DailyDurationMinutes - The daily duration of this time off request in minutes
	DailyDurationMinutes *int32 `json:"dailyDurationMinutes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Createagenttimeoffrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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
	DailyDurationMinutes *int `json:"dailyDurationMinutes,omitempty"`

}

func (u *Createagenttimeoffrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createagenttimeoffrequest

	

	return json.Marshal(&struct { 
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		FullDayManagementUnitDates *[]string `json:"fullDayManagementUnitDates,omitempty"`
		
		PartialDayStartDateTimes *[]time.Time `json:"partialDayStartDateTimes,omitempty"`
		
		DailyDurationMinutes *int `json:"dailyDurationMinutes,omitempty"`
		*Alias
	}{ 
		ActivityCodeId: u.ActivityCodeId,
		
		Notes: u.Notes,
		
		FullDayManagementUnitDates: u.FullDayManagementUnitDates,
		
		PartialDayStartDateTimes: u.PartialDayStartDateTimes,
		
		DailyDurationMinutes: u.DailyDurationMinutes,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Createagenttimeoffrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

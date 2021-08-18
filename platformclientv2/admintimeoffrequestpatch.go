package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Admintimeoffrequestpatch
type Admintimeoffrequestpatch struct { 
	// Status - The status of this time off request
	Status *string `json:"status,omitempty"`


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


	// Metadata - Version metadata for the time off request
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

func (u *Admintimeoffrequestpatch) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Admintimeoffrequestpatch

	

	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		FullDayManagementUnitDates *[]string `json:"fullDayManagementUnitDates,omitempty"`
		
		PartialDayStartDateTimes *[]time.Time `json:"partialDayStartDateTimes,omitempty"`
		
		DailyDurationMinutes *int `json:"dailyDurationMinutes,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		*Alias
	}{ 
		Status: u.Status,
		
		ActivityCodeId: u.ActivityCodeId,
		
		Notes: u.Notes,
		
		FullDayManagementUnitDates: u.FullDayManagementUnitDates,
		
		PartialDayStartDateTimes: u.PartialDayStartDateTimes,
		
		DailyDurationMinutes: u.DailyDurationMinutes,
		
		Metadata: u.Metadata,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Admintimeoffrequestpatch) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

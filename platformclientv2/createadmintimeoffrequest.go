package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createadmintimeoffrequest
type Createadmintimeoffrequest struct { 
	// Status - The status of this time off request
	Status *string `json:"status,omitempty"`


	// Users - A set of IDs for users to associate with this time off request
	Users *[]Userreference `json:"users,omitempty"`


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

func (o *Createadmintimeoffrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createadmintimeoffrequest
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		Users *[]Userreference `json:"users,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		FullDayManagementUnitDates *[]string `json:"fullDayManagementUnitDates,omitempty"`
		
		PartialDayStartDateTimes *[]time.Time `json:"partialDayStartDateTimes,omitempty"`
		
		DailyDurationMinutes *int `json:"dailyDurationMinutes,omitempty"`
		*Alias
	}{ 
		Status: o.Status,
		
		Users: o.Users,
		
		ActivityCodeId: o.ActivityCodeId,
		
		Notes: o.Notes,
		
		FullDayManagementUnitDates: o.FullDayManagementUnitDates,
		
		PartialDayStartDateTimes: o.PartialDayStartDateTimes,
		
		DailyDurationMinutes: o.DailyDurationMinutes,
		Alias:    (*Alias)(o),
	})
}

func (o *Createadmintimeoffrequest) UnmarshalJSON(b []byte) error {
	var CreateadmintimeoffrequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreateadmintimeoffrequestMap)
	if err != nil {
		return err
	}
	
	if Status, ok := CreateadmintimeoffrequestMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if Users, ok := CreateadmintimeoffrequestMap["users"].([]interface{}); ok {
		UsersString, _ := json.Marshal(Users)
		json.Unmarshal(UsersString, &o.Users)
	}
	
	if ActivityCodeId, ok := CreateadmintimeoffrequestMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
	
	if Notes, ok := CreateadmintimeoffrequestMap["notes"].(string); ok {
		o.Notes = &Notes
	}
	
	if FullDayManagementUnitDates, ok := CreateadmintimeoffrequestMap["fullDayManagementUnitDates"].([]interface{}); ok {
		FullDayManagementUnitDatesString, _ := json.Marshal(FullDayManagementUnitDates)
		json.Unmarshal(FullDayManagementUnitDatesString, &o.FullDayManagementUnitDates)
	}
	
	if PartialDayStartDateTimes, ok := CreateadmintimeoffrequestMap["partialDayStartDateTimes"].([]interface{}); ok {
		PartialDayStartDateTimesString, _ := json.Marshal(PartialDayStartDateTimes)
		json.Unmarshal(PartialDayStartDateTimesString, &o.PartialDayStartDateTimes)
	}
	
	if DailyDurationMinutes, ok := CreateadmintimeoffrequestMap["dailyDurationMinutes"].(float64); ok {
		DailyDurationMinutesInt := int(DailyDurationMinutes)
		o.DailyDurationMinutes = &DailyDurationMinutesInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createadmintimeoffrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

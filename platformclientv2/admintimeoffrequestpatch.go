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

func (o *Admintimeoffrequestpatch) MarshalJSON() ([]byte, error) {
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
		Status: o.Status,
		
		ActivityCodeId: o.ActivityCodeId,
		
		Notes: o.Notes,
		
		FullDayManagementUnitDates: o.FullDayManagementUnitDates,
		
		PartialDayStartDateTimes: o.PartialDayStartDateTimes,
		
		DailyDurationMinutes: o.DailyDurationMinutes,
		
		Metadata: o.Metadata,
		Alias:    (*Alias)(o),
	})
}

func (o *Admintimeoffrequestpatch) UnmarshalJSON(b []byte) error {
	var AdmintimeoffrequestpatchMap map[string]interface{}
	err := json.Unmarshal(b, &AdmintimeoffrequestpatchMap)
	if err != nil {
		return err
	}
	
	if Status, ok := AdmintimeoffrequestpatchMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if ActivityCodeId, ok := AdmintimeoffrequestpatchMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
	
	if Notes, ok := AdmintimeoffrequestpatchMap["notes"].(string); ok {
		o.Notes = &Notes
	}
	
	if FullDayManagementUnitDates, ok := AdmintimeoffrequestpatchMap["fullDayManagementUnitDates"].([]interface{}); ok {
		FullDayManagementUnitDatesString, _ := json.Marshal(FullDayManagementUnitDates)
		json.Unmarshal(FullDayManagementUnitDatesString, &o.FullDayManagementUnitDates)
	}
	
	if PartialDayStartDateTimes, ok := AdmintimeoffrequestpatchMap["partialDayStartDateTimes"].([]interface{}); ok {
		PartialDayStartDateTimesString, _ := json.Marshal(PartialDayStartDateTimes)
		json.Unmarshal(PartialDayStartDateTimesString, &o.PartialDayStartDateTimes)
	}
	
	if DailyDurationMinutes, ok := AdmintimeoffrequestpatchMap["dailyDurationMinutes"].(float64); ok {
		DailyDurationMinutesInt := int(DailyDurationMinutes)
		o.DailyDurationMinutes = &DailyDurationMinutesInt
	}
	
	if Metadata, ok := AdmintimeoffrequestpatchMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Admintimeoffrequestpatch) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

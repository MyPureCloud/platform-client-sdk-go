package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmtimeoffrequestupdatetopictimeoffrequestupdate
type Wfmtimeoffrequestupdatetopictimeoffrequestupdate struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// User
	User *Wfmtimeoffrequestupdatetopicuserreference `json:"user,omitempty"`


	// IsFullDayRequest
	IsFullDayRequest *bool `json:"isFullDayRequest,omitempty"`


	// MarkedAsRead
	MarkedAsRead *bool `json:"markedAsRead,omitempty"`


	// ActivityCodeId
	ActivityCodeId *string `json:"activityCodeId,omitempty"`


	// Paid
	Paid *bool `json:"paid,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`


	// Substatus
	Substatus *string `json:"substatus,omitempty"`


	// PartialDayStartDateTimes
	PartialDayStartDateTimes *[]string `json:"partialDayStartDateTimes,omitempty"`


	// FullDayManagementUnitDates
	FullDayManagementUnitDates *[]string `json:"fullDayManagementUnitDates,omitempty"`


	// DailyDurationMinutes
	DailyDurationMinutes *int `json:"dailyDurationMinutes,omitempty"`


	// Notes
	Notes *string `json:"notes,omitempty"`


	// ReviewedDate
	ReviewedDate *string `json:"reviewedDate,omitempty"`


	// ReviewedBy
	ReviewedBy *string `json:"reviewedBy,omitempty"`


	// SubmittedDate
	SubmittedDate *string `json:"submittedDate,omitempty"`


	// SubmittedBy
	SubmittedBy *string `json:"submittedBy,omitempty"`


	// ModifiedDate
	ModifiedDate *string `json:"modifiedDate,omitempty"`


	// ModifiedBy
	ModifiedBy *string `json:"modifiedBy,omitempty"`

}

func (o *Wfmtimeoffrequestupdatetopictimeoffrequestupdate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmtimeoffrequestupdatetopictimeoffrequestupdate
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		User *Wfmtimeoffrequestupdatetopicuserreference `json:"user,omitempty"`
		
		IsFullDayRequest *bool `json:"isFullDayRequest,omitempty"`
		
		MarkedAsRead *bool `json:"markedAsRead,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		Paid *bool `json:"paid,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Substatus *string `json:"substatus,omitempty"`
		
		PartialDayStartDateTimes *[]string `json:"partialDayStartDateTimes,omitempty"`
		
		FullDayManagementUnitDates *[]string `json:"fullDayManagementUnitDates,omitempty"`
		
		DailyDurationMinutes *int `json:"dailyDurationMinutes,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		ReviewedDate *string `json:"reviewedDate,omitempty"`
		
		ReviewedBy *string `json:"reviewedBy,omitempty"`
		
		SubmittedDate *string `json:"submittedDate,omitempty"`
		
		SubmittedBy *string `json:"submittedBy,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		ModifiedBy *string `json:"modifiedBy,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		User: o.User,
		
		IsFullDayRequest: o.IsFullDayRequest,
		
		MarkedAsRead: o.MarkedAsRead,
		
		ActivityCodeId: o.ActivityCodeId,
		
		Paid: o.Paid,
		
		Status: o.Status,
		
		Substatus: o.Substatus,
		
		PartialDayStartDateTimes: o.PartialDayStartDateTimes,
		
		FullDayManagementUnitDates: o.FullDayManagementUnitDates,
		
		DailyDurationMinutes: o.DailyDurationMinutes,
		
		Notes: o.Notes,
		
		ReviewedDate: o.ReviewedDate,
		
		ReviewedBy: o.ReviewedBy,
		
		SubmittedDate: o.SubmittedDate,
		
		SubmittedBy: o.SubmittedBy,
		
		ModifiedDate: o.ModifiedDate,
		
		ModifiedBy: o.ModifiedBy,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmtimeoffrequestupdatetopictimeoffrequestupdate) UnmarshalJSON(b []byte) error {
	var WfmtimeoffrequestupdatetopictimeoffrequestupdateMap map[string]interface{}
	err := json.Unmarshal(b, &WfmtimeoffrequestupdatetopictimeoffrequestupdateMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if User, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if IsFullDayRequest, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["isFullDayRequest"].(bool); ok {
		o.IsFullDayRequest = &IsFullDayRequest
	}
	
	if MarkedAsRead, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["markedAsRead"].(bool); ok {
		o.MarkedAsRead = &MarkedAsRead
	}
	
	if ActivityCodeId, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
	
	if Paid, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["paid"].(bool); ok {
		o.Paid = &Paid
	}
	
	if Status, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if Substatus, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["substatus"].(string); ok {
		o.Substatus = &Substatus
	}
	
	if PartialDayStartDateTimes, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["partialDayStartDateTimes"].([]interface{}); ok {
		PartialDayStartDateTimesString, _ := json.Marshal(PartialDayStartDateTimes)
		json.Unmarshal(PartialDayStartDateTimesString, &o.PartialDayStartDateTimes)
	}
	
	if FullDayManagementUnitDates, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["fullDayManagementUnitDates"].([]interface{}); ok {
		FullDayManagementUnitDatesString, _ := json.Marshal(FullDayManagementUnitDates)
		json.Unmarshal(FullDayManagementUnitDatesString, &o.FullDayManagementUnitDates)
	}
	
	if DailyDurationMinutes, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["dailyDurationMinutes"].(float64); ok {
		DailyDurationMinutesInt := int(DailyDurationMinutes)
		o.DailyDurationMinutes = &DailyDurationMinutesInt
	}
	
	if Notes, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["notes"].(string); ok {
		o.Notes = &Notes
	}
	
	if ReviewedDate, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["reviewedDate"].(string); ok {
		o.ReviewedDate = &ReviewedDate
	}
	
	if ReviewedBy, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["reviewedBy"].(string); ok {
		o.ReviewedBy = &ReviewedBy
	}
	
	if SubmittedDate, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["submittedDate"].(string); ok {
		o.SubmittedDate = &SubmittedDate
	}
	
	if SubmittedBy, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["submittedBy"].(string); ok {
		o.SubmittedBy = &SubmittedBy
	}
	
	if ModifiedDate, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["modifiedDate"].(string); ok {
		o.ModifiedDate = &ModifiedDate
	}
	
	if ModifiedBy, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmtimeoffrequestupdatetopictimeoffrequestupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

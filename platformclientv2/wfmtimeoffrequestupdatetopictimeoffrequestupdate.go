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

func (u *Wfmtimeoffrequestupdatetopictimeoffrequestupdate) MarshalJSON() ([]byte, error) {
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
		Id: u.Id,
		
		User: u.User,
		
		IsFullDayRequest: u.IsFullDayRequest,
		
		MarkedAsRead: u.MarkedAsRead,
		
		ActivityCodeId: u.ActivityCodeId,
		
		Paid: u.Paid,
		
		Status: u.Status,
		
		Substatus: u.Substatus,
		
		PartialDayStartDateTimes: u.PartialDayStartDateTimes,
		
		FullDayManagementUnitDates: u.FullDayManagementUnitDates,
		
		DailyDurationMinutes: u.DailyDurationMinutes,
		
		Notes: u.Notes,
		
		ReviewedDate: u.ReviewedDate,
		
		ReviewedBy: u.ReviewedBy,
		
		SubmittedDate: u.SubmittedDate,
		
		SubmittedBy: u.SubmittedBy,
		
		ModifiedDate: u.ModifiedDate,
		
		ModifiedBy: u.ModifiedBy,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmtimeoffrequestupdatetopictimeoffrequestupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

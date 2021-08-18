package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeoffrequest
type Timeoffrequest struct { 
	// Id - The id of the time off request
	Id *string `json:"id,omitempty"`


	// User - The user that the time off request belongs to
	User *Userreference `json:"user,omitempty"`


	// IsFullDayRequest - Whether this is a full day request (false means partial day)
	IsFullDayRequest *bool `json:"isFullDayRequest,omitempty"`


	// MarkedAsRead - Whether this request has been marked as read by the agent
	MarkedAsRead *bool `json:"markedAsRead,omitempty"`


	// ActivityCodeId - The ID of the activity code associated with this time off request. Activity code must be of the TimeOff category
	ActivityCodeId *string `json:"activityCodeId,omitempty"`


	// Status - The status of this time off request
	Status *string `json:"status,omitempty"`


	// PartialDayStartDateTimes - A set of start date-times in ISO-8601 format for partial day requests.  Will be not empty if isFullDayRequest == false
	PartialDayStartDateTimes *[]time.Time `json:"partialDayStartDateTimes,omitempty"`


	// FullDayManagementUnitDates - A set of dates in yyyy-MM-dd format.  Should be interpreted in the management unit's configured time zone.  Will be not empty if isFullDayRequest == true
	FullDayManagementUnitDates *[]string `json:"fullDayManagementUnitDates,omitempty"`


	// DailyDurationMinutes - The daily duration of this time off request in minutes
	DailyDurationMinutes *int `json:"dailyDurationMinutes,omitempty"`


	// Notes - Notes about the time off request
	Notes *string `json:"notes,omitempty"`


	// SubmittedBy - The user who submitted this time off request
	SubmittedBy *Userreference `json:"submittedBy,omitempty"`


	// SubmittedDate - The timestamp when this request was submitted. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	SubmittedDate *time.Time `json:"submittedDate,omitempty"`


	// ReviewedBy - The user who reviewed this time off request
	ReviewedBy *Userreference `json:"reviewedBy,omitempty"`


	// ReviewedDate - The timestamp when this request was reviewed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ReviewedDate *time.Time `json:"reviewedDate,omitempty"`


	// Metadata - The version metadata of the time off request
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Timeoffrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Timeoffrequest

	
	SubmittedDate := new(string)
	if u.SubmittedDate != nil {
		
		*SubmittedDate = timeutil.Strftime(u.SubmittedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		SubmittedDate = nil
	}
	
	ReviewedDate := new(string)
	if u.ReviewedDate != nil {
		
		*ReviewedDate = timeutil.Strftime(u.ReviewedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReviewedDate = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		User *Userreference `json:"user,omitempty"`
		
		IsFullDayRequest *bool `json:"isFullDayRequest,omitempty"`
		
		MarkedAsRead *bool `json:"markedAsRead,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		PartialDayStartDateTimes *[]time.Time `json:"partialDayStartDateTimes,omitempty"`
		
		FullDayManagementUnitDates *[]string `json:"fullDayManagementUnitDates,omitempty"`
		
		DailyDurationMinutes *int `json:"dailyDurationMinutes,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		SubmittedBy *Userreference `json:"submittedBy,omitempty"`
		
		SubmittedDate *string `json:"submittedDate,omitempty"`
		
		ReviewedBy *Userreference `json:"reviewedBy,omitempty"`
		
		ReviewedDate *string `json:"reviewedDate,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		User: u.User,
		
		IsFullDayRequest: u.IsFullDayRequest,
		
		MarkedAsRead: u.MarkedAsRead,
		
		ActivityCodeId: u.ActivityCodeId,
		
		Status: u.Status,
		
		PartialDayStartDateTimes: u.PartialDayStartDateTimes,
		
		FullDayManagementUnitDates: u.FullDayManagementUnitDates,
		
		DailyDurationMinutes: u.DailyDurationMinutes,
		
		Notes: u.Notes,
		
		SubmittedBy: u.SubmittedBy,
		
		SubmittedDate: SubmittedDate,
		
		ReviewedBy: u.ReviewedBy,
		
		ReviewedDate: ReviewedDate,
		
		Metadata: u.Metadata,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Timeoffrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"encoding/json"
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


	// Status
	Status *string `json:"status,omitempty"`


	// PartialDayStartDateTimes
	PartialDayStartDateTimes *[]string `json:"partialDayStartDateTimes,omitempty"`


	// FullDayManagementUnitDates
	FullDayManagementUnitDates *[]string `json:"fullDayManagementUnitDates,omitempty"`


	// DailyDurationMinutes
	DailyDurationMinutes *int32 `json:"dailyDurationMinutes,omitempty"`


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

// String returns a JSON representation of the model
func (o *Wfmtimeoffrequestupdatetopictimeoffrequestupdate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

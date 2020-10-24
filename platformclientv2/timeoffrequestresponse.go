package platformclientv2
import (
	"time"
	"encoding/json"
)

// Timeoffrequestresponse
type Timeoffrequestresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// User - The user associated with this time off request
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
	DailyDurationMinutes *int32 `json:"dailyDurationMinutes,omitempty"`


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


	// ModifiedBy - The user who last modified this TimeOffRequestResponse
	ModifiedBy *Userreference `json:"modifiedBy,omitempty"`


	// ModifiedDate - The timestamp when this request was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`


	// Metadata - The version metadata of the time off request
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Timeoffrequestresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

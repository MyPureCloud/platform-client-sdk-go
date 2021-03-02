package platformclientv2
import (
	"time"
	"encoding/json"
)

// Timeoffrequestnotification
type Timeoffrequestnotification struct { 
	// TimeOffRequestId - The ID of this time off request
	TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`


	// User - The user associated with this time off request
	User *Userreference `json:"user,omitempty"`


	// IsFullDayRequest - Whether this is a full day request (false means partial day)
	IsFullDayRequest *bool `json:"isFullDayRequest,omitempty"`


	// Status - The status of this time off request
	Status *string `json:"status,omitempty"`


	// PartialDayStartDateTimes - A set of start date-times in ISO-8601 format for partial day requests.  Will be not empty if isFullDayRequest == false
	PartialDayStartDateTimes *[]time.Time `json:"partialDayStartDateTimes,omitempty"`


	// FullDayManagementUnitDates - A set of dates in yyyy-MM-dd format.  Should be interpreted in the management unit's configured time zone.  Will be not empty if isFullDayRequest == true
	FullDayManagementUnitDates *[]string `json:"fullDayManagementUnitDates,omitempty"`

}

// String returns a JSON representation of the model
func (o *Timeoffrequestnotification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

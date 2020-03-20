package platformclientv2
import (
	"time"
	"encoding/json"
)

// Wfmusernotificationtopictimeoffrequestnotification
type Wfmusernotificationtopictimeoffrequestnotification struct { 
	// TimeOffRequestId
	TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`


	// User
	User *Wfmusernotificationtopicuserreference `json:"user,omitempty"`


	// IsFullDayRequest
	IsFullDayRequest *bool `json:"isFullDayRequest,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`


	// PartialDayStartDateTimes
	PartialDayStartDateTimes *[]time.Time `json:"partialDayStartDateTimes,omitempty"`


	// FullDayManagementUnitDates
	FullDayManagementUnitDates *[]string `json:"fullDayManagementUnitDates,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmusernotificationtopictimeoffrequestnotification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

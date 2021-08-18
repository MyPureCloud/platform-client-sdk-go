package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Wfmusernotificationtopictimeoffrequestnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmusernotificationtopictimeoffrequestnotification

	

	return json.Marshal(&struct { 
		TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`
		
		User *Wfmusernotificationtopicuserreference `json:"user,omitempty"`
		
		IsFullDayRequest *bool `json:"isFullDayRequest,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		PartialDayStartDateTimes *[]time.Time `json:"partialDayStartDateTimes,omitempty"`
		
		FullDayManagementUnitDates *[]string `json:"fullDayManagementUnitDates,omitempty"`
		*Alias
	}{ 
		TimeOffRequestId: u.TimeOffRequestId,
		
		User: u.User,
		
		IsFullDayRequest: u.IsFullDayRequest,
		
		Status: u.Status,
		
		PartialDayStartDateTimes: u.PartialDayStartDateTimes,
		
		FullDayManagementUnitDates: u.FullDayManagementUnitDates,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmusernotificationtopictimeoffrequestnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

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

func (o *Wfmusernotificationtopictimeoffrequestnotification) MarshalJSON() ([]byte, error) {
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
		TimeOffRequestId: o.TimeOffRequestId,
		
		User: o.User,
		
		IsFullDayRequest: o.IsFullDayRequest,
		
		Status: o.Status,
		
		PartialDayStartDateTimes: o.PartialDayStartDateTimes,
		
		FullDayManagementUnitDates: o.FullDayManagementUnitDates,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmusernotificationtopictimeoffrequestnotification) UnmarshalJSON(b []byte) error {
	var WfmusernotificationtopictimeoffrequestnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmusernotificationtopictimeoffrequestnotificationMap)
	if err != nil {
		return err
	}
	
	if TimeOffRequestId, ok := WfmusernotificationtopictimeoffrequestnotificationMap["timeOffRequestId"].(string); ok {
		o.TimeOffRequestId = &TimeOffRequestId
	}
    
	if User, ok := WfmusernotificationtopictimeoffrequestnotificationMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if IsFullDayRequest, ok := WfmusernotificationtopictimeoffrequestnotificationMap["isFullDayRequest"].(bool); ok {
		o.IsFullDayRequest = &IsFullDayRequest
	}
    
	if Status, ok := WfmusernotificationtopictimeoffrequestnotificationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if PartialDayStartDateTimes, ok := WfmusernotificationtopictimeoffrequestnotificationMap["partialDayStartDateTimes"].([]interface{}); ok {
		PartialDayStartDateTimesString, _ := json.Marshal(PartialDayStartDateTimes)
		json.Unmarshal(PartialDayStartDateTimesString, &o.PartialDayStartDateTimes)
	}
	
	if FullDayManagementUnitDates, ok := WfmusernotificationtopictimeoffrequestnotificationMap["fullDayManagementUnitDates"].([]interface{}); ok {
		FullDayManagementUnitDatesString, _ := json.Marshal(FullDayManagementUnitDates)
		json.Unmarshal(FullDayManagementUnitDatesString, &o.FullDayManagementUnitDates)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmusernotificationtopictimeoffrequestnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

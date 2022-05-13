package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createtimeoffplanrequest
type Createtimeoffplanrequest struct { 
	// Name - The name of this time off plan.
	Name *string `json:"name,omitempty"`


	// ActivityCodeIds - The set of activity code IDs to associate with this time off plan.
	ActivityCodeIds *[]string `json:"activityCodeIds,omitempty"`


	// TimeOffLimitIds - The set of time off limit IDs to associate with this time off plan.
	TimeOffLimitIds *[]string `json:"timeOffLimitIds,omitempty"`


	// AutoApprovalRule - Auto approval rule for the time off plan.
	AutoApprovalRule *string `json:"autoApprovalRule,omitempty"`


	// DaysBeforeStartToExpireFromWaitlist - The number of days before the time off request start date for when the request will be expired from the waitlist.
	DaysBeforeStartToExpireFromWaitlist *int `json:"daysBeforeStartToExpireFromWaitlist,omitempty"`


	// Active - Whether this time off plan should be used by agents.
	Active *bool `json:"active,omitempty"`

}

func (o *Createtimeoffplanrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createtimeoffplanrequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		ActivityCodeIds *[]string `json:"activityCodeIds,omitempty"`
		
		TimeOffLimitIds *[]string `json:"timeOffLimitIds,omitempty"`
		
		AutoApprovalRule *string `json:"autoApprovalRule,omitempty"`
		
		DaysBeforeStartToExpireFromWaitlist *int `json:"daysBeforeStartToExpireFromWaitlist,omitempty"`
		
		Active *bool `json:"active,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		ActivityCodeIds: o.ActivityCodeIds,
		
		TimeOffLimitIds: o.TimeOffLimitIds,
		
		AutoApprovalRule: o.AutoApprovalRule,
		
		DaysBeforeStartToExpireFromWaitlist: o.DaysBeforeStartToExpireFromWaitlist,
		
		Active: o.Active,
		Alias:    (*Alias)(o),
	})
}

func (o *Createtimeoffplanrequest) UnmarshalJSON(b []byte) error {
	var CreatetimeoffplanrequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreatetimeoffplanrequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := CreatetimeoffplanrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ActivityCodeIds, ok := CreatetimeoffplanrequestMap["activityCodeIds"].([]interface{}); ok {
		ActivityCodeIdsString, _ := json.Marshal(ActivityCodeIds)
		json.Unmarshal(ActivityCodeIdsString, &o.ActivityCodeIds)
	}
	
	if TimeOffLimitIds, ok := CreatetimeoffplanrequestMap["timeOffLimitIds"].([]interface{}); ok {
		TimeOffLimitIdsString, _ := json.Marshal(TimeOffLimitIds)
		json.Unmarshal(TimeOffLimitIdsString, &o.TimeOffLimitIds)
	}
	
	if AutoApprovalRule, ok := CreatetimeoffplanrequestMap["autoApprovalRule"].(string); ok {
		o.AutoApprovalRule = &AutoApprovalRule
	}
    
	if DaysBeforeStartToExpireFromWaitlist, ok := CreatetimeoffplanrequestMap["daysBeforeStartToExpireFromWaitlist"].(float64); ok {
		DaysBeforeStartToExpireFromWaitlistInt := int(DaysBeforeStartToExpireFromWaitlist)
		o.DaysBeforeStartToExpireFromWaitlist = &DaysBeforeStartToExpireFromWaitlistInt
	}
	
	if Active, ok := CreatetimeoffplanrequestMap["active"].(bool); ok {
		o.Active = &Active
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Createtimeoffplanrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

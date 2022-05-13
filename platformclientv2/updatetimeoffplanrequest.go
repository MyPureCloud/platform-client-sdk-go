package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Updatetimeoffplanrequest
type Updatetimeoffplanrequest struct { 
	// Name - The name of this time off plan.
	Name *string `json:"name,omitempty"`


	// ActivityCodeIds - The set of activity code IDs to associate with this time off plan.
	ActivityCodeIds *Setwrapperstring `json:"activityCodeIds,omitempty"`


	// TimeOffLimitIds - The set of time off limit IDs to associate with this time off plan.
	TimeOffLimitIds *Setwrapperstring `json:"timeOffLimitIds,omitempty"`


	// AutoApprovalRule - Auto approval rule for the time off plan.
	AutoApprovalRule *string `json:"autoApprovalRule,omitempty"`


	// DaysBeforeStartToExpireFromWaitlist - The number of days before the time off request start date for when the request will be expired from the waitlist.
	DaysBeforeStartToExpireFromWaitlist *int `json:"daysBeforeStartToExpireFromWaitlist,omitempty"`


	// Active - Whether this time off plan should be used by agents.
	Active *bool `json:"active,omitempty"`


	// Metadata - Version metadata for the time off plan
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

func (o *Updatetimeoffplanrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Updatetimeoffplanrequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		ActivityCodeIds *Setwrapperstring `json:"activityCodeIds,omitempty"`
		
		TimeOffLimitIds *Setwrapperstring `json:"timeOffLimitIds,omitempty"`
		
		AutoApprovalRule *string `json:"autoApprovalRule,omitempty"`
		
		DaysBeforeStartToExpireFromWaitlist *int `json:"daysBeforeStartToExpireFromWaitlist,omitempty"`
		
		Active *bool `json:"active,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		ActivityCodeIds: o.ActivityCodeIds,
		
		TimeOffLimitIds: o.TimeOffLimitIds,
		
		AutoApprovalRule: o.AutoApprovalRule,
		
		DaysBeforeStartToExpireFromWaitlist: o.DaysBeforeStartToExpireFromWaitlist,
		
		Active: o.Active,
		
		Metadata: o.Metadata,
		Alias:    (*Alias)(o),
	})
}

func (o *Updatetimeoffplanrequest) UnmarshalJSON(b []byte) error {
	var UpdatetimeoffplanrequestMap map[string]interface{}
	err := json.Unmarshal(b, &UpdatetimeoffplanrequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := UpdatetimeoffplanrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ActivityCodeIds, ok := UpdatetimeoffplanrequestMap["activityCodeIds"].(map[string]interface{}); ok {
		ActivityCodeIdsString, _ := json.Marshal(ActivityCodeIds)
		json.Unmarshal(ActivityCodeIdsString, &o.ActivityCodeIds)
	}
	
	if TimeOffLimitIds, ok := UpdatetimeoffplanrequestMap["timeOffLimitIds"].(map[string]interface{}); ok {
		TimeOffLimitIdsString, _ := json.Marshal(TimeOffLimitIds)
		json.Unmarshal(TimeOffLimitIdsString, &o.TimeOffLimitIds)
	}
	
	if AutoApprovalRule, ok := UpdatetimeoffplanrequestMap["autoApprovalRule"].(string); ok {
		o.AutoApprovalRule = &AutoApprovalRule
	}
    
	if DaysBeforeStartToExpireFromWaitlist, ok := UpdatetimeoffplanrequestMap["daysBeforeStartToExpireFromWaitlist"].(float64); ok {
		DaysBeforeStartToExpireFromWaitlistInt := int(DaysBeforeStartToExpireFromWaitlist)
		o.DaysBeforeStartToExpireFromWaitlist = &DaysBeforeStartToExpireFromWaitlistInt
	}
	
	if Active, ok := UpdatetimeoffplanrequestMap["active"].(bool); ok {
		o.Active = &Active
	}
    
	if Metadata, ok := UpdatetimeoffplanrequestMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Updatetimeoffplanrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

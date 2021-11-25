package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeoffplan
type Timeoffplan struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of this time off plan.
	Name *string `json:"name,omitempty"`


	// ActivityCodeIds - The set of activity code IDs associated with this time off plan.
	ActivityCodeIds *[]string `json:"activityCodeIds,omitempty"`


	// TimeOffLimits - The set of time off limit IDs associated with this time off plan.
	TimeOffLimits *[]Timeofflimitreference `json:"timeOffLimits,omitempty"`


	// AutoApprovalRule - Auto approval rule for this time off plan
	AutoApprovalRule *string `json:"autoApprovalRule,omitempty"`


	// DaysBeforeStartToExpireFromWaitlist - The number of days before the time off request start date for when the request will be expired from the waitlist.
	DaysBeforeStartToExpireFromWaitlist *int `json:"daysBeforeStartToExpireFromWaitlist,omitempty"`


	// Active - Whether this time off plan is currently being used by agents.
	Active *bool `json:"active,omitempty"`


	// Metadata - Version metadata for the time off plan.
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Timeoffplan) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Timeoffplan
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ActivityCodeIds *[]string `json:"activityCodeIds,omitempty"`
		
		TimeOffLimits *[]Timeofflimitreference `json:"timeOffLimits,omitempty"`
		
		AutoApprovalRule *string `json:"autoApprovalRule,omitempty"`
		
		DaysBeforeStartToExpireFromWaitlist *int `json:"daysBeforeStartToExpireFromWaitlist,omitempty"`
		
		Active *bool `json:"active,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ActivityCodeIds: o.ActivityCodeIds,
		
		TimeOffLimits: o.TimeOffLimits,
		
		AutoApprovalRule: o.AutoApprovalRule,
		
		DaysBeforeStartToExpireFromWaitlist: o.DaysBeforeStartToExpireFromWaitlist,
		
		Active: o.Active,
		
		Metadata: o.Metadata,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Timeoffplan) UnmarshalJSON(b []byte) error {
	var TimeoffplanMap map[string]interface{}
	err := json.Unmarshal(b, &TimeoffplanMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TimeoffplanMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := TimeoffplanMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if ActivityCodeIds, ok := TimeoffplanMap["activityCodeIds"].([]interface{}); ok {
		ActivityCodeIdsString, _ := json.Marshal(ActivityCodeIds)
		json.Unmarshal(ActivityCodeIdsString, &o.ActivityCodeIds)
	}
	
	if TimeOffLimits, ok := TimeoffplanMap["timeOffLimits"].([]interface{}); ok {
		TimeOffLimitsString, _ := json.Marshal(TimeOffLimits)
		json.Unmarshal(TimeOffLimitsString, &o.TimeOffLimits)
	}
	
	if AutoApprovalRule, ok := TimeoffplanMap["autoApprovalRule"].(string); ok {
		o.AutoApprovalRule = &AutoApprovalRule
	}
	
	if DaysBeforeStartToExpireFromWaitlist, ok := TimeoffplanMap["daysBeforeStartToExpireFromWaitlist"].(float64); ok {
		DaysBeforeStartToExpireFromWaitlistInt := int(DaysBeforeStartToExpireFromWaitlist)
		o.DaysBeforeStartToExpireFromWaitlist = &DaysBeforeStartToExpireFromWaitlistInt
	}
	
	if Active, ok := TimeoffplanMap["active"].(bool); ok {
		o.Active = &Active
	}
	
	if Metadata, ok := TimeoffplanMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if SelfUri, ok := TimeoffplanMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Timeoffplan) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

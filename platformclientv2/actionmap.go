package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Actionmap
type Actionmap struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Version - The version of the action map.
	Version *int `json:"version,omitempty"`


	// IsActive - Whether the action map is active.
	IsActive *bool `json:"isActive,omitempty"`


	// DisplayName - Display name of the action map.
	DisplayName *string `json:"displayName,omitempty"`


	// TriggerWithSegments - Trigger action map if any segment in the list is assigned to a given customer.
	TriggerWithSegments *[]string `json:"triggerWithSegments,omitempty"`


	// TriggerWithEventConditions - List of event conditions that must be satisfied to trigger the action map.
	TriggerWithEventConditions *[]Eventcondition `json:"triggerWithEventConditions,omitempty"`


	// TriggerWithOutcomeProbabilityConditions - Probability conditions for outcomes that must be satisfied to trigger the action map.
	TriggerWithOutcomeProbabilityConditions *[]Outcomeprobabilitycondition `json:"triggerWithOutcomeProbabilityConditions,omitempty"`


	// PageUrlConditions - URL conditions that a page must match for web actions to be displayable.
	PageUrlConditions *[]Urlcondition `json:"pageUrlConditions,omitempty"`


	// Activation - Type of activation.
	Activation *Activation `json:"activation,omitempty"`


	// Weight - Weight of the action map with higher number denoting higher weight.
	Weight *int `json:"weight,omitempty"`


	// Action - The action that will be executed if this action map is triggered.
	Action *Actionmapaction `json:"action,omitempty"`


	// ActionMapScheduleGroups - The action map's associated schedule groups.
	ActionMapScheduleGroups *Actionmapschedulegroups `json:"actionMapScheduleGroups,omitempty"`


	// IgnoreFrequencyCap - Override organization-level frequency cap and always offer web engagements from this action map.
	IgnoreFrequencyCap *bool `json:"ignoreFrequencyCap,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// CreatedDate - Timestamp indicating when the action map was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// ModifiedDate - Timestamp indicating when the action map was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`


	// StartDate - Timestamp at which the action map is scheduled to start firing. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - Timestamp at which the action map is scheduled to stop firing. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndDate *time.Time `json:"endDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Actionmap) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

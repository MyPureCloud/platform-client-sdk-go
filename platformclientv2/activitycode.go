package platformclientv2
import (
	"encoding/json"
)

// Activitycode - Activity code data
type Activitycode struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// Name - The name of the activity code. Default activity codes will be created with an empty name
	Name *string `json:"name,omitempty"`


	// IsActive - Whether this activity code is active or has been deleted
	IsActive *bool `json:"isActive,omitempty"`


	// IsDefault - Whether this is a default activity code
	IsDefault *bool `json:"isDefault,omitempty"`


	// Category - The activity code's category.
	Category *string `json:"category,omitempty"`


	// LengthInMinutes - The default length of the activity in minutes
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`


	// CountsAsPaidTime - Whether an agent is paid while performing this activity
	CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`


	// CountsAsWorkTime - Indicates whether or not the activity should be counted as contiguous work time for calculating daily constraints
	CountsAsWorkTime *bool `json:"countsAsWorkTime,omitempty"`


	// AgentTimeOffSelectable - Whether an agent can select this activity code when creating or editing a time off request. Null if the activity's category is not time off.
	AgentTimeOffSelectable *bool `json:"agentTimeOffSelectable,omitempty"`


	// Metadata - Version metadata for the associated management unit's list of activity codes
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

// String returns a JSON representation of the model
func (o *Activitycode) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

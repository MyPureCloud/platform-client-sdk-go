package platformclientv2
import (
	"time"
	"encoding/json"
)

// Trustgroup
type Trustgroup struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The group name.
	Name *string `json:"name,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// DateModified - Last modified date/time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// MemberCount - Number of members.
	MemberCount *int `json:"memberCount,omitempty"`


	// State - Active, inactive, or deleted state.
	State *string `json:"state,omitempty"`


	// Version - Current version for this resource.
	Version *int `json:"version,omitempty"`


	// VarType - Type of group.
	VarType *string `json:"type,omitempty"`


	// Images
	Images *[]Userimage `json:"images,omitempty"`


	// Addresses
	Addresses *[]Groupcontact `json:"addresses,omitempty"`


	// RulesVisible - Are membership rules visible to the person requesting to view the group
	RulesVisible *bool `json:"rulesVisible,omitempty"`


	// Visibility - Who can view this group
	Visibility *string `json:"visibility,omitempty"`


	// Owners - Owners of the group
	Owners *[]User `json:"owners,omitempty"`


	// DateCreated - The date on which the trusted group was added. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// CreatedBy - The user that added trusted group.
	CreatedBy *Orguser `json:"createdBy,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trustgroup) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

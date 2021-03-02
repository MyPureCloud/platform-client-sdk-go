package platformclientv2
import (
	"encoding/json"
)

// Updateuser
type Updateuser struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Chat
	Chat *Chat `json:"chat,omitempty"`


	// Department
	Department *string `json:"department,omitempty"`


	// Email
	Email *string `json:"email,omitempty"`


	// PrimaryContactInfo - The address(s) used for primary contact. Updates to the corresponding address in the addresses list will be reflected here.
	PrimaryContactInfo *[]Contact `json:"primaryContactInfo,omitempty"`


	// Addresses - Email address, phone number, and/or extension for this user. One entry is allowed per media type
	Addresses *[]Contact `json:"addresses,omitempty"`


	// Title
	Title *string `json:"title,omitempty"`


	// Username
	Username *string `json:"username,omitempty"`


	// Manager
	Manager *string `json:"manager,omitempty"`


	// Images
	Images *[]Userimage `json:"images,omitempty"`


	// Version - This value should be the current version of the user. The current version can be obtained with a GET on the user before doing a PATCH.
	Version *int `json:"version,omitempty"`


	// ProfileSkills - Profile skills possessed by the user
	ProfileSkills *[]string `json:"profileSkills,omitempty"`


	// Locations - The user placement at each site location.
	Locations *[]Location `json:"locations,omitempty"`


	// Groups - The groups the user is a member of
	Groups *[]Group `json:"groups,omitempty"`


	// State - The state of the user. This property can be used to restore a deleted user or transition between active and inactive. If specified, it is the only modifiable field.
	State *string `json:"state,omitempty"`


	// AcdAutoAnswer - The value that denotes if acdAutoAnswer is set on the user
	AcdAutoAnswer *bool `json:"acdAutoAnswer,omitempty"`


	// Certifications
	Certifications *[]string `json:"certifications,omitempty"`


	// Biography
	Biography *Biography `json:"biography,omitempty"`


	// EmployerInfo
	EmployerInfo *Employerinfo `json:"employerInfo,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Updateuser) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

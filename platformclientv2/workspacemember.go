package platformclientv2
import (
	"encoding/json"
)

// Workspacemember
type Workspacemember struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Workspace
	Workspace *Domainentityref `json:"workspace,omitempty"`


	// MemberType - The workspace member type.
	MemberType *string `json:"memberType,omitempty"`


	// Member
	Member *Domainentityref `json:"member,omitempty"`


	// User
	User *User `json:"user,omitempty"`


	// Group
	Group *Group `json:"group,omitempty"`


	// SecurityProfile
	SecurityProfile *Securityprofile `json:"securityProfile,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Workspacemember) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

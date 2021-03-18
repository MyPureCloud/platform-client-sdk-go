package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Share
type Share struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// SharedEntityType
	SharedEntityType *string `json:"sharedEntityType,omitempty"`


	// SharedEntity
	SharedEntity *Domainentityref `json:"sharedEntity,omitempty"`


	// MemberType
	MemberType *string `json:"memberType,omitempty"`


	// Member
	Member *Domainentityref `json:"member,omitempty"`


	// SharedBy
	SharedBy *Domainentityref `json:"sharedBy,omitempty"`


	// Workspace
	Workspace *Domainentityref `json:"workspace,omitempty"`


	// User
	User *User `json:"user,omitempty"`


	// Group
	Group *Group `json:"group,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Share) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"encoding/json"
)

// Createshareresponse
type Createshareresponse struct { 
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


	// Succeeded
	Succeeded *[]Share `json:"succeeded,omitempty"`


	// Failed
	Failed *[]Share `json:"failed,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Createshareresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

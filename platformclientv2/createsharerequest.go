package platformclientv2
import (
	"encoding/json"
)

// Createsharerequest
type Createsharerequest struct { 
	// SharedEntityType - The share entity type
	SharedEntityType *string `json:"sharedEntityType,omitempty"`


	// SharedEntity - The entity that will be shared
	SharedEntity *Sharedentity `json:"sharedEntity,omitempty"`


	// MemberType
	MemberType *string `json:"memberType,omitempty"`


	// Member - The member that will have access to this share. Only required if a list of members is not provided.
	Member *Sharedentity `json:"member,omitempty"`


	// Members
	Members *[]Createsharerequestmember `json:"members,omitempty"`

}

// String returns a JSON representation of the model
func (o *Createsharerequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

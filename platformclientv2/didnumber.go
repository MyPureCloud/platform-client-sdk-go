package platformclientv2
import (
	"encoding/json"
)

// Didnumber - Represents an unassigned or assigned DID in a DID Pool.
type Didnumber struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Number - The number of the DID formatted as E164.
	Number *string `json:"number,omitempty"`


	// Assigned - True if this DID is assigned to an entity.  False otherwise.
	Assigned *bool `json:"assigned,omitempty"`


	// DidPool - A Uri reference to the DID Pool this DID is a part of.
	DidPool *Addressableentityref `json:"didPool,omitempty"`


	// Owner - A Uri reference to the owner of this DID.  The owner's type can be found in ownerType.  If the DID is unassigned, this will be NULL.
	Owner *Domainentityref `json:"owner,omitempty"`


	// OwnerType - The type of the entity that owns this DID.  If the DID is unassigned, this will be NULL.
	OwnerType *string `json:"ownerType,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Didnumber) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Didnumber) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Didnumber

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Number *string `json:"number,omitempty"`
		
		Assigned *bool `json:"assigned,omitempty"`
		
		DidPool *Addressableentityref `json:"didPool,omitempty"`
		
		Owner *Domainentityref `json:"owner,omitempty"`
		
		OwnerType *string `json:"ownerType,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Number: u.Number,
		
		Assigned: u.Assigned,
		
		DidPool: u.DidPool,
		
		Owner: u.Owner,
		
		OwnerType: u.OwnerType,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Didnumber) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

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

func (o *Didnumber) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		Name: o.Name,
		
		Number: o.Number,
		
		Assigned: o.Assigned,
		
		DidPool: o.DidPool,
		
		Owner: o.Owner,
		
		OwnerType: o.OwnerType,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Didnumber) UnmarshalJSON(b []byte) error {
	var DidnumberMap map[string]interface{}
	err := json.Unmarshal(b, &DidnumberMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DidnumberMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := DidnumberMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Number, ok := DidnumberMap["number"].(string); ok {
		o.Number = &Number
	}
	
	if Assigned, ok := DidnumberMap["assigned"].(bool); ok {
		o.Assigned = &Assigned
	}
	
	if DidPool, ok := DidnumberMap["didPool"].(map[string]interface{}); ok {
		DidPoolString, _ := json.Marshal(DidPool)
		json.Unmarshal(DidPoolString, &o.DidPool)
	}
	
	if Owner, ok := DidnumberMap["owner"].(map[string]interface{}); ok {
		OwnerString, _ := json.Marshal(Owner)
		json.Unmarshal(OwnerString, &o.Owner)
	}
	
	if OwnerType, ok := DidnumberMap["ownerType"].(string); ok {
		o.OwnerType = &OwnerType
	}
	
	if SelfUri, ok := DidnumberMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Didnumber) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

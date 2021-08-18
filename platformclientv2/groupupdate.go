package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Groupupdate
type Groupupdate struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The group name.
	Name *string `json:"name,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// State - State of the group.
	State *string `json:"state,omitempty"`


	// Version - Current version for this resource.
	Version *int `json:"version,omitempty"`


	// Images
	Images *[]Userimage `json:"images,omitempty"`


	// Addresses
	Addresses *[]Groupcontact `json:"addresses,omitempty"`


	// RulesVisible - Are membership rules visible to the person requesting to view the group
	RulesVisible *bool `json:"rulesVisible,omitempty"`


	// Visibility - Who can view this group
	Visibility *string `json:"visibility,omitempty"`


	// OwnerIds - Owners of the group
	OwnerIds *[]string `json:"ownerIds,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Groupupdate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Groupupdate

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		Images *[]Userimage `json:"images,omitempty"`
		
		Addresses *[]Groupcontact `json:"addresses,omitempty"`
		
		RulesVisible *bool `json:"rulesVisible,omitempty"`
		
		Visibility *string `json:"visibility,omitempty"`
		
		OwnerIds *[]string `json:"ownerIds,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Description: u.Description,
		
		State: u.State,
		
		Version: u.Version,
		
		Images: u.Images,
		
		Addresses: u.Addresses,
		
		RulesVisible: u.RulesVisible,
		
		Visibility: u.Visibility,
		
		OwnerIds: u.OwnerIds,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Groupupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

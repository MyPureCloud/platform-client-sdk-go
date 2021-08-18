package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Locationdefinition
type Locationdefinition struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ContactUser - Site contact for the location entity
	ContactUser *Addressableentityref `json:"contactUser,omitempty"`


	// EmergencyNumber - Emergency number for the location entity
	EmergencyNumber *Locationemergencynumber `json:"emergencyNumber,omitempty"`


	// Address
	Address *Locationaddress `json:"address,omitempty"`


	// State - Current state of the location entity
	State *string `json:"state,omitempty"`


	// Notes - Notes for the location entity
	Notes *string `json:"notes,omitempty"`


	// Version - Current version of the location entity, value to be supplied should be retrieved by a GET or on create/update response
	Version *int `json:"version,omitempty"`


	// Path - A list of ancestor IDs in order
	Path *[]string `json:"path,omitempty"`


	// ProfileImage - Profile image of the location entity, retrieved with ?expand=images query parameter
	ProfileImage *[]Locationimage `json:"profileImage,omitempty"`


	// FloorplanImage - Floorplan images of the location entity, retrieved with ?expand=images query parameter
	FloorplanImage *[]Locationimage `json:"floorplanImage,omitempty"`


	// AddressVerificationDetails - Address verification information, retrieve dwith the ?expand=addressVerificationDetails query parameter
	AddressVerificationDetails *Locationaddressverificationdetails `json:"addressVerificationDetails,omitempty"`


	// AddressVerified - Boolean field which states if the address has been verified as an actual address
	AddressVerified *bool `json:"addressVerified,omitempty"`


	// AddressStored - Boolean field which states if the address has been stored for E911
	AddressStored *bool `json:"addressStored,omitempty"`


	// Images
	Images *string `json:"images,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Locationdefinition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Locationdefinition

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ContactUser *Addressableentityref `json:"contactUser,omitempty"`
		
		EmergencyNumber *Locationemergencynumber `json:"emergencyNumber,omitempty"`
		
		Address *Locationaddress `json:"address,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		Path *[]string `json:"path,omitempty"`
		
		ProfileImage *[]Locationimage `json:"profileImage,omitempty"`
		
		FloorplanImage *[]Locationimage `json:"floorplanImage,omitempty"`
		
		AddressVerificationDetails *Locationaddressverificationdetails `json:"addressVerificationDetails,omitempty"`
		
		AddressVerified *bool `json:"addressVerified,omitempty"`
		
		AddressStored *bool `json:"addressStored,omitempty"`
		
		Images *string `json:"images,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		ContactUser: u.ContactUser,
		
		EmergencyNumber: u.EmergencyNumber,
		
		Address: u.Address,
		
		State: u.State,
		
		Notes: u.Notes,
		
		Version: u.Version,
		
		Path: u.Path,
		
		ProfileImage: u.ProfileImage,
		
		FloorplanImage: u.FloorplanImage,
		
		AddressVerificationDetails: u.AddressVerificationDetails,
		
		AddressVerified: u.AddressVerified,
		
		AddressStored: u.AddressStored,
		
		Images: u.Images,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Locationdefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

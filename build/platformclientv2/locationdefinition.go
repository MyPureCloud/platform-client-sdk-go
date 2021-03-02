package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Locationdefinition) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

package platformclientv2
import (
	"encoding/json"
)

// Smsaddressprovision
type Smsaddressprovision struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - Name associated with this address
	Name *string `json:"name,omitempty"`


	// Street - The number and street address where this address is located.
	Street *string `json:"street,omitempty"`


	// City - The city in which this address is in
	City *string `json:"city,omitempty"`


	// Region - The state or region this address is in
	Region *string `json:"region,omitempty"`


	// PostalCode - The postal code this address is in
	PostalCode *string `json:"postalCode,omitempty"`


	// CountryCode - The ISO country code of this address
	CountryCode *string `json:"countryCode,omitempty"`


	// AutoCorrectAddress - This is used when the address is created. If the value is not set or true, then the system will, if necessary, auto-correct the address you provide. Set this value to false if the system should not auto-correct the address.
	AutoCorrectAddress *bool `json:"autoCorrectAddress,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Smsaddressprovision) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

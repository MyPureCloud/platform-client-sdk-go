package platformclientv2
import (
	"encoding/json"
)

// Smsavailablephonenumber
type Smsavailablephonenumber struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// PhoneNumber - A phone number available for provisioning in E.164 format. E.g. +13175555555 or +34234234234
	PhoneNumber *string `json:"phoneNumber,omitempty"`


	// CountryCode - The ISO 3166-1 alpha-2 country code of the country this phone number is associated with.
	CountryCode *string `json:"countryCode,omitempty"`


	// Region - The region/province/state the phone number is associated with.
	Region *string `json:"region,omitempty"`


	// City - The city the phone number is associated with.
	City *string `json:"city,omitempty"`


	// Capabilities - The capabilities of the phone number available for provisioning.
	Capabilities *[]string `json:"capabilities,omitempty"`


	// PhoneNumberType - The type of phone number available for provisioning.
	PhoneNumberType *string `json:"phoneNumberType,omitempty"`


	// AddressRequirement - The address requirement needed for provisioning this number. If there is a requirement, the address must be the residence or place of business of the individual or entity using the phone number.
	AddressRequirement *string `json:"addressRequirement,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Smsavailablephonenumber) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

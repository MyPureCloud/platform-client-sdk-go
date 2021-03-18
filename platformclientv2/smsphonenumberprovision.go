package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Smsphonenumberprovision
type Smsphonenumberprovision struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// PhoneNumber - A phone number to be used for SMS communications. E.g. +13175555555 or +34234234234
	PhoneNumber *string `json:"phoneNumber,omitempty"`


	// PhoneNumberType - Type of the phone number provisioned.
	PhoneNumberType *string `json:"phoneNumberType,omitempty"`


	// CountryCode - The ISO 3166-1 alpha-2 country code of the country this phone number is associated with.
	CountryCode *string `json:"countryCode,omitempty"`


	// AddressId - The id of an address added on your account. Due to regulatory requirements in some countries, an address may be required when provisioning a sms number. In those cases you should provide the provisioned sms address id here
	AddressId *string `json:"addressId,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Smsphonenumberprovision) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

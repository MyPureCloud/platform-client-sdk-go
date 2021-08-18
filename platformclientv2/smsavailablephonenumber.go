package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Smsavailablephonenumber) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Smsavailablephonenumber

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		PhoneNumber *string `json:"phoneNumber,omitempty"`
		
		CountryCode *string `json:"countryCode,omitempty"`
		
		Region *string `json:"region,omitempty"`
		
		City *string `json:"city,omitempty"`
		
		Capabilities *[]string `json:"capabilities,omitempty"`
		
		PhoneNumberType *string `json:"phoneNumberType,omitempty"`
		
		AddressRequirement *string `json:"addressRequirement,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		PhoneNumber: u.PhoneNumber,
		
		CountryCode: u.CountryCode,
		
		Region: u.Region,
		
		City: u.City,
		
		Capabilities: u.Capabilities,
		
		PhoneNumberType: u.PhoneNumberType,
		
		AddressRequirement: u.AddressRequirement,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Smsavailablephonenumber) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (o *Smsphonenumberprovision) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Smsphonenumberprovision
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		PhoneNumber *string `json:"phoneNumber,omitempty"`
		
		PhoneNumberType *string `json:"phoneNumberType,omitempty"`
		
		CountryCode *string `json:"countryCode,omitempty"`
		
		AddressId *string `json:"addressId,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		PhoneNumber: o.PhoneNumber,
		
		PhoneNumberType: o.PhoneNumberType,
		
		CountryCode: o.CountryCode,
		
		AddressId: o.AddressId,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Smsphonenumberprovision) UnmarshalJSON(b []byte) error {
	var SmsphonenumberprovisionMap map[string]interface{}
	err := json.Unmarshal(b, &SmsphonenumberprovisionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SmsphonenumberprovisionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := SmsphonenumberprovisionMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if PhoneNumber, ok := SmsphonenumberprovisionMap["phoneNumber"].(string); ok {
		o.PhoneNumber = &PhoneNumber
	}
    
	if PhoneNumberType, ok := SmsphonenumberprovisionMap["phoneNumberType"].(string); ok {
		o.PhoneNumberType = &PhoneNumberType
	}
    
	if CountryCode, ok := SmsphonenumberprovisionMap["countryCode"].(string); ok {
		o.CountryCode = &CountryCode
	}
    
	if AddressId, ok := SmsphonenumberprovisionMap["addressId"].(string); ok {
		o.AddressId = &AddressId
	}
    
	if SelfUri, ok := SmsphonenumberprovisionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Smsphonenumberprovision) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

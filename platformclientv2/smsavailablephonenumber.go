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

func (o *Smsavailablephonenumber) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		Name: o.Name,
		
		PhoneNumber: o.PhoneNumber,
		
		CountryCode: o.CountryCode,
		
		Region: o.Region,
		
		City: o.City,
		
		Capabilities: o.Capabilities,
		
		PhoneNumberType: o.PhoneNumberType,
		
		AddressRequirement: o.AddressRequirement,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Smsavailablephonenumber) UnmarshalJSON(b []byte) error {
	var SmsavailablephonenumberMap map[string]interface{}
	err := json.Unmarshal(b, &SmsavailablephonenumberMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SmsavailablephonenumberMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := SmsavailablephonenumberMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if PhoneNumber, ok := SmsavailablephonenumberMap["phoneNumber"].(string); ok {
		o.PhoneNumber = &PhoneNumber
	}
	
	if CountryCode, ok := SmsavailablephonenumberMap["countryCode"].(string); ok {
		o.CountryCode = &CountryCode
	}
	
	if Region, ok := SmsavailablephonenumberMap["region"].(string); ok {
		o.Region = &Region
	}
	
	if City, ok := SmsavailablephonenumberMap["city"].(string); ok {
		o.City = &City
	}
	
	if Capabilities, ok := SmsavailablephonenumberMap["capabilities"].([]interface{}); ok {
		CapabilitiesString, _ := json.Marshal(Capabilities)
		json.Unmarshal(CapabilitiesString, &o.Capabilities)
	}
	
	if PhoneNumberType, ok := SmsavailablephonenumberMap["phoneNumberType"].(string); ok {
		o.PhoneNumberType = &PhoneNumberType
	}
	
	if AddressRequirement, ok := SmsavailablephonenumberMap["addressRequirement"].(string); ok {
		o.AddressRequirement = &AddressRequirement
	}
	
	if SelfUri, ok := SmsavailablephonenumberMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Smsavailablephonenumber) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

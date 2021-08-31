package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Smsaddress
type Smsaddress struct { 
	// Id - The id of this address.
	Id *string `json:"id,omitempty"`


	// Name
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


	// Validated - In some countries, addresses are validated to comply with local regulation. In those countries, if the address you provide does not pass validation, it will not be accepted as an Address. This value will be true if the Address has been validated, or false for countries that don't require validation or if the Address is non-compliant.
	Validated *bool `json:"validated,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Smsaddress) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Smsaddress
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Street *string `json:"street,omitempty"`
		
		City *string `json:"city,omitempty"`
		
		Region *string `json:"region,omitempty"`
		
		PostalCode *string `json:"postalCode,omitempty"`
		
		CountryCode *string `json:"countryCode,omitempty"`
		
		Validated *bool `json:"validated,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Street: o.Street,
		
		City: o.City,
		
		Region: o.Region,
		
		PostalCode: o.PostalCode,
		
		CountryCode: o.CountryCode,
		
		Validated: o.Validated,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Smsaddress) UnmarshalJSON(b []byte) error {
	var SmsaddressMap map[string]interface{}
	err := json.Unmarshal(b, &SmsaddressMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SmsaddressMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := SmsaddressMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Street, ok := SmsaddressMap["street"].(string); ok {
		o.Street = &Street
	}
	
	if City, ok := SmsaddressMap["city"].(string); ok {
		o.City = &City
	}
	
	if Region, ok := SmsaddressMap["region"].(string); ok {
		o.Region = &Region
	}
	
	if PostalCode, ok := SmsaddressMap["postalCode"].(string); ok {
		o.PostalCode = &PostalCode
	}
	
	if CountryCode, ok := SmsaddressMap["countryCode"].(string); ok {
		o.CountryCode = &CountryCode
	}
	
	if Validated, ok := SmsaddressMap["validated"].(bool); ok {
		o.Validated = &Validated
	}
	
	if SelfUri, ok := SmsaddressMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Smsaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

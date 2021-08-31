package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Locationaddress
type Locationaddress struct { 
	// City
	City *string `json:"city,omitempty"`


	// Country
	Country *string `json:"country,omitempty"`


	// CountryName
	CountryName *string `json:"countryName,omitempty"`


	// State
	State *string `json:"state,omitempty"`


	// Street1
	Street1 *string `json:"street1,omitempty"`


	// Street2
	Street2 *string `json:"street2,omitempty"`


	// Zipcode
	Zipcode *string `json:"zipcode,omitempty"`

}

func (o *Locationaddress) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Locationaddress
	
	return json.Marshal(&struct { 
		City *string `json:"city,omitempty"`
		
		Country *string `json:"country,omitempty"`
		
		CountryName *string `json:"countryName,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Street1 *string `json:"street1,omitempty"`
		
		Street2 *string `json:"street2,omitempty"`
		
		Zipcode *string `json:"zipcode,omitempty"`
		*Alias
	}{ 
		City: o.City,
		
		Country: o.Country,
		
		CountryName: o.CountryName,
		
		State: o.State,
		
		Street1: o.Street1,
		
		Street2: o.Street2,
		
		Zipcode: o.Zipcode,
		Alias:    (*Alias)(o),
	})
}

func (o *Locationaddress) UnmarshalJSON(b []byte) error {
	var LocationaddressMap map[string]interface{}
	err := json.Unmarshal(b, &LocationaddressMap)
	if err != nil {
		return err
	}
	
	if City, ok := LocationaddressMap["city"].(string); ok {
		o.City = &City
	}
	
	if Country, ok := LocationaddressMap["country"].(string); ok {
		o.Country = &Country
	}
	
	if CountryName, ok := LocationaddressMap["countryName"].(string); ok {
		o.CountryName = &CountryName
	}
	
	if State, ok := LocationaddressMap["state"].(string); ok {
		o.State = &State
	}
	
	if Street1, ok := LocationaddressMap["street1"].(string); ok {
		o.Street1 = &Street1
	}
	
	if Street2, ok := LocationaddressMap["street2"].(string); ok {
		o.Street2 = &Street2
	}
	
	if Zipcode, ok := LocationaddressMap["zipcode"].(string); ok {
		o.Zipcode = &Zipcode
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Locationaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalcontactsunresolvedcontactchangedtopiccontactaddress
type Externalcontactsunresolvedcontactchangedtopiccontactaddress struct { 
	// Address1
	Address1 *string `json:"address1,omitempty"`


	// Address2
	Address2 *string `json:"address2,omitempty"`


	// City
	City *string `json:"city,omitempty"`


	// State
	State *string `json:"state,omitempty"`


	// PostalCode
	PostalCode *string `json:"postalCode,omitempty"`


	// CountryCode
	CountryCode *string `json:"countryCode,omitempty"`

}

func (o *Externalcontactsunresolvedcontactchangedtopiccontactaddress) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalcontactsunresolvedcontactchangedtopiccontactaddress
	
	return json.Marshal(&struct { 
		Address1 *string `json:"address1,omitempty"`
		
		Address2 *string `json:"address2,omitempty"`
		
		City *string `json:"city,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		PostalCode *string `json:"postalCode,omitempty"`
		
		CountryCode *string `json:"countryCode,omitempty"`
		*Alias
	}{ 
		Address1: o.Address1,
		
		Address2: o.Address2,
		
		City: o.City,
		
		State: o.State,
		
		PostalCode: o.PostalCode,
		
		CountryCode: o.CountryCode,
		Alias:    (*Alias)(o),
	})
}

func (o *Externalcontactsunresolvedcontactchangedtopiccontactaddress) UnmarshalJSON(b []byte) error {
	var ExternalcontactsunresolvedcontactchangedtopiccontactaddressMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalcontactsunresolvedcontactchangedtopiccontactaddressMap)
	if err != nil {
		return err
	}
	
	if Address1, ok := ExternalcontactsunresolvedcontactchangedtopiccontactaddressMap["address1"].(string); ok {
		o.Address1 = &Address1
	}
    
	if Address2, ok := ExternalcontactsunresolvedcontactchangedtopiccontactaddressMap["address2"].(string); ok {
		o.Address2 = &Address2
	}
    
	if City, ok := ExternalcontactsunresolvedcontactchangedtopiccontactaddressMap["city"].(string); ok {
		o.City = &City
	}
    
	if State, ok := ExternalcontactsunresolvedcontactchangedtopiccontactaddressMap["state"].(string); ok {
		o.State = &State
	}
    
	if PostalCode, ok := ExternalcontactsunresolvedcontactchangedtopiccontactaddressMap["postalCode"].(string); ok {
		o.PostalCode = &PostalCode
	}
    
	if CountryCode, ok := ExternalcontactsunresolvedcontactchangedtopiccontactaddressMap["countryCode"].(string); ok {
		o.CountryCode = &CountryCode
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externalcontactsunresolvedcontactchangedtopiccontactaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

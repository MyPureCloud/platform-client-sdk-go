package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalcontactscontactchangedtopiccontactaddress
type Externalcontactscontactchangedtopiccontactaddress struct { 
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

func (o *Externalcontactscontactchangedtopiccontactaddress) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalcontactscontactchangedtopiccontactaddress
	
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

func (o *Externalcontactscontactchangedtopiccontactaddress) UnmarshalJSON(b []byte) error {
	var ExternalcontactscontactchangedtopiccontactaddressMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalcontactscontactchangedtopiccontactaddressMap)
	if err != nil {
		return err
	}
	
	if Address1, ok := ExternalcontactscontactchangedtopiccontactaddressMap["address1"].(string); ok {
		o.Address1 = &Address1
	}
    
	if Address2, ok := ExternalcontactscontactchangedtopiccontactaddressMap["address2"].(string); ok {
		o.Address2 = &Address2
	}
    
	if City, ok := ExternalcontactscontactchangedtopiccontactaddressMap["city"].(string); ok {
		o.City = &City
	}
    
	if State, ok := ExternalcontactscontactchangedtopiccontactaddressMap["state"].(string); ok {
		o.State = &State
	}
    
	if PostalCode, ok := ExternalcontactscontactchangedtopiccontactaddressMap["postalCode"].(string); ok {
		o.PostalCode = &PostalCode
	}
    
	if CountryCode, ok := ExternalcontactscontactchangedtopiccontactaddressMap["countryCode"].(string); ok {
		o.CountryCode = &CountryCode
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externalcontactscontactchangedtopiccontactaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

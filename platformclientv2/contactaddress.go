package platformclientv2
import (
	"encoding/json"
)

// Contactaddress
type Contactaddress struct { 
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

// String returns a JSON representation of the model
func (o *Contactaddress) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

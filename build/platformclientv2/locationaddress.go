package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Locationaddress) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Contactaddress) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactaddress

	

	return json.Marshal(&struct { 
		Address1 *string `json:"address1,omitempty"`
		
		Address2 *string `json:"address2,omitempty"`
		
		City *string `json:"city,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		PostalCode *string `json:"postalCode,omitempty"`
		
		CountryCode *string `json:"countryCode,omitempty"`
		*Alias
	}{ 
		Address1: u.Address1,
		
		Address2: u.Address2,
		
		City: u.City,
		
		State: u.State,
		
		PostalCode: u.PostalCode,
		
		CountryCode: u.CountryCode,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Contactaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

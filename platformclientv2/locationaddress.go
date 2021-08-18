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

func (u *Locationaddress) MarshalJSON() ([]byte, error) {
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
		City: u.City,
		
		Country: u.Country,
		
		CountryName: u.CountryName,
		
		State: u.State,
		
		Street1: u.Street1,
		
		Street2: u.Street2,
		
		Zipcode: u.Zipcode,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Locationaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contact
type Contact struct { 
	// Address - Email address or phone number for this contact type
	Address *string `json:"address,omitempty"`


	// Display - Formatted version of the address property
	Display *string `json:"display,omitempty"`


	// MediaType
	MediaType *string `json:"mediaType,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// Extension - Use internal extension instead of address. Mutually exclusive with the address field.
	Extension *string `json:"extension,omitempty"`


	// CountryCode
	CountryCode *string `json:"countryCode,omitempty"`


	// Integration - Integration tag value if this number is associated with an external integration.
	Integration *string `json:"integration,omitempty"`

}

func (u *Contact) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contact

	

	return json.Marshal(&struct { 
		Address *string `json:"address,omitempty"`
		
		Display *string `json:"display,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Extension *string `json:"extension,omitempty"`
		
		CountryCode *string `json:"countryCode,omitempty"`
		
		Integration *string `json:"integration,omitempty"`
		*Alias
	}{ 
		Address: u.Address,
		
		Display: u.Display,
		
		MediaType: u.MediaType,
		
		VarType: u.VarType,
		
		Extension: u.Extension,
		
		CountryCode: u.CountryCode,
		
		Integration: u.Integration,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Contact) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Phonenumber
type Phonenumber struct { 
	// Display
	Display *string `json:"display,omitempty"`


	// Extension
	Extension *int `json:"extension,omitempty"`


	// AcceptsSMS
	AcceptsSMS *bool `json:"acceptsSMS,omitempty"`


	// UserInput
	UserInput *string `json:"userInput,omitempty"`


	// E164
	E164 *string `json:"e164,omitempty"`


	// CountryCode
	CountryCode *string `json:"countryCode,omitempty"`

}

func (u *Phonenumber) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Phonenumber

	

	return json.Marshal(&struct { 
		Display *string `json:"display,omitempty"`
		
		Extension *int `json:"extension,omitempty"`
		
		AcceptsSMS *bool `json:"acceptsSMS,omitempty"`
		
		UserInput *string `json:"userInput,omitempty"`
		
		E164 *string `json:"e164,omitempty"`
		
		CountryCode *string `json:"countryCode,omitempty"`
		*Alias
	}{ 
		Display: u.Display,
		
		Extension: u.Extension,
		
		AcceptsSMS: u.AcceptsSMS,
		
		UserInput: u.UserInput,
		
		E164: u.E164,
		
		CountryCode: u.CountryCode,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Phonenumber) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

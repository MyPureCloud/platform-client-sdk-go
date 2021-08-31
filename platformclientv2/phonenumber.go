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

func (o *Phonenumber) MarshalJSON() ([]byte, error) {
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
		Display: o.Display,
		
		Extension: o.Extension,
		
		AcceptsSMS: o.AcceptsSMS,
		
		UserInput: o.UserInput,
		
		E164: o.E164,
		
		CountryCode: o.CountryCode,
		Alias:    (*Alias)(o),
	})
}

func (o *Phonenumber) UnmarshalJSON(b []byte) error {
	var PhonenumberMap map[string]interface{}
	err := json.Unmarshal(b, &PhonenumberMap)
	if err != nil {
		return err
	}
	
	if Display, ok := PhonenumberMap["display"].(string); ok {
		o.Display = &Display
	}
	
	if Extension, ok := PhonenumberMap["extension"].(float64); ok {
		ExtensionInt := int(Extension)
		o.Extension = &ExtensionInt
	}
	
	if AcceptsSMS, ok := PhonenumberMap["acceptsSMS"].(bool); ok {
		o.AcceptsSMS = &AcceptsSMS
	}
	
	if UserInput, ok := PhonenumberMap["userInput"].(string); ok {
		o.UserInput = &UserInput
	}
	
	if E164, ok := PhonenumberMap["e164"].(string); ok {
		o.E164 = &E164
	}
	
	if CountryCode, ok := PhonenumberMap["countryCode"].(string); ok {
		o.CountryCode = &CountryCode
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Phonenumber) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

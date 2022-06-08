package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalcontactsunresolvedcontactchangedtopicphonenumber
type Externalcontactsunresolvedcontactchangedtopicphonenumber struct { 
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

func (o *Externalcontactsunresolvedcontactchangedtopicphonenumber) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalcontactsunresolvedcontactchangedtopicphonenumber
	
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

func (o *Externalcontactsunresolvedcontactchangedtopicphonenumber) UnmarshalJSON(b []byte) error {
	var ExternalcontactsunresolvedcontactchangedtopicphonenumberMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalcontactsunresolvedcontactchangedtopicphonenumberMap)
	if err != nil {
		return err
	}
	
	if Display, ok := ExternalcontactsunresolvedcontactchangedtopicphonenumberMap["display"].(string); ok {
		o.Display = &Display
	}
    
	if Extension, ok := ExternalcontactsunresolvedcontactchangedtopicphonenumberMap["extension"].(float64); ok {
		ExtensionInt := int(Extension)
		o.Extension = &ExtensionInt
	}
	
	if AcceptsSMS, ok := ExternalcontactsunresolvedcontactchangedtopicphonenumberMap["acceptsSMS"].(bool); ok {
		o.AcceptsSMS = &AcceptsSMS
	}
    
	if UserInput, ok := ExternalcontactsunresolvedcontactchangedtopicphonenumberMap["userInput"].(string); ok {
		o.UserInput = &UserInput
	}
    
	if E164, ok := ExternalcontactsunresolvedcontactchangedtopicphonenumberMap["e164"].(string); ok {
		o.E164 = &E164
	}
    
	if CountryCode, ok := ExternalcontactsunresolvedcontactchangedtopicphonenumberMap["countryCode"].(string); ok {
		o.CountryCode = &CountryCode
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externalcontactsunresolvedcontactchangedtopicphonenumber) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

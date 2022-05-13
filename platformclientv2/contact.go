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

func (o *Contact) MarshalJSON() ([]byte, error) {
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
		Address: o.Address,
		
		Display: o.Display,
		
		MediaType: o.MediaType,
		
		VarType: o.VarType,
		
		Extension: o.Extension,
		
		CountryCode: o.CountryCode,
		
		Integration: o.Integration,
		Alias:    (*Alias)(o),
	})
}

func (o *Contact) UnmarshalJSON(b []byte) error {
	var ContactMap map[string]interface{}
	err := json.Unmarshal(b, &ContactMap)
	if err != nil {
		return err
	}
	
	if Address, ok := ContactMap["address"].(string); ok {
		o.Address = &Address
	}
    
	if Display, ok := ContactMap["display"].(string); ok {
		o.Display = &Display
	}
    
	if MediaType, ok := ContactMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if VarType, ok := ContactMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Extension, ok := ContactMap["extension"].(string); ok {
		o.Extension = &Extension
	}
    
	if CountryCode, ok := ContactMap["countryCode"].(string); ok {
		o.CountryCode = &CountryCode
	}
    
	if Integration, ok := ContactMap["integration"].(string); ok {
		o.Integration = &Integration
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contact) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

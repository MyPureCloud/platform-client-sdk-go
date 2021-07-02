package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Contact) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Groupcontact
type Groupcontact struct { 
	// Address - Phone number for this contact type
	Address *string `json:"address,omitempty"`


	// Extension - Extension is set if the number is e164 valid
	Extension *string `json:"extension,omitempty"`


	// Display - Formatted version of the address property
	Display *string `json:"display,omitempty"`


	// VarType - Contact type of the address
	VarType *string `json:"type,omitempty"`


	// MediaType - Media type of the address
	MediaType *string `json:"mediaType,omitempty"`

}

// String returns a JSON representation of the model
func (o *Groupcontact) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

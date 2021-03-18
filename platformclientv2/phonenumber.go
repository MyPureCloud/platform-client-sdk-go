package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Phonenumber) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

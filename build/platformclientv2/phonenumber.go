package platformclientv2
import (
	"encoding/json"
)

// Phonenumber
type Phonenumber struct { 
	// Display
	Display *string `json:"display,omitempty"`


	// Extension
	Extension *int64 `json:"extension,omitempty"`


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
	return string(j)
}

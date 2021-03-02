package platformclientv2
import (
	"encoding/json"
)

// Supportedlanguage
type Supportedlanguage struct { 
	// Language - Architect supported language tag, e.g. en-us, es-us
	Language *string `json:"language,omitempty"`


	// IsDefault - Whether or not this language is the default language
	IsDefault *bool `json:"isDefault,omitempty"`

}

// String returns a JSON representation of the model
func (o *Supportedlanguage) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

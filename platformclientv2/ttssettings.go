package platformclientv2
import (
	"encoding/json"
)

// Ttssettings
type Ttssettings struct { 
	// DefaultEngine - ID of the global default TTS engine
	DefaultEngine *string `json:"defaultEngine,omitempty"`


	// LanguageOverrides - The list of default overrides for specific languages
	LanguageOverrides *[]Languageoverride `json:"languageOverrides,omitempty"`

}

// String returns a JSON representation of the model
func (o *Ttssettings) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

package platformclientv2
import (
	"encoding/json"
)

// Languageoverride
type Languageoverride struct { 
	// Language - The language code of the language being overridden
	Language *string `json:"language,omitempty"`


	// Engine - The ID of the TTS engine to use for this language override
	Engine *string `json:"engine,omitempty"`


	// Voice - The ID of the voice to use for this language override. The voice must be supported by the chosen engine.
	Voice *string `json:"voice,omitempty"`

}

// String returns a JSON representation of the model
func (o *Languageoverride) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

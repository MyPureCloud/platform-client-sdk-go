package platformclientv2
import (
	"encoding/json"
)

// Ttsvoiceentity
type Ttsvoiceentity struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Gender - The gender of the TTS voice
	Gender *string `json:"gender,omitempty"`


	// Language - The language supported by the TTS voice
	Language *string `json:"language,omitempty"`


	// Engine - Ths TTS engine this voice belongs to
	Engine *Ttsengineentity `json:"engine,omitempty"`


	// IsDefault - The voice is the default voice for its language
	IsDefault *bool `json:"isDefault,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Ttsvoiceentity) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

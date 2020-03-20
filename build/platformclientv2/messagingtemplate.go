package platformclientv2
import (
	"encoding/json"
)

// Messagingtemplate - The messaging template identifies a structured message templates supported by a messaging channel. For example, WhatsApp
type Messagingtemplate struct { 
	// Name - The messaging template name.
	Name *string `json:"name,omitempty"`


	// Namespace - The messaging template namespace.
	Namespace *string `json:"namespace,omitempty"`


	// Language - The messaging template language. For example, 'en-US'
	Language *string `json:"language,omitempty"`

}

// String returns a JSON representation of the model
func (o *Messagingtemplate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

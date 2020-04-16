package platformclientv2
import (
	"encoding/json"
)

// Messagingtemplate - The messaging template identifies a structured message templates supported by a messaging channel.
type Messagingtemplate struct { 
	// WhatsApp - Defines a messaging template for a WhatsApp messaging channel
	WhatsApp *Whatsappdefinition `json:"whatsApp,omitempty"`

}

// String returns a JSON representation of the model
func (o *Messagingtemplate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

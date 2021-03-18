package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Messagingtemplate - The messaging template identifies a structured message templates supported by a messaging channel.
type Messagingtemplate struct { 
	// WhatsApp - Defines a messaging template for a WhatsApp messaging channel
	WhatsApp *Whatsappdefinition `json:"whatsApp,omitempty"`

}

// String returns a JSON representation of the model
func (o *Messagingtemplate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

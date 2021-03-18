package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Whatsappdefinition - A WhatsApp messaging template definition as defined in the WhatsApp Business Manager
type Whatsappdefinition struct { 
	// Name - The messaging template name.
	Name *string `json:"name,omitempty"`


	// Namespace - The messaging template namespace.
	Namespace *string `json:"namespace,omitempty"`


	// Language - The messaging template language configured for this template. This is a WhatsApp specific value. For example, 'en_US'
	Language *string `json:"language,omitempty"`

}

// String returns a JSON representation of the model
func (o *Whatsappdefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

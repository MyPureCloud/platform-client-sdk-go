package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Whatsappintegrationrequest
type Whatsappintegrationrequest struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the WhatsApp Integration
	Name *string `json:"name,omitempty"`


	// PhoneNumber - The phone number associated to the whatsApp integration
	PhoneNumber *string `json:"phoneNumber,omitempty"`


	// WabaCertificate - The waba(WhatsApp Business Manager) certificate associated to the WhatsApp integration phone number
	WabaCertificate *string `json:"wabaCertificate,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Whatsappintegrationrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

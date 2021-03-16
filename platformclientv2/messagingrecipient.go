package platformclientv2
import (
	"encoding/json"
)

// Messagingrecipient - This is used to identify who the message is sent to, as well as who it was sent from. This information is channel specific - depends on capabilities to describe party by the platform
type Messagingrecipient struct { 
	// Nickname - Nickname/user name
	Nickname *string `json:"nickname,omitempty"`


	// Id - The recipient identifier specific for particular channel/integration. This is required when sending a message.
	Id *string `json:"id,omitempty"`


	// IdType - The recipient identifier type. This is used to indicate the format used by the recipient identifier.
	IdType *string `json:"idType,omitempty"`


	// Image - Avatar image
	Image *string `json:"image,omitempty"`


	// FirstName - Sender's first name
	FirstName *string `json:"firstName,omitempty"`


	// LastName - Sender's last name
	LastName *string `json:"lastName,omitempty"`


	// Email - Sender's email address
	Email *string `json:"email,omitempty"`

}

// String returns a JSON representation of the model
func (o *Messagingrecipient) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

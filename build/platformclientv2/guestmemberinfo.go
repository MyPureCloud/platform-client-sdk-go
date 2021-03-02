package platformclientv2
import (
	"encoding/json"
)

// Guestmemberinfo
type Guestmemberinfo struct { 
	// DisplayName - The display name to use for the guest member in the conversation.
	DisplayName *string `json:"displayName,omitempty"`


	// FirstName - The first name to use for the guest member in the conversation.
	FirstName *string `json:"firstName,omitempty"`


	// LastName - The last name to use for the guest member in the conversation.
	LastName *string `json:"lastName,omitempty"`


	// Email - The email address to use for the guest member in the conversation.
	Email *string `json:"email,omitempty"`


	// PhoneNumber - The phone number to use for the guest member in the conversation.
	PhoneNumber *string `json:"phoneNumber,omitempty"`


	// AvatarImageUrl - The URL to the avatar image to use for the guest member in the conversation, if any.
	AvatarImageUrl *string `json:"avatarImageUrl,omitempty"`


	// CustomFields - Any custom fields of information, in key-value format, to attach to the guest member in the conversation.
	CustomFields *map[string]string `json:"customFields,omitempty"`

}

// String returns a JSON representation of the model
func (o *Guestmemberinfo) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

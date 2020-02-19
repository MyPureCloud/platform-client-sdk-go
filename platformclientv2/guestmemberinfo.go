package platformclientv2
import (
	"encoding/json"
)

// Guestmemberinfo
type Guestmemberinfo struct { 
	// DisplayName - The display name to use for the guest member in the conversation.
	DisplayName *string `json:"displayName,omitempty"`


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

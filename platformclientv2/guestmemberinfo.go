package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Guestmemberinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Guestmemberinfo

	

	return json.Marshal(&struct { 
		DisplayName *string `json:"displayName,omitempty"`
		
		FirstName *string `json:"firstName,omitempty"`
		
		LastName *string `json:"lastName,omitempty"`
		
		Email *string `json:"email,omitempty"`
		
		PhoneNumber *string `json:"phoneNumber,omitempty"`
		
		AvatarImageUrl *string `json:"avatarImageUrl,omitempty"`
		
		CustomFields *map[string]string `json:"customFields,omitempty"`
		*Alias
	}{ 
		DisplayName: u.DisplayName,
		
		FirstName: u.FirstName,
		
		LastName: u.LastName,
		
		Email: u.Email,
		
		PhoneNumber: u.PhoneNumber,
		
		AvatarImageUrl: u.AvatarImageUrl,
		
		CustomFields: u.CustomFields,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Guestmemberinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

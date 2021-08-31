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

func (o *Guestmemberinfo) MarshalJSON() ([]byte, error) {
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
		DisplayName: o.DisplayName,
		
		FirstName: o.FirstName,
		
		LastName: o.LastName,
		
		Email: o.Email,
		
		PhoneNumber: o.PhoneNumber,
		
		AvatarImageUrl: o.AvatarImageUrl,
		
		CustomFields: o.CustomFields,
		Alias:    (*Alias)(o),
	})
}

func (o *Guestmemberinfo) UnmarshalJSON(b []byte) error {
	var GuestmemberinfoMap map[string]interface{}
	err := json.Unmarshal(b, &GuestmemberinfoMap)
	if err != nil {
		return err
	}
	
	if DisplayName, ok := GuestmemberinfoMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
	
	if FirstName, ok := GuestmemberinfoMap["firstName"].(string); ok {
		o.FirstName = &FirstName
	}
	
	if LastName, ok := GuestmemberinfoMap["lastName"].(string); ok {
		o.LastName = &LastName
	}
	
	if Email, ok := GuestmemberinfoMap["email"].(string); ok {
		o.Email = &Email
	}
	
	if PhoneNumber, ok := GuestmemberinfoMap["phoneNumber"].(string); ok {
		o.PhoneNumber = &PhoneNumber
	}
	
	if AvatarImageUrl, ok := GuestmemberinfoMap["avatarImageUrl"].(string); ok {
		o.AvatarImageUrl = &AvatarImageUrl
	}
	
	if CustomFields, ok := GuestmemberinfoMap["customFields"].(map[string]interface{}); ok {
		CustomFieldsString, _ := json.Marshal(CustomFields)
		json.Unmarshal(CustomFieldsString, &o.CustomFields)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Guestmemberinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

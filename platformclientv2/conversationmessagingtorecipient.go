package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmessagingtorecipient - Information about the recipient the message is sent to.
type Conversationmessagingtorecipient struct { 
	// Nickname - Nickname or display name of the recipient.
	Nickname *string `json:"nickname,omitempty"`


	// Id - The recipient ID specific to the provider.
	Id *string `json:"id,omitempty"`


	// IdType - The recipient ID type. This is used to indicate the format used for the ID.
	IdType *string `json:"idType,omitempty"`


	// Image - URL of an image that represents the recipient.
	Image *string `json:"image,omitempty"`


	// FirstName - First name of the recipient.
	FirstName *string `json:"firstName,omitempty"`


	// LastName - Last name of the recipient.
	LastName *string `json:"lastName,omitempty"`


	// Email - E-mail address of the recipient.
	Email *string `json:"email,omitempty"`

}

func (u *Conversationmessagingtorecipient) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationmessagingtorecipient

	

	return json.Marshal(&struct { 
		Nickname *string `json:"nickname,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		IdType *string `json:"idType,omitempty"`
		
		Image *string `json:"image,omitempty"`
		
		FirstName *string `json:"firstName,omitempty"`
		
		LastName *string `json:"lastName,omitempty"`
		
		Email *string `json:"email,omitempty"`
		*Alias
	}{ 
		Nickname: u.Nickname,
		
		Id: u.Id,
		
		IdType: u.IdType,
		
		Image: u.Image,
		
		FirstName: u.FirstName,
		
		LastName: u.LastName,
		
		Email: u.Email,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Conversationmessagingtorecipient) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

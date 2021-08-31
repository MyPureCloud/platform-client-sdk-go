package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmessagingfromrecipient - Information about the recipient the message is received from.
type Conversationmessagingfromrecipient struct { 
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

func (o *Conversationmessagingfromrecipient) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationmessagingfromrecipient
	
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
		Nickname: o.Nickname,
		
		Id: o.Id,
		
		IdType: o.IdType,
		
		Image: o.Image,
		
		FirstName: o.FirstName,
		
		LastName: o.LastName,
		
		Email: o.Email,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationmessagingfromrecipient) UnmarshalJSON(b []byte) error {
	var ConversationmessagingfromrecipientMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationmessagingfromrecipientMap)
	if err != nil {
		return err
	}
	
	if Nickname, ok := ConversationmessagingfromrecipientMap["nickname"].(string); ok {
		o.Nickname = &Nickname
	}
	
	if Id, ok := ConversationmessagingfromrecipientMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if IdType, ok := ConversationmessagingfromrecipientMap["idType"].(string); ok {
		o.IdType = &IdType
	}
	
	if Image, ok := ConversationmessagingfromrecipientMap["image"].(string); ok {
		o.Image = &Image
	}
	
	if FirstName, ok := ConversationmessagingfromrecipientMap["firstName"].(string); ok {
		o.FirstName = &FirstName
	}
	
	if LastName, ok := ConversationmessagingfromrecipientMap["lastName"].(string); ok {
		o.LastName = &LastName
	}
	
	if Email, ok := ConversationmessagingfromrecipientMap["email"].(string); ok {
		o.Email = &Email
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationmessagingfromrecipient) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

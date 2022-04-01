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


	// AdditionalIds - List of recipient additional identifiers
	AdditionalIds *[]Conversationrecipientadditionalidentifier `json:"additionalIds,omitempty"`

}

func (o *Conversationmessagingtorecipient) MarshalJSON() ([]byte, error) {
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
		
		AdditionalIds *[]Conversationrecipientadditionalidentifier `json:"additionalIds,omitempty"`
		*Alias
	}{ 
		Nickname: o.Nickname,
		
		Id: o.Id,
		
		IdType: o.IdType,
		
		Image: o.Image,
		
		FirstName: o.FirstName,
		
		LastName: o.LastName,
		
		Email: o.Email,
		
		AdditionalIds: o.AdditionalIds,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationmessagingtorecipient) UnmarshalJSON(b []byte) error {
	var ConversationmessagingtorecipientMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationmessagingtorecipientMap)
	if err != nil {
		return err
	}
	
	if Nickname, ok := ConversationmessagingtorecipientMap["nickname"].(string); ok {
		o.Nickname = &Nickname
	}
	
	if Id, ok := ConversationmessagingtorecipientMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if IdType, ok := ConversationmessagingtorecipientMap["idType"].(string); ok {
		o.IdType = &IdType
	}
	
	if Image, ok := ConversationmessagingtorecipientMap["image"].(string); ok {
		o.Image = &Image
	}
	
	if FirstName, ok := ConversationmessagingtorecipientMap["firstName"].(string); ok {
		o.FirstName = &FirstName
	}
	
	if LastName, ok := ConversationmessagingtorecipientMap["lastName"].(string); ok {
		o.LastName = &LastName
	}
	
	if Email, ok := ConversationmessagingtorecipientMap["email"].(string); ok {
		o.Email = &Email
	}
	
	if AdditionalIds, ok := ConversationmessagingtorecipientMap["additionalIds"].([]interface{}); ok {
		AdditionalIdsString, _ := json.Marshal(AdditionalIds)
		json.Unmarshal(AdditionalIdsString, &o.AdditionalIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationmessagingtorecipient) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

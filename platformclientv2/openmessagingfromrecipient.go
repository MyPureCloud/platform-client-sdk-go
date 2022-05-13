package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Openmessagingfromrecipient - Information about the recipient the message is received from.
type Openmessagingfromrecipient struct { 
	// Nickname - Nickname or display name of the recipient.
	Nickname *string `json:"nickname,omitempty"`


	// Id - The recipient ID specific to the provider.
	Id *string `json:"id,omitempty"`


	// IdType - The recipient ID type. This is used to indicate the format used for the ID.
	IdType *string `json:"idType,omitempty"`


	// FirstName - First name of the recipient.
	FirstName *string `json:"firstName,omitempty"`


	// LastName - Last name of the recipient.
	LastName *string `json:"lastName,omitempty"`


	// Image - URL of an image that represents the recipient.
	Image *string `json:"image,omitempty"`


	// Email - E-mail address of the recipient.
	Email *string `json:"email,omitempty"`

}

func (o *Openmessagingfromrecipient) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Openmessagingfromrecipient
	
	return json.Marshal(&struct { 
		Nickname *string `json:"nickname,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		IdType *string `json:"idType,omitempty"`
		
		FirstName *string `json:"firstName,omitempty"`
		
		LastName *string `json:"lastName,omitempty"`
		
		Image *string `json:"image,omitempty"`
		
		Email *string `json:"email,omitempty"`
		*Alias
	}{ 
		Nickname: o.Nickname,
		
		Id: o.Id,
		
		IdType: o.IdType,
		
		FirstName: o.FirstName,
		
		LastName: o.LastName,
		
		Image: o.Image,
		
		Email: o.Email,
		Alias:    (*Alias)(o),
	})
}

func (o *Openmessagingfromrecipient) UnmarshalJSON(b []byte) error {
	var OpenmessagingfromrecipientMap map[string]interface{}
	err := json.Unmarshal(b, &OpenmessagingfromrecipientMap)
	if err != nil {
		return err
	}
	
	if Nickname, ok := OpenmessagingfromrecipientMap["nickname"].(string); ok {
		o.Nickname = &Nickname
	}
    
	if Id, ok := OpenmessagingfromrecipientMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if IdType, ok := OpenmessagingfromrecipientMap["idType"].(string); ok {
		o.IdType = &IdType
	}
    
	if FirstName, ok := OpenmessagingfromrecipientMap["firstName"].(string); ok {
		o.FirstName = &FirstName
	}
    
	if LastName, ok := OpenmessagingfromrecipientMap["lastName"].(string); ok {
		o.LastName = &LastName
	}
    
	if Image, ok := OpenmessagingfromrecipientMap["image"].(string); ok {
		o.Image = &Image
	}
    
	if Email, ok := OpenmessagingfromrecipientMap["email"].(string); ok {
		o.Email = &Email
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Openmessagingfromrecipient) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

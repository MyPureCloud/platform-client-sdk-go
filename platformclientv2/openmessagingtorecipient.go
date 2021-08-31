package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Openmessagingtorecipient - Information about the recipient the message is sent to.
type Openmessagingtorecipient struct { 
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

}

func (o *Openmessagingtorecipient) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Openmessagingtorecipient
	
	return json.Marshal(&struct { 
		Nickname *string `json:"nickname,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		IdType *string `json:"idType,omitempty"`
		
		FirstName *string `json:"firstName,omitempty"`
		
		LastName *string `json:"lastName,omitempty"`
		*Alias
	}{ 
		Nickname: o.Nickname,
		
		Id: o.Id,
		
		IdType: o.IdType,
		
		FirstName: o.FirstName,
		
		LastName: o.LastName,
		Alias:    (*Alias)(o),
	})
}

func (o *Openmessagingtorecipient) UnmarshalJSON(b []byte) error {
	var OpenmessagingtorecipientMap map[string]interface{}
	err := json.Unmarshal(b, &OpenmessagingtorecipientMap)
	if err != nil {
		return err
	}
	
	if Nickname, ok := OpenmessagingtorecipientMap["nickname"].(string); ok {
		o.Nickname = &Nickname
	}
	
	if Id, ok := OpenmessagingtorecipientMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if IdType, ok := OpenmessagingtorecipientMap["idType"].(string); ok {
		o.IdType = &IdType
	}
	
	if FirstName, ok := OpenmessagingtorecipientMap["firstName"].(string); ok {
		o.FirstName = &FirstName
	}
	
	if LastName, ok := OpenmessagingtorecipientMap["lastName"].(string); ok {
		o.LastName = &LastName
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Openmessagingtorecipient) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforusertopicconversationmessagingfromrecipient
type V2conversationmessagetypingeventforusertopicconversationmessagingfromrecipient struct { 
	// Nickname
	Nickname *string `json:"nickname,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`


	// IdType
	IdType *string `json:"idType,omitempty"`


	// Image
	Image *string `json:"image,omitempty"`


	// FirstName
	FirstName *string `json:"firstName,omitempty"`


	// LastName
	LastName *string `json:"lastName,omitempty"`


	// Email
	Email *string `json:"email,omitempty"`


	// AdditionalIds
	AdditionalIds *[]V2conversationmessagetypingeventforusertopicconversationrecipientadditionalidentifier `json:"additionalIds,omitempty"`

}

func (o *V2conversationmessagetypingeventforusertopicconversationmessagingfromrecipient) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforusertopicconversationmessagingfromrecipient
	
	return json.Marshal(&struct { 
		Nickname *string `json:"nickname,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		IdType *string `json:"idType,omitempty"`
		
		Image *string `json:"image,omitempty"`
		
		FirstName *string `json:"firstName,omitempty"`
		
		LastName *string `json:"lastName,omitempty"`
		
		Email *string `json:"email,omitempty"`
		
		AdditionalIds *[]V2conversationmessagetypingeventforusertopicconversationrecipientadditionalidentifier `json:"additionalIds,omitempty"`
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

func (o *V2conversationmessagetypingeventforusertopicconversationmessagingfromrecipient) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforusertopicconversationmessagingfromrecipientMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforusertopicconversationmessagingfromrecipientMap)
	if err != nil {
		return err
	}
	
	if Nickname, ok := V2conversationmessagetypingeventforusertopicconversationmessagingfromrecipientMap["nickname"].(string); ok {
		o.Nickname = &Nickname
	}
	
	if Id, ok := V2conversationmessagetypingeventforusertopicconversationmessagingfromrecipientMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if IdType, ok := V2conversationmessagetypingeventforusertopicconversationmessagingfromrecipientMap["idType"].(string); ok {
		o.IdType = &IdType
	}
	
	if Image, ok := V2conversationmessagetypingeventforusertopicconversationmessagingfromrecipientMap["image"].(string); ok {
		o.Image = &Image
	}
	
	if FirstName, ok := V2conversationmessagetypingeventforusertopicconversationmessagingfromrecipientMap["firstName"].(string); ok {
		o.FirstName = &FirstName
	}
	
	if LastName, ok := V2conversationmessagetypingeventforusertopicconversationmessagingfromrecipientMap["lastName"].(string); ok {
		o.LastName = &LastName
	}
	
	if Email, ok := V2conversationmessagetypingeventforusertopicconversationmessagingfromrecipientMap["email"].(string); ok {
		o.Email = &Email
	}
	
	if AdditionalIds, ok := V2conversationmessagetypingeventforusertopicconversationmessagingfromrecipientMap["additionalIds"].([]interface{}); ok {
		AdditionalIdsString, _ := json.Marshal(AdditionalIds)
		json.Unmarshal(AdditionalIdsString, &o.AdditionalIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforusertopicconversationmessagingfromrecipient) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

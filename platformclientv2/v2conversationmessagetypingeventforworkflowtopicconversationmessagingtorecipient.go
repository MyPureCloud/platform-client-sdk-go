package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipient
type V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipient struct { 
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
	AdditionalIds *[]V2conversationmessagetypingeventforworkflowtopicconversationrecipientadditionalidentifier `json:"additionalIds,omitempty"`

}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipient) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipient
	
	return json.Marshal(&struct { 
		Nickname *string `json:"nickname,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		IdType *string `json:"idType,omitempty"`
		
		Image *string `json:"image,omitempty"`
		
		FirstName *string `json:"firstName,omitempty"`
		
		LastName *string `json:"lastName,omitempty"`
		
		Email *string `json:"email,omitempty"`
		
		AdditionalIds *[]V2conversationmessagetypingeventforworkflowtopicconversationrecipientadditionalidentifier `json:"additionalIds,omitempty"`
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

func (o *V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipient) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipientMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipientMap)
	if err != nil {
		return err
	}
	
	if Nickname, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipientMap["nickname"].(string); ok {
		o.Nickname = &Nickname
	}
    
	if Id, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipientMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if IdType, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipientMap["idType"].(string); ok {
		o.IdType = &IdType
	}
    
	if Image, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipientMap["image"].(string); ok {
		o.Image = &Image
	}
    
	if FirstName, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipientMap["firstName"].(string); ok {
		o.FirstName = &FirstName
	}
    
	if LastName, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipientMap["lastName"].(string); ok {
		o.LastName = &LastName
	}
    
	if Email, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipientMap["email"].(string); ok {
		o.Email = &Email
	}
    
	if AdditionalIds, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipientMap["additionalIds"].([]interface{}); ok {
		AdditionalIdsString, _ := json.Marshal(AdditionalIds)
		json.Unmarshal(AdditionalIdsString, &o.AdditionalIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipient) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

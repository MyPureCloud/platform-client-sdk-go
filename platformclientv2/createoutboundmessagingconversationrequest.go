package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createoutboundmessagingconversationrequest
type Createoutboundmessagingconversationrequest struct { 
	// QueueId - The ID of the queue to be associated with the message. This will determine the fromAddress of the message.
	QueueId *string `json:"queueId,omitempty"`


	// ToAddress - The messaging address of the recipient of the message. For an SMS messenger type, the phone number address must be in E.164 format. E.g. +13175555555 or +34234234234
	ToAddress *string `json:"toAddress,omitempty"`


	// ToAddressMessengerType - The messaging address messenger type.
	ToAddressMessengerType *string `json:"toAddressMessengerType,omitempty"`


	// UseExistingConversation - An override to use an existing conversation.  If set to true, an existing conversation will be used if there is one within the conversation window.  If set to false, create request fails if there is a conversation within the conversation window.
	UseExistingConversation *bool `json:"useExistingConversation,omitempty"`


	// ExternalContactId - The external contact with which the message will be associated.
	ExternalContactId *string `json:"externalContactId,omitempty"`

}

func (o *Createoutboundmessagingconversationrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createoutboundmessagingconversationrequest
	
	return json.Marshal(&struct { 
		QueueId *string `json:"queueId,omitempty"`
		
		ToAddress *string `json:"toAddress,omitempty"`
		
		ToAddressMessengerType *string `json:"toAddressMessengerType,omitempty"`
		
		UseExistingConversation *bool `json:"useExistingConversation,omitempty"`
		
		ExternalContactId *string `json:"externalContactId,omitempty"`
		*Alias
	}{ 
		QueueId: o.QueueId,
		
		ToAddress: o.ToAddress,
		
		ToAddressMessengerType: o.ToAddressMessengerType,
		
		UseExistingConversation: o.UseExistingConversation,
		
		ExternalContactId: o.ExternalContactId,
		Alias:    (*Alias)(o),
	})
}

func (o *Createoutboundmessagingconversationrequest) UnmarshalJSON(b []byte) error {
	var CreateoutboundmessagingconversationrequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreateoutboundmessagingconversationrequestMap)
	if err != nil {
		return err
	}
	
	if QueueId, ok := CreateoutboundmessagingconversationrequestMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
	
	if ToAddress, ok := CreateoutboundmessagingconversationrequestMap["toAddress"].(string); ok {
		o.ToAddress = &ToAddress
	}
	
	if ToAddressMessengerType, ok := CreateoutboundmessagingconversationrequestMap["toAddressMessengerType"].(string); ok {
		o.ToAddressMessengerType = &ToAddressMessengerType
	}
	
	if UseExistingConversation, ok := CreateoutboundmessagingconversationrequestMap["useExistingConversation"].(bool); ok {
		o.UseExistingConversation = &UseExistingConversation
	}
	
	if ExternalContactId, ok := CreateoutboundmessagingconversationrequestMap["externalContactId"].(string); ok {
		o.ExternalContactId = &ExternalContactId
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createoutboundmessagingconversationrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

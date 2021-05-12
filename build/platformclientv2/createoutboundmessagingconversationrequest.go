package platformclientv2
import (
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


	// ExternalContactId - The external contact Id of the recipient of the message.
	ExternalContactId *string `json:"externalContactId,omitempty"`


	// ExternalOrganizationId - The external organization Id of the recipient of the message.
	ExternalOrganizationId *string `json:"externalOrganizationId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Createoutboundmessagingconversationrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

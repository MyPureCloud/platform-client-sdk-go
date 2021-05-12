package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Sendagentlessoutboundmessagerequest
type Sendagentlessoutboundmessagerequest struct { 
	// FromAddress - The messaging address of the sender of the message. For an SMS messenger type, this must be a currently provisioned SMS phone number. For a WhatsApp messenger type use the provisioned WhatsApp integration’s ID
	FromAddress *string `json:"fromAddress,omitempty"`


	// ToAddress - The messaging address of the recipient of the message. For an SMS messenger type, the phone number address must be in E.164 format. E.g. +13175555555 or +34234234234.
	ToAddress *string `json:"toAddress,omitempty"`


	// ToAddressMessengerType - The recipient messaging address messenger type. Currently SMS is the only one supported. WhatsApp will be supported in a future release.
	ToAddressMessengerType *string `json:"toAddressMessengerType,omitempty"`


	// TextBody - The text of the message to send. This field is required in the case of SMS messenger type
	TextBody *string `json:"textBody,omitempty"`


	// MessagingTemplate - The messaging template to use in the case of WhatsApp messenger type. This field is required when using WhatsApp messenger type
	MessagingTemplate *Messagingtemplaterequest `json:"messagingTemplate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Sendagentlessoutboundmessagerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

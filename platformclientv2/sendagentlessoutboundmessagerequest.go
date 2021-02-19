package platformclientv2
import (
	"encoding/json"
)

// Sendagentlessoutboundmessagerequest
type Sendagentlessoutboundmessagerequest struct { 
	// FromAddress - The messaging address of the sender of the message. For an SMS messenger type, this must be a currently provisioned sms phone number.
	FromAddress *string `json:"fromAddress,omitempty"`


	// ToAddress - The messaging address of the recipient of the message. For an SMS messenger type, the phone number address must be in E.164 format. E.g. +13175555555 or +34234234234.
	ToAddress *string `json:"toAddress,omitempty"`


	// ToAddressMessengerType - The recipient messaging address messenger type. Currently SMS is the only supported type.
	ToAddressMessengerType *string `json:"toAddressMessengerType,omitempty"`


	// TextBody - The text of the message to send
	TextBody *string `json:"textBody,omitempty"`

}

// String returns a JSON representation of the model
func (o *Sendagentlessoutboundmessagerequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

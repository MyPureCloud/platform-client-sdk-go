package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationnormalizedmessage - General rich media message structure with normalized feature support across many messaging channels.
type Conversationnormalizedmessage struct { 
	// Id - Unique ID of the message. Message receipts will have the same ID as the message they reference.
	Id *string `json:"id,omitempty"`


	// Channel - Channel-specific information that describes the message and the message channel/provider.
	Channel *Conversationmessagingchannel `json:"channel,omitempty"`


	// VarType - Message type.
	VarType *string `json:"type,omitempty"`


	// Text - Message text.
	Text *string `json:"text,omitempty"`


	// Content - List of content elements
	Content *[]Conversationmessagecontent `json:"content,omitempty"`


	// Status - Message receipt status, only used with type Receipt.
	Status *string `json:"status,omitempty"`


	// Reasons - List of reasons for a message receipt that indicates the message has failed. Only used with Failed status.
	Reasons *[]Conversationreason `json:"reasons,omitempty"`


	// OriginatingEntity - Specifies if this message was sent by a human agent or bot. The platform may use this to apply appropriate provider policies.
	OriginatingEntity *string `json:"originatingEntity,omitempty"`


	// IsFinalReceipt - Indicates if this is the last message receipt for this message, or if another message receipt can be expected.
	IsFinalReceipt *bool `json:"isFinalReceipt,omitempty"`


	// Direction - The direction of the message.
	Direction *string `json:"direction,omitempty"`


	// Metadata - Additional metadata about this message.
	Metadata *map[string]string `json:"metadata,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationnormalizedmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

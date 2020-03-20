package platformclientv2
import (
	"time"
	"encoding/json"
)

// Messagedata
type Messagedata struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ProviderMessageId - The unique identifier of the message from provider
	ProviderMessageId *string `json:"providerMessageId,omitempty"`


	// Timestamp - The time when the message was received or sent. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	Timestamp *time.Time `json:"timestamp,omitempty"`


	// FromAddress - The sender of the text message.
	FromAddress *string `json:"fromAddress,omitempty"`


	// ToAddress - The recipient of the text message.
	ToAddress *string `json:"toAddress,omitempty"`


	// Direction - The direction of the message.
	Direction *string `json:"direction,omitempty"`


	// MessengerType - Type of text messenger.
	MessengerType *string `json:"messengerType,omitempty"`


	// TextBody - The body of the text message.
	TextBody *string `json:"textBody,omitempty"`


	// Status - The status of the message.
	Status *string `json:"status,omitempty"`


	// Media - The media details associated to a message.
	Media *[]Messagemedia `json:"media,omitempty"`


	// Stickers - The sticker details associated to a message.
	Stickers *[]Messagesticker `json:"stickers,omitempty"`


	// CreatedBy - User who sent this message.
	CreatedBy *User `json:"createdBy,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Messagedata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

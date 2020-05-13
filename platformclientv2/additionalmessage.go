package platformclientv2
import (
	"encoding/json"
)

// Additionalmessage
type Additionalmessage struct { 
	// TextBody - The body of the text message.
	TextBody *string `json:"textBody,omitempty"`


	// MediaIds - The media ids associated with the text message.
	MediaIds *[]string `json:"mediaIds,omitempty"`


	// StickerIds - The sticker ids associated with the text message.
	StickerIds *[]string `json:"stickerIds,omitempty"`


	// MessagingTemplate - The messaging template use to send a predefined canned response with the message
	MessagingTemplate *Messagingtemplaterequest `json:"messagingTemplate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Additionalmessage) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

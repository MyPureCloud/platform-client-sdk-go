package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Additionalmessage
type Additionalmessage struct { 
	// TextBody - The body of the text message.
	TextBody *string `json:"textBody,omitempty"`


	// MediaIds - The media ids associated with the text message. See https://developer.genesys.cloud/api/rest/v2/conversations/messaging-media-upload for example usage.
	MediaIds *[]string `json:"mediaIds,omitempty"`


	// StickerIds - The sticker ids associated with the text message.
	StickerIds *[]string `json:"stickerIds,omitempty"`


	// MessagingTemplate - The messaging template use to send a predefined canned response with the message
	MessagingTemplate *Messagingtemplaterequest `json:"messagingTemplate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Additionalmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

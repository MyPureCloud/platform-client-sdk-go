package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (o *Additionalmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Additionalmessage
	
	return json.Marshal(&struct { 
		TextBody *string `json:"textBody,omitempty"`
		
		MediaIds *[]string `json:"mediaIds,omitempty"`
		
		StickerIds *[]string `json:"stickerIds,omitempty"`
		
		MessagingTemplate *Messagingtemplaterequest `json:"messagingTemplate,omitempty"`
		*Alias
	}{ 
		TextBody: o.TextBody,
		
		MediaIds: o.MediaIds,
		
		StickerIds: o.StickerIds,
		
		MessagingTemplate: o.MessagingTemplate,
		Alias:    (*Alias)(o),
	})
}

func (o *Additionalmessage) UnmarshalJSON(b []byte) error {
	var AdditionalmessageMap map[string]interface{}
	err := json.Unmarshal(b, &AdditionalmessageMap)
	if err != nil {
		return err
	}
	
	if TextBody, ok := AdditionalmessageMap["textBody"].(string); ok {
		o.TextBody = &TextBody
	}
	
	if MediaIds, ok := AdditionalmessageMap["mediaIds"].([]interface{}); ok {
		MediaIdsString, _ := json.Marshal(MediaIds)
		json.Unmarshal(MediaIdsString, &o.MediaIds)
	}
	
	if StickerIds, ok := AdditionalmessageMap["stickerIds"].([]interface{}); ok {
		StickerIdsString, _ := json.Marshal(StickerIds)
		json.Unmarshal(StickerIdsString, &o.StickerIds)
	}
	
	if MessagingTemplate, ok := AdditionalmessageMap["messagingTemplate"].(map[string]interface{}); ok {
		MessagingTemplateString, _ := json.Marshal(MessagingTemplate)
		json.Unmarshal(MessagingTemplateString, &o.MessagingTemplate)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Additionalmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

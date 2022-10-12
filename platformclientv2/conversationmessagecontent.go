package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmessagecontent - Message content element. If contentType = \"Attachment\" only one item is allowed.
type Conversationmessagecontent struct { 
	// ContentType - Type of this content element.
	ContentType *string `json:"contentType,omitempty"`


	// Location - Location content.
	Location *Conversationcontentlocation `json:"location,omitempty"`


	// Attachment - Attachment content.
	Attachment *Conversationcontentattachment `json:"attachment,omitempty"`


	// QuickReply - Quick reply content.
	QuickReply *Conversationcontentquickreply `json:"quickReply,omitempty"`


	// ButtonResponse - Button response content.
	ButtonResponse *Conversationcontentbuttonresponse `json:"buttonResponse,omitempty"`


	// Template - Template notification content.
	Template *Conversationcontentnotificationtemplate `json:"template,omitempty"`


	// Story - Ephemeral story content.
	Story *Conversationcontentstory `json:"story,omitempty"`


	// Card - Card content
	Card *Conversationcontentcard `json:"card,omitempty"`


	// Carousel - Carousel content
	Carousel *Conversationcontentcarousel `json:"carousel,omitempty"`


	// Text - Text content.
	Text *Conversationcontenttext `json:"text,omitempty"`


	// QuickReplyV2 - Quick reply V2 content.
	QuickReplyV2 *Conversationcontentquickreplyv2 `json:"quickReplyV2,omitempty"`

}

func (o *Conversationmessagecontent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationmessagecontent
	
	return json.Marshal(&struct { 
		ContentType *string `json:"contentType,omitempty"`
		
		Location *Conversationcontentlocation `json:"location,omitempty"`
		
		Attachment *Conversationcontentattachment `json:"attachment,omitempty"`
		
		QuickReply *Conversationcontentquickreply `json:"quickReply,omitempty"`
		
		ButtonResponse *Conversationcontentbuttonresponse `json:"buttonResponse,omitempty"`
		
		Template *Conversationcontentnotificationtemplate `json:"template,omitempty"`
		
		Story *Conversationcontentstory `json:"story,omitempty"`
		
		Card *Conversationcontentcard `json:"card,omitempty"`
		
		Carousel *Conversationcontentcarousel `json:"carousel,omitempty"`
		
		Text *Conversationcontenttext `json:"text,omitempty"`
		
		QuickReplyV2 *Conversationcontentquickreplyv2 `json:"quickReplyV2,omitempty"`
		*Alias
	}{ 
		ContentType: o.ContentType,
		
		Location: o.Location,
		
		Attachment: o.Attachment,
		
		QuickReply: o.QuickReply,
		
		ButtonResponse: o.ButtonResponse,
		
		Template: o.Template,
		
		Story: o.Story,
		
		Card: o.Card,
		
		Carousel: o.Carousel,
		
		Text: o.Text,
		
		QuickReplyV2: o.QuickReplyV2,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationmessagecontent) UnmarshalJSON(b []byte) error {
	var ConversationmessagecontentMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationmessagecontentMap)
	if err != nil {
		return err
	}
	
	if ContentType, ok := ConversationmessagecontentMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if Location, ok := ConversationmessagecontentMap["location"].(map[string]interface{}); ok {
		LocationString, _ := json.Marshal(Location)
		json.Unmarshal(LocationString, &o.Location)
	}
	
	if Attachment, ok := ConversationmessagecontentMap["attachment"].(map[string]interface{}); ok {
		AttachmentString, _ := json.Marshal(Attachment)
		json.Unmarshal(AttachmentString, &o.Attachment)
	}
	
	if QuickReply, ok := ConversationmessagecontentMap["quickReply"].(map[string]interface{}); ok {
		QuickReplyString, _ := json.Marshal(QuickReply)
		json.Unmarshal(QuickReplyString, &o.QuickReply)
	}
	
	if ButtonResponse, ok := ConversationmessagecontentMap["buttonResponse"].(map[string]interface{}); ok {
		ButtonResponseString, _ := json.Marshal(ButtonResponse)
		json.Unmarshal(ButtonResponseString, &o.ButtonResponse)
	}
	
	if Template, ok := ConversationmessagecontentMap["template"].(map[string]interface{}); ok {
		TemplateString, _ := json.Marshal(Template)
		json.Unmarshal(TemplateString, &o.Template)
	}
	
	if Story, ok := ConversationmessagecontentMap["story"].(map[string]interface{}); ok {
		StoryString, _ := json.Marshal(Story)
		json.Unmarshal(StoryString, &o.Story)
	}
	
	if Card, ok := ConversationmessagecontentMap["card"].(map[string]interface{}); ok {
		CardString, _ := json.Marshal(Card)
		json.Unmarshal(CardString, &o.Card)
	}
	
	if Carousel, ok := ConversationmessagecontentMap["carousel"].(map[string]interface{}); ok {
		CarouselString, _ := json.Marshal(Carousel)
		json.Unmarshal(CarouselString, &o.Carousel)
	}
	
	if Text, ok := ConversationmessagecontentMap["text"].(map[string]interface{}); ok {
		TextString, _ := json.Marshal(Text)
		json.Unmarshal(TextString, &o.Text)
	}
	
	if QuickReplyV2, ok := ConversationmessagecontentMap["quickReplyV2"].(map[string]interface{}); ok {
		QuickReplyV2String, _ := json.Marshal(QuickReplyV2)
		json.Unmarshal(QuickReplyV2String, &o.QuickReplyV2)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationmessagecontent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

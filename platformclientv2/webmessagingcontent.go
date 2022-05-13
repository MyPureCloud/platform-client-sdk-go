package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webmessagingcontent - Message content element.
type Webmessagingcontent struct { 
	// ContentType - Type of this content element. If contentType = \"Attachment\" only one item is allowed.
	ContentType *string `json:"contentType,omitempty"`


	// Attachment - Attachment content.
	Attachment *Webmessagingattachment `json:"attachment,omitempty"`


	// QuickReply - Quick reply content.
	QuickReply *Webmessagingquickreply `json:"quickReply,omitempty"`


	// ButtonResponse - Button response content.
	ButtonResponse *Webmessagingbuttonresponse `json:"buttonResponse,omitempty"`


	// Generic - Generic content.
	Generic *Webmessaginggeneric `json:"generic,omitempty"`


	// Card - Card content
	Card *Contentcard `json:"card,omitempty"`


	// Carousel - Carousel content
	Carousel *Contentcarousel `json:"carousel,omitempty"`

}

func (o *Webmessagingcontent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webmessagingcontent
	
	return json.Marshal(&struct { 
		ContentType *string `json:"contentType,omitempty"`
		
		Attachment *Webmessagingattachment `json:"attachment,omitempty"`
		
		QuickReply *Webmessagingquickreply `json:"quickReply,omitempty"`
		
		ButtonResponse *Webmessagingbuttonresponse `json:"buttonResponse,omitempty"`
		
		Generic *Webmessaginggeneric `json:"generic,omitempty"`
		
		Card *Contentcard `json:"card,omitempty"`
		
		Carousel *Contentcarousel `json:"carousel,omitempty"`
		*Alias
	}{ 
		ContentType: o.ContentType,
		
		Attachment: o.Attachment,
		
		QuickReply: o.QuickReply,
		
		ButtonResponse: o.ButtonResponse,
		
		Generic: o.Generic,
		
		Card: o.Card,
		
		Carousel: o.Carousel,
		Alias:    (*Alias)(o),
	})
}

func (o *Webmessagingcontent) UnmarshalJSON(b []byte) error {
	var WebmessagingcontentMap map[string]interface{}
	err := json.Unmarshal(b, &WebmessagingcontentMap)
	if err != nil {
		return err
	}
	
	if ContentType, ok := WebmessagingcontentMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if Attachment, ok := WebmessagingcontentMap["attachment"].(map[string]interface{}); ok {
		AttachmentString, _ := json.Marshal(Attachment)
		json.Unmarshal(AttachmentString, &o.Attachment)
	}
	
	if QuickReply, ok := WebmessagingcontentMap["quickReply"].(map[string]interface{}); ok {
		QuickReplyString, _ := json.Marshal(QuickReply)
		json.Unmarshal(QuickReplyString, &o.QuickReply)
	}
	
	if ButtonResponse, ok := WebmessagingcontentMap["buttonResponse"].(map[string]interface{}); ok {
		ButtonResponseString, _ := json.Marshal(ButtonResponse)
		json.Unmarshal(ButtonResponseString, &o.ButtonResponse)
	}
	
	if Generic, ok := WebmessagingcontentMap["generic"].(map[string]interface{}); ok {
		GenericString, _ := json.Marshal(Generic)
		json.Unmarshal(GenericString, &o.Generic)
	}
	
	if Card, ok := WebmessagingcontentMap["card"].(map[string]interface{}); ok {
		CardString, _ := json.Marshal(Card)
		json.Unmarshal(CardString, &o.Card)
	}
	
	if Carousel, ok := WebmessagingcontentMap["carousel"].(map[string]interface{}); ok {
		CarouselString, _ := json.Marshal(Carousel)
		json.Unmarshal(CarouselString, &o.Carousel)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Webmessagingcontent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

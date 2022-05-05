package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messagecontent - Message content element. If contentType = \"Attachment\" only one item is allowed.
type Messagecontent struct { 
	// ContentType - Type of this content element.
	ContentType *string `json:"contentType,omitempty"`


	// Location - Location content.
	Location *Contentlocation `json:"location,omitempty"`


	// Attachment - Attachment content.
	Attachment *Contentattachment `json:"attachment,omitempty"`


	// QuickReply - Quick reply content.
	QuickReply *Contentquickreply `json:"quickReply,omitempty"`


	// ButtonResponse - Button response content.
	ButtonResponse *Contentbuttonresponse `json:"buttonResponse,omitempty"`


	// Generic - Generic content (Deprecated).
	Generic *Contentgeneric `json:"generic,omitempty"`


	// List - List content (Deprecated).
	List *Contentlist `json:"list,omitempty"`


	// Template - Template notification content.
	Template *Contentnotificationtemplate `json:"template,omitempty"`


	// Reactions - A set of reactions to a message.
	Reactions *[]Contentreaction `json:"reactions,omitempty"`


	// Mention - Mention content.
	Mention *Messagingrecipient `json:"mention,omitempty"`


	// Postback - Structured message postback (Deprecated).
	Postback *Contentpostback `json:"postback,omitempty"`


	// Story - Ephemeral story content.
	Story *Contentstory `json:"story,omitempty"`


	// Card - Card content
	Card *Contentcard `json:"card,omitempty"`


	// Carousel - Carousel content
	Carousel *Contentcarousel `json:"carousel,omitempty"`

}

func (o *Messagecontent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messagecontent
	
	return json.Marshal(&struct { 
		ContentType *string `json:"contentType,omitempty"`
		
		Location *Contentlocation `json:"location,omitempty"`
		
		Attachment *Contentattachment `json:"attachment,omitempty"`
		
		QuickReply *Contentquickreply `json:"quickReply,omitempty"`
		
		ButtonResponse *Contentbuttonresponse `json:"buttonResponse,omitempty"`
		
		Generic *Contentgeneric `json:"generic,omitempty"`
		
		List *Contentlist `json:"list,omitempty"`
		
		Template *Contentnotificationtemplate `json:"template,omitempty"`
		
		Reactions *[]Contentreaction `json:"reactions,omitempty"`
		
		Mention *Messagingrecipient `json:"mention,omitempty"`
		
		Postback *Contentpostback `json:"postback,omitempty"`
		
		Story *Contentstory `json:"story,omitempty"`
		
		Card *Contentcard `json:"card,omitempty"`
		
		Carousel *Contentcarousel `json:"carousel,omitempty"`
		*Alias
	}{ 
		ContentType: o.ContentType,
		
		Location: o.Location,
		
		Attachment: o.Attachment,
		
		QuickReply: o.QuickReply,
		
		ButtonResponse: o.ButtonResponse,
		
		Generic: o.Generic,
		
		List: o.List,
		
		Template: o.Template,
		
		Reactions: o.Reactions,
		
		Mention: o.Mention,
		
		Postback: o.Postback,
		
		Story: o.Story,
		
		Card: o.Card,
		
		Carousel: o.Carousel,
		Alias:    (*Alias)(o),
	})
}

func (o *Messagecontent) UnmarshalJSON(b []byte) error {
	var MessagecontentMap map[string]interface{}
	err := json.Unmarshal(b, &MessagecontentMap)
	if err != nil {
		return err
	}
	
	if ContentType, ok := MessagecontentMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
	
	if Location, ok := MessagecontentMap["location"].(map[string]interface{}); ok {
		LocationString, _ := json.Marshal(Location)
		json.Unmarshal(LocationString, &o.Location)
	}
	
	if Attachment, ok := MessagecontentMap["attachment"].(map[string]interface{}); ok {
		AttachmentString, _ := json.Marshal(Attachment)
		json.Unmarshal(AttachmentString, &o.Attachment)
	}
	
	if QuickReply, ok := MessagecontentMap["quickReply"].(map[string]interface{}); ok {
		QuickReplyString, _ := json.Marshal(QuickReply)
		json.Unmarshal(QuickReplyString, &o.QuickReply)
	}
	
	if ButtonResponse, ok := MessagecontentMap["buttonResponse"].(map[string]interface{}); ok {
		ButtonResponseString, _ := json.Marshal(ButtonResponse)
		json.Unmarshal(ButtonResponseString, &o.ButtonResponse)
	}
	
	if Generic, ok := MessagecontentMap["generic"].(map[string]interface{}); ok {
		GenericString, _ := json.Marshal(Generic)
		json.Unmarshal(GenericString, &o.Generic)
	}
	
	if List, ok := MessagecontentMap["list"].(map[string]interface{}); ok {
		ListString, _ := json.Marshal(List)
		json.Unmarshal(ListString, &o.List)
	}
	
	if Template, ok := MessagecontentMap["template"].(map[string]interface{}); ok {
		TemplateString, _ := json.Marshal(Template)
		json.Unmarshal(TemplateString, &o.Template)
	}
	
	if Reactions, ok := MessagecontentMap["reactions"].([]interface{}); ok {
		ReactionsString, _ := json.Marshal(Reactions)
		json.Unmarshal(ReactionsString, &o.Reactions)
	}
	
	if Mention, ok := MessagecontentMap["mention"].(map[string]interface{}); ok {
		MentionString, _ := json.Marshal(Mention)
		json.Unmarshal(MentionString, &o.Mention)
	}
	
	if Postback, ok := MessagecontentMap["postback"].(map[string]interface{}); ok {
		PostbackString, _ := json.Marshal(Postback)
		json.Unmarshal(PostbackString, &o.Postback)
	}
	
	if Story, ok := MessagecontentMap["story"].(map[string]interface{}); ok {
		StoryString, _ := json.Marshal(Story)
		json.Unmarshal(StoryString, &o.Story)
	}
	
	if Card, ok := MessagecontentMap["card"].(map[string]interface{}); ok {
		CardString, _ := json.Marshal(Card)
		json.Unmarshal(CardString, &o.Card)
	}
	
	if Carousel, ok := MessagecontentMap["carousel"].(map[string]interface{}); ok {
		CarouselString, _ := json.Marshal(Carousel)
		json.Unmarshal(CarouselString, &o.Carousel)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Messagecontent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmessagecontent - Message content element.
type Conversationmessagecontent struct { 
	// ContentType - Type of this content element. If contentType = \"Attachment\" only one item is allowed.
	ContentType *string `json:"contentType,omitempty"`


	// Location - Location content.
	Location *Conversationcontentlocation `json:"location,omitempty"`


	// Attachment - Attachment content.
	Attachment *Conversationcontentattachment `json:"attachment,omitempty"`


	// QuickReply - Quick reply content.
	QuickReply *Conversationcontentquickreply `json:"quickReply,omitempty"`


	// Template - Template notification content.
	Template *Conversationcontentnotificationtemplate `json:"template,omitempty"`


	// ButtonResponse - Button response content.
	ButtonResponse *Conversationcontentbuttonresponse `json:"buttonResponse,omitempty"`


	// Generic - Generic Template Object
	Generic *Conversationcontentgeneric `json:"generic,omitempty"`

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
		
		Template *Conversationcontentnotificationtemplate `json:"template,omitempty"`
		
		ButtonResponse *Conversationcontentbuttonresponse `json:"buttonResponse,omitempty"`
		
		Generic *Conversationcontentgeneric `json:"generic,omitempty"`
		*Alias
	}{ 
		ContentType: o.ContentType,
		
		Location: o.Location,
		
		Attachment: o.Attachment,
		
		QuickReply: o.QuickReply,
		
		Template: o.Template,
		
		ButtonResponse: o.ButtonResponse,
		
		Generic: o.Generic,
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
	
	if Template, ok := ConversationmessagecontentMap["template"].(map[string]interface{}); ok {
		TemplateString, _ := json.Marshal(Template)
		json.Unmarshal(TemplateString, &o.Template)
	}
	
	if ButtonResponse, ok := ConversationmessagecontentMap["buttonResponse"].(map[string]interface{}); ok {
		ButtonResponseString, _ := json.Marshal(ButtonResponse)
		json.Unmarshal(ButtonResponseString, &o.ButtonResponse)
	}
	
	if Generic, ok := ConversationmessagecontentMap["generic"].(map[string]interface{}); ok {
		GenericString, _ := json.Marshal(Generic)
		json.Unmarshal(GenericString, &o.Generic)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationmessagecontent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

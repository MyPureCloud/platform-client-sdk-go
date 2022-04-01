package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforworkflowtopicconversationmessagecontent
type V2conversationmessagetypingeventforworkflowtopicconversationmessagecontent struct { 
	// ContentType
	ContentType *string `json:"contentType,omitempty"`


	// Location
	Location *V2conversationmessagetypingeventforworkflowtopicconversationcontentlocation `json:"location,omitempty"`


	// Story
	Story *V2conversationmessagetypingeventforworkflowtopicconversationcontentstory `json:"story,omitempty"`


	// Attachment
	Attachment *V2conversationmessagetypingeventforworkflowtopicconversationcontentattachment `json:"attachment,omitempty"`


	// QuickReply
	QuickReply *V2conversationmessagetypingeventforworkflowtopicconversationcontentquickreply `json:"quickReply,omitempty"`


	// Template
	Template *V2conversationmessagetypingeventforworkflowtopicconversationcontentnotificationtemplate `json:"template,omitempty"`


	// ButtonResponse
	ButtonResponse *V2conversationmessagetypingeventforworkflowtopicconversationcontentbuttonresponse `json:"buttonResponse,omitempty"`


	// Generic
	Generic *V2conversationmessagetypingeventforworkflowtopicconversationcontentgeneric `json:"generic,omitempty"`

}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationmessagecontent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforworkflowtopicconversationmessagecontent
	
	return json.Marshal(&struct { 
		ContentType *string `json:"contentType,omitempty"`
		
		Location *V2conversationmessagetypingeventforworkflowtopicconversationcontentlocation `json:"location,omitempty"`
		
		Story *V2conversationmessagetypingeventforworkflowtopicconversationcontentstory `json:"story,omitempty"`
		
		Attachment *V2conversationmessagetypingeventforworkflowtopicconversationcontentattachment `json:"attachment,omitempty"`
		
		QuickReply *V2conversationmessagetypingeventforworkflowtopicconversationcontentquickreply `json:"quickReply,omitempty"`
		
		Template *V2conversationmessagetypingeventforworkflowtopicconversationcontentnotificationtemplate `json:"template,omitempty"`
		
		ButtonResponse *V2conversationmessagetypingeventforworkflowtopicconversationcontentbuttonresponse `json:"buttonResponse,omitempty"`
		
		Generic *V2conversationmessagetypingeventforworkflowtopicconversationcontentgeneric `json:"generic,omitempty"`
		*Alias
	}{ 
		ContentType: o.ContentType,
		
		Location: o.Location,
		
		Story: o.Story,
		
		Attachment: o.Attachment,
		
		QuickReply: o.QuickReply,
		
		Template: o.Template,
		
		ButtonResponse: o.ButtonResponse,
		
		Generic: o.Generic,
		Alias:    (*Alias)(o),
	})
}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationmessagecontent) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforworkflowtopicconversationmessagecontentMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforworkflowtopicconversationmessagecontentMap)
	if err != nil {
		return err
	}
	
	if ContentType, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagecontentMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
	
	if Location, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagecontentMap["location"].(map[string]interface{}); ok {
		LocationString, _ := json.Marshal(Location)
		json.Unmarshal(LocationString, &o.Location)
	}
	
	if Story, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagecontentMap["story"].(map[string]interface{}); ok {
		StoryString, _ := json.Marshal(Story)
		json.Unmarshal(StoryString, &o.Story)
	}
	
	if Attachment, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagecontentMap["attachment"].(map[string]interface{}); ok {
		AttachmentString, _ := json.Marshal(Attachment)
		json.Unmarshal(AttachmentString, &o.Attachment)
	}
	
	if QuickReply, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagecontentMap["quickReply"].(map[string]interface{}); ok {
		QuickReplyString, _ := json.Marshal(QuickReply)
		json.Unmarshal(QuickReplyString, &o.QuickReply)
	}
	
	if Template, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagecontentMap["template"].(map[string]interface{}); ok {
		TemplateString, _ := json.Marshal(Template)
		json.Unmarshal(TemplateString, &o.Template)
	}
	
	if ButtonResponse, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagecontentMap["buttonResponse"].(map[string]interface{}); ok {
		ButtonResponseString, _ := json.Marshal(ButtonResponse)
		json.Unmarshal(ButtonResponseString, &o.ButtonResponse)
	}
	
	if Generic, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagecontentMap["generic"].(map[string]interface{}); ok {
		GenericString, _ := json.Marshal(Generic)
		json.Unmarshal(GenericString, &o.Generic)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforworkflowtopicconversationmessagecontent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

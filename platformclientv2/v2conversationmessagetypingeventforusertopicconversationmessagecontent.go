package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforusertopicconversationmessagecontent
type V2conversationmessagetypingeventforusertopicconversationmessagecontent struct { 
	// ContentType
	ContentType *string `json:"contentType,omitempty"`


	// Location
	Location *V2conversationmessagetypingeventforusertopicconversationcontentlocation `json:"location,omitempty"`


	// Story
	Story *V2conversationmessagetypingeventforusertopicconversationcontentstory `json:"story,omitempty"`


	// Attachment
	Attachment *V2conversationmessagetypingeventforusertopicconversationcontentattachment `json:"attachment,omitempty"`


	// QuickReply
	QuickReply *V2conversationmessagetypingeventforusertopicconversationcontentquickreply `json:"quickReply,omitempty"`


	// Template
	Template *V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplate `json:"template,omitempty"`


	// ButtonResponse
	ButtonResponse *V2conversationmessagetypingeventforusertopicconversationcontentbuttonresponse `json:"buttonResponse,omitempty"`


	// Generic
	Generic *V2conversationmessagetypingeventforusertopicconversationcontentgeneric `json:"generic,omitempty"`

}

func (o *V2conversationmessagetypingeventforusertopicconversationmessagecontent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforusertopicconversationmessagecontent
	
	return json.Marshal(&struct { 
		ContentType *string `json:"contentType,omitempty"`
		
		Location *V2conversationmessagetypingeventforusertopicconversationcontentlocation `json:"location,omitempty"`
		
		Story *V2conversationmessagetypingeventforusertopicconversationcontentstory `json:"story,omitempty"`
		
		Attachment *V2conversationmessagetypingeventforusertopicconversationcontentattachment `json:"attachment,omitempty"`
		
		QuickReply *V2conversationmessagetypingeventforusertopicconversationcontentquickreply `json:"quickReply,omitempty"`
		
		Template *V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplate `json:"template,omitempty"`
		
		ButtonResponse *V2conversationmessagetypingeventforusertopicconversationcontentbuttonresponse `json:"buttonResponse,omitempty"`
		
		Generic *V2conversationmessagetypingeventforusertopicconversationcontentgeneric `json:"generic,omitempty"`
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

func (o *V2conversationmessagetypingeventforusertopicconversationmessagecontent) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforusertopicconversationmessagecontentMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforusertopicconversationmessagecontentMap)
	if err != nil {
		return err
	}
	
	if ContentType, ok := V2conversationmessagetypingeventforusertopicconversationmessagecontentMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if Location, ok := V2conversationmessagetypingeventforusertopicconversationmessagecontentMap["location"].(map[string]interface{}); ok {
		LocationString, _ := json.Marshal(Location)
		json.Unmarshal(LocationString, &o.Location)
	}
	
	if Story, ok := V2conversationmessagetypingeventforusertopicconversationmessagecontentMap["story"].(map[string]interface{}); ok {
		StoryString, _ := json.Marshal(Story)
		json.Unmarshal(StoryString, &o.Story)
	}
	
	if Attachment, ok := V2conversationmessagetypingeventforusertopicconversationmessagecontentMap["attachment"].(map[string]interface{}); ok {
		AttachmentString, _ := json.Marshal(Attachment)
		json.Unmarshal(AttachmentString, &o.Attachment)
	}
	
	if QuickReply, ok := V2conversationmessagetypingeventforusertopicconversationmessagecontentMap["quickReply"].(map[string]interface{}); ok {
		QuickReplyString, _ := json.Marshal(QuickReply)
		json.Unmarshal(QuickReplyString, &o.QuickReply)
	}
	
	if Template, ok := V2conversationmessagetypingeventforusertopicconversationmessagecontentMap["template"].(map[string]interface{}); ok {
		TemplateString, _ := json.Marshal(Template)
		json.Unmarshal(TemplateString, &o.Template)
	}
	
	if ButtonResponse, ok := V2conversationmessagetypingeventforusertopicconversationmessagecontentMap["buttonResponse"].(map[string]interface{}); ok {
		ButtonResponseString, _ := json.Marshal(ButtonResponse)
		json.Unmarshal(ButtonResponseString, &o.ButtonResponse)
	}
	
	if Generic, ok := V2conversationmessagetypingeventforusertopicconversationmessagecontentMap["generic"].(map[string]interface{}); ok {
		GenericString, _ := json.Marshal(Generic)
		json.Unmarshal(GenericString, &o.Generic)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforusertopicconversationmessagecontent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

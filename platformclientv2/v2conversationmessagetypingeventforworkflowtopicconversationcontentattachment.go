package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforworkflowtopicconversationcontentattachment
type V2conversationmessagetypingeventforworkflowtopicconversationcontentattachment struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// MediaType
	MediaType *string `json:"mediaType,omitempty"`


	// Url
	Url *string `json:"url,omitempty"`


	// Mime
	Mime *string `json:"mime,omitempty"`


	// Text
	Text *string `json:"text,omitempty"`


	// Sha256
	Sha256 *string `json:"sha256,omitempty"`


	// Filename
	Filename *string `json:"filename,omitempty"`

}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationcontentattachment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforworkflowtopicconversationcontentattachment
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		Url *string `json:"url,omitempty"`
		
		Mime *string `json:"mime,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Sha256 *string `json:"sha256,omitempty"`
		
		Filename *string `json:"filename,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		MediaType: o.MediaType,
		
		Url: o.Url,
		
		Mime: o.Mime,
		
		Text: o.Text,
		
		Sha256: o.Sha256,
		
		Filename: o.Filename,
		Alias:    (*Alias)(o),
	})
}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationcontentattachment) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforworkflowtopicconversationcontentattachmentMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforworkflowtopicconversationcontentattachmentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := V2conversationmessagetypingeventforworkflowtopicconversationcontentattachmentMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if MediaType, ok := V2conversationmessagetypingeventforworkflowtopicconversationcontentattachmentMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
	
	if Url, ok := V2conversationmessagetypingeventforworkflowtopicconversationcontentattachmentMap["url"].(string); ok {
		o.Url = &Url
	}
	
	if Mime, ok := V2conversationmessagetypingeventforworkflowtopicconversationcontentattachmentMap["mime"].(string); ok {
		o.Mime = &Mime
	}
	
	if Text, ok := V2conversationmessagetypingeventforworkflowtopicconversationcontentattachmentMap["text"].(string); ok {
		o.Text = &Text
	}
	
	if Sha256, ok := V2conversationmessagetypingeventforworkflowtopicconversationcontentattachmentMap["sha256"].(string); ok {
		o.Sha256 = &Sha256
	}
	
	if Filename, ok := V2conversationmessagetypingeventforworkflowtopicconversationcontentattachmentMap["filename"].(string); ok {
		o.Filename = &Filename
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforworkflowtopicconversationcontentattachment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

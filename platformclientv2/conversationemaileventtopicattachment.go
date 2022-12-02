package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationemaileventtopicattachment
type Conversationemaileventtopicattachment struct { 
	// AttachmentId - The unique identifier for the attachment.
	AttachmentId *string `json:"attachmentId,omitempty"`


	// Name - The name of the attachment.
	Name *string `json:"name,omitempty"`


	// ContentUri - The content uri of the attachment. If set, this is commonly a public api download location.
	ContentUri *string `json:"contentUri,omitempty"`


	// ContentType - The type of file the attachment is.
	ContentType *string `json:"contentType,omitempty"`


	// ContentLength - The length of the attachment file.
	ContentLength *int `json:"contentLength,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

func (o *Conversationemaileventtopicattachment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationemaileventtopicattachment
	
	return json.Marshal(&struct { 
		AttachmentId *string `json:"attachmentId,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ContentUri *string `json:"contentUri,omitempty"`
		
		ContentType *string `json:"contentType,omitempty"`
		
		ContentLength *int `json:"contentLength,omitempty"`
		
		AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		AttachmentId: o.AttachmentId,
		
		Name: o.Name,
		
		ContentUri: o.ContentUri,
		
		ContentType: o.ContentType,
		
		ContentLength: o.ContentLength,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationemaileventtopicattachment) UnmarshalJSON(b []byte) error {
	var ConversationemaileventtopicattachmentMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationemaileventtopicattachmentMap)
	if err != nil {
		return err
	}
	
	if AttachmentId, ok := ConversationemaileventtopicattachmentMap["attachmentId"].(string); ok {
		o.AttachmentId = &AttachmentId
	}
    
	if Name, ok := ConversationemaileventtopicattachmentMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ContentUri, ok := ConversationemaileventtopicattachmentMap["contentUri"].(string); ok {
		o.ContentUri = &ContentUri
	}
    
	if ContentType, ok := ConversationemaileventtopicattachmentMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if ContentLength, ok := ConversationemaileventtopicattachmentMap["contentLength"].(float64); ok {
		ContentLengthInt := int(ContentLength)
		o.ContentLength = &ContentLengthInt
	}
	
	if AdditionalProperties, ok := ConversationemaileventtopicattachmentMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationemaileventtopicattachment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

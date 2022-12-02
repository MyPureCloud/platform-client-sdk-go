package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationvideoeventtopicattachment
type Queueconversationvideoeventtopicattachment struct { 
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

func (o *Queueconversationvideoeventtopicattachment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationvideoeventtopicattachment
	
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

func (o *Queueconversationvideoeventtopicattachment) UnmarshalJSON(b []byte) error {
	var QueueconversationvideoeventtopicattachmentMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationvideoeventtopicattachmentMap)
	if err != nil {
		return err
	}
	
	if AttachmentId, ok := QueueconversationvideoeventtopicattachmentMap["attachmentId"].(string); ok {
		o.AttachmentId = &AttachmentId
	}
    
	if Name, ok := QueueconversationvideoeventtopicattachmentMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ContentUri, ok := QueueconversationvideoeventtopicattachmentMap["contentUri"].(string); ok {
		o.ContentUri = &ContentUri
	}
    
	if ContentType, ok := QueueconversationvideoeventtopicattachmentMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if ContentLength, ok := QueueconversationvideoeventtopicattachmentMap["contentLength"].(float64); ok {
		ContentLengthInt := int(ContentLength)
		o.ContentLength = &ContentLengthInt
	}
	
	if AdditionalProperties, ok := QueueconversationvideoeventtopicattachmentMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicattachment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Attachment
type Attachment struct { 
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


	// InlineImage - Whether or not the attachment was attached inline.,
	InlineImage *bool `json:"inlineImage,omitempty"`

}

func (o *Attachment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Attachment
	
	return json.Marshal(&struct { 
		AttachmentId *string `json:"attachmentId,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ContentUri *string `json:"contentUri,omitempty"`
		
		ContentType *string `json:"contentType,omitempty"`
		
		ContentLength *int `json:"contentLength,omitempty"`
		
		InlineImage *bool `json:"inlineImage,omitempty"`
		*Alias
	}{ 
		AttachmentId: o.AttachmentId,
		
		Name: o.Name,
		
		ContentUri: o.ContentUri,
		
		ContentType: o.ContentType,
		
		ContentLength: o.ContentLength,
		
		InlineImage: o.InlineImage,
		Alias:    (*Alias)(o),
	})
}

func (o *Attachment) UnmarshalJSON(b []byte) error {
	var AttachmentMap map[string]interface{}
	err := json.Unmarshal(b, &AttachmentMap)
	if err != nil {
		return err
	}
	
	if AttachmentId, ok := AttachmentMap["attachmentId"].(string); ok {
		o.AttachmentId = &AttachmentId
	}
	
	if Name, ok := AttachmentMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if ContentUri, ok := AttachmentMap["contentUri"].(string); ok {
		o.ContentUri = &ContentUri
	}
	
	if ContentType, ok := AttachmentMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
	
	if ContentLength, ok := AttachmentMap["contentLength"].(float64); ok {
		ContentLengthInt := int(ContentLength)
		o.ContentLength = &ContentLengthInt
	}
	
	if InlineImage, ok := AttachmentMap["inlineImage"].(bool); ok {
		o.InlineImage = &InlineImage
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Attachment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

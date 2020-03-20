package platformclientv2
import (
	"encoding/json"
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
	ContentLength *int32 `json:"contentLength,omitempty"`


	// InlineImage - Whether or not the attachment was attached inline.,
	InlineImage *bool `json:"inlineImage,omitempty"`

}

// String returns a JSON representation of the model
func (o *Attachment) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

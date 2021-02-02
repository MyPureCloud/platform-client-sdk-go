package platformclientv2
import (
	"encoding/json"
)

// Emailattachment
type Emailattachment struct { 
	// Name
	Name *string `json:"name,omitempty"`


	// ContentPath
	ContentPath *string `json:"contentPath,omitempty"`


	// ContentType
	ContentType *string `json:"contentType,omitempty"`


	// AttachmentId
	AttachmentId *string `json:"attachmentId,omitempty"`


	// ContentLength
	ContentLength *int `json:"contentLength,omitempty"`

}

// String returns a JSON representation of the model
func (o *Emailattachment) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

package platformclientv2
import (
	"encoding/json"
)

// Queueconversationeventtopicattachment
type Queueconversationeventtopicattachment struct { 
	// AttachmentId
	AttachmentId *string `json:"attachmentId,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ContentUri
	ContentUri *string `json:"contentUri,omitempty"`


	// ContentType
	ContentType *string `json:"contentType,omitempty"`


	// ContentLength
	ContentLength *int `json:"contentLength,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationeventtopicattachment) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

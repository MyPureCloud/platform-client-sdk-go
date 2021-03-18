package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationemaileventtopicattachment
type Conversationemaileventtopicattachment struct { 
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
func (o *Conversationemaileventtopicattachment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

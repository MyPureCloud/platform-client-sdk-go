package platformclientv2
import (
	"encoding/json"
)

// Batchdownloadrequest
type Batchdownloadrequest struct { 
	// ConversationId - Conversation id requested
	ConversationId *string `json:"conversationId,omitempty"`


	// RecordingId - Recording id requested, optional.  Leave null for all recordings on the conversation
	RecordingId *string `json:"recordingId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Batchdownloadrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

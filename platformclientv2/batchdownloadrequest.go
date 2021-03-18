package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

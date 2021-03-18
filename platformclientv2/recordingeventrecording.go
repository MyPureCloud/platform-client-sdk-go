package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingeventrecording
type Recordingeventrecording struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`


	// FileState
	FileState *string `json:"fileState,omitempty"`


	// MediaUris
	MediaUris *[]Recordingeventmediaresult `json:"mediaUris,omitempty"`


	// EstimatedTranscodeTimeMs
	EstimatedTranscodeTimeMs *float32 `json:"estimatedTranscodeTimeMs,omitempty"`


	// ActualTranscodeTimeMs
	ActualTranscodeTimeMs *float32 `json:"actualTranscodeTimeMs,omitempty"`

}

// String returns a JSON representation of the model
func (o *Recordingeventrecording) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

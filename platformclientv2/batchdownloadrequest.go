package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (o *Batchdownloadrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Batchdownloadrequest
	
	return json.Marshal(&struct { 
		ConversationId *string `json:"conversationId,omitempty"`
		
		RecordingId *string `json:"recordingId,omitempty"`
		*Alias
	}{ 
		ConversationId: o.ConversationId,
		
		RecordingId: o.RecordingId,
		Alias:    (*Alias)(o),
	})
}

func (o *Batchdownloadrequest) UnmarshalJSON(b []byte) error {
	var BatchdownloadrequestMap map[string]interface{}
	err := json.Unmarshal(b, &BatchdownloadrequestMap)
	if err != nil {
		return err
	}
	
	if ConversationId, ok := BatchdownloadrequestMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if RecordingId, ok := BatchdownloadrequestMap["recordingId"].(string); ok {
		o.RecordingId = &RecordingId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Batchdownloadrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingtranscodecompletetopicrecording
type Recordingtranscodecompletetopicrecording struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`


	// FileState
	FileState *string `json:"fileState,omitempty"`


	// MediaUris
	MediaUris *[]Recordingtranscodecompletetopicmediaresult `json:"mediaUris,omitempty"`


	// EstimatedTranscodeTimeMs
	EstimatedTranscodeTimeMs *float32 `json:"estimatedTranscodeTimeMs,omitempty"`


	// ActualTranscodeTimeMs
	ActualTranscodeTimeMs *float32 `json:"actualTranscodeTimeMs,omitempty"`

}

func (u *Recordingtranscodecompletetopicrecording) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recordingtranscodecompletetopicrecording

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		FileState *string `json:"fileState,omitempty"`
		
		MediaUris *[]Recordingtranscodecompletetopicmediaresult `json:"mediaUris,omitempty"`
		
		EstimatedTranscodeTimeMs *float32 `json:"estimatedTranscodeTimeMs,omitempty"`
		
		ActualTranscodeTimeMs *float32 `json:"actualTranscodeTimeMs,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		ConversationId: u.ConversationId,
		
		FileState: u.FileState,
		
		MediaUris: u.MediaUris,
		
		EstimatedTranscodeTimeMs: u.EstimatedTranscodeTimeMs,
		
		ActualTranscodeTimeMs: u.ActualTranscodeTimeMs,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Recordingtranscodecompletetopicrecording) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

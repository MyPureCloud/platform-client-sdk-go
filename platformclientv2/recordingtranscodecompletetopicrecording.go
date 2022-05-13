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
	EstimatedTranscodeTimeMs *int `json:"estimatedTranscodeTimeMs,omitempty"`


	// ActualTranscodeTimeMs
	ActualTranscodeTimeMs *int `json:"actualTranscodeTimeMs,omitempty"`

}

func (o *Recordingtranscodecompletetopicrecording) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recordingtranscodecompletetopicrecording
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		FileState *string `json:"fileState,omitempty"`
		
		MediaUris *[]Recordingtranscodecompletetopicmediaresult `json:"mediaUris,omitempty"`
		
		EstimatedTranscodeTimeMs *int `json:"estimatedTranscodeTimeMs,omitempty"`
		
		ActualTranscodeTimeMs *int `json:"actualTranscodeTimeMs,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		ConversationId: o.ConversationId,
		
		FileState: o.FileState,
		
		MediaUris: o.MediaUris,
		
		EstimatedTranscodeTimeMs: o.EstimatedTranscodeTimeMs,
		
		ActualTranscodeTimeMs: o.ActualTranscodeTimeMs,
		Alias:    (*Alias)(o),
	})
}

func (o *Recordingtranscodecompletetopicrecording) UnmarshalJSON(b []byte) error {
	var RecordingtranscodecompletetopicrecordingMap map[string]interface{}
	err := json.Unmarshal(b, &RecordingtranscodecompletetopicrecordingMap)
	if err != nil {
		return err
	}
	
	if Id, ok := RecordingtranscodecompletetopicrecordingMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if ConversationId, ok := RecordingtranscodecompletetopicrecordingMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if FileState, ok := RecordingtranscodecompletetopicrecordingMap["fileState"].(string); ok {
		o.FileState = &FileState
	}
    
	if MediaUris, ok := RecordingtranscodecompletetopicrecordingMap["mediaUris"].([]interface{}); ok {
		MediaUrisString, _ := json.Marshal(MediaUris)
		json.Unmarshal(MediaUrisString, &o.MediaUris)
	}
	
	if EstimatedTranscodeTimeMs, ok := RecordingtranscodecompletetopicrecordingMap["estimatedTranscodeTimeMs"].(float64); ok {
		EstimatedTranscodeTimeMsInt := int(EstimatedTranscodeTimeMs)
		o.EstimatedTranscodeTimeMs = &EstimatedTranscodeTimeMsInt
	}
	
	if ActualTranscodeTimeMs, ok := RecordingtranscodecompletetopicrecordingMap["actualTranscodeTimeMs"].(float64); ok {
		ActualTranscodeTimeMsInt := int(ActualTranscodeTimeMs)
		o.ActualTranscodeTimeMs = &ActualTranscodeTimeMsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Recordingtranscodecompletetopicrecording) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

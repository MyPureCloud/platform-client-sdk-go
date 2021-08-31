package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (o *Recordingeventrecording) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recordingeventrecording
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		FileState *string `json:"fileState,omitempty"`
		
		MediaUris *[]Recordingeventmediaresult `json:"mediaUris,omitempty"`
		
		EstimatedTranscodeTimeMs *float32 `json:"estimatedTranscodeTimeMs,omitempty"`
		
		ActualTranscodeTimeMs *float32 `json:"actualTranscodeTimeMs,omitempty"`
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

func (o *Recordingeventrecording) UnmarshalJSON(b []byte) error {
	var RecordingeventrecordingMap map[string]interface{}
	err := json.Unmarshal(b, &RecordingeventrecordingMap)
	if err != nil {
		return err
	}
	
	if Id, ok := RecordingeventrecordingMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if ConversationId, ok := RecordingeventrecordingMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
	
	if FileState, ok := RecordingeventrecordingMap["fileState"].(string); ok {
		o.FileState = &FileState
	}
	
	if MediaUris, ok := RecordingeventrecordingMap["mediaUris"].([]interface{}); ok {
		MediaUrisString, _ := json.Marshal(MediaUris)
		json.Unmarshal(MediaUrisString, &o.MediaUris)
	}
	
	if EstimatedTranscodeTimeMs, ok := RecordingeventrecordingMap["estimatedTranscodeTimeMs"].(float64); ok {
		EstimatedTranscodeTimeMsFloat32 := float32(EstimatedTranscodeTimeMs)
		o.EstimatedTranscodeTimeMs = &EstimatedTranscodeTimeMsFloat32
	}
	
	if ActualTranscodeTimeMs, ok := RecordingeventrecordingMap["actualTranscodeTimeMs"].(float64); ok {
		ActualTranscodeTimeMsFloat32 := float32(ActualTranscodeTimeMs)
		o.ActualTranscodeTimeMs = &ActualTranscodeTimeMsFloat32
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Recordingeventrecording) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

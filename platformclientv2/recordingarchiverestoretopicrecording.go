package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingarchiverestoretopicrecording
type Recordingarchiverestoretopicrecording struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`


	// FileState
	FileState *string `json:"fileState,omitempty"`


	// MediaUris
	MediaUris *[]Recordingarchiverestoretopicmediaresult `json:"mediaUris,omitempty"`


	// EstimatedTranscodeTimeMs
	EstimatedTranscodeTimeMs *int `json:"estimatedTranscodeTimeMs,omitempty"`


	// ActualTranscodeTimeMs
	ActualTranscodeTimeMs *int `json:"actualTranscodeTimeMs,omitempty"`

}

func (o *Recordingarchiverestoretopicrecording) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recordingarchiverestoretopicrecording
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		FileState *string `json:"fileState,omitempty"`
		
		MediaUris *[]Recordingarchiverestoretopicmediaresult `json:"mediaUris,omitempty"`
		
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

func (o *Recordingarchiverestoretopicrecording) UnmarshalJSON(b []byte) error {
	var RecordingarchiverestoretopicrecordingMap map[string]interface{}
	err := json.Unmarshal(b, &RecordingarchiverestoretopicrecordingMap)
	if err != nil {
		return err
	}
	
	if Id, ok := RecordingarchiverestoretopicrecordingMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if ConversationId, ok := RecordingarchiverestoretopicrecordingMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if FileState, ok := RecordingarchiverestoretopicrecordingMap["fileState"].(string); ok {
		o.FileState = &FileState
	}
    
	if MediaUris, ok := RecordingarchiverestoretopicrecordingMap["mediaUris"].([]interface{}); ok {
		MediaUrisString, _ := json.Marshal(MediaUris)
		json.Unmarshal(MediaUrisString, &o.MediaUris)
	}
	
	if EstimatedTranscodeTimeMs, ok := RecordingarchiverestoretopicrecordingMap["estimatedTranscodeTimeMs"].(float64); ok {
		EstimatedTranscodeTimeMsInt := int(EstimatedTranscodeTimeMs)
		o.EstimatedTranscodeTimeMs = &EstimatedTranscodeTimeMsInt
	}
	
	if ActualTranscodeTimeMs, ok := RecordingarchiverestoretopicrecordingMap["actualTranscodeTimeMs"].(float64); ok {
		ActualTranscodeTimeMsInt := int(ActualTranscodeTimeMs)
		o.ActualTranscodeTimeMs = &ActualTranscodeTimeMsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Recordingarchiverestoretopicrecording) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

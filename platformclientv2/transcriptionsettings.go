package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Transcriptionsettings
type Transcriptionsettings struct { 
	// Transcription - Setting to enable/disable transcription capability
	Transcription *string `json:"transcription,omitempty"`


	// TranscriptionConfidenceThreshold - Configure confidence threshold. The possible values are from 1 to 100.
	TranscriptionConfidenceThreshold *int `json:"transcriptionConfidenceThreshold,omitempty"`


	// LowLatencyTranscriptionEnabled - Boolean flag indicating whether low latency transcription via Notification API is enabled
	LowLatencyTranscriptionEnabled *bool `json:"lowLatencyTranscriptionEnabled,omitempty"`


	// ContentSearchEnabled - Setting to enable/disable content search
	ContentSearchEnabled *bool `json:"contentSearchEnabled,omitempty"`

}

func (o *Transcriptionsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Transcriptionsettings
	
	return json.Marshal(&struct { 
		Transcription *string `json:"transcription,omitempty"`
		
		TranscriptionConfidenceThreshold *int `json:"transcriptionConfidenceThreshold,omitempty"`
		
		LowLatencyTranscriptionEnabled *bool `json:"lowLatencyTranscriptionEnabled,omitempty"`
		
		ContentSearchEnabled *bool `json:"contentSearchEnabled,omitempty"`
		*Alias
	}{ 
		Transcription: o.Transcription,
		
		TranscriptionConfidenceThreshold: o.TranscriptionConfidenceThreshold,
		
		LowLatencyTranscriptionEnabled: o.LowLatencyTranscriptionEnabled,
		
		ContentSearchEnabled: o.ContentSearchEnabled,
		Alias:    (*Alias)(o),
	})
}

func (o *Transcriptionsettings) UnmarshalJSON(b []byte) error {
	var TranscriptionsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &TranscriptionsettingsMap)
	if err != nil {
		return err
	}
	
	if Transcription, ok := TranscriptionsettingsMap["transcription"].(string); ok {
		o.Transcription = &Transcription
	}
    
	if TranscriptionConfidenceThreshold, ok := TranscriptionsettingsMap["transcriptionConfidenceThreshold"].(float64); ok {
		TranscriptionConfidenceThresholdInt := int(TranscriptionConfidenceThreshold)
		o.TranscriptionConfidenceThreshold = &TranscriptionConfidenceThresholdInt
	}
	
	if LowLatencyTranscriptionEnabled, ok := TranscriptionsettingsMap["lowLatencyTranscriptionEnabled"].(bool); ok {
		o.LowLatencyTranscriptionEnabled = &LowLatencyTranscriptionEnabled
	}
    
	if ContentSearchEnabled, ok := TranscriptionsettingsMap["contentSearchEnabled"].(bool); ok {
		o.ContentSearchEnabled = &ContentSearchEnabled
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Transcriptionsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

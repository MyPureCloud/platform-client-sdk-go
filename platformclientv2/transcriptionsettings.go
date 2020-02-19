package platformclientv2
import (
	"encoding/json"
)

// Transcriptionsettings
type Transcriptionsettings struct { 
	// Transcription - Setting to enable/disable transcription capability
	Transcription *string `json:"transcription,omitempty"`


	// TranscriptionConfidenceThreshold - Configure confidence threshold. The possible values are from 1 to 100.
	TranscriptionConfidenceThreshold *int32 `json:"transcriptionConfidenceThreshold,omitempty"`

}

// String returns a JSON representation of the model
func (o *Transcriptionsettings) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

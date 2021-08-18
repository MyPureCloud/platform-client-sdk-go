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


	// ContentSearchEnabled - Setting to enable/disable content search
	ContentSearchEnabled *bool `json:"contentSearchEnabled,omitempty"`

}

func (u *Transcriptionsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Transcriptionsettings

	

	return json.Marshal(&struct { 
		Transcription *string `json:"transcription,omitempty"`
		
		TranscriptionConfidenceThreshold *int `json:"transcriptionConfidenceThreshold,omitempty"`
		
		ContentSearchEnabled *bool `json:"contentSearchEnabled,omitempty"`
		*Alias
	}{ 
		Transcription: u.Transcription,
		
		TranscriptionConfidenceThreshold: u.TranscriptionConfidenceThreshold,
		
		ContentSearchEnabled: u.ContentSearchEnabled,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Transcriptionsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

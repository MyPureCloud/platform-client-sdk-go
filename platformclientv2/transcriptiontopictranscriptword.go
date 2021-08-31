package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Transcriptiontopictranscriptword
type Transcriptiontopictranscriptword struct { 
	// Confidence
	Confidence *float32 `json:"confidence,omitempty"`


	// StartTimeMs
	StartTimeMs *int `json:"startTimeMs,omitempty"`


	// OffsetMs
	OffsetMs *int `json:"offsetMs,omitempty"`


	// DurationMs
	DurationMs *int `json:"durationMs,omitempty"`


	// Word
	Word *string `json:"word,omitempty"`

}

func (o *Transcriptiontopictranscriptword) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Transcriptiontopictranscriptword
	
	return json.Marshal(&struct { 
		Confidence *float32 `json:"confidence,omitempty"`
		
		StartTimeMs *int `json:"startTimeMs,omitempty"`
		
		OffsetMs *int `json:"offsetMs,omitempty"`
		
		DurationMs *int `json:"durationMs,omitempty"`
		
		Word *string `json:"word,omitempty"`
		*Alias
	}{ 
		Confidence: o.Confidence,
		
		StartTimeMs: o.StartTimeMs,
		
		OffsetMs: o.OffsetMs,
		
		DurationMs: o.DurationMs,
		
		Word: o.Word,
		Alias:    (*Alias)(o),
	})
}

func (o *Transcriptiontopictranscriptword) UnmarshalJSON(b []byte) error {
	var TranscriptiontopictranscriptwordMap map[string]interface{}
	err := json.Unmarshal(b, &TranscriptiontopictranscriptwordMap)
	if err != nil {
		return err
	}
	
	if Confidence, ok := TranscriptiontopictranscriptwordMap["confidence"].(float64); ok {
		ConfidenceFloat32 := float32(Confidence)
		o.Confidence = &ConfidenceFloat32
	}
	
	if StartTimeMs, ok := TranscriptiontopictranscriptwordMap["startTimeMs"].(float64); ok {
		StartTimeMsInt := int(StartTimeMs)
		o.StartTimeMs = &StartTimeMsInt
	}
	
	if OffsetMs, ok := TranscriptiontopictranscriptwordMap["offsetMs"].(float64); ok {
		OffsetMsInt := int(OffsetMs)
		o.OffsetMs = &OffsetMsInt
	}
	
	if DurationMs, ok := TranscriptiontopictranscriptwordMap["durationMs"].(float64); ok {
		DurationMsInt := int(DurationMs)
		o.DurationMs = &DurationMsInt
	}
	
	if Word, ok := TranscriptiontopictranscriptwordMap["word"].(string); ok {
		o.Word = &Word
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Transcriptiontopictranscriptword) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

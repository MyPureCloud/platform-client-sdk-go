package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Transcriptiontopictranscriptalternative
type Transcriptiontopictranscriptalternative struct { 
	// Confidence
	Confidence *float32 `json:"confidence,omitempty"`


	// OffsetMs
	OffsetMs *int `json:"offsetMs,omitempty"`


	// DurationMs
	DurationMs *int `json:"durationMs,omitempty"`


	// Transcript
	Transcript *string `json:"transcript,omitempty"`


	// Words
	Words *[]Transcriptiontopictranscriptword `json:"words,omitempty"`

}

func (o *Transcriptiontopictranscriptalternative) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Transcriptiontopictranscriptalternative
	
	return json.Marshal(&struct { 
		Confidence *float32 `json:"confidence,omitempty"`
		
		OffsetMs *int `json:"offsetMs,omitempty"`
		
		DurationMs *int `json:"durationMs,omitempty"`
		
		Transcript *string `json:"transcript,omitempty"`
		
		Words *[]Transcriptiontopictranscriptword `json:"words,omitempty"`
		*Alias
	}{ 
		Confidence: o.Confidence,
		
		OffsetMs: o.OffsetMs,
		
		DurationMs: o.DurationMs,
		
		Transcript: o.Transcript,
		
		Words: o.Words,
		Alias:    (*Alias)(o),
	})
}

func (o *Transcriptiontopictranscriptalternative) UnmarshalJSON(b []byte) error {
	var TranscriptiontopictranscriptalternativeMap map[string]interface{}
	err := json.Unmarshal(b, &TranscriptiontopictranscriptalternativeMap)
	if err != nil {
		return err
	}
	
	if Confidence, ok := TranscriptiontopictranscriptalternativeMap["confidence"].(float64); ok {
		ConfidenceFloat32 := float32(Confidence)
		o.Confidence = &ConfidenceFloat32
	}
	
	if OffsetMs, ok := TranscriptiontopictranscriptalternativeMap["offsetMs"].(float64); ok {
		OffsetMsInt := int(OffsetMs)
		o.OffsetMs = &OffsetMsInt
	}
	
	if DurationMs, ok := TranscriptiontopictranscriptalternativeMap["durationMs"].(float64); ok {
		DurationMsInt := int(DurationMs)
		o.DurationMs = &DurationMsInt
	}
	
	if Transcript, ok := TranscriptiontopictranscriptalternativeMap["transcript"].(string); ok {
		o.Transcript = &Transcript
	}
	
	if Words, ok := TranscriptiontopictranscriptalternativeMap["words"].([]interface{}); ok {
		WordsString, _ := json.Marshal(Words)
		json.Unmarshal(WordsString, &o.Words)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Transcriptiontopictranscriptalternative) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

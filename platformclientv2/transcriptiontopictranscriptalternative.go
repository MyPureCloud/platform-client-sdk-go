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

func (u *Transcriptiontopictranscriptalternative) MarshalJSON() ([]byte, error) {
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
		Confidence: u.Confidence,
		
		OffsetMs: u.OffsetMs,
		
		DurationMs: u.DurationMs,
		
		Transcript: u.Transcript,
		
		Words: u.Words,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Transcriptiontopictranscriptalternative) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

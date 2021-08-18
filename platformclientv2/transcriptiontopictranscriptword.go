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

func (u *Transcriptiontopictranscriptword) MarshalJSON() ([]byte, error) {
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
		Confidence: u.Confidence,
		
		StartTimeMs: u.StartTimeMs,
		
		OffsetMs: u.OffsetMs,
		
		DurationMs: u.DurationMs,
		
		Word: u.Word,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Transcriptiontopictranscriptword) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Transcriptiontopictranscriptword) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

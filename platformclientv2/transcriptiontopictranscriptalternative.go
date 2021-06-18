package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Transcriptiontopictranscriptalternative) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

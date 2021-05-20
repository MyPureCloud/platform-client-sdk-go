package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Transcriptionstopictranscriptalternative
type Transcriptionstopictranscriptalternative struct { 
	// Confidence
	Confidence *float32 `json:"confidence,omitempty"`


	// Transcript
	Transcript *string `json:"transcript,omitempty"`


	// Words
	Words *[]Transcriptionstopictranscriptword `json:"words,omitempty"`

}

// String returns a JSON representation of the model
func (o *Transcriptionstopictranscriptalternative) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

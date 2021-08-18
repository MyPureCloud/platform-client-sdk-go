package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbottranscript - Data for a single bot flow transcript.
type Textbottranscript struct { 
	// Text - The text of the transcript item.
	Text *string `json:"text,omitempty"`


	// Confidence - The confidence factor, expressed as a decimal between 0.0 and 1.0, of the transcript item.
	Confidence *float32 `json:"confidence,omitempty"`

}

func (u *Textbottranscript) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbottranscript

	

	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		Confidence *float32 `json:"confidence,omitempty"`
		*Alias
	}{ 
		Text: u.Text,
		
		Confidence: u.Confidence,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Textbottranscript) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

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

func (o *Textbottranscript) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbottranscript
	
	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		Confidence *float32 `json:"confidence,omitempty"`
		*Alias
	}{ 
		Text: o.Text,
		
		Confidence: o.Confidence,
		Alias:    (*Alias)(o),
	})
}

func (o *Textbottranscript) UnmarshalJSON(b []byte) error {
	var TextbottranscriptMap map[string]interface{}
	err := json.Unmarshal(b, &TextbottranscriptMap)
	if err != nil {
		return err
	}
	
	if Text, ok := TextbottranscriptMap["text"].(string); ok {
		o.Text = &Text
	}
	
	if Confidence, ok := TextbottranscriptMap["confidence"].(float64); ok {
		ConfidenceFloat32 := float32(Confidence)
		o.Confidence = &ConfidenceFloat32
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Textbottranscript) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotuserinputalternative - User input data used in a bot flow turn.
type Textbotuserinputalternative struct { 
	// Transcript - The user input transcript.
	Transcript *Textbottranscript `json:"transcript,omitempty"`

}

func (o *Textbotuserinputalternative) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbotuserinputalternative
	
	return json.Marshal(&struct { 
		Transcript *Textbottranscript `json:"transcript,omitempty"`
		*Alias
	}{ 
		Transcript: o.Transcript,
		Alias:    (*Alias)(o),
	})
}

func (o *Textbotuserinputalternative) UnmarshalJSON(b []byte) error {
	var TextbotuserinputalternativeMap map[string]interface{}
	err := json.Unmarshal(b, &TextbotuserinputalternativeMap)
	if err != nil {
		return err
	}
	
	if Transcript, ok := TextbotuserinputalternativeMap["transcript"].(map[string]interface{}); ok {
		TranscriptString, _ := json.Marshal(Transcript)
		json.Unmarshal(TranscriptString, &o.Transcript)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Textbotuserinputalternative) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

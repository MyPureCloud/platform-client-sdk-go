package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Nludetectioninput
type Nludetectioninput struct { 
	// Text - The text to perform NLU detection on.
	Text *string `json:"text,omitempty"`


	// Language - Language of the version for multilingual detection, e.g. `en-us`, `de-de`
	Language *string `json:"language,omitempty"`

}

func (o *Nludetectioninput) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nludetectioninput
	
	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		Language *string `json:"language,omitempty"`
		*Alias
	}{ 
		Text: o.Text,
		
		Language: o.Language,
		Alias:    (*Alias)(o),
	})
}

func (o *Nludetectioninput) UnmarshalJSON(b []byte) error {
	var NludetectioninputMap map[string]interface{}
	err := json.Unmarshal(b, &NludetectioninputMap)
	if err != nil {
		return err
	}
	
	if Text, ok := NludetectioninputMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Language, ok := NludetectioninputMap["language"].(string); ok {
		o.Language = &Language
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Nludetectioninput) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

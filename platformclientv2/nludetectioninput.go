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

}

func (o *Nludetectioninput) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nludetectioninput
	
	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		*Alias
	}{ 
		Text: o.Text,
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
	

	return nil
}

// String returns a JSON representation of the model
func (o *Nludetectioninput) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

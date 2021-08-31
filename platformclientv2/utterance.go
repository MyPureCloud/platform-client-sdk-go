package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Utterance
type Utterance struct { 
	// UtteranceText - Utterance text
	UtteranceText *string `json:"utteranceText,omitempty"`

}

func (o *Utterance) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Utterance
	
	return json.Marshal(&struct { 
		UtteranceText *string `json:"utteranceText,omitempty"`
		*Alias
	}{ 
		UtteranceText: o.UtteranceText,
		Alias:    (*Alias)(o),
	})
}

func (o *Utterance) UnmarshalJSON(b []byte) error {
	var UtteranceMap map[string]interface{}
	err := json.Unmarshal(b, &UtteranceMap)
	if err != nil {
		return err
	}
	
	if UtteranceText, ok := UtteranceMap["utteranceText"].(string); ok {
		o.UtteranceText = &UtteranceText
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Utterance) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

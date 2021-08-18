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

func (u *Utterance) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Utterance

	

	return json.Marshal(&struct { 
		UtteranceText *string `json:"utteranceText,omitempty"`
		*Alias
	}{ 
		UtteranceText: u.UtteranceText,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Utterance) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

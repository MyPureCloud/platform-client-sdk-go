package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Utterance
type Utterance struct { 
	// UtteranceText - Utterance text
	UtteranceText *string `json:"utteranceText,omitempty"`

}

// String returns a JSON representation of the model
func (o *Utterance) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

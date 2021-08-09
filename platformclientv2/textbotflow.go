package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotflow - Description of the Bot Flow.
type Textbotflow struct { 
	// Id - The Bot Flow ID.
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Textbotflow) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

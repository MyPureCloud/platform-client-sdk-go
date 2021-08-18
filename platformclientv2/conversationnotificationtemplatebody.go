package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationnotificationtemplatebody - Template body object.
type Conversationnotificationtemplatebody struct { 
	// Text - Body text. For WhatsApp, ignored.
	Text *string `json:"text,omitempty"`


	// Parameters - Template parameters for placeholders in template.
	Parameters *[]Conversationnotificationtemplateparameter `json:"parameters,omitempty"`

}

func (u *Conversationnotificationtemplatebody) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationnotificationtemplatebody

	

	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		Parameters *[]Conversationnotificationtemplateparameter `json:"parameters,omitempty"`
		*Alias
	}{ 
		Text: u.Text,
		
		Parameters: u.Parameters,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Conversationnotificationtemplatebody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

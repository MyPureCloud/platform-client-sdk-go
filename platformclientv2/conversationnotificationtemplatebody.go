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

func (o *Conversationnotificationtemplatebody) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationnotificationtemplatebody
	
	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		Parameters *[]Conversationnotificationtemplateparameter `json:"parameters,omitempty"`
		*Alias
	}{ 
		Text: o.Text,
		
		Parameters: o.Parameters,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationnotificationtemplatebody) UnmarshalJSON(b []byte) error {
	var ConversationnotificationtemplatebodyMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationnotificationtemplatebodyMap)
	if err != nil {
		return err
	}
	
	if Text, ok := ConversationnotificationtemplatebodyMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Parameters, ok := ConversationnotificationtemplatebodyMap["parameters"].([]interface{}); ok {
		ParametersString, _ := json.Marshal(Parameters)
		json.Unmarshal(ParametersString, &o.Parameters)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationnotificationtemplatebody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

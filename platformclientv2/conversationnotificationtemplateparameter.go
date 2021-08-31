package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationnotificationtemplateparameter - Template parameters for placeholders in template.
type Conversationnotificationtemplateparameter struct { 
	// Name - Parameter name.
	Name *string `json:"name,omitempty"`


	// Text - Parameter text value.
	Text *string `json:"text,omitempty"`

}

func (o *Conversationnotificationtemplateparameter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationnotificationtemplateparameter
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Text *string `json:"text,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Text: o.Text,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationnotificationtemplateparameter) UnmarshalJSON(b []byte) error {
	var ConversationnotificationtemplateparameterMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationnotificationtemplateparameterMap)
	if err != nil {
		return err
	}
	
	if Name, ok := ConversationnotificationtemplateparameterMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Text, ok := ConversationnotificationtemplateparameterMap["text"].(string); ok {
		o.Text = &Text
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationnotificationtemplateparameter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

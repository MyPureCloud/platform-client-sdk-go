package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Notificationtemplatebody - Template body object.
type Notificationtemplatebody struct { 
	// Text - Body text. For WhatsApp, ignored.
	Text *string `json:"text,omitempty"`


	// Parameters - Template parameters for placeholders in template.
	Parameters *[]Notificationtemplateparameter `json:"parameters,omitempty"`

}

func (o *Notificationtemplatebody) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Notificationtemplatebody
	
	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		Parameters *[]Notificationtemplateparameter `json:"parameters,omitempty"`
		*Alias
	}{ 
		Text: o.Text,
		
		Parameters: o.Parameters,
		Alias:    (*Alias)(o),
	})
}

func (o *Notificationtemplatebody) UnmarshalJSON(b []byte) error {
	var NotificationtemplatebodyMap map[string]interface{}
	err := json.Unmarshal(b, &NotificationtemplatebodyMap)
	if err != nil {
		return err
	}
	
	if Text, ok := NotificationtemplatebodyMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Parameters, ok := NotificationtemplatebodyMap["parameters"].([]interface{}); ok {
		ParametersString, _ := json.Marshal(Parameters)
		json.Unmarshal(ParametersString, &o.Parameters)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Notificationtemplatebody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

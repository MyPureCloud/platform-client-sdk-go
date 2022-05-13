package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Notificationtemplateparameter - Template parameters for placeholders in template.
type Notificationtemplateparameter struct { 
	// Name - Parameter name.
	Name *string `json:"name,omitempty"`


	// Text - Parameter text value.
	Text *string `json:"text,omitempty"`

}

func (o *Notificationtemplateparameter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Notificationtemplateparameter
	
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

func (o *Notificationtemplateparameter) UnmarshalJSON(b []byte) error {
	var NotificationtemplateparameterMap map[string]interface{}
	err := json.Unmarshal(b, &NotificationtemplateparameterMap)
	if err != nil {
		return err
	}
	
	if Name, ok := NotificationtemplateparameterMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Text, ok := NotificationtemplateparameterMap["text"].(string); ok {
		o.Text = &Text
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Notificationtemplateparameter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

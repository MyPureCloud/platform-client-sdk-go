package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Quickreply
type Quickreply struct { 
	// Text - Text to show inside the quick reply. This is also used as the response text after clicking on the quick reply.
	Text *string `json:"text,omitempty"`


	// Payload - Content of the textback payload after clicking a quick reply
	Payload *string `json:"payload,omitempty"`


	// Url - The location of the image file associated with quick reply
	Url *string `json:"url,omitempty"`


	// Action - Specifies the type of action that is triggered upon clicking the quick reply. Currently, the only supported action is \"Message\" which sends a message using the quick reply text.
	Action *string `json:"action,omitempty"`


	// IsSelected - Indicates if the quick reply option is selected by end customer
	IsSelected *bool `json:"isSelected,omitempty"`

}

func (o *Quickreply) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Quickreply
	
	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		Payload *string `json:"payload,omitempty"`
		
		Url *string `json:"url,omitempty"`
		
		Action *string `json:"action,omitempty"`
		
		IsSelected *bool `json:"isSelected,omitempty"`
		*Alias
	}{ 
		Text: o.Text,
		
		Payload: o.Payload,
		
		Url: o.Url,
		
		Action: o.Action,
		
		IsSelected: o.IsSelected,
		Alias:    (*Alias)(o),
	})
}

func (o *Quickreply) UnmarshalJSON(b []byte) error {
	var QuickreplyMap map[string]interface{}
	err := json.Unmarshal(b, &QuickreplyMap)
	if err != nil {
		return err
	}
	
	if Text, ok := QuickreplyMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Payload, ok := QuickreplyMap["payload"].(string); ok {
		o.Payload = &Payload
	}
    
	if Url, ok := QuickreplyMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if Action, ok := QuickreplyMap["action"].(string); ok {
		o.Action = &Action
	}
    
	if IsSelected, ok := QuickreplyMap["isSelected"].(bool); ok {
		o.IsSelected = &IsSelected
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Quickreply) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

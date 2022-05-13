package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforusertopicconversationnotificationtemplateparameter
type V2conversationmessagetypingeventforusertopicconversationnotificationtemplateparameter struct { 
	// Name
	Name *string `json:"name,omitempty"`


	// Text
	Text *string `json:"text,omitempty"`

}

func (o *V2conversationmessagetypingeventforusertopicconversationnotificationtemplateparameter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforusertopicconversationnotificationtemplateparameter
	
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

func (o *V2conversationmessagetypingeventforusertopicconversationnotificationtemplateparameter) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforusertopicconversationnotificationtemplateparameterMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforusertopicconversationnotificationtemplateparameterMap)
	if err != nil {
		return err
	}
	
	if Name, ok := V2conversationmessagetypingeventforusertopicconversationnotificationtemplateparameterMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Text, ok := V2conversationmessagetypingeventforusertopicconversationnotificationtemplateparameterMap["text"].(string); ok {
		o.Text = &Text
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforusertopicconversationnotificationtemplateparameter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

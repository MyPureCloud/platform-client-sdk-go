package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplateparameter
type V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplateparameter struct { 
	// Name
	Name *string `json:"name,omitempty"`


	// Text
	Text *string `json:"text,omitempty"`

}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplateparameter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplateparameter
	
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

func (o *V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplateparameter) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplateparameterMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplateparameterMap)
	if err != nil {
		return err
	}
	
	if Name, ok := V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplateparameterMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Text, ok := V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplateparameterMap["text"].(string); ok {
		o.Text = &Text
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplateparameter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

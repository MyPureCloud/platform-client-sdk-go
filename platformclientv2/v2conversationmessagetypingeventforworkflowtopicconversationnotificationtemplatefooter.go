package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplatefooter
type V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplatefooter struct { 
	// Text
	Text *string `json:"text,omitempty"`

}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplatefooter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplatefooter
	
	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		*Alias
	}{ 
		Text: o.Text,
		Alias:    (*Alias)(o),
	})
}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplatefooter) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplatefooterMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplatefooterMap)
	if err != nil {
		return err
	}
	
	if Text, ok := V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplatefooterMap["text"].(string); ok {
		o.Text = &Text
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplatefooter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

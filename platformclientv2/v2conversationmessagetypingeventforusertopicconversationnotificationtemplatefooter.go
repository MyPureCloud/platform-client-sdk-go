package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforusertopicconversationnotificationtemplatefooter
type V2conversationmessagetypingeventforusertopicconversationnotificationtemplatefooter struct { 
	// Text
	Text *string `json:"text,omitempty"`

}

func (o *V2conversationmessagetypingeventforusertopicconversationnotificationtemplatefooter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforusertopicconversationnotificationtemplatefooter
	
	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		*Alias
	}{ 
		Text: o.Text,
		Alias:    (*Alias)(o),
	})
}

func (o *V2conversationmessagetypingeventforusertopicconversationnotificationtemplatefooter) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforusertopicconversationnotificationtemplatefooterMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforusertopicconversationnotificationtemplatefooterMap)
	if err != nil {
		return err
	}
	
	if Text, ok := V2conversationmessagetypingeventforusertopicconversationnotificationtemplatefooterMap["text"].(string); ok {
		o.Text = &Text
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforusertopicconversationnotificationtemplatefooter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

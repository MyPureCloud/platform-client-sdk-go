package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforusertopicconversationnotificationtemplatebody
type V2conversationmessagetypingeventforusertopicconversationnotificationtemplatebody struct { 
	// Text
	Text *string `json:"text,omitempty"`


	// Parameters
	Parameters *[]V2conversationmessagetypingeventforusertopicconversationnotificationtemplateparameter `json:"parameters,omitempty"`

}

func (o *V2conversationmessagetypingeventforusertopicconversationnotificationtemplatebody) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforusertopicconversationnotificationtemplatebody
	
	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		Parameters *[]V2conversationmessagetypingeventforusertopicconversationnotificationtemplateparameter `json:"parameters,omitempty"`
		*Alias
	}{ 
		Text: o.Text,
		
		Parameters: o.Parameters,
		Alias:    (*Alias)(o),
	})
}

func (o *V2conversationmessagetypingeventforusertopicconversationnotificationtemplatebody) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforusertopicconversationnotificationtemplatebodyMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforusertopicconversationnotificationtemplatebodyMap)
	if err != nil {
		return err
	}
	
	if Text, ok := V2conversationmessagetypingeventforusertopicconversationnotificationtemplatebodyMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Parameters, ok := V2conversationmessagetypingeventforusertopicconversationnotificationtemplatebodyMap["parameters"].([]interface{}); ok {
		ParametersString, _ := json.Marshal(Parameters)
		json.Unmarshal(ParametersString, &o.Parameters)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforusertopicconversationnotificationtemplatebody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

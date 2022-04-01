package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforworkflowtopicconversationcontentbuttonresponse
type V2conversationmessagetypingeventforworkflowtopicconversationcontentbuttonresponse struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// Text
	Text *string `json:"text,omitempty"`


	// Payload
	Payload *string `json:"payload,omitempty"`

}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationcontentbuttonresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforworkflowtopicconversationcontentbuttonresponse
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Payload *string `json:"payload,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Text: o.Text,
		
		Payload: o.Payload,
		Alias:    (*Alias)(o),
	})
}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationcontentbuttonresponse) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforworkflowtopicconversationcontentbuttonresponseMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforworkflowtopicconversationcontentbuttonresponseMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := V2conversationmessagetypingeventforworkflowtopicconversationcontentbuttonresponseMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Text, ok := V2conversationmessagetypingeventforworkflowtopicconversationcontentbuttonresponseMap["text"].(string); ok {
		o.Text = &Text
	}
	
	if Payload, ok := V2conversationmessagetypingeventforworkflowtopicconversationcontentbuttonresponseMap["payload"].(string); ok {
		o.Payload = &Payload
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforworkflowtopicconversationcontentbuttonresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

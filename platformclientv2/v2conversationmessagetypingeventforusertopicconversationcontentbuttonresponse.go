package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforusertopicconversationcontentbuttonresponse
type V2conversationmessagetypingeventforusertopicconversationcontentbuttonresponse struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// Text
	Text *string `json:"text,omitempty"`


	// Payload
	Payload *string `json:"payload,omitempty"`

}

func (o *V2conversationmessagetypingeventforusertopicconversationcontentbuttonresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforusertopicconversationcontentbuttonresponse
	
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

func (o *V2conversationmessagetypingeventforusertopicconversationcontentbuttonresponse) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforusertopicconversationcontentbuttonresponseMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforusertopicconversationcontentbuttonresponseMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := V2conversationmessagetypingeventforusertopicconversationcontentbuttonresponseMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Text, ok := V2conversationmessagetypingeventforusertopicconversationcontentbuttonresponseMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Payload, ok := V2conversationmessagetypingeventforusertopicconversationcontentbuttonresponseMap["payload"].(string); ok {
		o.Payload = &Payload
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforusertopicconversationcontentbuttonresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

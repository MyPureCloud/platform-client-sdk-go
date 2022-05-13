package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforusertopicconversationreason
type V2conversationmessagetypingeventforusertopicconversationreason struct { 
	// Code
	Code *string `json:"code,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`

}

func (o *V2conversationmessagetypingeventforusertopicconversationreason) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforusertopicconversationreason
	
	return json.Marshal(&struct { 
		Code *string `json:"code,omitempty"`
		
		Message *string `json:"message,omitempty"`
		*Alias
	}{ 
		Code: o.Code,
		
		Message: o.Message,
		Alias:    (*Alias)(o),
	})
}

func (o *V2conversationmessagetypingeventforusertopicconversationreason) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforusertopicconversationreasonMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforusertopicconversationreasonMap)
	if err != nil {
		return err
	}
	
	if Code, ok := V2conversationmessagetypingeventforusertopicconversationreasonMap["code"].(string); ok {
		o.Code = &Code
	}
    
	if Message, ok := V2conversationmessagetypingeventforusertopicconversationreasonMap["message"].(string); ok {
		o.Message = &Message
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforusertopicconversationreason) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

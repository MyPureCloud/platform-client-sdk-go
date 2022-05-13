package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforworkflowtopicmessagedata
type V2conversationmessagetypingeventforworkflowtopicmessagedata struct { 
	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`


	// NormalizedMessage
	NormalizedMessage *V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessage `json:"normalizedMessage,omitempty"`

}

func (o *V2conversationmessagetypingeventforworkflowtopicmessagedata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforworkflowtopicmessagedata
	
	return json.Marshal(&struct { 
		ConversationId *string `json:"conversationId,omitempty"`
		
		NormalizedMessage *V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessage `json:"normalizedMessage,omitempty"`
		*Alias
	}{ 
		ConversationId: o.ConversationId,
		
		NormalizedMessage: o.NormalizedMessage,
		Alias:    (*Alias)(o),
	})
}

func (o *V2conversationmessagetypingeventforworkflowtopicmessagedata) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforworkflowtopicmessagedataMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforworkflowtopicmessagedataMap)
	if err != nil {
		return err
	}
	
	if ConversationId, ok := V2conversationmessagetypingeventforworkflowtopicmessagedataMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if NormalizedMessage, ok := V2conversationmessagetypingeventforworkflowtopicmessagedataMap["normalizedMessage"].(map[string]interface{}); ok {
		NormalizedMessageString, _ := json.Marshal(NormalizedMessage)
		json.Unmarshal(NormalizedMessageString, &o.NormalizedMessage)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforworkflowtopicmessagedata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

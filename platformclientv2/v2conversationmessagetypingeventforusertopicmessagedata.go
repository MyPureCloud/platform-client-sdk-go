package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforusertopicmessagedata
type V2conversationmessagetypingeventforusertopicmessagedata struct { 
	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`


	// NormalizedMessage
	NormalizedMessage *V2conversationmessagetypingeventforusertopicconversationnormalizedmessage `json:"normalizedMessage,omitempty"`

}

func (o *V2conversationmessagetypingeventforusertopicmessagedata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforusertopicmessagedata
	
	return json.Marshal(&struct { 
		ConversationId *string `json:"conversationId,omitempty"`
		
		NormalizedMessage *V2conversationmessagetypingeventforusertopicconversationnormalizedmessage `json:"normalizedMessage,omitempty"`
		*Alias
	}{ 
		ConversationId: o.ConversationId,
		
		NormalizedMessage: o.NormalizedMessage,
		Alias:    (*Alias)(o),
	})
}

func (o *V2conversationmessagetypingeventforusertopicmessagedata) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforusertopicmessagedataMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforusertopicmessagedataMap)
	if err != nil {
		return err
	}
	
	if ConversationId, ok := V2conversationmessagetypingeventforusertopicmessagedataMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if NormalizedMessage, ok := V2conversationmessagetypingeventforusertopicmessagedataMap["normalizedMessage"].(map[string]interface{}); ok {
		NormalizedMessageString, _ := json.Marshal(NormalizedMessage)
		json.Unmarshal(NormalizedMessageString, &o.NormalizedMessage)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforusertopicmessagedata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

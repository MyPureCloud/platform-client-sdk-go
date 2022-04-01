package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Callbackdisconnectidentifier
type Callbackdisconnectidentifier struct { 
	// ConversationId - The Conversation Id.
	ConversationId *string `json:"conversationId,omitempty"`


	// CallbackId - The callback id.
	CallbackId *string `json:"callbackId,omitempty"`

}

func (o *Callbackdisconnectidentifier) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Callbackdisconnectidentifier
	
	return json.Marshal(&struct { 
		ConversationId *string `json:"conversationId,omitempty"`
		
		CallbackId *string `json:"callbackId,omitempty"`
		*Alias
	}{ 
		ConversationId: o.ConversationId,
		
		CallbackId: o.CallbackId,
		Alias:    (*Alias)(o),
	})
}

func (o *Callbackdisconnectidentifier) UnmarshalJSON(b []byte) error {
	var CallbackdisconnectidentifierMap map[string]interface{}
	err := json.Unmarshal(b, &CallbackdisconnectidentifierMap)
	if err != nil {
		return err
	}
	
	if ConversationId, ok := CallbackdisconnectidentifierMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
	
	if CallbackId, ok := CallbackdisconnectidentifierMap["callbackId"].(string); ok {
		o.CallbackId = &CallbackId
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Callbackdisconnectidentifier) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

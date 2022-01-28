package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationappsettings - Conversation settings that handles chats within the messenger
type Conversationappsettings struct { 
	// ShowAgentTypingIndicator - The toggle to enable or disable typing indicator for messenger
	ShowAgentTypingIndicator *bool `json:"showAgentTypingIndicator,omitempty"`


	// ShowUserTypingIndicator - The toggle to enable or disable typing indicator for messenger
	ShowUserTypingIndicator *bool `json:"showUserTypingIndicator,omitempty"`

}

func (o *Conversationappsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationappsettings
	
	return json.Marshal(&struct { 
		ShowAgentTypingIndicator *bool `json:"showAgentTypingIndicator,omitempty"`
		
		ShowUserTypingIndicator *bool `json:"showUserTypingIndicator,omitempty"`
		*Alias
	}{ 
		ShowAgentTypingIndicator: o.ShowAgentTypingIndicator,
		
		ShowUserTypingIndicator: o.ShowUserTypingIndicator,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationappsettings) UnmarshalJSON(b []byte) error {
	var ConversationappsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationappsettingsMap)
	if err != nil {
		return err
	}
	
	if ShowAgentTypingIndicator, ok := ConversationappsettingsMap["showAgentTypingIndicator"].(bool); ok {
		o.ShowAgentTypingIndicator = &ShowAgentTypingIndicator
	}
	
	if ShowUserTypingIndicator, ok := ConversationappsettingsMap["showUserTypingIndicator"].(bool); ok {
		o.ShowUserTypingIndicator = &ShowUserTypingIndicator
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationappsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

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


	// AutoStartType - Deprecated. The auto start type for the messenger conversation
	AutoStartType *string `json:"autoStartType,omitempty"`


	// AutoStart - The auto start for the messenger conversation
	AutoStart *Autostart `json:"autoStart,omitempty"`


	// Markdown - The markdown for the messenger app
	Markdown *Markdown `json:"markdown,omitempty"`


	// ConversationDisconnect - The conversation disconnect settings for the messenger app
	ConversationDisconnect *Conversationdisconnectsettings `json:"conversationDisconnect,omitempty"`

}

func (o *Conversationappsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationappsettings
	
	return json.Marshal(&struct { 
		ShowAgentTypingIndicator *bool `json:"showAgentTypingIndicator,omitempty"`
		
		ShowUserTypingIndicator *bool `json:"showUserTypingIndicator,omitempty"`
		
		AutoStartType *string `json:"autoStartType,omitempty"`
		
		AutoStart *Autostart `json:"autoStart,omitempty"`
		
		Markdown *Markdown `json:"markdown,omitempty"`
		
		ConversationDisconnect *Conversationdisconnectsettings `json:"conversationDisconnect,omitempty"`
		*Alias
	}{ 
		ShowAgentTypingIndicator: o.ShowAgentTypingIndicator,
		
		ShowUserTypingIndicator: o.ShowUserTypingIndicator,
		
		AutoStartType: o.AutoStartType,
		
		AutoStart: o.AutoStart,
		
		Markdown: o.Markdown,
		
		ConversationDisconnect: o.ConversationDisconnect,
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
    
	if AutoStartType, ok := ConversationappsettingsMap["autoStartType"].(string); ok {
		o.AutoStartType = &AutoStartType
	}
    
	if AutoStart, ok := ConversationappsettingsMap["autoStart"].(map[string]interface{}); ok {
		AutoStartString, _ := json.Marshal(AutoStart)
		json.Unmarshal(AutoStartString, &o.AutoStart)
	}
	
	if Markdown, ok := ConversationappsettingsMap["markdown"].(map[string]interface{}); ok {
		MarkdownString, _ := json.Marshal(Markdown)
		json.Unmarshal(MarkdownString, &o.Markdown)
	}
	
	if ConversationDisconnect, ok := ConversationappsettingsMap["conversationDisconnect"].(map[string]interface{}); ok {
		ConversationDisconnectString, _ := json.Marshal(ConversationDisconnect)
		json.Unmarshal(ConversationDisconnectString, &o.ConversationDisconnect)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationappsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

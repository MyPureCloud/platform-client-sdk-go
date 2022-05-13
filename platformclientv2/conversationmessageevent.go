package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmessageevent - Message event element.  Examples include: system messages, typing indicators, cobrowse offerings.
type Conversationmessageevent struct { 
	// EventType - Type of this event element
	EventType *string `json:"eventType,omitempty"`


	// CoBrowse - CoBrowse event.
	CoBrowse *Conversationeventcobrowse `json:"coBrowse,omitempty"`


	// Typing - Typing event.
	Typing *Conversationeventtyping `json:"typing,omitempty"`


	// Presence - Presence event.
	Presence *Conversationeventpresence `json:"presence,omitempty"`

}

func (o *Conversationmessageevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationmessageevent
	
	return json.Marshal(&struct { 
		EventType *string `json:"eventType,omitempty"`
		
		CoBrowse *Conversationeventcobrowse `json:"coBrowse,omitempty"`
		
		Typing *Conversationeventtyping `json:"typing,omitempty"`
		
		Presence *Conversationeventpresence `json:"presence,omitempty"`
		*Alias
	}{ 
		EventType: o.EventType,
		
		CoBrowse: o.CoBrowse,
		
		Typing: o.Typing,
		
		Presence: o.Presence,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationmessageevent) UnmarshalJSON(b []byte) error {
	var ConversationmessageeventMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationmessageeventMap)
	if err != nil {
		return err
	}
	
	if EventType, ok := ConversationmessageeventMap["eventType"].(string); ok {
		o.EventType = &EventType
	}
    
	if CoBrowse, ok := ConversationmessageeventMap["coBrowse"].(map[string]interface{}); ok {
		CoBrowseString, _ := json.Marshal(CoBrowse)
		json.Unmarshal(CoBrowseString, &o.CoBrowse)
	}
	
	if Typing, ok := ConversationmessageeventMap["typing"].(map[string]interface{}); ok {
		TypingString, _ := json.Marshal(Typing)
		json.Unmarshal(TypingString, &o.Typing)
	}
	
	if Presence, ok := ConversationmessageeventMap["presence"].(map[string]interface{}); ok {
		PresenceString, _ := json.Marshal(Presence)
		json.Unmarshal(PresenceString, &o.Presence)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationmessageevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

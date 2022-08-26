package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messageevent - Message event element.  Examples include: system messages, typing indicators, cobrowse offerings.
type Messageevent struct { 
	// EventType - Type of this event element
	EventType *string `json:"eventType,omitempty"`


	// CoBrowse - CoBrowse event.
	CoBrowse *Eventcobrowse `json:"coBrowse,omitempty"`


	// Typing - Typing event.
	Typing *Eventtyping `json:"typing,omitempty"`


	// Presence - Presence event.
	Presence *Eventpresence `json:"presence,omitempty"`

}

func (o *Messageevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messageevent
	
	return json.Marshal(&struct { 
		EventType *string `json:"eventType,omitempty"`
		
		CoBrowse *Eventcobrowse `json:"coBrowse,omitempty"`
		
		Typing *Eventtyping `json:"typing,omitempty"`
		
		Presence *Eventpresence `json:"presence,omitempty"`
		*Alias
	}{ 
		EventType: o.EventType,
		
		CoBrowse: o.CoBrowse,
		
		Typing: o.Typing,
		
		Presence: o.Presence,
		Alias:    (*Alias)(o),
	})
}

func (o *Messageevent) UnmarshalJSON(b []byte) error {
	var MessageeventMap map[string]interface{}
	err := json.Unmarshal(b, &MessageeventMap)
	if err != nil {
		return err
	}
	
	if EventType, ok := MessageeventMap["eventType"].(string); ok {
		o.EventType = &EventType
	}
    
	if CoBrowse, ok := MessageeventMap["coBrowse"].(map[string]interface{}); ok {
		CoBrowseString, _ := json.Marshal(CoBrowse)
		json.Unmarshal(CoBrowseString, &o.CoBrowse)
	}
	
	if Typing, ok := MessageeventMap["typing"].(map[string]interface{}); ok {
		TypingString, _ := json.Marshal(Typing)
		json.Unmarshal(TypingString, &o.Typing)
	}
	
	if Presence, ok := MessageeventMap["presence"].(map[string]interface{}); ok {
		PresenceString, _ := json.Marshal(Presence)
		json.Unmarshal(PresenceString, &o.Presence)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Messageevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

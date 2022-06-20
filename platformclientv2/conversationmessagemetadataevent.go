package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmessagemetadataevent - Metadata information about a message event.
type Conversationmessagemetadataevent struct { 
	// EventType - Type of this event element
	EventType *string `json:"eventType,omitempty"`


	// SubType - Event subtype
	SubType *string `json:"subType,omitempty"`

}

func (o *Conversationmessagemetadataevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationmessagemetadataevent
	
	return json.Marshal(&struct { 
		EventType *string `json:"eventType,omitempty"`
		
		SubType *string `json:"subType,omitempty"`
		*Alias
	}{ 
		EventType: o.EventType,
		
		SubType: o.SubType,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationmessagemetadataevent) UnmarshalJSON(b []byte) error {
	var ConversationmessagemetadataeventMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationmessagemetadataeventMap)
	if err != nil {
		return err
	}
	
	if EventType, ok := ConversationmessagemetadataeventMap["eventType"].(string); ok {
		o.EventType = &EventType
	}
    
	if SubType, ok := ConversationmessagemetadataeventMap["subType"].(string); ok {
		o.SubType = &SubType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationmessagemetadataevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

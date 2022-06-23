package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmessageeventtopicmessagemetadataevent
type Conversationmessageeventtopicmessagemetadataevent struct { 
	// EventType - Type of this event element.
	EventType *string `json:"eventType,omitempty"`


	// SubType - Event subtype, if any
	SubType *string `json:"subType,omitempty"`

}

func (o *Conversationmessageeventtopicmessagemetadataevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationmessageeventtopicmessagemetadataevent
	
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

func (o *Conversationmessageeventtopicmessagemetadataevent) UnmarshalJSON(b []byte) error {
	var ConversationmessageeventtopicmessagemetadataeventMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationmessageeventtopicmessagemetadataeventMap)
	if err != nil {
		return err
	}
	
	if EventType, ok := ConversationmessageeventtopicmessagemetadataeventMap["eventType"].(string); ok {
		o.EventType = &EventType
	}
    
	if SubType, ok := ConversationmessageeventtopicmessagemetadataeventMap["subType"].(string); ok {
		o.SubType = &SubType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationmessageeventtopicmessagemetadataevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

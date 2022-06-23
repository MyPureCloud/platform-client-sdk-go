package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationsocialexpressioneventtopicmessagemetadataevent
type Queueconversationsocialexpressioneventtopicmessagemetadataevent struct { 
	// EventType - Type of this event element.
	EventType *string `json:"eventType,omitempty"`


	// SubType - Event subtype, if any
	SubType *string `json:"subType,omitempty"`

}

func (o *Queueconversationsocialexpressioneventtopicmessagemetadataevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationsocialexpressioneventtopicmessagemetadataevent
	
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

func (o *Queueconversationsocialexpressioneventtopicmessagemetadataevent) UnmarshalJSON(b []byte) error {
	var QueueconversationsocialexpressioneventtopicmessagemetadataeventMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationsocialexpressioneventtopicmessagemetadataeventMap)
	if err != nil {
		return err
	}
	
	if EventType, ok := QueueconversationsocialexpressioneventtopicmessagemetadataeventMap["eventType"].(string); ok {
		o.EventType = &EventType
	}
    
	if SubType, ok := QueueconversationsocialexpressioneventtopicmessagemetadataeventMap["subType"].(string); ok {
		o.SubType = &SubType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicmessagemetadataevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

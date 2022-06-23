package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationvideoeventtopicmessagemetadataevent
type Queueconversationvideoeventtopicmessagemetadataevent struct { 
	// EventType - Type of this event element.
	EventType *string `json:"eventType,omitempty"`


	// SubType - Event subtype, if any
	SubType *string `json:"subType,omitempty"`

}

func (o *Queueconversationvideoeventtopicmessagemetadataevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationvideoeventtopicmessagemetadataevent
	
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

func (o *Queueconversationvideoeventtopicmessagemetadataevent) UnmarshalJSON(b []byte) error {
	var QueueconversationvideoeventtopicmessagemetadataeventMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationvideoeventtopicmessagemetadataeventMap)
	if err != nil {
		return err
	}
	
	if EventType, ok := QueueconversationvideoeventtopicmessagemetadataeventMap["eventType"].(string); ok {
		o.EventType = &EventType
	}
    
	if SubType, ok := QueueconversationvideoeventtopicmessagemetadataeventMap["subType"].(string); ok {
		o.SubType = &SubType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicmessagemetadataevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationeventtopicmessagemetadataevent
type Queueconversationeventtopicmessagemetadataevent struct { 
	// EventType - Type of this event element.
	EventType *string `json:"eventType,omitempty"`


	// SubType - Event subtype, if any
	SubType *string `json:"subType,omitempty"`

}

func (o *Queueconversationeventtopicmessagemetadataevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationeventtopicmessagemetadataevent
	
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

func (o *Queueconversationeventtopicmessagemetadataevent) UnmarshalJSON(b []byte) error {
	var QueueconversationeventtopicmessagemetadataeventMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationeventtopicmessagemetadataeventMap)
	if err != nil {
		return err
	}
	
	if EventType, ok := QueueconversationeventtopicmessagemetadataeventMap["eventType"].(string); ok {
		o.EventType = &EventType
	}
    
	if SubType, ok := QueueconversationeventtopicmessagemetadataeventMap["subType"].(string); ok {
		o.SubType = &SubType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationeventtopicmessagemetadataevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

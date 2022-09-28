package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Openmessageevent - Message event element.
type Openmessageevent struct { 
	// EventType - Type of this event element
	EventType *string `json:"eventType,omitempty"`


	// Typing - Typing event.
	Typing *Eventtyping `json:"typing,omitempty"`

}

func (o *Openmessageevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Openmessageevent
	
	return json.Marshal(&struct { 
		EventType *string `json:"eventType,omitempty"`
		
		Typing *Eventtyping `json:"typing,omitempty"`
		*Alias
	}{ 
		EventType: o.EventType,
		
		Typing: o.Typing,
		Alias:    (*Alias)(o),
	})
}

func (o *Openmessageevent) UnmarshalJSON(b []byte) error {
	var OpenmessageeventMap map[string]interface{}
	err := json.Unmarshal(b, &OpenmessageeventMap)
	if err != nil {
		return err
	}
	
	if EventType, ok := OpenmessageeventMap["eventType"].(string); ok {
		o.EventType = &EventType
	}
    
	if Typing, ok := OpenmessageeventMap["typing"].(map[string]interface{}); ok {
		TypingString, _ := json.Marshal(Typing)
		json.Unmarshal(TypingString, &o.Typing)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Openmessageevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

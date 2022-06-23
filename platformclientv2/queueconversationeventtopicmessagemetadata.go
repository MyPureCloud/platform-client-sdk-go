package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationeventtopicmessagemetadata - Metadata information about a message.
type Queueconversationeventtopicmessagemetadata struct { 
	// VarType - Message type.
	VarType *string `json:"type,omitempty"`


	// Events - List of message events, if any
	Events *[]Queueconversationeventtopicmessagemetadataevent `json:"events,omitempty"`


	// Content - List of message content, if any
	Content *[]Queueconversationeventtopicmessagemetadatacontent `json:"content,omitempty"`

}

func (o *Queueconversationeventtopicmessagemetadata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationeventtopicmessagemetadata
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Events *[]Queueconversationeventtopicmessagemetadataevent `json:"events,omitempty"`
		
		Content *[]Queueconversationeventtopicmessagemetadatacontent `json:"content,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Events: o.Events,
		
		Content: o.Content,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationeventtopicmessagemetadata) UnmarshalJSON(b []byte) error {
	var QueueconversationeventtopicmessagemetadataMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationeventtopicmessagemetadataMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := QueueconversationeventtopicmessagemetadataMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Events, ok := QueueconversationeventtopicmessagemetadataMap["events"].([]interface{}); ok {
		EventsString, _ := json.Marshal(Events)
		json.Unmarshal(EventsString, &o.Events)
	}
	
	if Content, ok := QueueconversationeventtopicmessagemetadataMap["content"].([]interface{}); ok {
		ContentString, _ := json.Marshal(Content)
		json.Unmarshal(ContentString, &o.Content)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationeventtopicmessagemetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
